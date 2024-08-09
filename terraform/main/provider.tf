provider "aws" {
  region  = var.region
  profile = var.profile
}


terraform {
  required_version = ">= 1.2.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  # the backend doesn't accept any var, so its necessary to hard code it here
  backend "s3" {
    bucket         = ""
    key            = ""
    region         = ""
    dynamodb_table = ""
    encrypt        = true
    profile = ""
  }
}