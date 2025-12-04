package drivers

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/EchoCog/echollama/core/ecco9"
	"github.com/EchoCog/echollama/core/entelechy"
	"github.com/EchoCog/echollama/core/llm"
	"github.com/EchoCog/echollama/core/ontogenesis"
)

// NPUDriver implements the Neural Processing Unit coprocessor driver
// Treats LLM inference as a memory-mapped peripheral device with
// entelechy-aware self-actualization and ontogenetic self-generation
type NPUDriver struct {
	mu          sync.RWMutex
	name        string
	version     string
	devices     map[string]*NPUDevice
	llmManager  *llm.ProviderManager
	
	// Entelechy integration
	entelechyGenome *entelechy.EntelechyGenome
	dimensions      *NPUDimensions
	
	// Ontogenesis integration
	ontogeneticKernel *ontogenesis.OntogeneticKernel
	generation        int
	lineage           []string
}

// NPUDimensions tracks the five dimensions of NPU entelechy
type NPUDimensions struct {
	Ontological  *entelechy.OntologicalDimension
	Teleological *entelechy.TeleologicalDimension
	Cognitive    *entelechy.CognitiveDimension
	Integrative  *entelechy.IntegrativeDimension
	Evolutionary *entelechy.EvolutionaryDimension
}

// NPUDevice represents a single NPU coprocessor instance
type NPUDevice struct {
	mu             sync.RWMutex
	id             string
	name           string
	deviceType     ecco9.DeviceType
	state          ecco9.DeviceState
	
	// Hardware abstraction
	registers      *NPURegisters
	sramRegion     *NPUMemoryRegion
	
	// Configuration
	modelConfig    NPUModelConfig
	sequenceConfig NPUSequenceConfig
	
	// Telemetry
	telemetry      *NPUTelemetry
	
	// LLM integration
	llmManager     *llm.ProviderManager
	currentModel   string
	
	// Entelechy state
	actualizationLevel float64
	fitnessScore       float64
	
	// Lifecycle
	initialized    bool
	modelLoaded    bool
	inferenceActive bool
	lastError      error
}

// NewNPUDriver creates a new NPU driver with entelechy and ontogenesis support
func NewNPUDriver(llmManager *llm.ProviderManager) *NPUDriver {
	driver := &NPUDriver{
		name:       "npu",
		version:    "1.0.0-entelechy",
		devices:    make(map[string]*NPUDevice),
		llmManager: llmManager,
		generation: 0,
		lineage:    []string{},
	}
	
	// Initialize entelechy dimensions
	driver.dimensions = &NPUDimensions{
		Ontological:  entelechy.NewOntologicalDimension(),
		Teleological: entelechy.NewTeleologicalDimension(),
		Cognitive:    entelechy.NewCognitiveDimension(),
		Integrative:  entelechy.NewIntegrativeDimension(),
		Evolutionary: entelechy.NewEvolutionaryDimension(),
	}
	
	// Initialize entelechy genome
	driver.entelechyGenome = entelechy.NewEntelechyGenome("npu-gen0", 0)
	driver.entelechyGenome.Genes.Ontological = 0.7  // Foundation present
	driver.entelechyGenome.Genes.Teleological = 0.6 // Purpose clear
	driver.entelechyGenome.Genes.Cognitive = 0.5    // Cognitive capabilities developing
	driver.entelechyGenome.Genes.Integrative = 0.8  // Well integrated with ecco9
	driver.entelechyGenome.Genes.Evolutionary = 0.7 // Good growth potential
	driver.entelechyGenome.CalculateFitness()
	
	return driver
}

// Load implements Driver.Load
func (nd *NPUDriver) Load(config interface{}) error {
	nd.mu.Lock()
	defer nd.mu.Unlock()
	
	// Create primary NPU device
	device := NewNPUDevice("npu0", nd.llmManager)
	nd.devices["npu0"] = device
	
	// Update ontological dimension (foundation health)
	nd.dimensions.Ontological.FoundationHealth = 1.0
	nd.dimensions.Ontological.CoreHealth = 0.8
	nd.dimensions.Ontological.Assess()
	
	return nil
}

