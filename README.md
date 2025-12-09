# particle41-devops-challenge
Senior DevOps Challenge submission focusing on Security, Infrastructure-as-Code modularity, and Cost-Efficiency using AWS, Terraform, and Go.

##  Project Overview

This repository contains a solution for the Senior DevOps Challenge. It consists of a lightweight Go application and a modular Terraform infrastructure to deploy it securely on AWS.

### Key Features
* **Microservice:** High-performance Go application returning JSON time/IP data.
* **Containerization:** Optimized ~8MB Docker image (Scratch base) running as a non-root user.
* **Infrastructure:** Production-ready AWS environment (VPC, ALB, ECS Fargate) defined in Terraform.
* **Security:** Least-privilege IAM roles, private subnets for compute, and strict security groups.
* **Automation:** GitHub Actions pipeline for continuous integration and plan validation.

---

##  Repository Structure

```text
.
├── .github/workflows   # CI/CD pipeline definitions
├── app/                # Task 1: Go Application code & Dockerfile
├── scripts/            # Automation scripts (e.g., backend bootstrapping)
├── terraform/          # Task 2: Infrastructure as Code definitions
├── .gitignore          # Security rules for git
└── README.md           # Master documentation (You are here)
```
## Task 1: Application (Local Development)
The application is a simple HTTP server written in Go.

### Prerequisites
* [Docker](https://docs.docker.com/get-started/get-docker/)

### Quick Start
1. Build the image:
    ```bash
    cd app
    docker build -t simpletime .
    ```
2. Run the container:
    ```bash
    docker run -p 8080:8080 simpletime
    ```
3. Test: Open http://localhost:8080 to see the JSON response:

    ```JSON
    {"timestamp": "2023-10-27T10:00:00Z", "ip": "172.17.0.1"}
    ```

Note: The Dockerfile uses a multi-stage build pattern to ensure the final image is extremely small (~8MB) and secure (runs as non-root user appuser).

## Task 2: Infrastructure (AWS Deployment)
The infrastructure is defined using Terraform modules to provision a scalable ECS Fargate environment.

### Prerequisites
* [Terraform](https://developer.hashicorp.com/terraform/downloads) (v1.5.0+)
* [AWS CLI](https://aws.amazon.com/cli/) (v2.0+)

### Deployment Steps
1. **Navigate to the directory:**
    ```Bash
    cd terraform
    ```
2. **Initialize Terraform:**
    ```Bash
    terraform init
    ```
3. R**eview the Plan:**
    ```Bash
    terraform plan
    ```
    This generates an execution plan showing the resources to be created (VPC, ALB, ECS, etc.).

4. **Deploy:**
    ```Bash
    terraform apply
    ```
### Mock Mode (Important)
By default, the `providers.tf` file is configured with **mock credentials** (`skip_credentials_validation = true`). This allows you to generate a valid **Execution Plan** to verify logic without needing an active AWS account or incurring costs.

To deploy to a real account:
1. Open `terraform/providers.tf`.
2. Remove or comment out the `skip_` validation flags.
3. Ensure your terminal is authenticated via `aws configure`.

## Extra Credit: Automation & Bootstrapping
1. **CI/CD Pipeline** (`.github/workflows/terraform.yml`)
    A fully functional GitHub Actions workflow is included. On every push to `main` or feature branches, it automatically:

    **Formats:** Checks Terraform code style (`terraform fmt`).

    **Validates:** Ensures syntax correctness (`terraform validate`).

    **Plans:** Runs a speculative `terraform plan` to catch infrastructure errors before deployment.

2. **Remote Backend Bootstrapping**
    This project solves the "Chicken and Egg" problem of remote state.

    **Bootstrapping:** `terraform/main.tf` includes code to provision the S3 Bucket and DynamoDB table for state locking.

    **Automation:** The `scripts/bootstrap_backend.sh` script demonstrates how to automate the migration from local state to the newly created remote backend.

    **Configuration:** `terraform/backend.conf` and `terraform/backend.tf` are pre-configured to support this switch.
