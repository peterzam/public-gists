terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
}

resource "aws_codecommit_repository" "repo" {
  repository_name = "test"
  description     = ""
  tags = {
    Name = "value"
  }
}

output "clone_url_ssh" {
  value       = aws_codecommit_repository.repo.clone_url_ssh
  description = "Clone URL - SSH"
}
