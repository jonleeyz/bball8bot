terraform {
  required_version = ">=1.7.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.35.0"
    }
  }

  backend "s3" {
  }
}

provider "aws" {
  region = "ap-southeast-1"
}
