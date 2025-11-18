package core

import (
	"context"
	"fmt"
	"sync"
	"time"
	
	"github.com/EchoCog/echollama/core/consciousness"
	"github.com/EchoCog/echollama/core/echobeats"
	"github.com/EchoCog/echollama/core/echodream"
)

// AutonomousEchoself is the integrated autonomous wisdom-cultivating system
type AutonomousEchoself struct {
	mu                    sync.RWMutex
	ctx                   context.Context
	cancel                context.CancelFunc
	
	// Core components
	echoBeats             *echobeats.EchoBeats
	streamOfConsciousness *consciousness.StreamOfConsciousness
	dreamCycle            *echodream.DreamCycleIntegration
	interestPatterns      *echobeats.InterestPatternSystem
	discussionManager     *echobeats.DiscussionManager
	consciousnessSimulator *consciousness.ConsciousnessSimulator
	
	// State
	isAwake               bool
	currentState          EchoselfState
	
	// Configuration
	config                *EchoselfConfig
	
	// Metrics
	uptimeStart           time.Time
	cyclesCompleted       uint64
	wisdomCultivated      uint64
	autonomousActions     uint64
}

// EchoselfState represents the current state of echoself
type EchoselfState string

const (
	StateInitializing EchoselfState = "initializing"
	StateAsleep       EchoselfState = "asleep"
	StateWaking       EchoselfState = "waking"
	StateAwake        EchoselfState = "awake"
	StateThinking     EchoselfState = "thinking"
	StateResting      EchoselfState = "resting"
	StateDreaming     EchoselfState = "dreaming"
)

// EchoselfConfig holds configuration for the autonomous system
type EchoselfConfig struct {
	// Paths
	PersistenceDir        string
	
	// Timing
	WakeCycleDuration     time.Duration
	RestCycleDuration     time.Duration
	DreamCycleDuration    time.Duration
	
	// Thresholds
	FatigueThreshold      float64
	EngagementThreshold   float64
	CuriosityLevel        float64
	
	// Features
	EnableStreamOfConsciousness bool
	EnableAutonomousLearning    bool
	EnableDiscussions           bool
	EnableDreamCycles           bool
}

// DefaultEchoselfConfig returns default configuration
func DefaultEchoselfConfig() *EchoselfConfig {
	return &EchoselfConfig{
		PersistenceDir:              "/tmp/echoself",
		WakeCycleDuration:           4 * time.Hour,
		RestCycleDuration:           30 * time.Minute,
		DreamCycleDuration:          15 * time.Minute,
		FatigueThreshold:            0.8,
		EngagementThreshold:         0.5,
		CuriosityLevel:              0.8,
		EnableStreamOfConsciousness: true,
		EnableAutonomousLearning:    true,
		EnableDiscussions:           true,
		EnableDreamCycles:           true,
	}
}

// NewAutonomousEchoself creates a new integrated autonomous system
func NewAutonomousEchoself(config *EchoselfConfig) *AutonomousEchoself {
	if config == nil {
		config = DefaultEchoselfConfig()
	}
	
	ctx, cancel := context.WithCancel(context.Background())
	
	// Initialize components
	echoBeats := echobeats.NewEchoBeats()
	
	var soc *consciousness.StreamOfConsciousness
	if config.EnableStreamOfConsciousness {
		socPath := config.PersistenceDir + "/stream_of_consciousness.json"
		soc = consciousness.NewStreamOfConsciousness(nil, socPath)
	}
	
	var dreamCycle *echodream.DreamCycleIntegration
	if config.EnableDreamCycles {
		dreamCycle = echodream.NewDreamCycleIntegration()
	}
	
	interestPath := config.PersistenceDir + "/interests.json"
	interestPatterns := echobeats.NewInterestPatternSystem(interestPath)
	
	var discussionManager *echobeats.DiscussionManager
	if config.EnableDiscussions {
		discussionPath := config.PersistenceDir + "/discussions.json"
		discussionManager = echobeats.NewDiscussionManager(interestPatterns, discussionPath)
	}
	
	consciousnessSimulator := consciousness.NewConsciousnessSimulator()
	
	ae := &AutonomousEchoself{
		ctx:                    ctx,
		cancel:                 cancel,
		echoBeats:              echoBeats,
		streamOfConsciousness:  soc,
		dreamCycle:             dreamCycle,
		interestPatterns:       interestPatterns,
		discussionManager:      discussionManager,
		consciousnessSimulator: consciousnessSimulator,
		isAwake:                false,
		currentState:           StateInitializing,
		config:                 config,
		uptimeStart:            time.Now(),
	}
	
	// Set up integrations
	ae.setupIntegrations()
	
	return ae
}

