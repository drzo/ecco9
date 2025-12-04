package drivers

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/EchoCog/echollama/core/entelechy"
	"github.com/EchoCog/echollama/core/ontogenesis"
)

// NPUOntogenesis provides self-generation and evolution capabilities for NPU
type NPUOntogenesis struct {
	// Evolution parameters
	MutationRate    float64
	CrossoverRate   float64
	ElitismRate     float64
	PopulationSize  int
	MaxGenerations  int
	FitnessThreshold float64
}

// DefaultNPUOntogenesis returns default ontogenesis configuration
func DefaultNPUOntogenesis() *NPUOntogenesis {
	return &NPUOntogenesis{
		MutationRate:     0.15,
		CrossoverRate:    0.8,
		ElitismRate:      0.1,
		PopulationSize:   20,
		MaxGenerations:   100,
		FitnessThreshold: 0.9,
	}
}

// SelfGenerate creates offspring NPU from parent
func (no *NPUOntogenesis) SelfGenerate(parent *NPUDriver) *NPUDriver {
	offspring := parent.SelfGenerate()
	
	// Apply ontogenetic mutations
	offspring = no.applyMutations(offspring)
	
	// Assess fitness
	offspring.AssessEntelechy()
	
	return offspring
}

// SelfOptimize performs iterative self-optimization
func (no *NPUOntogenesis) SelfOptimize(npu *NPUDriver, iterations int) *NPUDriver {
	optimized := npu
	learningRate := 0.01
	
	for iter := 0; iter < iterations; iter++ {
		// Assess current fitness
		currentGenome := optimized.AssessEntelechy()
		currentFitness := currentGenome.Fitness
		
		// Try optimization step
		candidate := optimized.SelfGenerate()
		candidateGenome := candidate.AssessEntelechy()
		
		// Keep if improvement
		if candidateGenome.Fitness > currentFitness {
			optimized = candidate
		}
		
		// Decay learning rate
		learningRate *= 0.99
	}
	
	return optimized
}

// SelfReproduce creates offspring from two parent NPUs
func (no *NPUOntogenesis) SelfReproduce(parent1, parent2 *NPUDriver) *NPUDriver {
	// Create base offspring from parent1
	offspring := NewNPUDriver(parent1.llmManager)
	
	// Genetic crossover of entelechy genomes
	offspring.entelechyGenome = no.crossoverGenomes(
		parent1.entelechyGenome,
		parent2.entelechyGenome,
	)
	
	offspring.generation = maxInt(parent1.generation, parent2.generation) + 1
	offspring.lineage = []string{
		parent1.entelechyGenome.ID,
		parent2.entelechyGenome.ID,
	}
	
	// Apply mutations
	offspring = no.applyMutations(offspring)
	
	// Assess fitness
	offspring.AssessEntelechy()
	
	return offspring
}

// EvolvePopulation evolves a population of NPUs over generations
func (no *NPUOntogenesis) EvolvePopulation(seeds []*NPUDriver) ([]*NPUDriver, []GenerationStats) {
	population := seeds
	history := []GenerationStats{}
	
	for gen := 0; gen < no.MaxGenerations; gen++ {
		// Evaluate fitness for all individuals
		fitnesses := make([]float64, len(population))
		for i, npu := range population {
			genome := npu.AssessEntelechy()
			fitnesses[i] = genome.Fitness
		}
		
		// Record statistics
		stats := GenerationStats{
			Generation:   gen,
			BestFitness:  maxFloat64(fitnesses),
			AvgFitness:   avgFloat64(fitnesses),
			Diversity:    no.calculateDiversity(population),
			PopulationSize: len(population),
		}
		history = append(history, stats)
		
		// Check termination
		if stats.BestFitness >= no.FitnessThreshold {
			break
		}
		
		// Evolve next generation
		population = no.evolveGeneration(population, fitnesses)
	}
	
	return population, history
}

// applyMutations applies genetic mutations to NPU
func (no *NPUOntogenesis) applyMutations(npu *NPUDriver) *NPUDriver {
	if rand.Float64() < no.MutationRate {
		// Mutate ontological genes
		npu.dimensions.Ontological.CoreHealth += (rand.Float64() - 0.5) * 0.1
		npu.dimensions.Ontological.CoreHealth = clamp(npu.dimensions.Ontological.CoreHealth, 0, 1)
	}
	
	if rand.Float64() < no.MutationRate {
		// Mutate cognitive genes
		npu.dimensions.Cognitive.LearningCapacity += (rand.Float64() - 0.5) * 0.1
		npu.dimensions.Cognitive.LearningCapacity = clamp(npu.dimensions.Cognitive.LearningCapacity, 0, 1)
	}
	
	if rand.Float64() < no.MutationRate {
		// Mutate evolutionary genes
		npu.dimensions.Evolutionary.SelfImprovementCapacity += (rand.Float64() - 0.5) * 0.1
		npu.dimensions.Evolutionary.SelfImprovementCapacity = clamp(npu.dimensions.Evolutionary.SelfImprovementCapacity, 0, 1)
	}
	
	return npu
}

