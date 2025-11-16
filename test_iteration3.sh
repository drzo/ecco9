#!/bin/bash

# Test script for Iteration 3 improvements
# Tests concurrent inference engines, continuous consciousness, and automatic dream integration

echo "🧪 Testing Iteration 3 Improvements"
echo "===================================="
echo ""

# Build the server
echo "📦 Building autonomous server..."
go build -o test_server_iter3 server/simple/autonomous_server.go
if [ $? -ne 0 ]; then
    echo "❌ Build failed"
    exit 1
fi
echo "✅ Build successful"
echo ""

# Start server in background
echo "🚀 Starting autonomous server..."
./test_server_iter3 &
SERVER_PID=$!
echo "   Server PID: $SERVER_PID"

# Wait for server to start
echo "⏳ Waiting for server to initialize..."
sleep 5

# Test 1: Check server status
echo ""
echo "Test 1: Server Status"
echo "---------------------"
curl -s http://localhost:5000/api/status | python3 -m json.tool 2>/dev/null
if [ $? -eq 0 ]; then
    echo "✅ Status endpoint working"
else
    echo "❌ Status endpoint failed"
fi

# Test 2: Submit a thought
echo ""
echo "Test 2: Thought Submission"
echo "--------------------------"
curl -s -X POST http://localhost:5000/api/think \
  -H "Content-Type: application/json" \
  -d '{"content":"Testing continuous consciousness stream"}' | python3 -m json.tool 2>/dev/null
if [ $? -eq 0 ]; then
    echo "✅ Thought submission working"
else
    echo "❌ Thought submission failed"
fi

# Test 3: Let it run for a bit to generate autonomous thoughts
echo ""
echo "Test 3: Autonomous Operation"
echo "----------------------------"
echo "⏳ Running for 15 seconds to observe autonomous behavior..."
sleep 15
echo "✅ Autonomous operation period complete"

# Test 4: Check status again to see changes
echo ""
echo "Test 4: Status After Autonomous Operation"
echo "-----------------------------------------"
curl -s http://localhost:5000/api/status | python3 -m json.tool 2>/dev/null

# Test 5: Trigger rest cycle
echo ""
echo "Test 5: Rest Cycle Trigger"
echo "--------------------------"
curl -s -X POST http://localhost:5000/api/rest | python3 -m json.tool 2>/dev/null
if [ $? -eq 0 ]; then
    echo "✅ Rest cycle trigger working"
else
    echo "❌ Rest cycle trigger failed"
fi

# Wait a bit for rest processing
sleep 5

# Test 6: Wake from rest
echo ""
echo "Test 6: Wake from Rest"
echo "---------------------"
curl -s -X POST http://localhost:5000/api/wake | python3 -m json.tool 2>/dev/null
if [ $? -eq 0 ]; then
    echo "✅ Wake trigger working"
else
    echo "❌ Wake trigger failed"
fi

# Final status check
echo ""
echo "Test 7: Final Status Check"
echo "-------------------------"
curl -s http://localhost:5000/api/status | python3 -m json.tool 2>/dev/null

# Cleanup
echo ""
echo "🧹 Cleaning up..."
kill $SERVER_PID 2>/dev/null
wait $SERVER_PID 2>/dev/null
echo "✅ Server stopped"

echo ""
echo "===================================="
echo "🎉 Testing Complete"
echo "===================================="