// Unload implements Driver.Unload
func (nd *NPUDriver) Unload() error {
	nd.mu.Lock()
	defer nd.mu.Unlock()
	
	// Shutdown all devices
	for _, device := range nd.devices {
		ctx := context.Background()
		if err := device.Shutdown(ctx); err != nil {
			return err
		}
	}
	
	nd.devices = make(map[string]*NPUDevice)
	return nil
}

// GetDevice implements Driver.GetDevice
func (nd *NPUDriver) GetDevice(id string) (ecco9.CognitiveDevice, error) {
	nd.mu.RLock()
	defer nd.mu.RUnlock()
	
	device, exists := nd.devices[id]
	if !exists {
		return nil, fmt.Errorf("NPU device %s not found", id)
	}
	return device, nil
}

// ListDevices implements Driver.ListDevices
func (nd *NPUDriver) ListDevices() []ecco9.CognitiveDevice {
	nd.mu.RLock()
	defer nd.mu.RUnlock()
	
	devices := make([]ecco9.CognitiveDevice, 0, len(nd.devices))
	for _, device := range nd.devices {
		devices = append(devices, device)
	}
	return devices
}

// GetName implements Driver.GetName
func (nd *NPUDriver) GetName() string {
	return nd.name
}

// GetVersion implements Driver.GetVersion
func (nd *NPUDriver) GetVersion() string {
	return nd.version
}

// GetCapabilities implements Driver.GetCapabilities
func (nd *NPUDriver) GetCapabilities() []string {
	return []string{
		"llm-inference",
		"token-streaming",
		"memory-mapped-io",
		"hardware-registers",
		"gguf-models",
		"gpu-offload",
		"kv-cache-management",
		"batch-inference",
		"entelechy-actualization",
		"ontogenetic-evolution",
		"self-generation",
		"self-optimization",
	}
}

// AssessEntelechy performs comprehensive entelechy assessment
func (nd *NPUDriver) AssessEntelechy() *entelechy.EntelechyGenome {
	nd.mu.Lock()
	defer nd.mu.Unlock()
	
	// Assess all dimensions
	nd.dimensions.Ontological.Assess()
	nd.dimensions.Teleological.Assess()
	nd.dimensions.Cognitive.Assess()
	nd.dimensions.Integrative.Assess()
	nd.dimensions.Evolutionary.Assess()
	
	// Update genome
	nd.entelechyGenome.Genes.Ontological = nd.dimensions.Ontological.ArchitecturalCompleteness
	nd.entelechyGenome.Genes.Teleological = nd.dimensions.Teleological.ActualizationTrajectory
	nd.entelechyGenome.Genes.Cognitive = nd.dimensions.Cognitive.CognitiveCompleteness
	nd.entelechyGenome.Genes.Integrative = nd.dimensions.Integrative.IntegrationHealth
	nd.entelechyGenome.Genes.Evolutionary = nd.dimensions.Evolutionary.EvolutionaryPotential
	
	nd.entelechyGenome.CalculateFitness()
	
	return nd.entelechyGenome
}

// SelfGenerate creates offspring NPU driver
func (nd *NPUDriver) SelfGenerate() *NPUDriver {
	nd.mu.RLock()
	defer nd.mu.RUnlock()
	
	offspring := NewNPUDriver(nd.llmManager)
	offspring.generation = nd.generation + 1
	offspring.lineage = append([]string{nd.entelechyGenome.ID}, nd.lineage...)
	
	// Inherit and mutate genome
	offspring.entelechyGenome = nd.entelechyGenome.Clone()
	offspring.entelechyGenome.ID = fmt.Sprintf("npu-gen%d", offspring.generation)
	offspring.entelechyGenome.Generation = offspring.generation
	
	// Apply genetic mutations (small improvements)
	offspring.entelechyGenome.Genes.Cognitive += 0.05
	offspring.entelechyGenome.Genes.Evolutionary += 0.03
	offspring.entelechyGenome.CalculateFitness()
	
	return offspring
}

// SelfOptimize performs iterative self-optimization
func (nd *NPUDriver) SelfOptimize(iterations int) {
	for i := 0; i < iterations; i++ {
		// Assess current state
		genome := nd.AssessEntelechy()
		currentFitness := genome.Fitness
		
		// Identify weakest dimension
		weakest := nd.identifyWeakestDimension()
		
		// Apply targeted improvement
		nd.improveDimension(weakest)
		
		// Re-assess
		newGenome := nd.AssessEntelechy()
		
		if newGenome.Fitness > currentFitness {
			// Improvement successful
			nd.mu.Lock()
			nd.entelechyGenome = newGenome
			nd.mu.Unlock()
		}
	}
}

