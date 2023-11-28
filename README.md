## How to Start the App

To start the application, follow these steps:

1. Clone the repository:

    ```bash
    git clone <repository-url>
    cd <repository-directory>
    ```

2. Set up the environment variables by creating a `.env` file based on the provided `.env.example`.

3. Build and run the application using the following commands:

```bash
go build ./cmd/user-api
./user-api
```

4. Access the application at the specified port (configurable via environment variables) and explore the defined API routes.

## Clean Build

To perform a clean build, you can use the following commands:

```bash
go clean -i ./...
go mod tidy
go mod download
go build ./...
```

```bash
go run cmd/user-api/main.go
```
This command starts the Go application, and it should be accessible at http://localhost:8080.


## Test
To rerun tests without using the cache, you can use the -count flag with the go test command. The -count flag specifies the number of times to run the tests. By setting it to 1, you ensure that the tests are not cached. Here's the command:
```bash
go test -count=1 ./internal/app/test/...
```

To generate swagger doc
```bash
swag init -g cmd/user-api/main.go
```