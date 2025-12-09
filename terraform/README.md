#  Infrastructure as Code (Terraform)

This directory contains the Terraform configuration to provision a secure, scalable AWS environment for the **SimpleTime** application.

The infrastructure is built using modular best practices and includes networking, compute (ECS Fargate), load balancing, and container storage.

##  Architecture Overview

The following resources are provisioned:
* **VPC Module:** Production-ready network with 2 Public and 2 Private subnets across multiple AZs.
* **Networking:** Internet Gateway (IGW) for public access and NAT Gateway for secure private egress.
* **Compute:** Amazon ECS Cluster running Fargate tasks (Serverless containers).
* **Load Balancing:** Application Load Balancer (ALB) acting as the ingress point.
* **Security:**
    * `lb_sg`: Allows HTTP (80) traffic from the internet.
    * `ecs_sg`: Allows traffic **only** from the Load Balancer (Zero Trust).
* **Storage:** ECR Repository for Docker images.
* **IAM:** Least-privilege execution roles for ECS tasks.

---

##  Usage Guide

### 1. Prerequisites
Ensure you have the following installed locally:
* [Terraform](https://developer.hashicorp.com/terraform/downloads) (v1.5.0+)
* [AWS CLI](https://aws.amazon.com/cli/) (v2.0+)

### 2. Configuration
The project uses `terraform.tfvars` to manage environment-specific variables.

| Variable | Description | Default |
| :--- | :--- | :--- |
| `aws_region` | Target AWS Region | `us-east-1` |
| `project_name` | Resource tagging prefix | `simpletime-app` |
| `vpc_cidr` | Network IP range | `10.0.0.0/16` |
| `app_port` | Application container port | `8080` |

### 3. Running Terraform (Mock Mode)
**Default Behavior:** To facilitate verification without incurring cloud costs, the `providers.tf` is currently configured to **skip authentication**. This allows you to generate a valid **Execution Plan** to verify logic and dependencies.

1.  **Initialize:**
    ```bash
    terraform init
    ```

2.  **Generate Plan:**
    ```bash
    terraform plan -out=tfplan
    ```

3.  **Review Resources:**
    ```bash
    terraform show tfplan
    ```

### 4. Deploying to Real AWS
To deploy this to a live AWS account:

1.  **Authenticate:**
    ```bash
    aws configure
    ```
2.  **Update Provider:** Open `providers.tf` and remove/comment out the `skip_` validation flags.
3.  **Deploy:**
    ```bash
    terraform apply tfplan
    ```

---

##  Remote Backend (State Management)
This project follows the **Bootstrapping Pattern** for state management.

1.  **Bootstrapping:** The `main.tf` file includes resources (`aws_s3_bucket`, `aws_dynamodb_table`) to provision the backend infrastructure automatically.
2.  **Configuration:** The `backend.conf` and `backend.tf` files are prepared to switch from local state to remote state immediately after the bucket is created.
3.  **Automation:** See `../scripts/bootstrap_backend.sh` for the automated migration script.

---

##  Outputs
After a successful apply, Terraform will output:
* `ecr_repository_url`: The address to push your Docker image.
* `load_balancer_url`: The public URL to access the running application.