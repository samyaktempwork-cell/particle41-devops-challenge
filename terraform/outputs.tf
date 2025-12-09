# terraform/outputs.tf

output "ecr_repository_url" {
  description = "URL of the ECR repository"
  value       = aws_ecr_repository.app_repo.repository_url
}

output "load_balancer_url" {
  description = "URL to access the application"
  value       = "http://${aws_lb.main.dns_name}"
}