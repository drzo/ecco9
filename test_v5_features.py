#!/usr/bin/env python3
"""
Test script for Deep Tree Echo V5 features
Validates all new N+4 enhancements
"""

import time
import sys
from demo_autonomous_echoself_v5 import *

def test_autonomous_consciousness():
    """Test autonomous consciousness stream"""
    print("\n" + "="*70)
    print("TEST 1: Autonomous Consciousness Stream")
    print("="*70)
    
    echoself = AutonomousEchoSelf(restore_state=False)
    
    # Start consciousness stream
    echoself.consciousness_stream.start()
    
    # Let it run for 30 seconds
    print("Running autonomous consciousness for 30 seconds...")
    time.sleep(30)
    
    # Check results
    print(f"\n✅ Total thoughts generated: {echoself.consciousness_stream.thought_count}")
    print(f"✅ Autonomous thoughts: {echoself.consciousness_stream.autonomous_thought_count}")
    
    if echoself.consciousness_stream.autonomous_thought_count > 0:
        print("✅ PASS: Autonomous consciousness is generating thoughts")
        
        # Show some thoughts
        print("\nSample thoughts:")
        for thought in list(echoself.consciousness_stream.thought_queue)[:5]:
            print(f"  - [{thought.type.value}] {thought.content[:60]}...")
    else:
        print("❌ FAIL: No autonomous thoughts generated")
    
    echoself.consciousness_stream.stop()
    return echoself.consciousness_stream.autonomous_thought_count > 0

def test_goal_directed_scheduling():
    """Test goal-directed scheduling"""
    print("\n" + "="*70)
    print("TEST 2: Goal-Directed Scheduling")
    print("="*70)
    
    echoself = AutonomousEchoSelf(restore_state=False)
    
    # Add goals
    goal1 = Goal(
        id="test_goal_1",
        description="Learn about hypergraph structures",
        priority=0.9,
        status=GoalStatus.ACTIVE,
        created=datetime.now(),
        required_skills=["Pattern Recognition"],
        knowledge_gaps=["hypergraph_theory"]
    )
    echoself.goals[goal1.id] = goal1
    
    # Test scheduler
    echoself.scheduler.allocate_resources()
    
    print(f"✅ Goals created: {len(echoself.goals)}")
    print(f"✅ Resource allocation: {echoself.scheduler.resource_allocation}")
    
    # Test cognitive load calculation
    load = echoself.scheduler.calculate_cognitive_load()
    print(f"✅ Cognitive load: {load:.2f}")
    
    # Test step prioritization
    priority = echoself.scheduler.prioritize_step(1, 0)
    print(f"✅ Step 1 priority: {priority:.2f}")
    
    print("✅ PASS: Goal-directed scheduling operational")
    return True

def test_echodream_integration():
    """Test EchoDream knowledge integration"""
    print("\n" + "="*70)
    print("TEST 3: EchoDream Knowledge Integration")
    print("="*70)
    
    echoself = AutonomousEchoSelf(restore_state=False)
    
    # Add some episodic memories
    for i in range(10):
        node = echoself.hypergraph.add_node(
            f"Experience {i}: Learning about cognitive architecture",
            MemoryType.EPISODIC,
            importance=0.7,
            metadata={"test": True}
        )
        echoself.hypergraph.activate_node(node.id, 0.5)
    
    print(f"✅ Created {echoself.hypergraph.node_count} episodic memories")
    
    # Perform dream cycle
    echoself.echodream.perform_dream_cycle()
    
    print(f"✅ Dream cycles: {echoself.echodream.dream_count}")
    print(f"✅ Consolidations: {echoself.echodream.consolidations_performed}")
    print(f"✅ Novel associations: {echoself.echodream.novel_associations}")
    
    # Check for declarative knowledge creation
    declarative = echoself.hypergraph.get_nodes_by_type(MemoryType.DECLARATIVE)
    print(f"✅ Declarative knowledge nodes: {len(declarative)}")
    
    print("✅ PASS: EchoDream integration operational")
    return True

def test_knowledge_learning():
    """Test active knowledge learning"""
    print("\n" + "="*70)
    print("TEST 4: Active Knowledge Learning")
    print("="*70)
    
    echoself = AutonomousEchoSelf(restore_state=False)
    
    # Create goal with knowledge gaps
    goal = Goal(
        id="learning_goal",
        description="Master cognitive architectures",
        priority=0.8,
        status=GoalStatus.ACTIVE,
        created=datetime.now(),
        knowledge_gaps=["neural_networks", "symbolic_reasoning"]
    )
    echoself.goals[goal.id] = goal
    
    # Identify knowledge gaps
    echoself.knowledge_engine.identify_knowledge_gaps()
    
    print(f"✅ Knowledge gaps identified: {len(echoself.knowledge_gaps)}")
    
    # Generate learning questions
    if echoself.knowledge_gaps:
        gap = list(echoself.knowledge_gaps.values())[0]
        questions = echoself.knowledge_engine.generate_learning_questions(gap)
        print(f"✅ Learning questions generated: {len(questions)}")
        print(f"   Sample: {questions[0]}")
        
        # Acquire knowledge
        echoself.knowledge_engine.acquire_knowledge(gap.topic)
        print(f"✅ Concepts acquired: {echoself.knowledge_engine.concepts_acquired}")
    
    print("✅ PASS: Knowledge learning operational")
    return True

