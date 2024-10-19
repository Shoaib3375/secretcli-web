# modules/networking/outputs.tf
output "kops_sg_id" {
  value = aws_security_group.kops_sg.id
}

output "kops_subnet_1_id" {
  value = aws_subnet.kops_subnet_1.id
}

output "kops_subnet_2_id" {
  value = aws_subnet.kops_subnet_2.id
}

output "vpc_id" {
  value = aws_vpc.kops_vpc.id
}