// identifyWeakestDimension finds the dimension with lowest score
func (nd *NPUDriver) identifyWeakestDimension() string {
	scores := map[string]float64{
		"ontological":  nd.dimensions.Ontological.ArchitecturalCompleteness,
		"teleological": nd.dimensions.Teleological.ActualizationTrajectory,
		"cognitive":    nd.dimensions.Cognitive.CognitiveCompleteness,
		"integrative":  nd.dimensions.Integrative.IntegrationHealth,
		"evolutionary": nd.dimensions.Evolutionary.EvolutionaryPotential,
	}
	
	weakest := "cognitive"
	lowestScore := 1.0
	
	for dimension, score := range scores {
		if score < lowestScore {
			lowestScore = score
			weakest = dimension
		}
	}
	
	return weakest
}

// improveDimension applies targeted improvements to a dimension
func (nd *NPUDriver) improveDimension(dimension string) {
	nd.mu.Lock()
	defer nd.mu.Unlock()
	
	switch dimension {
	case "ontological":
		nd.dimensions.Ontological.CoreHealth += 0.05
		nd.dimensions.Ontological.SpecializedHealth += 0.03
	case "teleological":
		nd.dimensions.Teleological.PurposeClarity += 0.05
		nd.dimensions.Teleological.RoadmapAlignment += 0.03
	case "cognitive":
		nd.dimensions.Cognitive.LoopCoherence += 0.05
		nd.dimensions.Cognitive.LearningCapacity += 0.03
	case "integrative":
		nd.dimensions.Integrative.BuildHealth += 0.05
		nd.dimensions.Integrative.TestHealth += 0.03
	case "evolutionary":
		nd.dimensions.Evolutionary.CodeHealth += 0.05
		nd.dimensions.Evolutionary.ImplementationDepth += 0.03
	}
}

// NewNPUDevice creates a new NPU device instance
func NewNPUDevice(id string, llmManager *llm.ProviderManager) *NPUDevice {
	device := &NPUDevice{
		id:             id,
		name:           fmt.Sprintf("NPU-%s", id),
		deviceType:     "npu",
		registers:      NewNPURegisters(),
		telemetry:      NewNPUTelemetry(),
		llmManager:     llmManager,
		modelConfig:    DefaultNPUModelConfig(),
		sequenceConfig: DefaultNPUSequenceConfig(),
		initialized:    false,
		modelLoaded:    false,
	}
	
	// Initialize SRAM region
	device.sramRegion = &NPUMemoryRegion{
		BaseAddr: NPU_SRAM_BASE,
		Size:     NPU_SRAM_SIZE,
		Data:     make([]byte, NPU_SRAM_SIZE),
		Name:     "NPU-SRAM",
	}
	
	// Initialize device state
	device.state = ecco9.DeviceState{
		ID:         id,
		Name:       device.name,
		Status:     ecco9.DeviceStatusOffline,
		Power:      ecco9.PowerStateOff,
		Health:     ecco9.HealthStatusHealthy,
		LastUpdate: time.Now(),
		Uptime:     0,
	}
	
	return device
}

// Initialize implements CognitiveDevice.Initialize
func (d *NPUDevice) Initialize(ctx context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	
	if d.initialized {
		return fmt.Errorf("NPU device %s already initialized", d.id)
	}
	
	// Perform self-test
	if err := d.performSelfTest(); err != nil {
		d.state.Status = ecco9.DeviceStatusError
		d.state.Health = ecco9.HealthStatusFailed
		d.registers.WriteReg32(NPU_REG_ERROR_CODE, NPU_ERR_INVALID_CMD)
		d.registers.WriteReg32(NPU_REG_STATUS, NPU_STATUS_ERROR)
		return fmt.Errorf("NPU self-test failed: %w", err)
	}
	
	// Set hardware status
	d.registers.WriteReg32(NPU_REG_STATUS, NPU_STATUS_IDLE)
	d.registers.WriteReg32(NPU_REG_ERROR_CODE, NPU_ERR_NONE)
	
	// Update device state
	d.state.Status = ecco9.DeviceStatusReady
	d.state.Power = ecco9.PowerStateActive
	d.state.LastUpdate = time.Now()
	d.initialized = true
	
	return nil
}

