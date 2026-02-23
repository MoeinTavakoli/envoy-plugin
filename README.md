# Envoy Header Mutation and Go Service Example

This repository demonstrates how to **modify HTTP headers** using **Envoy** as a reverse proxy and **Go** as a backend service that logs and shows the mutated headers.

The **Envoy proxy** acts as an intermediary between the client and the Go service. It uses **Lua scripts** (or WASM if preferred) to mutate HTTP headers, while the **Go service** simply logs and displays the mutated headers for debugging or verification.

---

## Project Structure

- **`show_headers.go`**: The Go service that logs and responds with the `Authorization` header received from the Envoy proxy.
- **`Dockerfile`**: Defines the Docker image for the Go service.
- **`docker-compose.yml`**: Configuration to run the Go service and Envoy together with Docker Compose.
- **`envoy-configs/`**: Contains Envoy's configuration files (`lds.yaml`, `cds.yaml`, `envoy.yaml`).
- **`go.mod`**: Go module file for dependency management.
- **`docker-compose-host.yml`**: Docker Compose file for running the service on the host machine (optional).
- **`README.md`**: This file.

---

## Overview

The goal of this project is to demonstrate how Envoy can be configured to **mutate HTTP headers** before forwarding the request to the backend Go service. In this case, we are particularly focused on the `Authorization` header:

1. **Envoy** intercepts incoming requests.
2. Envoy applies a **Lua filter** to replace the `Authorization` header.
3. The mutated header is forwarded to the **Go service**, which logs it and responds with the header value.

---

## How It Works

1. **Envoy Configuration**:
   - Envoy listens on port 80.
   - The Lua filter within Envoy replaces the `Authorization` header with a static value (`abcd123`).
   
2. **Go Service**:
   - The Go service listens on port `8080`.
   - It logs the `Authorization` header it receives from Envoy.
   - It responds with the `Authorization` header it received.

---

## Deployment

### Prerequisites

Make sure you have the following installed:
- **Docker**: For containerized deployment.
- **Docker Compose**: To manage multi-container applications.
