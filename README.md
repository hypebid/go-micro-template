# go-micro-template
A Golang micro-service template. The boilerplate code for all of Hype Bid's micro-services coded in Go.

## Features
- gRPC Server
  - HealthCheck function
- gRPC Inteceptors / middleware
  - Tranaction Id
  - Logrus logger
  - ctx tags
  - Recover from panics
  - Hash validation check
- Gorm integration with Postgres
- .env file support
- Docker file
- Docker-Compose file
- Go modules
- Github Actions
- Prometheus metrics integrated

### Todo
- Go tests
- Move code to go-kit
- Cue integration / validation ?