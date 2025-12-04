package drivers

import (
	"context"
	"testing"
	"time"

	"github.com/EchoCog/echollama/core/ecco9"
	"github.com/EchoCog/echollama/core/llm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNPUDriverCreation tests NPU driver creation
func TestNPUDriverCreation(t *testing.T) {
	manager := llm.NewProviderManager()
	driver := NewNPUDriver(manager)
	
	assert.NotNil(t, driver)
	assert.Equal(t, "npu", driver.GetName())
	assert.Equal(t, "1.0.0-entelechy", driver.GetVersion())
	assert.NotNil(t, driver.dimensions)
	assert.NotNil(t, driver.entelechyGenome)
}

// TestNPUDriverLoad tests driver loading
func TestNPUDriverLoad(t *testing.T) {
	manager := llm.NewProviderManager()
	driver := NewNPUDriver(manager)
	
	err := driver.Load(nil)
	require.NoError(t, err)
	
	devices := driver.ListDevices()
	assert.Len(t, devices, 1)
	
	device, err := driver.GetDevice("npu0")
	require.NoError(t, err)
	assert.NotNil(t, device)
}

// TestNPUDriverCapabilities tests driver capabilities
func TestNPUDriverCapabilities(t *testing.T) {
	manager := llm.NewProviderManager()
	driver := NewNPUDriver(manager)
	
	capabilities := driver.GetCapabilities()
	assert.Contains(t, capabilities, "llm-inference")
	assert.Contains(t, capabilities, "token-streaming")
	assert.Contains(t, capabilities, "memory-mapped-io")
	assert.Contains(t, capabilities, "entelechy-actualization")
	assert.Contains(t, capabilities, "ontogenetic-evolution")
	assert.Contains(t, capabilities, "self-generation")
	assert.Contains(t, capabilities, "self-optimization")
}

// TestNPUDeviceInitialization tests device initialization
func TestNPUDeviceInitialization(t *testing.T) {
	manager := llm.NewProviderManager()
	driver := NewNPUDriver(manager)
	err := driver.Load(nil)
	require.NoError(t, err)
	
	device, err := driver.GetDevice("npu0")
	require.NoError(t, err)
	
	npuDevice := device.(*NPUDevice)
	
	// Initialize device
	ctx := context.Background()
	err = npuDevice.Initialize(ctx)
	require.NoError(t, err)
	
	// Check state
	state, err := npuDevice.GetState()
	require.NoError(t, err)
	assert.Equal(t, ecco9.DeviceStatusReady, state.Status)
	assert.Equal(t, ecco9.PowerStateActive, state.Power)
	assert.Equal(t, ecco9.HealthStatusHealthy, state.Health)
}

// TestNPURegisterAccess tests register read/write
func TestNPURegisterAccess(t *testing.T) {
	registers := NewNPURegisters()
	
	// Test 32-bit register access
	registers.WriteReg32(NPU_REG_MODEL_ID, 0xDEADBEEF)
	value := registers.ReadReg32(NPU_REG_MODEL_ID)
	assert.Equal(t, uint32(0xDEADBEEF), value)
	
	// Test 64-bit register access
	registers.WriteReg64(NPU_REG_PROMPT_ADDR, 0x0000DEAD0000BEEF)
	value64 := registers.ReadReg64(NPU_REG_PROMPT_ADDR)
	assert.Equal(t, uint64(0x0000DEAD0000BEEF), value64)
	
	// Test status register
	registers.WriteReg32(NPU_REG_STATUS, NPU_STATUS_IDLE|NPU_STATUS_MODEL_READY)
	status := registers.ReadReg32(NPU_REG_STATUS)
	assert.Equal(t, uint32(NPU_STATUS_IDLE|NPU_STATUS_MODEL_READY), status)
}

// TestNPUMemoryRegion tests SRAM memory region
func TestNPUMemoryRegion(t *testing.T) {
	manager := llm.NewProviderManager()
	device := NewNPUDevice("test", manager)
	
	// Check SRAM region initialized
	assert.NotNil(t, device.sramRegion)
	assert.Equal(t, uint64(NPU_SRAM_BASE), device.sramRegion.BaseAddr)
	assert.Equal(t, uint64(NPU_SRAM_SIZE), device.sramRegion.Size)
	assert.Len(t, device.sramRegion.Data, int(NPU_SRAM_SIZE))
}

// TestNPUTelemetry tests telemetry tracking
func TestNPUTelemetry(t *testing.T) {
	telemetry := NewNPUTelemetry()
	
	// Update prompt
	telemetry.UpdatePrompt(50)
	stats := telemetry.GetStats()
	assert.Equal(t, uint64(1), stats.TotalPrompts)
	assert.Equal(t, uint64(50), stats.LastPromptTokens)
	
	// Update token generation
	duration := 2 * time.Second
	telemetry.UpdateTokenGeneration(100, duration)
	stats = telemetry.GetStats()
	assert.Equal(t, uint64(100), stats.TotalTokensGenerated)
	assert.Equal(t, float64(50), stats.TokensPerSecond) // 100 tokens / 2 seconds
	assert.Equal(t, float64(50), stats.AverageTokensPerSecond)
	assert.Equal(t, float64(50), stats.PeakTokensPerSecond)
}

// TestNPUEntelechyAssessment tests entelechy self-assessment
func TestNPUEntelechyAssessment(t *testing.T) {
	manager := llm.NewProviderManager()
	driver := NewNPUDriver(manager)
	
	genome := driver.AssessEntelechy()
	require.NotNil(t, genome)
	
	// Check genome structure
	assert.NotEmpty(t, genome.ID)
	assert.GreaterOrEqual(t, genome.Fitness, 0.0)
	assert.LessOrEqual(t, genome.Fitness, 1.0)
	assert.NotEmpty(t, genome.ActualizationLevel)
	
	// Check dimensional genes
	assert.GreaterOrEqual(t, genome.Genes.Ontological, 0.0)
	assert.GreaterOrEqual(t, genome.Genes.Teleological, 0.0)
	assert.GreaterOrEqual(t, genome.Genes.Cognitive, 0.0)
	assert.GreaterOrEqual(t, genome.Genes.Integrative, 0.0)
	assert.GreaterOrEqual(t, genome.Genes.Evolutionary, 0.0)
}

// TestNPUDeviceSelfAssessment tests device-level self-assessment
func TestNPUDeviceSelfAssessment(t *testing.T) {
	manager := llm.NewProviderManager()
	driver := NewNPUDriver(manager)
	err := driver.Load(nil)
	require.NoError(t, err)
	
	device, err := driver.GetDevice("npu0")
	require.NoError(t, err)
	
	npuDevice := device.(*NPUDevice)
	ctx := context.Background()
	err = npuDevice.Initialize(ctx)
	require.NoError(t, err)
	
	// Perform self-assessment
	assessment := npuDevice.AssessSelf()
	require.NotNil(t, assessment)
	
	// Check assessment components
	assert.GreaterOrEqual(t, assessment.OverallActualization, 0.0)
	assert.LessOrEqual(t, assessment.OverallActualization, 1.0)
	assert.NotEmpty(t, assessment.ActualizationStage)
	
	// Check dimensional assessments
	assert.GreaterOrEqual(t, assessment.OntologicalHealth.ArchitecturalCoherence, 0.0)
	assert.GreaterOrEqual(t, assessment.TeleologicalAlignment.RoadmapAlignment, 0.0)
	assert.GreaterOrEqual(t, assessment.CognitiveCompleteness.OverallCognition, 0.0)
	assert.GreaterOrEqual(t, assessment.IntegrativeHealth.OverallIntegration, 0.0)
	assert.GreaterOrEqual(t, assessment.EvolutionaryPotential.EvolutionaryFitness, 0.0)
	
	// Check recommendations generated
	assert.NotNil(t, assessment.ImprovementRecommendations)
	assert.NotNil(t, assessment.CriticalIssues)
	assert.NotNil(t, assessment.Strengths)
	
	// Verify assessment string format
	assessmentStr := assessment.String()
	assert.Contains(t, assessmentStr, "NPU Self-Assessment Report")
	assert.Contains(t, assessmentStr, "Actualization Level")
	assert.Contains(t, assessmentStr, "Ontological")
	assert.Contains(t, assessmentStr, "Teleological")
	assert.Contains(t, assessmentStr, "Cognitive")
	assert.Contains(t, assessmentStr, "Integrative")
	assert.Contains(t, assessmentStr, "Evolutionary")
}

// TestNPUSelfGeneration tests self-generation capability
func TestNPUSelfGeneration(t *testing.T) {
	manager := llm.NewProviderManager()
	parent := NewNPUDriver(manager)
	
	// Generate offspring
	offspring := parent.SelfGenerate()
	require.NotNil(t, offspring)
	
	// Check offspring properties
	assert.Equal(t, parent.generation+1, offspring.generation)
	assert.Contains(t, offspring.lineage, parent.entelechyGenome.ID)
	assert.NotEqual(t, parent.entelechyGenome.ID, offspring.entelechyGenome.ID)
}

// TestNPUSelfOptimization tests self-optimization
func TestNPUSelfOptimization(t *testing.T) {
	manager := llm.NewProviderManager()
	driver := NewNPUDriver(manager)
	
	// Get initial fitness
	initialGenome := driver.AssessEntelechy()
	initialFitness := initialGenome.Fitness
	
	// Optimize
	iterations := 5
	driver.SelfOptimize(iterations)
	
	// Check fitness improved or stayed same
	optimizedGenome := driver.AssessEntelechy()
	assert.GreaterOrEqual(t, optimizedGenome.Fitness, initialFitness)
}

// TestNPUOntogenesis tests ontogenesis operations
func TestNPUOntogenesis(t *testing.T) {
	manager := llm.NewProviderManager()
	parent1 := NewNPUDriver(manager)
	parent2 := NewNPUDriver(manager)
	
	onto := DefaultNPUOntogenesis()
	
	// Test self-generation
	offspring1 := onto.SelfGenerate(parent1)
	require.NotNil(t, offspring1)
	assert.Equal(t, parent1.generation+1, offspring1.generation)
	
	// Test reproduction
	offspring2 := onto.SelfReproduce(parent1, parent2)
	require.NotNil(t, offspring2)
	assert.Contains(t, offspring2.lineage, parent1.entelechyGenome.ID)
	assert.Contains(t, offspring2.lineage, parent2.entelechyGenome.ID)
}

// TestNPUEvolution tests population evolution
func TestNPUEvolution(t *testing.T) {
	manager := llm.NewProviderManager()
	
	// Create seed population
	seeds := []*NPUDriver{
		NewNPUDriver(manager),
		NewNPUDriver(manager),
		NewNPUDriver(manager),
	}
	
	onto := DefaultNPUOntogenesis()
	onto.PopulationSize = 5
	onto.MaxGenerations = 3
	
	// Evolve population
	population, history := onto.EvolvePopulation(seeds)
	
	// Check results
	require.NotNil(t, population)
	require.NotNil(t, history)
	assert.LessOrEqual(t, len(history), 3) // Max 3 generations
	assert.Len(t, population, 5) // Population size maintained
	
	// Check history stats
	for _, stats := range history {
		assert.GreaterOrEqual(t, stats.BestFitness, 0.0)
		assert.LessOrEqual(t, stats.BestFitness, 1.0)
		assert.GreaterOrEqual(t, stats.AvgFitness, 0.0)
		assert.LessOrEqual(t, stats.AvgFitness, 1.0)
	}
}

// TestNPUModelConfiguration tests model configuration
func TestNPUModelConfiguration(t *testing.T) {
	config := DefaultNPUModelConfig()
	
	assert.Equal(t, int32(4096), config.NCtx)
	assert.Equal(t, int32(4), config.NThreads)
	assert.Equal(t, int32(0), config.NGPULayers)
	assert.Equal(t, float32(0.7), config.Temperature)
	assert.Equal(t, float32(0.9), config.TopP)
}

// TestNPUSequenceConfiguration tests sequence configuration
func TestNPUSequenceConfiguration(t *testing.T) {
	config := DefaultNPUSequenceConfig()
	
	assert.Equal(t, int32(128), config.NPredict)
	assert.Equal(t, int32(4096), config.MaxCtx)
	assert.False(t, config.EchoPrompt)
	assert.True(t, config.StreamTokens)
}

// TestNPUDeviceReset tests device reset
func TestNPUDeviceReset(t *testing.T) {
	manager := llm.NewProviderManager()
	device := NewNPUDevice("test", manager)
	
	ctx := context.Background()
	err := device.Initialize(ctx)
	require.NoError(t, err)
	
	// Modify some state
	device.registers.WriteReg32(NPU_REG_MODEL_ID, 123)
	device.telemetry.UpdatePrompt(50)
	
	// Reset
	err = device.Reset(ctx)
	require.NoError(t, err)
	
	// Check state cleared
	assert.Equal(t, uint32(0), device.registers.ReadReg32(NPU_REG_MODEL_ID))
	stats := device.telemetry.GetStats()
	assert.Equal(t, uint64(0), stats.TotalPrompts)
}

// TestNPUDeviceShutdown tests device shutdown
func TestNPUDeviceShutdown(t *testing.T) {
	manager := llm.NewProviderManager()
	device := NewNPUDevice("test", manager)
	
	ctx := context.Background()
	err := device.Initialize(ctx)
	require.NoError(t, err)
	
	// Shutdown
	err = device.Shutdown(ctx)
	require.NoError(t, err)
	
	// Check state
	state, err := device.GetState()
	require.NoError(t, err)
	assert.Equal(t, ecco9.DeviceStatusOffline, state.Status)
	assert.Equal(t, ecco9.PowerStateOff, state.Power)
}

// TestNPUDeviceMetrics tests metrics collection
func TestNPUDeviceMetrics(t *testing.T) {
	manager := llm.NewProviderManager()
	device := NewNPUDevice("test", manager)
	
	ctx := context.Background()
	err := device.Initialize(ctx)
	require.NoError(t, err)
	
	// Update telemetry
	device.telemetry.UpdatePrompt(50)
	device.telemetry.UpdateTokenGeneration(100, 2*time.Second)
	
	// Get metrics
	metrics, err := device.GetMetrics()
	require.NoError(t, err)
	assert.Equal(t, uint64(1), metrics.OperationCount)
}

// TestNPUDeviceHealth tests health status
func TestNPUDeviceHealth(t *testing.T) {
	manager := llm.NewProviderManager()
	device := NewNPUDevice("test", manager)
	
	ctx := context.Background()
	err := device.Initialize(ctx)
	require.NoError(t, err)
	
	health, err := device.GetHealth()
	require.NoError(t, err)
	assert.Equal(t, ecco9.HealthStatusHealthy, health)
}

// TestNPUDeviceInfo tests device information
func TestNPUDeviceInfo(t *testing.T) {
	manager := llm.NewProviderManager()
	device := NewNPUDevice("test123", manager)
	
	assert.Equal(t, "test123", device.GetID())
	assert.Contains(t, device.GetName(), "NPU")
	assert.Equal(t, ecco9.DeviceType("npu"), device.GetType())
}

// TestNPUDriverUnload tests driver unloading
func TestNPUDriverUnload(t *testing.T) {
	manager := llm.NewProviderManager()
	driver := NewNPUDriver(manager)
	
	err := driver.Load(nil)
	require.NoError(t, err)
	
	devices := driver.ListDevices()
	assert.Len(t, devices, 1)
	
	err = driver.Unload()
	require.NoError(t, err)
	
	devices = driver.ListDevices()
	assert.Len(t, devices, 0)
}

// BenchmarkNPURegisterAccess benchmarks register access
func BenchmarkNPURegisterAccess(b *testing.B) {
	registers := NewNPURegisters()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		registers.WriteReg32(NPU_REG_STATUS, uint32(i))
		_ = registers.ReadReg32(NPU_REG_STATUS)
	}
}

// BenchmarkNPUSelfAssessment benchmarks self-assessment
func BenchmarkNPUSelfAssessment(b *testing.B) {
	manager := llm.NewProviderManager()
	driver := NewNPUDriver(manager)
	driver.Load(nil)
	device, _ := driver.GetDevice("npu0")
	npuDevice := device.(*NPUDevice)
	npuDevice.Initialize(context.Background())
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = npuDevice.AssessSelf()
	}
}

// BenchmarkNPUSelfGeneration benchmarks self-generation
func BenchmarkNPUSelfGeneration(b *testing.B) {
	manager := llm.NewProviderManager()
	parent := NewNPUDriver(manager)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parent.SelfGenerate()
	}
}
