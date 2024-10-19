# infra/environments/dev/main.tf
provider "aws" {
  profile = var.profile
  region  = var.region
}

module "s3_bucket" {
  source      = "../../modules/s3"
  bucket_name = var.bucket_name
  tags        = var.tags
}

module "networking" {
  source              = "../../modules/networking"
  vpc_cidr            = var.vpc_cidr
  subnet_cidr_1       = var.subnet_cidr_1
  subnet_cidr_2       = var.subnet_cidr_2
  availability_zone_1 = var.availability_zone_1
  availability_zone_2 = var.availability_zone_2
}

module "ec2" {
  source           = "../../modules/ec2"
  ami_id           = var.ami_id
  instance_type    = var.instance_type
  public_key_path  = var.public_key_path
  security_group_id = module.networking.kops_sg_id
  subnet_id_1      = module.networking.kops_subnet_1_id
  subnet_id_2      = module.networking.kops_subnet_2_id
}
