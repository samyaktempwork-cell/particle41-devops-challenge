#!/bin/bash
set -e

PROJECT_NAME="simpletime-app"
REGION="us-east-1"

echo "Starting Backend Bootstrap for $PROJECT_NAME..."
echo "Initializing Terraform (Local State)..."

cd ../terraform
terraform init

echo "Provisioning S3 Bucket and DynamoDB Table..."
terraform apply -target=aws_s3_bucket.terraform_state \
                -target=aws_dynamodb_table.terraform_locks \
                -auto-approve

echo "Backend Infrastructure Created!"
echo "Migrating state to S3..."
terraform init -force-copy -backend-config=backend.conf

echo "Bootstrap Complete! State is now stored securely in S3."