provider "aws" {
  region = "ap-south-1"
  profile = "new"
}

terraform {
  backend "s3" {
    bucket = "kops-s3-bucket-1"
    key    = "kops/ap-south-1/dev/terraform.tfstate"
    region = "ap-south-1"
  }
}