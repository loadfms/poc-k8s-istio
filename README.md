# POC K8s

This repository demonstrates the usage of Kubernetes and Istio for validating JWTs on incoming traffic using Istio ingress.

## Getting Started

To get started with this POC, follow the instructions below.

### Prerequisites

Make sure you have the following dependencies installed:

- Minikube
- kubectl
- Istio

### Installation

1. Start minikube:

   ```shell
    minikube start --cpus 6 --memory 8192
   ```

2. Apply kubernates infra folder:

   ```
   kubectl apply -f infra/
   ```

3. Execute a request with any invalid token, adding to headers:

   ```
   Authorization: Bearer bla-bla-bla
   ```

3. Execute a request with a valid token, adding to headers:

   ```
   Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMyMzMyMDEsImlhdCI6MTY5MzIzMjMwMSwiaXNzIjoiaHR0cHM6Ly9sb2FkZm1zLmxvY2FsIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMTIzNDU2Nzg5MCJ9.vYbI9YD3_v5f7cM_rTwl9ZcOQ9deRQMejJhvl0o6xTk
   ```

## Usage

Once the application is running, you can access it through the Istio ingress. All incoming traffic will be validated using JWTs.

## License

This project is licensed under the [MIT License](LICENSE).
