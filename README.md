
# gRPC-Based Marketplace Application

## Project Overview
This project, a part of the *Distributed Systems Concepts & Designs (CSE530)* course, is a marketplace application developed using the[ **Go programming language**](https://go.dev). It leverages [**gRPC**](https://grpc.io), a high-performance, open-source universal RPC framework, to manage procedural calls between the Market (Server) and the Buyers/Sellers (Clients). Since this was a course project with a strict deadline, even though error handling has been implemented, there may be potential bugs in the scripts.

## Setup Instructions

### System Requirements
- A Linux machine with bash shell access

### Installation Steps
1. Clone the [project repository](https://github.com/adityaahuja7/go-grpc-simple-application/tree/master).
2. Execute the `install_script.sh` script to set up the Go environment through the following command `source ./install_script.sh`.
3. Compile and execute the `server.go`, `seller_client.go`, and `buyer_client.go` Go programs using the command `go run <program>.go`.

Please note that these instructions assume familiarity with the Go programming language and gRPC. If you're new to these technologies, you might want to explore some tutorials or documentation first.

