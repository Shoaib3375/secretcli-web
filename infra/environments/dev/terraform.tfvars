# infra/environments/dev/terraform.tfvars
profile    = "dev"
region     = "ap-south-1"

# s3 
bucket_name   = "my-dev-s3-bucket-1996"
tags = {
  Environment = "development"
  Owner       = "mahin"
  Project     = "secretcli-web"
}

# networking
vpc_cidr          = "172.16.0.0/16"

# ec2 
ami_id            = "ami-0dee22c13ea7a9a67"
instance_type     = "t2.micro"
public_key_path   = "~/.ssh/id_rsa.pub"
