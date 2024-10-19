# modules/ec2/variables.tf
variable "ami_id" {
  description = "AMI ID for the EC2 instance"
  type        = string
}

variable "instance_type" {
  description = "Instance type for the EC2 instance"
  type        = string
}

variable "public_key_path" {
  description = "Path to the public SSH key"
  type        = string
}

variable "security_group_id" {
  description = "Security Group ID to attach to the EC2 instance and ALB"
  type        = string
}

variable "subnet_id_1" {
  description = "First Subnet ID for the ALB and EC2 instance"
  type        = string
}

variable "subnet_id_2" {
  description = "Second Subnet ID for the ALB"
  type        = string
}