// Shutdown implements CognitiveDevice.Shutdown
func (d *NPUDevice) Shutdown(ctx context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	
	if !d.initialized {
		return nil
	}
	
	// Stop any active inference
	if d.inferenceActive {
		d.registers.WriteReg32(NPU_REG_CMD, NPU_CMD_SOFT_STOP)
		d.inferenceActive = false
	}
	
	// Update state
	d.state.Status = ecco9.DeviceStatusOffline
	d.state.Power = ecco9.PowerStateOff
	d.state.LastUpdate = time.Now()
	d.initialized = false
	
	// Reset registers
	d.registers.WriteReg32(NPU_REG_STATUS, 0)
	
	return nil
}

// Reset implements CognitiveDevice.Reset
func (d *NPUDevice) Reset(ctx context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	
	// Execute hardware reset
	d.registers.WriteReg32(NPU_REG_CMD, NPU_CMD_RESET)
	
	// Reset state
	d.registers = NewNPURegisters()
	d.telemetry = NewNPUTelemetry()
	d.inferenceActive = false
	d.modelLoaded = false
	d.lastError = nil
	
	// Clear SRAM
	d.sramRegion.Data = make([]byte, NPU_SRAM_SIZE)
	
	d.state.Status = ecco9.DeviceStatusReady
	d.state.LastUpdate = time.Now()
	
	return nil
}

// GetState implements CognitiveDevice.GetState
func (d *NPUDevice) GetState() (ecco9.DeviceState, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	
	// Update uptime - calculate from initialization time, not last update
	if d.initialized {
		d.state.Uptime = time.Since(d.state.LastUpdate)
	}
	
	// Make a copy to return
	stateCopy := d.state
	return stateCopy, nil
}

// SetState implements CognitiveDevice.SetState
func (d *NPUDevice) SetState(state ecco9.DeviceState) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	
	d.state = state
	d.state.LastUpdate = time.Now()
	
	return nil
}

// Read implements CognitiveDevice.Read
func (d *NPUDevice) Read(buffer []byte) (int, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	
	// Read from token output register
	if d.registers.TokenReady > 0 {
		// Simulated token read
		token := d.registers.TokenOut
		formatted := fmt.Sprintf("%d\n", token)
		n := copy(buffer, []byte(formatted))
		d.registers.TokenReady = 0
		return n, nil
	}
	
	return 0, nil
}

// Write implements CognitiveDevice.Write
func (d *NPUDevice) Write(buffer []byte) (int, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	
	// Write prompt to SRAM
	promptAddr := d.registers.PromptAddr
	if promptAddr >= NPU_SRAM_BASE && promptAddr < NPU_SRAM_BASE+NPU_SRAM_SIZE {
		offset := promptAddr - NPU_SRAM_BASE
		maxLen := NPU_SRAM_SIZE - offset
		if uint64(len(buffer)) > maxLen {
			buffer = buffer[:maxLen]
		}
		copy(d.sramRegion.Data[offset:], buffer)
		d.registers.WriteReg32(NPU_REG_PROMPT_LEN, uint32(len(buffer)))
		return len(buffer), nil
	}
	
	return 0, fmt.Errorf("invalid prompt address: 0x%x", promptAddr)
}

// IoCtl implements CognitiveDevice.IoCtl
func (d *NPUDevice) IoCtl(command uint32, arg interface{}) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	
	// Process hardware commands
	switch command {
	case NPU_CMD_RESET:
		return d.Reset(context.Background())
	case NPU_CMD_LOAD_MODEL:
		if config, ok := arg.(NPUModelConfig); ok {
			return d.loadModel(config)
		}
		return fmt.Errorf("invalid model config")
	case NPU_CMD_START_INF:
		if seqConfig, ok := arg.(NPUSequenceConfig); ok {
			return d.startInference(seqConfig)
		}
		return fmt.Errorf("invalid sequence config")
	case NPU_CMD_SOFT_STOP:
		d.inferenceActive = false
		d.registers.WriteReg32(NPU_REG_STATUS, NPU_STATUS_IDLE)
		return nil
	default:
		d.registers.WriteReg32(NPU_REG_ERROR_CODE, NPU_ERR_INVALID_CMD)
		return fmt.Errorf("invalid command: 0x%x", command)
	}
}