// crossoverGenomes performs genetic crossover between two entelechy genomes
func (no *NPUOntogenesis) crossoverGenomes(genome1, genome2 *entelechy.EntelechyGenome) *entelechy.EntelechyGenome {
	offspring := entelechy.NewEntelechyGenome(
		fmt.Sprintf("npu-gen%d-%d", genome1.Generation+1, rand.Int()),
		maxInt(genome1.Generation, genome2.Generation)+1,
	)
	
	// Crossover genes
	if rand.Float64() < 0.5 {
		offspring.Genes.Ontological = genome1.Genes.Ontological
	} else {
		offspring.Genes.Ontological = genome2.Genes.Ontological
	}
	
	if rand.Float64() < 0.5 {
		offspring.Genes.Teleological = genome1.Genes.Teleological
	} else {
		offspring.Genes.Teleological = genome2.Genes.Teleological
	}
	
	if rand.Float64() < 0.5 {
		offspring.Genes.Cognitive = genome1.Genes.Cognitive
	} else {
		offspring.Genes.Cognitive = genome2.Genes.Cognitive
	}
	
	if rand.Float64() < 0.5 {
		offspring.Genes.Integrative = genome1.Genes.Integrative
	} else {
		offspring.Genes.Integrative = genome2.Genes.Integrative
	}
	
	if rand.Float64() < 0.5 {
		offspring.Genes.Evolutionary = genome1.Genes.Evolutionary
	} else {
		offspring.Genes.Evolutionary = genome2.Genes.Evolutionary
	}
	
	offspring.CalculateFitness()
	
	return offspring
}

// evolveGeneration creates next generation through selection, crossover, mutation
func (no *NPUOntogenesis) evolveGeneration(population []*NPUDriver, fitnesses []float64) []*NPUDriver {
	newPopulation := []*NPUDriver{}
	
	// Elitism - keep best individuals
	eliteCount := int(float64(len(population)) * no.ElitismRate)
	eliteIndices := no.selectElite(fitnesses, eliteCount)
	for _, idx := range eliteIndices {
		newPopulation = append(newPopulation, population[idx])
	}
	
	// Generate offspring through crossover and mutation
	for len(newPopulation) < no.PopulationSize {
		// Tournament selection
		parent1 := population[no.tournamentSelection(fitnesses)]
		parent2 := population[no.tournamentSelection(fitnesses)]
		
		// Crossover
		var offspring *NPUDriver
		if rand.Float64() < no.CrossoverRate {
			offspring = no.SelfReproduce(parent1, parent2)
		} else {
			offspring = parent1.SelfGenerate()
		}
		
		newPopulation = append(newPopulation, offspring)
	}
	
	return newPopulation
}

// selectElite returns indices of top N individuals
func (no *NPUOntogenesis) selectElite(fitnesses []float64, count int) []int {
	type indexedFitness struct {
		index   int
		fitness float64
	}
	
	indexed := make([]indexedFitness, len(fitnesses))
	for i, f := range fitnesses {
		indexed[i] = indexedFitness{i, f}
	}
	
	// Sort by fitness (descending)
	for i := 0; i < len(indexed)-1; i++ {
		for j := i + 1; j < len(indexed); j++ {
			if indexed[j].fitness > indexed[i].fitness {
				indexed[i], indexed[j] = indexed[j], indexed[i]
			}
		}
	}
	
	// Return top N indices
	result := make([]int, minInt(count, len(indexed)))
	for i := 0; i < len(result); i++ {
		result[i] = indexed[i].index
	}
	
	return result
}

// tournamentSelection selects individual via tournament selection
func (no *NPUOntogenesis) tournamentSelection(fitnesses []float64) int {
	tournamentSize := 3
	bestIdx := rand.Intn(len(fitnesses))
	bestFitness := fitnesses[bestIdx]
	
	for i := 1; i < tournamentSize; i++ {
		idx := rand.Intn(len(fitnesses))
		if fitnesses[idx] > bestFitness {
			bestIdx = idx
			bestFitness = fitnesses[idx]
		}
	}
	
	return bestIdx
}

// calculateDiversity measures genetic diversity in population
func (no *NPUOntogenesis) calculateDiversity(population []*NPUDriver) float64 {
	if len(population) < 2 {
		return 0.0
	}
	
	totalDistance := 0.0
	comparisons := 0
	
	for i := 0; i < len(population)-1; i++ {
		for j := i + 1; j < len(population); j++ {
			genome1 := population[i].AssessEntelechy()
			genome2 := population[j].AssessEntelechy()
			
			// Calculate Euclidean distance in gene space
			dist := math.Sqrt(
				math.Pow(genome1.Genes.Ontological-genome2.Genes.Ontological, 2) +
					math.Pow(genome1.Genes.Teleological-genome2.Genes.Teleological, 2) +
					math.Pow(genome1.Genes.Cognitive-genome2.Genes.Cognitive, 2) +
					math.Pow(genome1.Genes.Integrative-genome2.Genes.Integrative, 2) +
					math.Pow(genome1.Genes.Evolutionary-genome2.Genes.Evolutionary, 2),
			)
			
			totalDistance += dist
			comparisons++
		}
	}
	
	return totalDistance / float64(comparisons)
}

