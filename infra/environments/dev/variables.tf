# infra/environments/dev/variables.tf
variable "profile" {
  description = "The AWS profile to use"
  type        = string
  default     = "dev"
}

variable "region" {
  description = "The AWS region to use"
  type        = string
  default     = "ap-south-1"
}

# s3 
variable "bucket_name" {
  description = "The name of the S3 bucket for the dev environment"
  type        = string
}

variable "tags" {
  description = "Tags to be applied to the S3 bucket in the dev environment"
  type        = map(string)
  default     = {
    Environment = "development"
    Owner       = "mahin"
  }
}

# networking
variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  type        = string
  default     = "172.16.0.0/16"
}

variable "subnet_cidr_1" {
  description = "CIDR block for the first subnet"
  type        = string
  default     = "172.16.10.0/24"
}

variable "subnet_cidr_2" {
  description = "CIDR block for the second subnet"
  type        = string
  default     = "172.16.20.0/24"
}

variable "availability_zone_1" {
  description = "Availability zone for the first subnet"
  type        = string
  default     = "ap-south-1a"
}

variable "availability_zone_2" {
  description = "Availability zone for the second subnet"
  type        = string
  default     = "ap-south-1b"
}

# ec2
variable "ami_id" {
  description = "AMI ID for the EC2 instance"
  type        = string
  default     = "ami-0f5ee92e2d63afc18"  # Amazon Linux 2 AMI in ap-south-1
}

variable "instance_type" {
  description = "Instance type for the EC2 instance"
  type        = string
  default     = "t2.micro"
}

variable "public_key_path" {
  description = "Path to the public SSH key"
  type        = string
  default     = "~/.ssh/id_rsa.pub"
}