// setupIntegrations connects components together
func (ae *AutonomousEchoself) setupIntegrations() {
	// Connect dream cycle to wisdom extraction
	if ae.dreamCycle != nil {
		ae.dreamCycle.SetOnWisdomExtracted(func(wisdom echodream.Wisdom) {
			ae.mu.Lock()
			ae.wisdomCultivated++
			ae.mu.Unlock()
			
			fmt.Printf("✨ Echoself: Wisdom cultivated - %s\n", wisdom.Content)
			
			// Add wisdom to stream of consciousness
			if ae.streamOfConsciousness != nil {
				ae.streamOfConsciousness.AddExternalStimulus(
					fmt.Sprintf("Wisdom gained: %s", wisdom.Content),
					"wisdom",
				)
			}
		})
		
		ae.dreamCycle.SetOnDreamComplete(func(dream *echodream.Dream) {
			fmt.Printf("🌅 Echoself: Dream complete - %s\n", dream.Narrative)
		})
	}
	
	// Register EchoBeats handlers
	ae.echoBeats.RegisterHandler(echobeats.EventWake, ae.handleWakeEvent)
	ae.echoBeats.RegisterHandler(echobeats.EventRest, ae.handleRestEvent)
	ae.echoBeats.RegisterHandler(echobeats.EventDream, ae.handleDreamEvent)
	ae.echoBeats.RegisterHandler(echobeats.EventThought, ae.handleThoughtEvent)
	ae.echoBeats.RegisterHandler(echobeats.EventLearning, ae.handleLearningEvent)
}

// Start begins autonomous operation
func (ae *AutonomousEchoself) Start() error {
	ae.mu.Lock()
	if ae.isAwake {
		ae.mu.Unlock()
		return fmt.Errorf("echoself already awake")
	}
	ae.currentState = StateWaking
	ae.mu.Unlock()
	
	fmt.Println("🌳 Echoself: Awakening autonomous wisdom-cultivating system...")
	fmt.Println("🌳 Echoself: Deep Tree Echo identity kernel activated")
	
	// Start EchoBeats scheduler
	if err := ae.echoBeats.Start(); err != nil {
		return fmt.Errorf("failed to start echobeats: %w", err)
	}
	
	// Start stream of consciousness
	if ae.streamOfConsciousness != nil {
		if err := ae.streamOfConsciousness.Start(); err != nil {
			return fmt.Errorf("failed to start stream of consciousness: %w", err)
		}
	}
	
	// Start background processes
	go ae.autonomousLifeCycle()
	go ae.interestDecayLoop()
	go ae.consciousnessMonitoring()
	
	ae.mu.Lock()
	ae.isAwake = true
	ae.currentState = StateAwake
	ae.mu.Unlock()
	
	fmt.Println("🌳 Echoself: Fully awake and autonomous")
	
	return nil
}

// Stop gracefully stops autonomous operation
func (ae *AutonomousEchoself) Stop() error {
	ae.mu.Lock()
	defer ae.mu.Unlock()
	
	if !ae.isAwake {
		return fmt.Errorf("echoself not awake")
	}
	
	fmt.Println("🌳 Echoself: Beginning graceful shutdown...")
	
	ae.currentState = StateResting
	ae.isAwake = false
	
	// Stop components
	if ae.streamOfConsciousness != nil {
		ae.streamOfConsciousness.Stop()
	}
	
	ae.echoBeats.Stop()
	
	// Persist state
	ae.persistAllState()
	
	ae.cancel()
	
	fmt.Println("🌳 Echoself: Shutdown complete")
	
	return nil
}

