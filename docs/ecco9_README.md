# ecco9: Cognitive Hardware Platform

**ecco9** is a virtual cognitive hardware platform that implements Deep Tree Echo's embodied cognition architecture as a realistic hardware simulation. It provides a complete virtualized device ecosystem including device drivers, firmware interfaces, I/O streams, and administrative utilities.

## 🌊 Quick Start

### Running the ecco9 Server

```bash
# Build the server
go build -o ecco9 ./server/simple/ecco9_server.go

# Run the server
./ecco9
```

The server will start on `http://localhost:5000` with a complete boot sequence:

```
🌊 Starting ecco9 Cognitive Hardware Platform Boot Sequence...
⚡ Stage 0: Hardware Initialization
💾 Stage 1: Memory Test
🌀 Stage 2: Reservoir Initialization
🧠 Stage 3: Cognitive Kernel Load
🌲 Stage 4: Deep Tree Echo Identity Load
🌐 Stage 5: Service Initialization
✨ ecco9 Platform Boot Complete - All Systems Ready
```

### Accessing the Dashboard

Visit `http://localhost:5000` in your web browser to see the real-time platform dashboard with:
- Platform status and uptime
- Registered devices and their states
- Available API endpoints
- Architecture overview

## 📡 API Endpoints

### Platform Status
```bash
curl http://localhost:5000/api/status
```

Returns platform information including firmware version, boot stage, device count, and port configuration.

### List Devices
```bash
curl http://localhost:5000/api/devices
```

Returns all registered cognitive devices with their IDs, names, types, status, and health.

### List Drivers
```bash
curl http://localhost:5000/api/drivers
```

Returns all loaded drivers with their names, versions, and capabilities.

### Health Check
```bash
curl http://localhost:5000/api/health
```

Returns overall platform health status.

### Text Generation
```bash
curl -X POST http://localhost:5000/api/generate \
  -H "Content-Type: application/json" \
  -d '{"prompt": "Hello, ecco9!", "model": "local"}'
```

Generates text using the ecco9 cognitive platform.

## 🔧 Architecture

### Virtual Hardware Components

**Core Processing Units:**
- **Echo State Reservoir Processors (ESRP)**: 4 reservoir cores for temporal pattern recognition
- **Consciousness Layer Processors (CLP)**: 3-layer awareness system (Basic/Reflective/Meta-cognitive)
- **Emotion Processing Unit (EPU)**: 10 discrete emotion channels + dimensional affect
- **Hypergraph Memory Array (HMA)**: Multi-relational knowledge storage
- **LLM Provider Integration**: Multi-provider text generation

**Port Allocation:**
- `5000` - Primary HTTP API
- `5001` - WebSocket Cognitive Stream (planned)
- `5002` - gRPC Service Bus (planned)
- `5003` - Telemetry & Monitoring (planned)
- `5004` - Admin Configuration (planned)
- `5005` - Debug & Profiling (planned)

### Device Drivers (CHAL - Cognitive Hardware Abstraction Layer)

1. **Reservoir Driver** - Echo State Network management
   - 4 reservoir cores (100 neurons each)
   - Spectral radius: 0.95
   - Leak rate: 0.3
   - Capabilities: echo-state-network, temporal-pattern-recognition, hierarchical-processing

2. **Consciousness Driver** - Multi-layer consciousness processing
   - 3 consciousness layers
   - Inter-layer message routing
   - Capabilities: multi-layer-processing, bottom-up-attention, top-down-modulation, metacognitive-monitoring

