# modules/ec2/main.tf
resource "aws_key_pair" "kops_key_pair" {
  key_name   = "kops-key-pair"
  public_key = file(var.public_key_path)
}

resource "aws_lb" "kops_alb" {
  name               = "kops-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [var.security_group_id]
  subnets            = [var.subnet_id_1, var.subnet_id_2]
  enable_deletion_protection = false

  tags = {
    Environment = "dev"
  }
}

resource "aws_instance" "web" {
  ami           = var.ami_id
  instance_type = var.instance_type
  key_name      = aws_key_pair.kops_key_pair.key_name
  vpc_security_group_ids = [var.security_group_id]
  subnet_id     = var.subnet_id_1

  tags = {
    Name = "kops-ec2"
  }
}
