# modules/networking/variables.tf
variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  type        = string
}

variable "subnet_cidr_1" {
  description = "CIDR block for the first subnet"
  type        = string
}

variable "subnet_cidr_2" {
  description = "CIDR block for the second subnet"
  type        = string
}

variable "availability_zone_1" {
  description = "Availability zone for the first subnet"
  type        = string
}

variable "availability_zone_2" {
  description = "Availability zone for the second subnet"
  type        = string
}
