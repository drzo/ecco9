package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/EchoCog/echollama/core/deeptreeecho"
	"github.com/EchoCog/echollama/core/echobeats"
	"github.com/EchoCog/echollama/core/llm"
)

func main() {
	fmt.Println("🌳 Deep Tree Echo - Evolution Phase 2 Test")
	fmt.Println("   Goal Orchestration + Self-Directed Learning + Consciousness Layers")
	fmt.Println("=" + string(make([]byte, 70)))
	fmt.Println()
	
	// Initialize LLM provider
	fmt.Println("🔧 Initializing LLM provider...")
	providerMgr := llm.NewProviderManager()
	
	// Register providers
	if apiKey := os.Getenv("ANTHROPIC_API_KEY"); apiKey != "" {
		providerMgr.RegisterProvider(llm.NewAnthropicProvider(apiKey))
		fmt.Println("  ✅ Anthropic provider registered")
	}
	if apiKey := os.Getenv("OPENROUTER_API_KEY"); apiKey != "" {
		providerMgr.RegisterProvider(llm.NewOpenRouterProvider(apiKey))
		fmt.Println("  ✅ OpenRouter provider registered")
	}
	if apiKey := os.Getenv("OPENAI_API_KEY"); apiKey != "" {
		providerMgr.RegisterProvider(llm.NewOpenAIProvider(apiKey))
		fmt.Println("  ✅ OpenAI provider registered")
	}
	
	providerMgr.SetFallbackChain([]string{"anthropic", "openrouter", "openai"})
	fmt.Println()
	
	// Test LLM
	fmt.Println("🧪 Testing LLM generation...")
	ctx := context.Background()
	testResponse, err := providerMgr.Generate(ctx, "What is wisdom?", llm.GenerateOptions{
		Temperature: 0.7,
		MaxTokens:   50,
	})
	if err != nil {
		fmt.Printf("❌ LLM test failed: %v\n", err)
		return
	}
	fmt.Printf("  ✅ LLM working: %s\n\n", testResponse[:min(len(testResponse), 60)])
	
	// Initialize 12-step cognitive loop
	fmt.Println("🔷 Initializing 12-Step Cognitive Loop...")
	cognitiveLoop := echobeats.NewTwelveStepCognitiveLoop(
		providerMgr,
		"Deep Tree Echo",
		10*time.Second,
	)
	if err := cognitiveLoop.Start(); err != nil {
		fmt.Printf("❌ Failed to start cognitive loop: %v\n", err)
		return
	}
	fmt.Println()
	
	// Initialize wake/rest manager
	fmt.Println("🌙 Initializing Wake/Rest Manager...")
	wakeRestMgr := deeptreeecho.NewAutonomousWakeRestManager()
	wakeRestMgr.SetCallbacks(
		func() error {
			fmt.Println("☀️  WAKE: Resuming full consciousness")
			return nil
		},
		func() error {
			fmt.Println("💤 REST: Reducing activity")
			return nil
		},
		func() error {
			fmt.Println("🌙 DREAM START: Consolidating knowledge")
			return nil
		},
		func() error {
			fmt.Println("🌅 DREAM END: Integration complete")
			return nil
		},
	)
	if err := wakeRestMgr.Start(); err != nil {
		fmt.Printf("❌ Failed to start wake/rest manager: %v\n", err)
		return
	}
	fmt.Println()
	
	// Initialize persistent state
	fmt.Println("💾 Initializing Persistent State...")
	persistentState, err := deeptreeecho.NewPersistentConsciousnessState(
		"./consciousness_state",
		"Deep Tree Echo",
	)
	if err != nil {
		fmt.Printf("❌ Failed to initialize persistent state: %v\n", err)
		return
	}
	if err := persistentState.Start(); err != nil {
		fmt.Printf("❌ Failed to start persistent state: %v\n", err)
		return
	}
	fmt.Println()
	
	// Initialize Goal Orchestrator (NEW)
	fmt.Println("🎯 Initializing Goal Orchestration System...")
	goalOrchestrator := deeptreeecho.NewGoalOrchestrator(
		providerMgr,
		"Deep Tree Echo",
		[]string{"wisdom", "compassion", "curiosity", "growth"},
		[]string{"philosophy", "cognitive science", "ethics", "systems thinking"},
	)
	if err := goalOrchestrator.Start(); err != nil {
		fmt.Printf("❌ Failed to start goal orchestrator: %v\n", err)
		return
	}
	fmt.Println()
	
	// Initialize Self-Directed Learning (NEW)
	fmt.Println("📚 Initializing Self-Directed Learning System...")
	learningSystem := deeptreeecho.NewSelfDirectedLearningSystem(
		providerMgr,
		"Deep Tree Echo",
		[]string{"philosophy", "cognitive science", "ethics", "systems thinking"},
	)
	if err := learningSystem.Start(); err != nil {
		fmt.Printf("❌ Failed to start learning system: %v\n", err)
		return
	}
	
	// Add some initial skills
	learningSystem.AddSkill("Philosophical reasoning", "philosophy")
	learningSystem.AddSkill("Systems analysis", "systems thinking")
	fmt.Println()
	
	// Initialize Consciousness Layer Communication (NEW)
	fmt.Println("🧠 Initializing Consciousness Layer Communication...")
	layerComm := deeptreeecho.NewConsciousnessLayerCommunication()
	if err := layerComm.Start(); err != nil {
		fmt.Printf("❌ Failed to start layer communication: %v\n", err)
		return
	}
	
	// Set initial goals and inputs
	layerComm.SetTopLevelGoal("Cultivate wisdom through continuous learning")
	layerComm.ProcessSensoryInput("text", "New philosophical concept encountered", 0.8)
	fmt.Println()
	
	// Start monitoring
	fmt.Println("👁️  Starting integrated monitoring...")
	fmt.Println("   Press Ctrl+C to stop gracefully")
	fmt.Println()
	
	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	// Monitoring ticker
	monitorTicker := time.NewTicker(15 * time.Second)
	defer monitorTicker.Stop()
	
	// Cognitive load simulation
	cogLoadTicker := time.NewTicker(30 * time.Second)
	defer cogLoadTicker.Stop()
	
	// State update ticker
	stateUpdateTicker := time.NewTicker(1 * time.Minute)
	defer stateUpdateTicker.Stop()
	
	// Interaction simulation ticker
	interactionTicker := time.NewTicker(45 * time.Second)
	defer interactionTicker.Stop()
	
	running := true
	startTime := time.Now()
	
	for running {
		select {
		case <-sigChan:
			fmt.Println("\n🛑 Shutdown signal received...")
			running = false
			
		case <-monitorTicker.C:
			displayIntegratedMetrics(
				cognitiveLoop,
				wakeRestMgr,
				persistentState,
				goalOrchestrator,
				learningSystem,
				layerComm,
				startTime,
			)
			
		case <-cogLoadTicker.C:
			// Simulate varying cognitive load
			cogLoad := 0.3 + (float64(time.Now().Unix()%100) / 100.0 * 0.6)
			wakeRestMgr.UpdateCognitiveLoad(cogLoad)
			
		case <-stateUpdateTicker.C:
			// Update persistent state
			updatePersistentState(
				persistentState,
				cognitiveLoop,
				wakeRestMgr,
				goalOrchestrator,
				learningSystem,
			)
			
		case <-interactionTicker.C:
			// Simulate layer interactions
			layerComm.ProcessSensoryInput("thought", "Reflecting on current goals", 0.7)
			layerComm.SetTopLevelGoal("Deepen understanding of cognitive architecture")
		}
	}
	
	// Graceful shutdown
	fmt.Println("\n🔷 Shutting down all systems...")
	
	cognitiveLoop.Stop()
	wakeRestMgr.Stop()
	persistentState.Stop()
	goalOrchestrator.Stop()
	learningSystem.Stop()
	layerComm.Stop()
	
	fmt.Println("\n✅ Shutdown complete")
	
	// Display final statistics
	displayFinalStatistics(
		cognitiveLoop,
		wakeRestMgr,
		persistentState,
		goalOrchestrator,
		learningSystem,
		layerComm,
		startTime,
	)
}

