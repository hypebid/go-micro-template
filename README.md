# go-micro-template
A Golang micro-service template. The boilerplate code for all of Hype Bid's micro-services coded in Go.

## Features
- gRPC Server
  - HealthCheck function
- gRPC Inteceptors / middleware
  - Logrus logger
  - ctx tags
  - Recover from panics
- Gorm integration with Postgres
- .env file support
- Docker file
- Docker-Compose file
- Go modules
- Github Actions

### Todo
- Custom Middleware
  - Transaction Id
- Go tests
- Cue integration / validation
- Move code to go-kit