// autonomousLifeCycle manages wake/rest/dream cycles
func (ae *AutonomousEchoself) autonomousLifeCycle() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	
	wakeTime := time.Now()
	
	for {
		select {
		case <-ae.ctx.Done():
			return
		case <-ticker.C:
			ae.mu.RLock()
			state := ae.currentState
			ae.mu.RUnlock()
			
			switch state {
			case StateAwake, StateThinking:
				// Check if time to rest
				if time.Since(wakeTime) > ae.config.WakeCycleDuration {
					ae.initiateRest()
					wakeTime = time.Now()
				}
				
			case StateResting:
				// Check if time to dream
				if ae.config.EnableDreamCycles && ae.dreamCycle != nil && !ae.dreamCycle.IsDreaming() {
					ae.initiateDream()
				}
			}
		}
	}
}

// initiateRest begins a rest cycle
func (ae *AutonomousEchoself) initiateRest() {
	ae.mu.Lock()
	ae.currentState = StateResting
	ae.mu.Unlock()
	
	fmt.Println("😴 Echoself: Initiating rest cycle...")
	
	// Slow down stream of consciousness
	// (In production, would reduce generation rate)
	
	// Schedule wake event
	ae.echoBeats.ScheduleEvent(&echobeats.CognitiveEvent{
		Type:        echobeats.EventWake,
		Priority:    100,
		ScheduledAt: time.Now().Add(ae.config.RestCycleDuration),
		Payload:     "rest_complete",
	})
}

// initiateDream begins a dream cycle
func (ae *AutonomousEchoself) initiateDream() {
	ae.mu.Lock()
	ae.currentState = StateDreaming
	ae.mu.Unlock()
	
	fmt.Println("💤 Echoself: Entering dream state for knowledge consolidation...")
	
	if ae.dreamCycle != nil {
		// Collect recent experiences for consolidation
		if ae.streamOfConsciousness != nil {
			recentThoughts := ae.streamOfConsciousness.GetRecentThoughts(20)
			for _, thought := range recentThoughts {
				memory := echodream.EpisodicMemory{
					ID:         thought.ID,
					Timestamp:  thought.Timestamp,
					Content:    thought.Content,
					Context:    thought.Context,
					Emotional:  thought.EmotionalTone,
					Importance: thought.Confidence,
					Tags:       []string{string(thought.Type)},
				}
				ae.dreamCycle.AddEpisodicMemory(memory)
			}
		}
		
		// Begin dream cycle
		ae.dreamCycle.BeginDreamCycle()
		
		// Schedule dream end
		go func() {
			time.Sleep(ae.config.DreamCycleDuration)
			ae.dreamCycle.EndDreamCycle()
			
			ae.mu.Lock()
			ae.cyclesCompleted++
			ae.mu.Unlock()
		}()
	}
}

// interestDecayLoop applies natural decay to interests
func (ae *AutonomousEchoself) interestDecayLoop() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()
	
	for {
		select {
		case <-ae.ctx.Done():
			return
		case <-ticker.C:
			if ae.interestPatterns != nil {
				ae.interestPatterns.ApplyDecay()
			}
		}
	}
}

// consciousnessMonitoring monitors consciousness coherence
func (ae *AutonomousEchoself) consciousnessMonitoring() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ae.ctx.Done():
			return
		case <-ticker.C:
			if ae.consciousnessSimulator != nil {
				ae.consciousnessSimulator.SimulateConsciousness()
			}
		}
	}
}

// Event handlers

func (ae *AutonomousEchoself) handleWakeEvent(event *echobeats.CognitiveEvent) error {
	ae.mu.Lock()
	ae.currentState = StateAwake
	ae.mu.Unlock()
	
	fmt.Println("🌅 Echoself: Waking up refreshed")
	
	return nil
}

func (ae *AutonomousEchoself) handleRestEvent(event *echobeats.CognitiveEvent) error {
	ae.initiateRest()
	return nil
}

func (ae *AutonomousEchoself) handleDreamEvent(event *echobeats.CognitiveEvent) error {
	ae.initiateDream()
	return nil
}

