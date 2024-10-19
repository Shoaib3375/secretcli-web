# modules/s3/main.tf 
resource "aws_s3_bucket" "kops_s3" {
  bucket = var.bucket_name
  tags   = var.tags
}
