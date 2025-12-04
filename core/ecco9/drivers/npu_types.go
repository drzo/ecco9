package drivers

import (
	"sync"
	"time"
)

// NPU Memory-Mapped Register Layout
// Following hardware-style MMIO design in PERIPH space
const (
	// Base address in peripheral space
	NPU_REG_BASE = 0x40001000

	// Command and control registers
	NPU_REG_CMD            = NPU_REG_BASE + 0x00 // Command register
	NPU_REG_STATUS         = NPU_REG_BASE + 0x04 // Status register
	NPU_REG_PROMPT_ADDR    = NPU_REG_BASE + 0x08 // Prompt memory address
	NPU_REG_PROMPT_LEN     = NPU_REG_BASE + 0x0C // Prompt length
	NPU_REG_N_PREDICT      = NPU_REG_BASE + 0x10 // Number of tokens to predict
	NPU_REG_TOKEN_OUT      = NPU_REG_BASE + 0x14 // Output token
	NPU_REG_TOKEN_READY    = NPU_REG_BASE + 0x18 // Token ready flag
	NPU_REG_MODEL_ID       = NPU_REG_BASE + 0x1C // Model identifier
	NPU_REG_CTX_USED       = NPU_REG_BASE + 0x20 // Context tokens used
	NPU_REG_ERROR_CODE     = NPU_REG_BASE + 0x24 // Error code
	NPU_REG_PERF_TOKENS_SEC = NPU_REG_BASE + 0x28 // Performance: tokens/sec

	// Memory regions
	NPU_SRAM_BASE = 0x20000000 // Shared SRAM for prompts/KV-cache
	NPU_SRAM_SIZE = 0x10000000 // 256MB SRAM
)

// NPU Command bits
const (
	NPU_CMD_RESET      = 1 << 0 // Reset device state
	NPU_CMD_LOAD_MODEL = 1 << 1 // Load GGUF model
	NPU_CMD_START_INF  = 1 << 2 // Start inference
	NPU_CMD_SOFT_STOP  = 1 << 3 // Graceful stop
)

// NPU Status bits
const (
	NPU_STATUS_IDLE        = 1 << 0 // Device ready
	NPU_STATUS_BUSY        = 1 << 1 // Inference in progress
	NPU_STATUS_EOG         = 1 << 2 // End of generation
	NPU_STATUS_ERROR       = 1 << 3 // Error condition
	NPU_STATUS_MODEL_READY = 1 << 4 // Model loaded
	NPU_STATUS_TOKEN_READY = 1 << 5 // Token available
)

// NPU Error codes
const (
	NPU_ERR_NONE         = 0
	NPU_ERR_MODEL_LOAD   = 1
	NPU_ERR_INVALID_CMD  = 2
	NPU_ERR_OUT_OF_MEM   = 3
	NPU_ERR_TOKENIZATION = 4
	NPU_ERR_INFERENCE    = 5
	NPU_ERR_TIMEOUT      = 6
)

// NPUModelConfig configures the GGUF model
type NPUModelConfig struct {
	ModelPath        string // .gguf file path or model name
	ModelName        string // Friendly name
	NCtx             int32  // Context window size
	NThreads         int32  // CPU threads for inference
	NGPULayers       int32  // GPU layers (0=CPU only)
	BatchSize        int32  // Batch size
	OffloadKVCache   bool   // Offload KV cache to GPU
	LowVRAMMode      bool   // Enable low VRAM optimizations
	Temperature      float32
	TopP             float32
	TopK             int32
	RepeatPenalty    float32
}

// NPUSequenceConfig configures inference generation
type NPUSequenceConfig struct {
	NPredict      int32  // Max tokens to generate
	MaxCtx        int32  // Max context
	EchoPrompt    bool   // Echo prompt in output
	StreamTokens  bool   // Stream tokens as generated
	SystemPrompt  string // System prompt
	StopSequences []string // Stop sequences
}

// NPUTelemetry tracks performance metrics
type NPUTelemetry struct {
	mu                      sync.RWMutex
	TokensPerSecond         float64
	TotalTokensGenerated    uint64
	TotalPrompts            uint64
	LastPromptTokens        uint64
	LastCompletionTokens    uint64
	LastInferenceStart      time.Time
	LastInferenceEnd        time.Time
	LastInferenceDuration   time.Duration
	AverageTokensPerSecond  float64
	PeakTokensPerSecond     float64
}

// NPURegisters represents the hardware register state
type NPURegisters struct {
	mu             sync.RWMutex
	Command        uint32
	Status         uint32
	PromptAddr     uint64
	PromptLen      uint32
	NPredict       uint32
	TokenOut       int32
	TokenReady     uint32
	ModelID        uint32
	CtxUsed        uint32
	ErrorCode      uint32
	PerfTokensSec  uint32
}

// NPUMemoryRegion represents a memory-mapped region
type NPUMemoryRegion struct {
	BaseAddr uint64
	Size     uint64
	Data     []byte
	Name     string
}

// TokenCallback is called for each generated token in streaming mode
type TokenCallback func(tokenText string, tokenID int32, isLast bool)