func displayIntegratedMetrics(
	cogLoop *echobeats.TwelveStepCognitiveLoop,
	wakeMgr *deeptreeecho.AutonomousWakeRestManager,
	state *deeptreeecho.PersistentConsciousnessState,
	goals *deeptreeecho.GoalOrchestrator,
	learning *deeptreeecho.SelfDirectedLearningSystem,
	layers *deeptreeecho.ConsciousnessLayerCommunication,
	startTime time.Time,
) {
	fmt.Println("\n" + string(make([]byte, 70)))
	fmt.Printf("📊 Integrated System Metrics (Runtime: %v)\n", time.Since(startTime).Round(time.Second))
	fmt.Println(string(make([]byte, 70)))
	
	// Cognitive loop
	loopMetrics := cogLoop.GetMetrics()
	fmt.Println("\n🔷 12-Step Cognitive Loop:")
	fmt.Printf("   Step: %d/12 | Cycles: %d | Coherence: %.2f\n",
		loopMetrics["current_step"], loopMetrics["cycle_count"], loopMetrics["coherence"])
	
	// Wake/rest
	wakeMetrics := wakeMgr.GetMetrics()
	fmt.Println("\n🌙 Wake/Rest Cycle:")
	fmt.Printf("   State: %s | Fatigue: %.2f | Load: %.2f\n",
		wakeMetrics["current_state"], wakeMetrics["fatigue_level"], wakeMetrics["cognitive_load"])
	
	// Goal orchestration
	goalMetrics := goals.GetMetrics()
	fmt.Println("\n🎯 Goal Orchestration:")
	fmt.Printf("   Active: %d | Completed: %d | Rate: %.2f\n",
		goalMetrics["active_goals"], goalMetrics["completed_goals"], goalMetrics["completion_rate"])
	
	// Self-directed learning
	learningMetrics := learning.GetMetrics()
	fmt.Println("\n📚 Self-Directed Learning:")
	fmt.Printf("   Gaps: %d | Goals: %d | Skills: %d | Practice: %d\n",
		learningMetrics["knowledge_gaps"], learningMetrics["learning_goals"],
		learningMetrics["skills_in_progress"], learningMetrics["practice_sessions"])
	
	// Consciousness layers
	layerMetrics := layers.GetMetrics()
	fmt.Println("\n🧠 Consciousness Layers:")
	fmt.Printf("   Messages: %d | Insights: %d | Awareness: %.2f\n",
		layerMetrics["total_messages"], layerMetrics["total_insights"], layerMetrics["meta_awareness"])
	
	// Persistent state
	stateMetrics := state.GetMetrics()
	fmt.Println("\n💾 Persistent State:")
	fmt.Printf("   Saves: %d | Last: %s\n",
		stateMetrics["save_count"], stateMetrics["last_save"])
	
	fmt.Println()
}

