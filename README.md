# Web API in Go

This is a simple web API built using Go. It listens on port 8080 and serves requests using the `net/http` package.

## Project Structure

├── main.go └── api └── server.go


- **main.go**: Entry point of the application.
- **api**: Directory containing the API server implementation.

## Getting Started

### Prerequisites

- Go 1.16+ installed on your system. You can download it from [Go's official website](https://golang.org/dl/).

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/your-repo.git
    cd your-repo
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

### Running the Server

To start the server, run:

```bash
go run main.go