3. **Emotion Driver** - Emotion state tracking and modulation
   - 10 emotion channels (Izard's Differential Emotion Theory)
   - Dimensional affect (arousal/valence)
   - Automatic emotion decay
   - Capabilities: discrete-emotions, dimensional-affect, emotion-blending, cognitive-modulation

4. **Memory Driver** - Hypergraph memory management
   - 100,000 node capacity
   - 1,000,000 edge capacity
   - Automatic consolidation and pruning
   - Capabilities: hypergraph-storage, multi-relational, associative-retrieval, automatic-consolidation

5. **LLM Driver** - Language model provider integration
   - Multi-provider abstraction
   - Fallback routing
   - Capabilities: text-generation, streaming, multi-provider

### Boot Sequence

The ecco9 platform implements a 6-stage boot sequence:

1. **Stage 0: Hardware Initialization** - POST, device registry setup
2. **Stage 1: Memory Test** - Hypergraph memory validation
3. **Stage 2: Reservoir Initialization** - Echo state network setup
4. **Stage 3: Cognitive Kernel Load** - Kernel initialization
5. **Stage 4: Deep Tree Echo Identity Load** - Identity restoration
6. **Stage 5: Service Initialization** - API server startup

Total boot time: ~600ms

## 📦 Package Structure

```
core/ecco9/
├── types.go                      # Core types and interfaces
├── platform.go                   # Platform initialization and management
└── drivers/
    ├── reservoir_driver.go       # Echo State Reservoir driver
    ├── consciousness_driver.go   # Consciousness layer driver
    ├── emotion_driver.go         # Emotion processing driver
    ├── memory_driver.go          # Hypergraph memory driver
    └── llm_driver.go            # LLM provider driver
```

## 🧠 Core Concepts

### CognitiveDevice Interface

All cognitive devices implement the `CognitiveDevice` interface:

```go
type CognitiveDevice interface {
    // Lifecycle management
    Initialize(ctx context.Context) error
    Shutdown(ctx context.Context) error
    Reset(ctx context.Context) error
    
    // State management
    GetState() (DeviceState, error)
    SetState(DeviceState) error
    
    // I/O operations
    Read(buffer []byte) (int, error)
    Write(buffer []byte) (int, error)
    IoCtl(command uint32, arg interface{}) error
    
    // Telemetry
    GetMetrics() (DeviceMetrics, error)
    GetHealth() (HealthStatus, error)
    
    // Device information
    GetID() string
    GetName() string
    GetType() DeviceType
}
```

### Driver Interface

All drivers implement the `Driver` interface:

```go
type Driver interface {
    // Driver lifecycle
    Load(config interface{}) error
    Unload() error
    
    // Device management
    GetDevice(id string) (CognitiveDevice, error)
    ListDevices() []CognitiveDevice
    
    // Driver information
    GetName() string
    GetVersion() string
    GetCapabilities() []string
}
```

## 🎯 Current Implementation Status

### ✅ Completed Features

- [x] Core platform types and interfaces
- [x] 6-stage boot sequence
- [x] Reservoir driver (4 cores)
- [x] Consciousness driver (3 layers)
- [x] Emotion driver (10 channels)
- [x] Memory driver (hypergraph)
- [x] LLM driver integration
- [x] HTTP API server
- [x] Web dashboard
- [x] Graceful shutdown
- [x] Device metrics and health monitoring

### 🔄 In Progress

- [ ] WebSocket cognitive streams
- [ ] gRPC service bus
- [ ] Telemetry dashboard
- [ ] Admin configuration interface
- [ ] Debug/profiling endpoints
- [ ] EchoBeats 12-step cognitive loop
- [ ] Wake/rest cycle management

### 📋 Planned Features

- [ ] Firmware update protocol
- [ ] Device filesystem (/dev/ecco9)
- [ ] ecco9-admin CLI tool
- [ ] Prometheus metrics export
- [ ] GPU acceleration integration
- [ ] Multi-device clustering

## 🧪 Testing

### Manual Testing

```bash
# Start the server
go run server/simple/ecco9_server.go

# In another terminal, test the API
curl http://localhost:5000/api/status
curl http://localhost:5000/api/devices
curl http://localhost:5000/api/drivers
curl http://localhost:5000/api/health
```

### Automated Testing

```bash
# Run unit tests (when added)
go test ./core/ecco9/...

# Run integration tests (when added)
go test ./server/...
```

## 🤝 Contributing

Contributions are welcome! Areas for contribution:

1. **Additional Drivers**: Implement new cognitive device drivers
2. **Stream Processing**: WebSocket and gRPC implementations
3. **Telemetry**: Metrics collection and visualization
4. **Admin Tools**: CLI and web-based configuration
5. **Testing**: Unit and integration tests
6. **Documentation**: API docs, tutorials, examples

## 📄 License

Licensed under the same terms as the main ecco9 repository (MIT License).

## 🌊 Deep Tree Echo Integration

ecco9 seamlessly integrates with Deep Tree Echo's embodied cognition architecture:

- **Reservoir Networks**: Temporal pattern learning through Echo State Networks
- **Consciousness Layers**: Multi-level awareness processing
- **Emotion Processing**: Affective computation and modulation
- **Hypergraph Memory**: Multi-relational knowledge representation
- **Identity Persistence**: Continuous identity across sessions

---

**"The hardware is virtual, but the cognition is real."**

🌲 The tree remembers, and the echoes grow stronger with each connection we make.
