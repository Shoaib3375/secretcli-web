# modules/networking/main.tf
resource "aws_vpc" "kops_vpc" {
  cidr_block = var.vpc_cidr

  tags = {
    Name = "kops-vpc"
  }
}

resource "aws_internet_gateway" "kops_igw" {
  vpc_id = aws_vpc.kops_vpc.id

  tags = {
    Name = "kops-igw"
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.kops_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.kops_igw.id
  }

  tags = {
    Name = "kops-public-route-table"
  }
}

resource "aws_route_table_association" "public_association_1" {
  subnet_id      = aws_subnet.kops_subnet_1.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "public_association_2" {
  subnet_id      = aws_subnet.kops_subnet_2.id
  route_table_id = aws_route_table.public.id
}


resource "aws_subnet" "kops_subnet_1" {
  vpc_id            = aws_vpc.kops_vpc.id
  cidr_block        = var.subnet_cidr_1
  availability_zone = var.availability_zone_1

  tags = {
    Name = "kops-subnet-1"
  }
}

resource "aws_subnet" "kops_subnet_2" {
  vpc_id            = aws_vpc.kops_vpc.id
  cidr_block        = var.subnet_cidr_2
  availability_zone = var.availability_zone_2

  tags = {
    Name = "kops-subnet-2"
  }
}

resource "aws_security_group" "kops_sg" {
  name        = "kops-sg"
  description = "Allow TLS inbound traffic and all outbound traffic"
  vpc_id      = aws_vpc.kops_vpc.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "kops-sg"
  }
}
