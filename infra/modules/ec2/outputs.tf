# modules/ec2/outputs.tf
output "instance_public_dns" {
  value = aws_instance.web.public_dns
}

output "alb_dns_name" {
  value = aws_lb.kops_alb.dns_name
}
