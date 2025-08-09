# Go Load Balancer

This project provides a **simple, round-robin load balancer** written in Go.  
It is designed to distribute incoming requests across a pool of backend servers and serves as a great example for learning how to build a **basic load balancer, reverse proxy, and concurrent HTTP servers** in Go.

## Features

- **Round-Robin Load Balancing** – Distributes requests sequentially to each backend server in a cyclical manner.
- **Reverse Proxy** – Forwards incoming requests from the load balancer to the appropriate backend server.
- **Concurrent Backend Servers** – Starts multiple backend servers on different ports to demonstrate the load balancing in action.
- **Simple HTTP Server** – Each backend server responds with a message indicating which server handled the request.

## Project Structure

The project is organized into three files:

- **`main.go`** – Entry point of the application. Initializes the backend servers, configures the load balancer, and starts the main load balancer server.
- **`loadbalancer.go`** – Defines the `LoadBalancer` and `Backend` structs, including the round-robin selection logic.
- **`backend.go`** – Contains the `newBackendServer` function, which creates a simple HTTP server that responds with its port number.

## How to Run

### Prerequisites
- Go installed on your machine ([Download Go](https://go.dev/dl/)).

### Instructions

1. **Save the files**  
   Create three files in the same directory:  
   - `main.go`
   - `loadbalancer.go`
   - `backend.go`  
   and paste the provided code into them.

2. **Run the application**  
   Open your terminal, navigate to the directory, and run:

   ```bash
   go run main.go loadbalancer.go backend.go
   ```

Expected Result:
  ```bash
  2025/08/08 12:49:11 Starting load balancer on port 8080...
  2025/08/08 12:49:11 Starting backend server on port 8081...
  2025/08/08 12:49:11 Starting backend server on port 8082...
  2025/08/08 12:49:11 Starting backend server on port 8083...
  ```

## Testing the Load Balancer

With the application running, you can send requests to the load balancer's address: http://localhost:8080

The load balancer will forward these requests to the backend servers in a **round-robin** manner.

You can use a **web browser**, **curl**, or any other HTTP client to test this.

### Using `curl`

Open a new terminal window and run the following command multiple times:

```bash
curl http://localhost:8080
```

**First Run:**
```bash
Hello from backend server on port 8082!
```

**Second Run:**
```bash
Hello from backend server on port 8083!
```

**Third Run:**
```bash
Hello from backend server on port 8081!
```
