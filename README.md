## SniffOps
A high-performance agent for streaming OS metrics to a gRPC server and performing time series forecasting.

### Setup Client
```sh
  git clone https://github.com/Sai-Adarsh/sniffops
  cd sniffops/client
  make install
```

### Setup Server
```bash
  cd sniffops/server
  make install
```

### Logging Client
```bash
  cd sniffops/client
  make run
```

### Stop Logging Server
```bash
  cd sniffops/server
  make stop
```

### Stream Logs from Client to Server
```bash
  cd sniffops/client
  make client-stream
```

### Stream Logs from Server to Client
```bash
  cd sniffops/server
  make server-stream
```

<p align="center">
  <img src="client/src/media/screenshot.png" height="20%">
</p>

### Metrics
  - [x] CPU
  - [x] Memory
  - [ ] Loadavg
  - [ ] Uptime
  - [ ] Network
  - [ ] Disk I/O

### Deliverables
  - [x] Logging & monitoring
  - [x] Stream logs to server
  - [ ] Real-time forecasting

### Binaries
  - [x] Linux
  - [ ] Windows
