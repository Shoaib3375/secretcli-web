provider "aws" {
  profile = "prod"
  region  = "us-west-2"
}

module "s3_bucket" {
  source           = "../../modules/s3"
  bucket_name      = var.bucket_name
  tags             = var.tags
}
