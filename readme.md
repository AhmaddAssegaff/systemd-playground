# Systemd Playground
A simple playground repository for learning how to run and manage Linux services using **systemd**.
The goal of this project is not to build a production-ready application, but to understand the fundamentals of:

* Creating a systemd service unit
* Managing services with `systemctl`
* Viewing logs with `journalctl`
* Organizing deployment-related files inside a repository

## Project Structure
```text
.
├── bin/
│   └── hello-world
├── etc/
│   └── systemd/
│       └── system/
│           └── hello-world.service
├── go.mod
├── main.go
└── readme.md
```

### Directory Overview
| Path                  | Description                                      |
| --------------------- | ------------------------------------------------ |
| `main.go`             | Sample Go application                            |
| `bin/`                | Compiled binaries                                |
| `etc/systemd/system/` | Systemd unit files managed inside the repository |
| `readme.md`           | Project documentation                            |

## Application
The sample application:
* Starts an HTTP server
* Returns `Hello World`
* Writes a log message every 5 seconds

This application exists only to provide a process that can be managed by systemd.

## How to create a unit systemd service

### Build the application
Compile the application:
```bash
go build -o bin/hello-world .
```

### Check the application working correctly
Run it manually:
```bash
./bin/hello-world
```

### Copy a service configuration into `/etc/systemd/system/`
```bash
sudo cp etc/systemd/system/hello-world.service /etc/systemd/system/
```

### Copy binaries file to `/usr/local/bin/hello-world`
```bash
sudo cp bin/hello-world /usr/local/bin/hello-world
```

### Reload daemon
```bash
sudo systemctl daemon-reload
```

### Start the service
```bash
sudo systemctl start hello-world
```

### Check the service status
```bash
sudo systemctl status hello-world
```

## Learning Objectives
After completing this playground, you should understand:

* What a systemd service is
* How Linux services are started and supervised
* How automatic restart policies work
* How to inspect service logs

## Next Steps
Possible improvements:

* Environment files (`EnvironmentFile`)
* Dedicated Linux user
* Graceful shutdown handling
* Health check endpoint
* Multiple services (API + Worker)
* Systemd timers
* Service sandboxing and security settings
