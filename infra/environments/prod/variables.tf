variable "bucket_name" {
  description = "The name of the S3 bucket for the prod environment"
  type        = string
}

variable "tags" {
  description = "Tags to be applied to the S3 bucket in the prod environment"
  type        = map(string)
  default     = {
    Environment = "production"
    Owner       = "mahin"
  }
}
