# Go-gRPC Example

This repository contains a basic example demonstrating the use of gRPC with Go. It showcases a simple service definition, implementation, and client interaction.

## Overview

This example implements a simple "Greeter" service where a client can send a name to the server, and the server will respond with a personalized greeting.

**Key Components:**

* **`proto/`:** Contains the Protocol Buffer (`.proto`) definition for the Greeter service.
    * `proto/greet.proto`: Defines the `Greeter` service with a single `SayHello` RPC method.
* **`server/`:** Contains the Go code for the gRPC server.
    * `server/main.go`: Implements the `Greeter` service defined in the `.proto` file.
* **`client/`:** Contains the Go code for the gRPC client.
    * `client/main.go`: Interacts with the gRPC server by sending a greeting request.

## Prerequisites

* **Go:** Version 1.16 or higher is recommended. You can download and install Go from [https://go.dev/dl/](https://go.dev/dl/).
* **Protocol Buffer Compiler (protoc):** Required to compile the `.proto` file into Go code. You can find installation instructions here: [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/).
* **Go gRPC and Protocol Buffer Go plugins:** These are necessary for generating Go gRPC code from the `.proto` file. Install them using:
    ```bash
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    ```

## Getting Started

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/maruki00/Go-grpc.git](https://github.com/maruki00/Go-grpc.git)
    cd Go-grpc
    ```

2.  **Generate Go gRPC code:**
    Navigate to the root of the repository and run the following command to compile the `greet.proto` file:
    ```bash
    protoc --go_out=. --go-grpc_out=. proto/greet.proto
    ```
    This command will generate the `greet.pb.go` and `greet_grpc.pb.go` files in the `proto` directory.

3.  **Build and Run the Server:**
    Navigate to the `server` directory:
    ```bash
    cd server
    go build -o server .
    ./server
    ```
    The server will start and listen on port `50051`.

4.  **Build and Run the Client:**
    Open a new terminal window and navigate to the `client` directory:
    ```bash
    cd ../client
    go build -o client .
    ./client <your_name>
    ```
    Replace `<your_name>` with the name you want to send to the server (e.g., `./client Alice`). The client will connect to the server and print the greeting received.

## Project Structure