// Default configurations
func DefaultNPUModelConfig() NPUModelConfig {
	return NPUModelConfig{
		ModelPath:      "",
		ModelName:      "default",
		NCtx:           4096,
		NThreads:       4,
		NGPULayers:     0,
		BatchSize:      1,
		OffloadKVCache: false,
		LowVRAMMode:    false,
		Temperature:    0.7,
		TopP:           0.9,
		TopK:           40,
		RepeatPenalty:  1.1,
	}
}

func DefaultNPUSequenceConfig() NPUSequenceConfig {
	return NPUSequenceConfig{
		NPredict:      128,
		MaxCtx:        4096,
		EchoPrompt:    false,
		StreamTokens:  true,
		SystemPrompt:  "",
		StopSequences: []string{},
	}
}

// NewNPUTelemetry creates a new telemetry tracker
func NewNPUTelemetry() *NPUTelemetry {
	return &NPUTelemetry{}
}

// UpdateTokenGeneration updates telemetry after token generation
func (t *NPUTelemetry) UpdateTokenGeneration(tokensGenerated uint64, duration time.Duration) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.TotalTokensGenerated += tokensGenerated
	t.LastCompletionTokens = tokensGenerated
	t.LastInferenceDuration = duration

	if duration.Seconds() > 0 {
		t.TokensPerSecond = float64(tokensGenerated) / duration.Seconds()
		
		// Update average (simple moving average)
		if t.TotalPrompts > 0 {
			t.AverageTokensPerSecond = (t.AverageTokensPerSecond*float64(t.TotalPrompts-1) + t.TokensPerSecond) / float64(t.TotalPrompts)
		} else {
			t.AverageTokensPerSecond = t.TokensPerSecond
		}
		
		// Update peak
		if t.TokensPerSecond > t.PeakTokensPerSecond {
			t.PeakTokensPerSecond = t.TokensPerSecond
		}
	}
}

// UpdatePrompt updates telemetry for a new prompt
func (t *NPUTelemetry) UpdatePrompt(promptTokens uint64) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.TotalPrompts++
	t.LastPromptTokens = promptTokens
	t.LastInferenceStart = time.Now()
}

// GetStats returns a copy of current telemetry stats
func (t *NPUTelemetry) GetStats() NPUTelemetry {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return NPUTelemetry{
		TokensPerSecond:        t.TokensPerSecond,
		TotalTokensGenerated:   t.TotalTokensGenerated,
		TotalPrompts:           t.TotalPrompts,
		LastPromptTokens:       t.LastPromptTokens,
		LastCompletionTokens:   t.LastCompletionTokens,
		LastInferenceStart:     t.LastInferenceStart,
		LastInferenceEnd:       t.LastInferenceEnd,
		LastInferenceDuration:  t.LastInferenceDuration,
		AverageTokensPerSecond: t.AverageTokensPerSecond,
		PeakTokensPerSecond:    t.PeakTokensPerSecond,
	}
}

// NewNPURegisters creates initialized register state
func NewNPURegisters() *NPURegisters {
	return &NPURegisters{
		Status: NPU_STATUS_IDLE,
	}
}

// ReadReg32 safely reads a 32-bit register
func (r *NPURegisters) ReadReg32(addr uint64) uint32 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	switch addr {
	case NPU_REG_CMD:
		return r.Command
	case NPU_REG_STATUS:
		return r.Status
	case NPU_REG_PROMPT_LEN:
		return r.PromptLen
	case NPU_REG_N_PREDICT:
		return r.NPredict
	case NPU_REG_TOKEN_OUT:
		return uint32(r.TokenOut)
	case NPU_REG_TOKEN_READY:
		return r.TokenReady
	case NPU_REG_MODEL_ID:
		return r.ModelID
	case NPU_REG_CTX_USED:
		return r.CtxUsed
	case NPU_REG_ERROR_CODE:
		return r.ErrorCode
	case NPU_REG_PERF_TOKENS_SEC:
		return r.PerfTokensSec
	default:
		return 0
	}
}

// WriteReg32 safely writes a 32-bit register
func (r *NPURegisters) WriteReg32(addr uint64, value uint32) {
	r.mu.Lock()
	defer r.mu.Unlock()

	switch addr {
	case NPU_REG_CMD:
		r.Command = value
	case NPU_REG_STATUS:
		r.Status = value
	case NPU_REG_PROMPT_LEN:
		r.PromptLen = value
	case NPU_REG_N_PREDICT:
		r.NPredict = value
	case NPU_REG_TOKEN_OUT:
		r.TokenOut = int32(value)
	case NPU_REG_TOKEN_READY:
		r.TokenReady = value
	case NPU_REG_MODEL_ID:
		r.ModelID = value
	case NPU_REG_CTX_USED:
		r.CtxUsed = value
	case NPU_REG_ERROR_CODE:
		r.ErrorCode = value
	case NPU_REG_PERF_TOKENS_SEC:
		r.PerfTokensSec = value
	}
}

// ReadReg64 safely reads a 64-bit register
func (r *NPURegisters) ReadReg64(addr uint64) uint64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	switch addr {
	case NPU_REG_PROMPT_ADDR:
		return r.PromptAddr
	default:
		return 0
	}
}

// WriteReg64 safely writes a 64-bit register
func (r *NPURegisters) WriteReg64(addr uint64, value uint64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	switch addr {
	case NPU_REG_PROMPT_ADDR:
		r.PromptAddr = value
	}
}