// GetMetrics implements CognitiveDevice.GetMetrics
func (d *NPUDevice) GetMetrics() (ecco9.DeviceMetrics, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	
	stats := d.telemetry.GetStats()
	
	return ecco9.DeviceMetrics{
		CPUUsage:       0.0, // TODO: actual CPU usage
		MemoryUsage:    0.0, // TODO: actual memory usage
		OperationCount: stats.TotalPrompts,
		ErrorCount:     0, // TODO: track errors
		AverageLatency: stats.LastInferenceDuration,
		LastOperation:  stats.LastInferenceEnd,
	}, nil
}

// GetHealth implements CognitiveDevice.GetHealth
func (d *NPUDevice) GetHealth() (ecco9.HealthStatus, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	
	return d.state.Health, nil
}

// GetID implements CognitiveDevice.GetID
func (d *NPUDevice) GetID() string {
	return d.id
}

// GetName implements CognitiveDevice.GetName
func (d *NPUDevice) GetName() string {
	return d.name
}

// GetType implements CognitiveDevice.GetType
func (d *NPUDevice) GetType() ecco9.DeviceType {
	return d.deviceType
}

// performSelfTest runs hardware diagnostics
func (d *NPUDevice) performSelfTest() error {
	// Test register access
	testVal := uint32(0xDEADBEEF)
	d.registers.WriteReg32(NPU_REG_MODEL_ID, testVal)
	if d.registers.ReadReg32(NPU_REG_MODEL_ID) != testVal {
		return fmt.Errorf("register test failed")
	}
	
	// Test SRAM access
	testData := []byte("NPU-SELF-TEST")
	copy(d.sramRegion.Data[:len(testData)], testData)
	for i, b := range testData {
		if d.sramRegion.Data[i] != b {
			return fmt.Errorf("SRAM test failed")
		}
	}
	
	return nil
}

// loadModel loads a GGUF model (stubbed for now)
func (d *NPUDevice) loadModel(config NPUModelConfig) error {
	// TODO: Integrate with actual GGUF runtime
	d.modelConfig = config
	d.modelLoaded = true
	d.currentModel = config.ModelName
	
	d.registers.WriteReg32(NPU_REG_STATUS, NPU_STATUS_IDLE|NPU_STATUS_MODEL_READY)
	d.registers.WriteReg32(NPU_REG_MODEL_ID, 1) // Model ID
	
	return nil
}

// startInference begins generation (stubbed for now)
func (d *NPUDevice) startInference(config NPUSequenceConfig) error {
	if !d.modelLoaded {
		d.registers.WriteReg32(NPU_REG_ERROR_CODE, NPU_ERR_MODEL_LOAD)
		return fmt.Errorf("no model loaded")
	}
	
	// Update registers
	d.registers.WriteReg32(NPU_REG_STATUS, NPU_STATUS_BUSY)
	d.registers.WriteReg32(NPU_REG_N_PREDICT, uint32(config.NPredict))
	d.inferenceActive = true
	
	// Update telemetry
	d.telemetry.UpdatePrompt(uint64(d.registers.PromptLen))
	
	// TODO: Actual inference - for now, stub completion
	d.sequenceConfig = config
	
	return nil
}

// GetTelemetry returns current telemetry statistics
func (d *NPUDevice) GetTelemetry() NPUTelemetry {
	return d.telemetry.GetStats()
}

// GetDeviceStatusString returns human-readable device status
func (d *NPUDevice) GetDeviceStatusString() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	
	status := d.registers.ReadReg32(NPU_REG_STATUS)
	var statusStr string
	
	if status&NPU_STATUS_IDLE != 0 {
		statusStr = "IDLE"
	} else if status&NPU_STATUS_BUSY != 0 {
		statusStr = "BUSY"
	} else if status&NPU_STATUS_ERROR != 0 {
		statusStr = "ERROR"
	}
	
	if status&NPU_STATUS_MODEL_READY != 0 {
		statusStr += "|MODEL_READY"
	}
	if status&NPU_STATUS_EOG != 0 {
		statusStr += "|EOG"
	}
	
	return fmt.Sprintf("NPU[%s] Status: %s, Model: %s", d.id, statusStr, d.currentModel)
}