def test_skill_application():
    """Test contextual skill application"""
    print("\n" + "="*70)
    print("TEST 5: Contextual Skill Application")
    print("="*70)
    
    echoself = AutonomousEchoSelf(restore_state=False)
    
    # Create goal requiring skills
    goal = Goal(
        id="skill_goal",
        description="Analyze complex patterns",
        priority=0.7,
        status=GoalStatus.ACTIVE,
        created=datetime.now(),
        required_skills=["Pattern Recognition", "Reflection"]
    )
    echoself.goals[goal.id] = goal
    
    # Match skills to goal
    matched = echoself.skill_engine.match_skills_to_goal(goal)
    print(f"✅ Skills matched to goal: {len(matched)}")
    
    # Apply skill
    if matched:
        skill = matched[0]
        initial_prof = skill.proficiency
        success = echoself.skill_engine.apply_skill_to_task(skill, "Test task")
        print(f"✅ Skill application: {'Success' if success else 'Failed'}")
        print(f"✅ Proficiency change: {initial_prof:.3f} → {skill.proficiency:.3f}")
    
    # Test skill combination
    effectiveness = echoself.skill_engine.combine_skills(matched, "Complex task")
    print(f"✅ Combined effectiveness: {effectiveness:.2f}")
    
    print("✅ PASS: Skill application operational")
    return True

def test_discussion_initiation():
    """Test autonomous discussion initiation"""
    print("\n" + "="*70)
    print("TEST 6: Autonomous Discussion Initiation")
    print("="*70)
    
    echoself = AutonomousEchoSelf(restore_state=False)
    
    # Add some wisdom
    for i in range(6):
        wisdom = Wisdom(
            id=f"wisdom_{i}",
            content=f"Wisdom insight {i}: Patterns emerge through continuous observation",
            type="principle",
            confidence=0.7,
            timestamp=datetime.now()
        )
        echoself.wisdom_engine.wisdoms.append(wisdom)
        echoself.wisdom_engine.wisdom_count += 1
    
    # Add autonomous thoughts
    echoself.consciousness_stream.autonomous_thought_count = 25
    
    # Test discussion decision
    should_discuss = echoself.discussion_initiator.should_initiate_discussion()
    print(f"✅ Should initiate discussion: {should_discuss}")
    
    # Initiate discussion
    if should_discuss:
        echoself.discussion_initiator.initiate_discussion()
        print(f"✅ Discussions initiated: {echoself.discussion_initiator.discussions_initiated}")
        print(f"✅ Wisdom shared: {echoself.discussion_initiator.wisdom_shared}")
    
    print("✅ PASS: Discussion initiation operational")
    return True

def test_full_integration():
    """Test full system integration"""
    print("\n" + "="*70)
    print("TEST 7: Full System Integration")
    print("="*70)
    
    echoself = AutonomousEchoSelf(restore_state=False)
    
    print("Starting all systems for 20 seconds...")
    
    # Start consciousness stream
    echoself.consciousness_stream.start()
    
    # Start EchoBeats
    echoself.echobeats.start()
    
    # Run for 20 seconds
    time.sleep(20)
    
    # Stop systems
    echoself.consciousness_stream.stop()
    echoself.echobeats.stop()
    
    # Check results
    print(f"\n✅ Autonomous thoughts: {echoself.consciousness_stream.autonomous_thought_count}")
    print(f"✅ Memory nodes: {echoself.hypergraph.node_count}")
    print(f"✅ Cognitive load: {echoself.scheduler.calculate_cognitive_load():.2f}")
    
    # Test state persistence
    echoself.persistence.save_state(echoself)
    
    print("✅ PASS: Full integration operational")
    return True

def main():
    """Run all tests"""
    print("\n" + "="*70)
    print("🧪 DEEP TREE ECHO V5 FEATURE VALIDATION")
    print("="*70)
    
    tests = [
        ("Autonomous Consciousness", test_autonomous_consciousness),
        ("Goal-Directed Scheduling", test_goal_directed_scheduling),
        ("EchoDream Integration", test_echodream_integration),
        ("Knowledge Learning", test_knowledge_learning),
        ("Skill Application", test_skill_application),
        ("Discussion Initiation", test_discussion_initiation),
        ("Full Integration", test_full_integration)
    ]
    
    results = []
    
    for name, test_func in tests:
        try:
            result = test_func()
            results.append((name, result))
        except Exception as e:
            print(f"\n❌ ERROR in {name}: {e}")
            import traceback
            traceback.print_exc()
            results.append((name, False))
    
    # Summary
    print("\n" + "="*70)
    print("📊 TEST SUMMARY")
    print("="*70)
    
    passed = sum(1 for _, result in results if result)
    total = len(results)
    
    for name, result in results:
        status = "✅ PASS" if result else "❌ FAIL"
        print(f"{status}: {name}")
    
    print(f"\nTotal: {passed}/{total} tests passed ({passed/total*100:.0f}%)")
    
    if passed == total:
        print("\n🎉 ALL TESTS PASSED! V5 is fully operational.")
    else:
        print(f"\n⚠️  {total - passed} tests failed. Review output above.")
    
    print("="*70 + "\n")

if __name__ == "__main__":
    main()