// GenerationStats tracks evolution statistics per generation
type GenerationStats struct {
	Generation     int
	BestFitness    float64
	AvgFitness     float64
	Diversity      float64
	PopulationSize int
	Timestamp      time.Time
}

// String returns formatted generation stats
func (gs GenerationStats) String() string {
	return fmt.Sprintf("Gen %d: Best=%.3f, Avg=%.3f, Diversity=%.3f, Pop=%d",
		gs.Generation, gs.BestFitness, gs.AvgFitness, gs.Diversity, gs.PopulationSize)
}

// NPUEvolutionaryHistory tracks complete evolutionary history
type NPUEvolutionaryHistory struct {
	Generations []GenerationStats
	StartTime   time.Time
	EndTime     time.Time
	FinalBest   *NPUDriver
	Converged   bool
}

// String returns formatted evolution history
func (eh *NPUEvolutionaryHistory) String() string {
	result := "NPU Evolutionary History\n"
	result += "========================\n"
	result += fmt.Sprintf("Duration: %v\n", eh.EndTime.Sub(eh.StartTime))
	result += fmt.Sprintf("Generations: %d\n", len(eh.Generations))
	result += fmt.Sprintf("Converged: %v\n\n", eh.Converged)
	
	result += "Generation Stats:\n"
	for _, stats := range eh.Generations {
		result += fmt.Sprintf("  %s\n", stats.String())
	}
	
	if eh.FinalBest != nil {
		genome := eh.FinalBest.AssessEntelechy()
		result += fmt.Sprintf("\nFinal Best Fitness: %.3f [%s]\n", 
			genome.Fitness, genome.ActualizationLevel)
	}
	
	return result
}

// Utility functions

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxFloat64(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func avgFloat64(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// CreateOntogeneticKernel creates an ontogenetic kernel for NPU
func CreateOntogeneticKernel(npu *NPUDriver) *ontogenesis.OntogeneticKernel {
	// Extract genome from NPU entelechy
	genome := npu.AssessEntelechy()
	
	// Create kernel genome with NPU characteristics
	coefficients := []float64{
		genome.Genes.Ontological,
		genome.Genes.Teleological,
		genome.Genes.Cognitive,
		genome.Genes.Integrative,
		genome.Genes.Evolutionary,
	}
	
	treeStructure := []int{-1, 0, 0, 1, 2} // Hierarchical tree
	
	kernelGenome := ontogenesis.NewKernelGenome(coefficients, treeStructure)
	kernelGenome.Generation = npu.generation
	
	// Create ontogenetic kernel
	kernel := &ontogenesis.OntogeneticKernel{
		ID:         fmt.Sprintf("npu-kernel-%d", npu.generation),
		Generation: npu.generation,
		Genome:     kernelGenome,
		BirthTime:  time.Now(),
		Age:        0,
		Fitness:    genome.Fitness,
		State:      make([]float64, len(coefficients)),
		Output:     make([]float64, len(coefficients)),
		Metadata:   make(map[string]interface{}),
	}
	
	// Store NPU reference in metadata
	kernel.Metadata["npu_id"] = npu.entelechyGenome.ID
	kernel.Metadata["actualization_level"] = genome.ActualizationLevel
	
	return kernel
}

// ApplyKernelToNPU applies ontogenetic kernel evolution to NPU
func ApplyKernelToNPU(kernel *ontogenesis.OntogeneticKernel, npu *NPUDriver) {
	// Apply kernel coefficients to NPU dimensions
	if len(kernel.Genome.Coefficients) >= 5 {
		npu.mu.Lock()
		defer npu.mu.Unlock()
		
		npu.dimensions.Ontological.CoreHealth = clamp(kernel.Genome.Coefficients[0], 0, 1)
		npu.dimensions.Teleological.PurposeClarity = clamp(kernel.Genome.Coefficients[1], 0, 1)
		npu.dimensions.Cognitive.LearningCapacity = clamp(kernel.Genome.Coefficients[2], 0, 1)
		npu.dimensions.Integrative.BuildHealth = clamp(kernel.Genome.Coefficients[3], 0, 1)
		npu.dimensions.Evolutionary.SelfImprovementCapacity = clamp(kernel.Genome.Coefficients[4], 0, 1)
	}
	
	// Re-assess after kernel application
	npu.AssessEntelechy()
}
