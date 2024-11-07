provider "aws" {
  region = "ap-south-1"
  profile = "dev"
}

terraform {
  backend "s3" {
    bucket = "test-kops-s3-bucket"
    key    = "kops/ap-south-1/dev/terraform.tfstate"
    region = "ap-south-1"
  }
}