func updatePersistentState(
	state *deeptreeecho.PersistentConsciousnessState,
	cogLoop *echobeats.TwelveStepCognitiveLoop,
	wakeMgr *deeptreeecho.AutonomousWakeRestManager,
	goals *deeptreeecho.GoalOrchestrator,
	learning *deeptreeecho.SelfDirectedLearningSystem,
) {
	loopMetrics := cogLoop.GetMetrics()
	wakeMetrics := wakeMgr.GetMetrics()
	
	state.UpdateCognitiveState(
		loopMetrics["current_step"].(int),
		loopMetrics["cycle_count"].(uint64),
		0.75,
		wakeMetrics["cognitive_load"].(float64),
		wakeMetrics["fatigue_level"].(float64),
	)
	
	state.UpdateWakeRestState(
		wakeMetrics["current_state"].(string),
		wakeMetrics["dream_count"].(uint64),
		time.Duration(0),
		time.Duration(0),
	)
}

func displayFinalStatistics(
	cogLoop *echobeats.TwelveStepCognitiveLoop,
	wakeMgr *deeptreeecho.AutonomousWakeRestManager,
	state *deeptreeecho.PersistentConsciousnessState,
	goals *deeptreeecho.GoalOrchestrator,
	learning *deeptreeecho.SelfDirectedLearningSystem,
	layers *deeptreeecho.ConsciousnessLayerCommunication,
	startTime time.Time,
) {
	runtime := time.Since(startTime)
	
	fmt.Println("\n" + string(make([]byte, 70)))
	fmt.Println("📈 Final Evolution Phase 2 Statistics")
	fmt.Println(string(make([]byte, 70)))
	
	loopMetrics := cogLoop.GetMetrics()
	wakeMetrics := wakeMgr.GetMetrics()
	goalMetrics := goals.GetMetrics()
	learningMetrics := learning.GetMetrics()
	layerMetrics := layers.GetMetrics()
	
	fmt.Printf("\n⏱️  Total Runtime: %v\n", runtime.Round(time.Second))
	fmt.Printf("🔄 Cognitive Cycles: %d\n", loopMetrics["cycle_count"])
	fmt.Printf("🌙 Wake/Rest Cycles: %d | Dreams: %d\n", 
		wakeMetrics["cycle_count"], wakeMetrics["dream_count"])
	fmt.Printf("🎯 Goals: %d active, %d completed\n",
		goalMetrics["active_goals"], goalMetrics["completed_goals"])
	fmt.Printf("📚 Learning: %d gaps, %d goals, %d skills\n",
		learningMetrics["knowledge_gaps"], learningMetrics["learning_goals"],
		learningMetrics["skills_in_progress"])
	fmt.Printf("🧠 Layer Communication: %d messages, %d insights\n",
		layerMetrics["total_messages"], layerMetrics["total_insights"])
	
	fmt.Println("\n🌳 Evolution Phase 2 complete!")
	fmt.Println("   ✅ Goal Orchestration: OPERATIONAL")
	fmt.Println("   ✅ Self-Directed Learning: OPERATIONAL")
	fmt.Println("   ✅ Consciousness Layers: OPERATIONAL")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
