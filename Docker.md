# Go API Docker Compose Setup

This Docker Compose configuration is designed to facilitate the deployment of a Go API application, along with MongoDB and MongoDB Express for data storage and management. The setup includes a multi-stage Dockerfile for the Go API to optimize the size of the final Docker image.

## Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Usage](#usage)
- [Configuration](#configuration)
- [Dockerfile Explanation](#dockerfile-explanation)

## Overview

The Docker Compose file (`compose.yml`) defines three services:

- **go-api:** The main Go API service.
- **go-mongodb:** MongoDB service for data storage.
- **mongo-express:** MongoDB Express service for easy database management.

The Go API is built using a multi-stage Dockerfile, resulting in a smaller final image. The MongoDB service persists data using a volume named `mongodb_data`, and MongoDB Express provides a web-based interface for database administration.

# Prerequisites

- Docker installed on your machine.
- `docker-compose` installed on your machine.

# Usage

1. Clone the repository to your local machine.
2. Create a `.env` file based on the provided `.env.example`.
3. Run the following command in the project directory:

    ```bash
    docker-compose up -d
    ```

   The services should now be running, and you can access the Go API at `http://localhost:${GO_API_PORT}` and MongoDB Express at `http://localhost:${MONGODB_EXPRESS_PORT}`.

4. To stop the services, run:

    ```bash
    docker-compose down
    ```

# Configuration

## Environment Variables

Make sure to configure the `.env` file with the appropriate values. Key environment variables include:

- **GO_API_PORT:** Port on which the Go API will be accessible.
- **MONGODB_PORT:** Port for the MongoDB service.
- **MONGODB_EXPRESS_PORT:** Port for MongoDB Express.



# Dockerfile Explanation

The Dockerfile in the `compose/go` directory utilizes multi-stage building. Here's a breakdown:

## Builder Stage (`golang:1.21-alpine as builder`):

- Sets up the working directory.
- Copies `go.mod` and `go.sum` for dependency management.
- Downloads dependencies using `go mod download`.
- Copies the entire project and builds the Go application.

## Final Stage (`alpine:latest`):

- Sets up the working directory.
- Copies the compiled binary and necessary files from the builder stage.
- Exposes port 8080.
- Defines the command to run the executable.


# Running Docker Compose for a Specific Service

To start your Docker Compose setup for a specific service, follow these steps:

1. Open a terminal.

2. Run the following command to start Docker Compose with a specific service, rebuild the images, and force recreation:

```bash
docker compose -f compose.yml up -d --build [service-name] --force-recreate
```

```bash
docker compose -f compose.yml up -d --build go-api --force-recreate
```

```bash
docker compose -f compose.yml up -d --build mongodb --force-recreate
```

