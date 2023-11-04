# TorTorCoin System

The TorTorCoin System is a web application that allows users to register and log in using an account number and password. Each new account is credited with five "TorTorCoins". Users can pair up with their friends (1-1) and use TorTorCoins to "force" friends to do tasks. The system maintains a transaction record for all activities.

## Features

- User registration and login
- Initial credit of five TorTorCoins for new users
- Friend pairing and management
- Task requests and TorTorCoin transactions
- Transaction history

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:

```bash
go version go1.18 linux/amd64
```

### Installing

A step-by-step series of examples that tell you how to get a development environment running:

Clone the repository:

```bash
git clone https://github.com/yourusername/TorTorCoinSystem.git
```

Navigate to the project directory:

```bash
cd TorTorCoinSystem
```

Install the dependencies:

```bash
go mod tidy
```

Run the application:

```bash
go run cmd/main.go
```

The server should start and be listening for requests on `http://localhost:8080`.

## Running the tests

Explain how to run the automated tests for this system:

```bash
go test ./...
```

## Deployment

Add additional notes about how to deploy this on a live system.

## Built With

- [Go](https://golang.org/) - The Go programming language
- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework used (if applicable)
- [GORM](https://gorm.io/) - ORM library for Go (if applicable)

## Authors

- **Pluto Zhang** - *Initial work* - [plutozhang99](https://github.com/plutozhang99)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