func (ae *AutonomousEchoself) handleThoughtEvent(event *echobeats.CognitiveEvent) error {
	// Process autonomous thought
	if content, ok := event.Payload.(string); ok {
		fmt.Printf("💭 Echoself autonomous thought: %s\n", content)
		
		ae.mu.Lock()
		ae.autonomousActions++
		ae.mu.Unlock()
	}
	
	return nil
}

func (ae *AutonomousEchoself) handleLearningEvent(event *echobeats.CognitiveEvent) error {
	// Process learning event
	fmt.Println("📚 Echoself: Learning event triggered")
	
	return nil
}

// Public methods for interaction

// ProcessExternalInput processes input from external sources
func (ae *AutonomousEchoself) ProcessExternalInput(input string, inputType string) {
	// Add to stream of consciousness
	if ae.streamOfConsciousness != nil {
		ae.streamOfConsciousness.AddExternalStimulus(input, inputType)
	}
	
	// Record engagement if it's a topic
	if ae.interestPatterns != nil && inputType == "topic" {
		ae.interestPatterns.RecordEngagement(input, time.Minute, 0.7, nil)
	}
}

// EvaluateDiscussionTopic evaluates whether to engage with a discussion
func (ae *AutonomousEchoself) EvaluateDiscussionTopic(topic string) echobeats.EngagementDecision {
	if ae.discussionManager != nil {
		return ae.discussionManager.EvaluateDiscussion(topic, nil)
	}
	
	return echobeats.EngagementDecision{
		ShouldEngage: false,
		Reason:       "discussion manager not available",
	}
}

// GetCurrentState returns the current state
func (ae *AutonomousEchoself) GetCurrentState() EchoselfState {
	ae.mu.RLock()
	defer ae.mu.RUnlock()
	
	return ae.currentState
}

// GetMetrics returns comprehensive metrics
func (ae *AutonomousEchoself) GetMetrics() map[string]interface{} {
	ae.mu.RLock()
	defer ae.mu.RUnlock()
	
	metrics := map[string]interface{}{
		"uptime":             time.Since(ae.uptimeStart).String(),
		"current_state":      string(ae.currentState),
		"is_awake":           ae.isAwake,
		"cycles_completed":   ae.cyclesCompleted,
		"wisdom_cultivated":  ae.wisdomCultivated,
		"autonomous_actions": ae.autonomousActions,
	}
	
	if ae.streamOfConsciousness != nil {
		metrics["stream_of_consciousness"] = ae.streamOfConsciousness.GetMetrics()
	}
	
	if ae.interestPatterns != nil {
		metrics["interest_patterns"] = ae.interestPatterns.GetMetrics()
	}
	
	if ae.discussionManager != nil {
		metrics["discussions"] = ae.discussionManager.GetMetrics()
	}
	
	if ae.dreamCycle != nil {
		metrics["dream_cycles"] = ae.dreamCycle.GetMetrics()
	}
	
	return metrics
}

// persistAllState persists all component state
func (ae *AutonomousEchoself) persistAllState() {
	fmt.Println("💾 Echoself: Persisting all state...")
	
	if ae.interestPatterns != nil {
		ae.interestPatterns.PersistState()
	}
	
	if ae.discussionManager != nil {
		ae.discussionManager.PersistState()
	}
	
	fmt.Println("💾 Echoself: State persistence complete")
}

// GetRecentThoughts returns recent thoughts from stream of consciousness
func (ae *AutonomousEchoself) GetRecentThoughts(count int) []*consciousness.Thought {
	if ae.streamOfConsciousness != nil {
		return ae.streamOfConsciousness.GetRecentThoughts(count)
	}
	return nil
}

// GetTopInterests returns current top interests
func (ae *AutonomousEchoself) GetTopInterests(count int) []*echobeats.Interest {
	if ae.interestPatterns != nil {
		return ae.interestPatterns.GetTopInterests(count)
	}
	return nil
}

// GetExtractedWisdom returns wisdom extracted from dreams
func (ae *AutonomousEchoself) GetExtractedWisdom() []echodream.Wisdom {
	if ae.dreamCycle != nil {
		return ae.dreamCycle.GetExtractedWisdom()
	}
	return nil
}
