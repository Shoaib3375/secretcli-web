locals {
  cluster_name                 = "kops-k8s.prodcrashed.live"
  master_autoscaling_group_ids = [aws_autoscaling_group.control-plane-ap-south-1a-masters-kops-k8s-prodcrashed-live.id]
  master_security_group_ids    = [aws_security_group.masters-kops-k8s-prodcrashed-live.id]
  masters_role_arn             = aws_iam_role.masters-kops-k8s-prodcrashed-live.arn
  masters_role_name            = aws_iam_role.masters-kops-k8s-prodcrashed-live.name
  node_autoscaling_group_ids   = [aws_autoscaling_group.nodes-ap-south-1a-kops-k8s-prodcrashed-live.id]
  node_security_group_ids      = [aws_security_group.nodes-kops-k8s-prodcrashed-live.id]
  node_subnet_ids              = [aws_subnet.ap-south-1a-kops-k8s-prodcrashed-live.id]
  nodes_role_arn               = aws_iam_role.nodes-kops-k8s-prodcrashed-live.arn
  nodes_role_name              = aws_iam_role.nodes-kops-k8s-prodcrashed-live.name
  region                       = "ap-south-1"
  route_table_public_id        = aws_route_table.kops-k8s-prodcrashed-live.id
  subnet_ap-south-1a_id        = aws_subnet.ap-south-1a-kops-k8s-prodcrashed-live.id
  vpc_cidr_block               = aws_vpc.kops-k8s-prodcrashed-live.cidr_block
  vpc_id                       = aws_vpc.kops-k8s-prodcrashed-live.id
  vpc_ipv6_cidr_block          = aws_vpc.kops-k8s-prodcrashed-live.ipv6_cidr_block
  vpc_ipv6_cidr_length         = local.vpc_ipv6_cidr_block == "" ? null : tonumber(regex(".*/(\\d+)", local.vpc_ipv6_cidr_block)[0])
}

output "cluster_name" {
  value = "kops-k8s.prodcrashed.live"
}

output "master_autoscaling_group_ids" {
  value = [aws_autoscaling_group.control-plane-ap-south-1a-masters-kops-k8s-prodcrashed-live.id]
}

output "master_security_group_ids" {
  value = [aws_security_group.masters-kops-k8s-prodcrashed-live.id]
}

output "masters_role_arn" {
  value = aws_iam_role.masters-kops-k8s-prodcrashed-live.arn
}

output "masters_role_name" {
  value = aws_iam_role.masters-kops-k8s-prodcrashed-live.name
}

output "node_autoscaling_group_ids" {
  value = [aws_autoscaling_group.nodes-ap-south-1a-kops-k8s-prodcrashed-live.id]
}

output "node_security_group_ids" {
  value = [aws_security_group.nodes-kops-k8s-prodcrashed-live.id]
}

output "node_subnet_ids" {
  value = [aws_subnet.ap-south-1a-kops-k8s-prodcrashed-live.id]
}

output "nodes_role_arn" {
  value = aws_iam_role.nodes-kops-k8s-prodcrashed-live.arn
}

output "nodes_role_name" {
  value = aws_iam_role.nodes-kops-k8s-prodcrashed-live.name
}

output "region" {
  value = "ap-south-1"
}

output "route_table_public_id" {
  value = aws_route_table.kops-k8s-prodcrashed-live.id
}

output "subnet_ap-south-1a_id" {
  value = aws_subnet.ap-south-1a-kops-k8s-prodcrashed-live.id
}

output "vpc_cidr_block" {
  value = aws_vpc.kops-k8s-prodcrashed-live.cidr_block
}

output "vpc_id" {
  value = aws_vpc.kops-k8s-prodcrashed-live.id
}

output "vpc_ipv6_cidr_block" {
  value = aws_vpc.kops-k8s-prodcrashed-live.ipv6_cidr_block
}

output "vpc_ipv6_cidr_length" {
  value = local.vpc_ipv6_cidr_block == "" ? null : tonumber(regex(".*/(\\d+)", local.vpc_ipv6_cidr_block)[0])
}

provider "aws" {
  alias  = "files"
  region = "ap-south-1"
}

resource "aws_autoscaling_group" "control-plane-ap-south-1a-masters-kops-k8s-prodcrashed-live" {
  enabled_metrics = ["GroupDesiredCapacity", "GroupInServiceInstances", "GroupMaxSize", "GroupMinSize", "GroupPendingInstances", "GroupStandbyInstances", "GroupTerminatingInstances", "GroupTotalInstances"]
  launch_template {
    id      = aws_launch_template.control-plane-ap-south-1a-masters-kops-k8s-prodcrashed-live.id
    version = aws_launch_template.control-plane-ap-south-1a-masters-kops-k8s-prodcrashed-live.latest_version
  }
  max_instance_lifetime = 0
  max_size              = 1
  metrics_granularity   = "1Minute"
  min_size              = 1
  name                  = "control-plane-ap-south-1a.masters.kops-k8s.prodcrashed.live"
  protect_from_scale_in = false
  tag {
    key                 = "KubernetesCluster"
    propagate_at_launch = true
    value               = "kops-k8s.prodcrashed.live"
  }
  tag {
    key                 = "Name"
    propagate_at_launch = true
    value               = "control-plane-ap-south-1a.masters.kops-k8s.prodcrashed.live"
  }
  tag {
    key                 = "aws-node-termination-handler/managed"
    propagate_at_launch = true
    value               = ""
  }
  tag {
    key                 = "k8s.io/cluster-autoscaler/node-template/label/kops.k8s.io/kops-controller-pki"
    propagate_at_launch = true
    value               = ""
  }
  tag {
    key                 = "k8s.io/cluster-autoscaler/node-template/label/node-role.kubernetes.io/control-plane"
    propagate_at_launch = true
    value               = ""
  }
  tag {
    key                 = "k8s.io/cluster-autoscaler/node-template/label/node.kubernetes.io/exclude-from-external-load-balancers"
    propagate_at_launch = true
    value               = ""
  }
  tag {
    key                 = "k8s.io/role/control-plane"
    propagate_at_launch = true
    value               = "1"
  }
  tag {
    key                 = "k8s.io/role/master"
    propagate_at_launch = true
    value               = "1"
  }
  tag {
    key                 = "kops.k8s.io/instancegroup"
    propagate_at_launch = true
    value               = "control-plane-ap-south-1a"
  }
  tag {
    key                 = "kubernetes.io/cluster/kops-k8s.prodcrashed.live"
    propagate_at_launch = true
    value               = "owned"
  }
  target_group_arns   = [aws_lb_target_group.kops-controller-kops-k8s--j959e1.id, aws_lb_target_group.tcp-kops-k8s-prodcrashed--i61hbp.id]
  vpc_zone_identifier = [aws_subnet.ap-south-1a-kops-k8s-prodcrashed-live.id]
}

resource "aws_autoscaling_group" "nodes-ap-south-1a-kops-k8s-prodcrashed-live" {
  enabled_metrics = ["GroupDesiredCapacity", "GroupInServiceInstances", "GroupMaxSize", "GroupMinSize", "GroupPendingInstances", "GroupStandbyInstances", "GroupTerminatingInstances", "GroupTotalInstances"]
  launch_template {
    id      = aws_launch_template.nodes-ap-south-1a-kops-k8s-prodcrashed-live.id
    version = aws_launch_template.nodes-ap-south-1a-kops-k8s-prodcrashed-live.latest_version
  }
  max_instance_lifetime = 0
  max_size              = 3
  metrics_granularity   = "1Minute"
  min_size              = 3
  name                  = "nodes-ap-south-1a.kops-k8s.prodcrashed.live"
  protect_from_scale_in = false
  tag {
    key                 = "KubernetesCluster"
    propagate_at_launch = true
    value               = "kops-k8s.prodcrashed.live"
  }
  tag {
    key                 = "Name"
    propagate_at_launch = true
    value               = "nodes-ap-south-1a.kops-k8s.prodcrashed.live"
  }
  tag {
    key                 = "aws-node-termination-handler/managed"
    propagate_at_launch = true
    value               = ""
  }
  tag {
    key                 = "k8s.io/cluster-autoscaler/node-template/label/node-role.kubernetes.io/node"
    propagate_at_launch = true
    value               = ""
  }
  tag {
    key                 = "k8s.io/role/node"
    propagate_at_launch = true
    value               = "1"
  }
  tag {
    key                 = "kops.k8s.io/instancegroup"
    propagate_at_launch = true
    value               = "nodes-ap-south-1a"
  }
  tag {
    key                 = "kubernetes.io/cluster/kops-k8s.prodcrashed.live"
    propagate_at_launch = true
    value               = "owned"
  }
  vpc_zone_identifier = [aws_subnet.ap-south-1a-kops-k8s-prodcrashed-live.id]
}

resource "aws_autoscaling_lifecycle_hook" "control-plane-ap-south-1a-NTHLifecycleHook" {
  autoscaling_group_name = aws_autoscaling_group.control-plane-ap-south-1a-masters-kops-k8s-prodcrashed-live.id
  default_result         = "CONTINUE"
  heartbeat_timeout      = 300
  lifecycle_transition   = "autoscaling:EC2_INSTANCE_TERMINATING"
  name                   = "control-plane-ap-south-1a-NTHLifecycleHook"
}

resource "aws_autoscaling_lifecycle_hook" "nodes-ap-south-1a-NTHLifecycleHook" {
  autoscaling_group_name = aws_autoscaling_group.nodes-ap-south-1a-kops-k8s-prodcrashed-live.id
  default_result         = "CONTINUE"
  heartbeat_timeout      = 300
  lifecycle_transition   = "autoscaling:EC2_INSTANCE_TERMINATING"
  name                   = "nodes-ap-south-1a-NTHLifecycleHook"
}

resource "aws_cloudwatch_event_rule" "kops-k8s-prodcrashed-live-ASGLifecycle" {
  event_pattern = file("${path.module}/data/aws_cloudwatch_event_rule_kops-k8s.prodcrashed.live-ASGLifecycle_event_pattern")
  name          = "kops-k8s.prodcrashed.live-ASGLifecycle"
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "kops-k8s.prodcrashed.live-ASGLifecycle"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_cloudwatch_event_rule" "kops-k8s-prodcrashed-live-InstanceScheduledChange" {
  event_pattern = file("${path.module}/data/aws_cloudwatch_event_rule_kops-k8s.prodcrashed.live-InstanceScheduledChange_event_pattern")
  name          = "kops-k8s.prodcrashed.live-InstanceScheduledChange"
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "kops-k8s.prodcrashed.live-InstanceScheduledChange"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_cloudwatch_event_rule" "kops-k8s-prodcrashed-live-InstanceStateChange" {
  event_pattern = file("${path.module}/data/aws_cloudwatch_event_rule_kops-k8s.prodcrashed.live-InstanceStateChange_event_pattern")
  name          = "kops-k8s.prodcrashed.live-InstanceStateChange"
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "kops-k8s.prodcrashed.live-InstanceStateChange"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_cloudwatch_event_rule" "kops-k8s-prodcrashed-live-SpotInterruption" {
  event_pattern = file("${path.module}/data/aws_cloudwatch_event_rule_kops-k8s.prodcrashed.live-SpotInterruption_event_pattern")
  name          = "kops-k8s.prodcrashed.live-SpotInterruption"
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "kops-k8s.prodcrashed.live-SpotInterruption"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_cloudwatch_event_target" "kops-k8s-prodcrashed-live-ASGLifecycle-Target" {
  arn  = aws_sqs_queue.kops-k8s-prodcrashed-live-nth.arn
  rule = aws_cloudwatch_event_rule.kops-k8s-prodcrashed-live-ASGLifecycle.id
}

resource "aws_cloudwatch_event_target" "kops-k8s-prodcrashed-live-InstanceScheduledChange-Target" {
  arn  = aws_sqs_queue.kops-k8s-prodcrashed-live-nth.arn
  rule = aws_cloudwatch_event_rule.kops-k8s-prodcrashed-live-InstanceScheduledChange.id
}

resource "aws_cloudwatch_event_target" "kops-k8s-prodcrashed-live-InstanceStateChange-Target" {
  arn  = aws_sqs_queue.kops-k8s-prodcrashed-live-nth.arn
  rule = aws_cloudwatch_event_rule.kops-k8s-prodcrashed-live-InstanceStateChange.id
}

resource "aws_cloudwatch_event_target" "kops-k8s-prodcrashed-live-SpotInterruption-Target" {
  arn  = aws_sqs_queue.kops-k8s-prodcrashed-live-nth.arn
  rule = aws_cloudwatch_event_rule.kops-k8s-prodcrashed-live-SpotInterruption.id
}

resource "aws_ebs_volume" "a-etcd-events-kops-k8s-prodcrashed-live" {
  availability_zone = "ap-south-1a"
  encrypted         = true
  iops              = 3000
  size              = 20
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "a.etcd-events.kops-k8s.prodcrashed.live"
    "k8s.io/etcd/events"                              = "a/a"
    "k8s.io/role/control-plane"                       = "1"
    "k8s.io/role/master"                              = "1"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
  throughput = 125
  type       = "gp3"
}

resource "aws_ebs_volume" "a-etcd-main-kops-k8s-prodcrashed-live" {
  availability_zone = "ap-south-1a"
  encrypted         = true
  iops              = 3000
  size              = 20
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "a.etcd-main.kops-k8s.prodcrashed.live"
    "k8s.io/etcd/main"                                = "a/a"
    "k8s.io/role/control-plane"                       = "1"
    "k8s.io/role/master"                              = "1"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
  throughput = 125
  type       = "gp3"
}

resource "aws_iam_instance_profile" "masters-kops-k8s-prodcrashed-live" {
  name = "masters.kops-k8s.prodcrashed.live"
  role = aws_iam_role.masters-kops-k8s-prodcrashed-live.name
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "masters.kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_iam_instance_profile" "nodes-kops-k8s-prodcrashed-live" {
  name = "nodes.kops-k8s.prodcrashed.live"
  role = aws_iam_role.nodes-kops-k8s-prodcrashed-live.name
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "nodes.kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_iam_role" "masters-kops-k8s-prodcrashed-live" {
  assume_role_policy = file("${path.module}/data/aws_iam_role_masters.kops-k8s.prodcrashed.live_policy")
  name               = "masters.kops-k8s.prodcrashed.live"
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "masters.kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_iam_role" "nodes-kops-k8s-prodcrashed-live" {
  assume_role_policy = file("${path.module}/data/aws_iam_role_nodes.kops-k8s.prodcrashed.live_policy")
  name               = "nodes.kops-k8s.prodcrashed.live"
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "nodes.kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_iam_role_policy" "masters-kops-k8s-prodcrashed-live" {
  name   = "masters.kops-k8s.prodcrashed.live"
  policy = file("${path.module}/data/aws_iam_role_policy_masters.kops-k8s.prodcrashed.live_policy")
  role   = aws_iam_role.masters-kops-k8s-prodcrashed-live.name
}

resource "aws_iam_role_policy" "nodes-kops-k8s-prodcrashed-live" {
  name   = "nodes.kops-k8s.prodcrashed.live"
  policy = file("${path.module}/data/aws_iam_role_policy_nodes.kops-k8s.prodcrashed.live_policy")
  role   = aws_iam_role.nodes-kops-k8s-prodcrashed-live.name
}

resource "aws_internet_gateway" "kops-k8s-prodcrashed-live" {
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
  vpc_id = aws_vpc.kops-k8s-prodcrashed-live.id
}

resource "aws_key_pair" "kubernetes-kops-k8s-prodcrashed-live-006765ba239e1535853258a6d4c7229f" {
  key_name   = "kubernetes.kops-k8s.prodcrashed.live-00:67:65:ba:23:9e:15:35:85:32:58:a6:d4:c7:22:9f"
  public_key = file("${path.module}/data/aws_key_pair_kubernetes.kops-k8s.prodcrashed.live-006765ba239e1535853258a6d4c7229f_public_key")
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_launch_template" "control-plane-ap-south-1a-masters-kops-k8s-prodcrashed-live" {
  block_device_mappings {
    device_name = "/dev/sda1"
    ebs {
      delete_on_termination = true
      encrypted             = true
      iops                  = 3000
      throughput            = 125
      volume_size           = 64
      volume_type           = "gp3"
    }
  }
  iam_instance_profile {
    name = aws_iam_instance_profile.masters-kops-k8s-prodcrashed-live.id
  }
  image_id      = "ami-0554f7fb41d511dd0"
  instance_type = "t3.micro"
  key_name      = aws_key_pair.kubernetes-kops-k8s-prodcrashed-live-006765ba239e1535853258a6d4c7229f.id
  lifecycle {
    create_before_destroy = true
  }
  metadata_options {
    http_endpoint               = "enabled"
    http_protocol_ipv6          = "disabled"
    http_put_response_hop_limit = 1
    http_tokens                 = "required"
  }
  monitoring {
    enabled = false
  }
  name = "control-plane-ap-south-1a.masters.kops-k8s.prodcrashed.live"
  network_interfaces {
    associate_public_ip_address = true
    delete_on_termination       = true
    ipv6_address_count          = 0
    security_groups             = [aws_security_group.masters-kops-k8s-prodcrashed-live.id]
  }
  tag_specifications {
    resource_type = "instance"
    tags = {
      "KubernetesCluster"                                                                                     = "kops-k8s.prodcrashed.live"
      "Name"                                                                                                  = "control-plane-ap-south-1a.masters.kops-k8s.prodcrashed.live"
      "aws-node-termination-handler/managed"                                                                  = ""
      "k8s.io/cluster-autoscaler/node-template/label/kops.k8s.io/kops-controller-pki"                         = ""
      "k8s.io/cluster-autoscaler/node-template/label/node-role.kubernetes.io/control-plane"                   = ""
      "k8s.io/cluster-autoscaler/node-template/label/node.kubernetes.io/exclude-from-external-load-balancers" = ""
      "k8s.io/role/control-plane"                                                                             = "1"
      "k8s.io/role/master"                                                                                    = "1"
      "kops.k8s.io/instancegroup"                                                                             = "control-plane-ap-south-1a"
      "kubernetes.io/cluster/kops-k8s.prodcrashed.live"                                                       = "owned"
    }
  }
  tag_specifications {
    resource_type = "volume"
    tags = {
      "KubernetesCluster"                                                                                     = "kops-k8s.prodcrashed.live"
      "Name"                                                                                                  = "control-plane-ap-south-1a.masters.kops-k8s.prodcrashed.live"
      "aws-node-termination-handler/managed"                                                                  = ""
      "k8s.io/cluster-autoscaler/node-template/label/kops.k8s.io/kops-controller-pki"                         = ""
      "k8s.io/cluster-autoscaler/node-template/label/node-role.kubernetes.io/control-plane"                   = ""
      "k8s.io/cluster-autoscaler/node-template/label/node.kubernetes.io/exclude-from-external-load-balancers" = ""
      "k8s.io/role/control-plane"                                                                             = "1"
      "k8s.io/role/master"                                                                                    = "1"
      "kops.k8s.io/instancegroup"                                                                             = "control-plane-ap-south-1a"
      "kubernetes.io/cluster/kops-k8s.prodcrashed.live"                                                       = "owned"
    }
  }
  tags = {
    "KubernetesCluster"                                                                                     = "kops-k8s.prodcrashed.live"
    "Name"                                                                                                  = "control-plane-ap-south-1a.masters.kops-k8s.prodcrashed.live"
    "aws-node-termination-handler/managed"                                                                  = ""
    "k8s.io/cluster-autoscaler/node-template/label/kops.k8s.io/kops-controller-pki"                         = ""
    "k8s.io/cluster-autoscaler/node-template/label/node-role.kubernetes.io/control-plane"                   = ""
    "k8s.io/cluster-autoscaler/node-template/label/node.kubernetes.io/exclude-from-external-load-balancers" = ""
    "k8s.io/role/control-plane"                                                                             = "1"
    "k8s.io/role/master"                                                                                    = "1"
    "kops.k8s.io/instancegroup"                                                                             = "control-plane-ap-south-1a"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live"                                                       = "owned"
  }
  user_data = filebase64("${path.module}/data/aws_launch_template_control-plane-ap-south-1a.masters.kops-k8s.prodcrashed.live_user_data")
}

resource "aws_launch_template" "nodes-ap-south-1a-kops-k8s-prodcrashed-live" {
  block_device_mappings {
    device_name = "/dev/sda1"
    ebs {
      delete_on_termination = true
      encrypted             = true
      iops                  = 3000
      throughput            = 125
      volume_size           = 128
      volume_type           = "gp3"
    }
  }
  iam_instance_profile {
    name = aws_iam_instance_profile.nodes-kops-k8s-prodcrashed-live.id
  }
  image_id      = "ami-0554f7fb41d511dd0"
  instance_type = "t3.micro"
  key_name      = aws_key_pair.kubernetes-kops-k8s-prodcrashed-live-006765ba239e1535853258a6d4c7229f.id
  lifecycle {
    create_before_destroy = true
  }
  metadata_options {
    http_endpoint               = "enabled"
    http_protocol_ipv6          = "disabled"
    http_put_response_hop_limit = 1
    http_tokens                 = "required"
  }
  monitoring {
    enabled = false
  }
  name = "nodes-ap-south-1a.kops-k8s.prodcrashed.live"
  network_interfaces {
    associate_public_ip_address = true
    delete_on_termination       = true
    ipv6_address_count          = 0
    security_groups             = [aws_security_group.nodes-kops-k8s-prodcrashed-live.id]
  }
  tag_specifications {
    resource_type = "instance"
    tags = {
      "KubernetesCluster"                                                          = "kops-k8s.prodcrashed.live"
      "Name"                                                                       = "nodes-ap-south-1a.kops-k8s.prodcrashed.live"
      "aws-node-termination-handler/managed"                                       = ""
      "k8s.io/cluster-autoscaler/node-template/label/node-role.kubernetes.io/node" = ""
      "k8s.io/role/node"                                                           = "1"
      "kops.k8s.io/instancegroup"                                                  = "nodes-ap-south-1a"
      "kubernetes.io/cluster/kops-k8s.prodcrashed.live"                            = "owned"
    }
  }
  tag_specifications {
    resource_type = "volume"
    tags = {
      "KubernetesCluster"                                                          = "kops-k8s.prodcrashed.live"
      "Name"                                                                       = "nodes-ap-south-1a.kops-k8s.prodcrashed.live"
      "aws-node-termination-handler/managed"                                       = ""
      "k8s.io/cluster-autoscaler/node-template/label/node-role.kubernetes.io/node" = ""
      "k8s.io/role/node"                                                           = "1"
      "kops.k8s.io/instancegroup"                                                  = "nodes-ap-south-1a"
      "kubernetes.io/cluster/kops-k8s.prodcrashed.live"                            = "owned"
    }
  }
  tags = {
    "KubernetesCluster"                                                          = "kops-k8s.prodcrashed.live"
    "Name"                                                                       = "nodes-ap-south-1a.kops-k8s.prodcrashed.live"
    "aws-node-termination-handler/managed"                                       = ""
    "k8s.io/cluster-autoscaler/node-template/label/node-role.kubernetes.io/node" = ""
    "k8s.io/role/node"                                                           = "1"
    "kops.k8s.io/instancegroup"                                                  = "nodes-ap-south-1a"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live"                            = "owned"
  }
  user_data = filebase64("${path.module}/data/aws_launch_template_nodes-ap-south-1a.kops-k8s.prodcrashed.live_user_data")
}

resource "aws_lb" "api-kops-k8s-prodcrashed-live" {
  enable_cross_zone_load_balancing = true
  internal                         = false
  load_balancer_type               = "network"
  name                             = "api-kops-k8s-prodcrashed--fsrnsu"
  security_groups                  = [aws_security_group.api-elb-kops-k8s-prodcrashed-live.id]
  subnet_mapping {
    subnet_id = aws_subnet.ap-south-1a-kops-k8s-prodcrashed-live.id
  }
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "api.kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_lb_listener" "api-kops-k8s-prodcrashed-live-3988" {
  default_action {
    target_group_arn = aws_lb_target_group.kops-controller-kops-k8s--j959e1.id
    type             = "forward"
  }
  load_balancer_arn = aws_lb.api-kops-k8s-prodcrashed-live.id
  port              = 3988
  protocol          = "TCP"
}

resource "aws_lb_listener" "api-kops-k8s-prodcrashed-live-443" {
  default_action {
    target_group_arn = aws_lb_target_group.tcp-kops-k8s-prodcrashed--i61hbp.id
    type             = "forward"
  }
  load_balancer_arn = aws_lb.api-kops-k8s-prodcrashed-live.id
  port              = 443
  protocol          = "TCP"
}

resource "aws_lb_target_group" "kops-controller-kops-k8s--j959e1" {
  connection_termination = "true"
  deregistration_delay   = "30"
  health_check {
    healthy_threshold   = 2
    interval            = 10
    protocol            = "TCP"
    unhealthy_threshold = 2
  }
  name     = "kops-controller-kops-k8s--j959e1"
  port     = 3988
  protocol = "TCP"
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "kops-controller-kops-k8s--j959e1"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
  vpc_id = aws_vpc.kops-k8s-prodcrashed-live.id
}

resource "aws_lb_target_group" "tcp-kops-k8s-prodcrashed--i61hbp" {
  connection_termination = "true"
  deregistration_delay   = "30"
  health_check {
    healthy_threshold   = 2
    interval            = 10
    protocol            = "TCP"
    unhealthy_threshold = 2
  }
  name     = "tcp-kops-k8s-prodcrashed--i61hbp"
  port     = 443
  protocol = "TCP"
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "tcp-kops-k8s-prodcrashed--i61hbp"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
  vpc_id = aws_vpc.kops-k8s-prodcrashed-live.id
}

resource "aws_route" "route-0-0-0-0--0" {
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.kops-k8s-prodcrashed-live.id
  route_table_id         = aws_route_table.kops-k8s-prodcrashed-live.id
}

resource "aws_route" "route-__--0" {
  destination_ipv6_cidr_block = "::/0"
  gateway_id                  = aws_internet_gateway.kops-k8s-prodcrashed-live.id
  route_table_id              = aws_route_table.kops-k8s-prodcrashed-live.id
}

resource "aws_route_table" "kops-k8s-prodcrashed-live" {
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
    "kubernetes.io/kops/role"                         = "public"
  }
  vpc_id = aws_vpc.kops-k8s-prodcrashed-live.id
}

resource "aws_route_table_association" "ap-south-1a-kops-k8s-prodcrashed-live" {
  route_table_id = aws_route_table.kops-k8s-prodcrashed-live.id
  subnet_id      = aws_subnet.ap-south-1a-kops-k8s-prodcrashed-live.id
}

resource "aws_s3_object" "cluster-completed-spec" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_cluster-completed.spec_content")
  key                    = "kops-k8s.prodcrashed.live/cluster-completed.spec"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "etcd-cluster-spec-events" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_etcd-cluster-spec-events_content")
  key                    = "kops-k8s.prodcrashed.live/backups/etcd/events/control/etcd-cluster-spec"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "etcd-cluster-spec-main" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_etcd-cluster-spec-main_content")
  key                    = "kops-k8s.prodcrashed.live/backups/etcd/main/control/etcd-cluster-spec"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "kops-k8s-prodcrashed-live-addons-aws-cloud-controller-addons-k8s-io-k8s-1-18" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_kops-k8s.prodcrashed.live-addons-aws-cloud-controller.addons.k8s.io-k8s-1.18_content")
  key                    = "kops-k8s.prodcrashed.live/addons/aws-cloud-controller.addons.k8s.io/k8s-1.18.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "kops-k8s-prodcrashed-live-addons-aws-ebs-csi-driver-addons-k8s-io-k8s-1-17" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_kops-k8s.prodcrashed.live-addons-aws-ebs-csi-driver.addons.k8s.io-k8s-1.17_content")
  key                    = "kops-k8s.prodcrashed.live/addons/aws-ebs-csi-driver.addons.k8s.io/k8s-1.17.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "kops-k8s-prodcrashed-live-addons-bootstrap" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_kops-k8s.prodcrashed.live-addons-bootstrap_content")
  key                    = "kops-k8s.prodcrashed.live/addons/bootstrap-channel.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "kops-k8s-prodcrashed-live-addons-coredns-addons-k8s-io-k8s-1-12" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_kops-k8s.prodcrashed.live-addons-coredns.addons.k8s.io-k8s-1.12_content")
  key                    = "kops-k8s.prodcrashed.live/addons/coredns.addons.k8s.io/k8s-1.12.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "kops-k8s-prodcrashed-live-addons-kops-controller-addons-k8s-io-k8s-1-16" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_kops-k8s.prodcrashed.live-addons-kops-controller.addons.k8s.io-k8s-1.16_content")
  key                    = "kops-k8s.prodcrashed.live/addons/kops-controller.addons.k8s.io/k8s-1.16.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "kops-k8s-prodcrashed-live-addons-kubelet-api-rbac-addons-k8s-io-k8s-1-9" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_kops-k8s.prodcrashed.live-addons-kubelet-api.rbac.addons.k8s.io-k8s-1.9_content")
  key                    = "kops-k8s.prodcrashed.live/addons/kubelet-api.rbac.addons.k8s.io/k8s-1.9.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "kops-k8s-prodcrashed-live-addons-leader-migration-rbac-addons-k8s-io-k8s-1-23" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_kops-k8s.prodcrashed.live-addons-leader-migration.rbac.addons.k8s.io-k8s-1.23_content")
  key                    = "kops-k8s.prodcrashed.live/addons/leader-migration.rbac.addons.k8s.io/k8s-1.23.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "kops-k8s-prodcrashed-live-addons-limit-range-addons-k8s-io" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_kops-k8s.prodcrashed.live-addons-limit-range.addons.k8s.io_content")
  key                    = "kops-k8s.prodcrashed.live/addons/limit-range.addons.k8s.io/v1.5.0.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "kops-k8s-prodcrashed-live-addons-networking-cilium-io-k8s-1-16" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_kops-k8s.prodcrashed.live-addons-networking.cilium.io-k8s-1.16_content")
  key                    = "kops-k8s.prodcrashed.live/addons/networking.cilium.io/k8s-1.16-v1.15.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "kops-k8s-prodcrashed-live-addons-node-termination-handler-aws-k8s-1-11" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_kops-k8s.prodcrashed.live-addons-node-termination-handler.aws-k8s-1.11_content")
  key                    = "kops-k8s.prodcrashed.live/addons/node-termination-handler.aws/k8s-1.11.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "kops-k8s-prodcrashed-live-addons-storage-aws-addons-k8s-io-v1-15-0" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_kops-k8s.prodcrashed.live-addons-storage-aws.addons.k8s.io-v1.15.0_content")
  key                    = "kops-k8s.prodcrashed.live/addons/storage-aws.addons.k8s.io/v1.15.0.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "kops-version-txt" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_kops-version.txt_content")
  key                    = "kops-k8s.prodcrashed.live/kops-version.txt"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "manifests-etcdmanager-events-control-plane-ap-south-1a" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_manifests-etcdmanager-events-control-plane-ap-south-1a_content")
  key                    = "kops-k8s.prodcrashed.live/manifests/etcd/events-control-plane-ap-south-1a.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "manifests-etcdmanager-main-control-plane-ap-south-1a" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_manifests-etcdmanager-main-control-plane-ap-south-1a_content")
  key                    = "kops-k8s.prodcrashed.live/manifests/etcd/main-control-plane-ap-south-1a.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "manifests-static-kube-apiserver-healthcheck" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_manifests-static-kube-apiserver-healthcheck_content")
  key                    = "kops-k8s.prodcrashed.live/manifests/static/kube-apiserver-healthcheck.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "nodeupconfig-control-plane-ap-south-1a" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_nodeupconfig-control-plane-ap-south-1a_content")
  key                    = "kops-k8s.prodcrashed.live/igconfig/control-plane/control-plane-ap-south-1a/nodeupconfig.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_s3_object" "nodeupconfig-nodes-ap-south-1a" {
  acl                    = "bucket-owner-full-control"
  bucket                 = "test-kops-s3-bucket"
  content                = file("${path.module}/data/aws_s3_object_nodeupconfig-nodes-ap-south-1a_content")
  key                    = "kops-k8s.prodcrashed.live/igconfig/node/nodes-ap-south-1a/nodeupconfig.yaml"
  provider               = aws.files
  server_side_encryption = "AES256"
}

resource "aws_security_group" "api-elb-kops-k8s-prodcrashed-live" {
  description = "Security group for api ELB"
  name        = "api-elb.kops-k8s.prodcrashed.live"
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "api-elb.kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
  vpc_id = aws_vpc.kops-k8s-prodcrashed-live.id
}

resource "aws_security_group" "masters-kops-k8s-prodcrashed-live" {
  description = "Security group for masters"
  name        = "masters.kops-k8s.prodcrashed.live"
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "masters.kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
  vpc_id = aws_vpc.kops-k8s-prodcrashed-live.id
}

resource "aws_security_group" "nodes-kops-k8s-prodcrashed-live" {
  description = "Security group for nodes"
  name        = "nodes.kops-k8s.prodcrashed.live"
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "nodes.kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
  vpc_id = aws_vpc.kops-k8s-prodcrashed-live.id
}

resource "aws_security_group_rule" "from-0-0-0-0--0-ingress-tcp-22to22-masters-kops-k8s-prodcrashed-live" {
  cidr_blocks       = ["0.0.0.0/0"]
  from_port         = 22
  protocol          = "tcp"
  security_group_id = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  to_port           = 22
  type              = "ingress"
}

resource "aws_security_group_rule" "from-0-0-0-0--0-ingress-tcp-22to22-nodes-kops-k8s-prodcrashed-live" {
  cidr_blocks       = ["0.0.0.0/0"]
  from_port         = 22
  protocol          = "tcp"
  security_group_id = aws_security_group.nodes-kops-k8s-prodcrashed-live.id
  to_port           = 22
  type              = "ingress"
}

resource "aws_security_group_rule" "from-0-0-0-0--0-ingress-tcp-443to443-api-elb-kops-k8s-prodcrashed-live" {
  cidr_blocks       = ["0.0.0.0/0"]
  from_port         = 443
  protocol          = "tcp"
  security_group_id = aws_security_group.api-elb-kops-k8s-prodcrashed-live.id
  to_port           = 443
  type              = "ingress"
}

resource "aws_security_group_rule" "from-__--0-ingress-tcp-22to22-masters-kops-k8s-prodcrashed-live" {
  from_port         = 22
  ipv6_cidr_blocks  = ["::/0"]
  protocol          = "tcp"
  security_group_id = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  to_port           = 22
  type              = "ingress"
}

resource "aws_security_group_rule" "from-__--0-ingress-tcp-22to22-nodes-kops-k8s-prodcrashed-live" {
  from_port         = 22
  ipv6_cidr_blocks  = ["::/0"]
  protocol          = "tcp"
  security_group_id = aws_security_group.nodes-kops-k8s-prodcrashed-live.id
  to_port           = 22
  type              = "ingress"
}

resource "aws_security_group_rule" "from-__--0-ingress-tcp-443to443-api-elb-kops-k8s-prodcrashed-live" {
  from_port         = 443
  ipv6_cidr_blocks  = ["::/0"]
  protocol          = "tcp"
  security_group_id = aws_security_group.api-elb-kops-k8s-prodcrashed-live.id
  to_port           = 443
  type              = "ingress"
}

resource "aws_security_group_rule" "from-api-elb-kops-k8s-prodcrashed-live-egress-all-0to0-0-0-0-0--0" {
  cidr_blocks       = ["0.0.0.0/0"]
  from_port         = 0
  protocol          = "-1"
  security_group_id = aws_security_group.api-elb-kops-k8s-prodcrashed-live.id
  to_port           = 0
  type              = "egress"
}

resource "aws_security_group_rule" "from-api-elb-kops-k8s-prodcrashed-live-egress-all-0to0-__--0" {
  from_port         = 0
  ipv6_cidr_blocks  = ["::/0"]
  protocol          = "-1"
  security_group_id = aws_security_group.api-elb-kops-k8s-prodcrashed-live.id
  to_port           = 0
  type              = "egress"
}

resource "aws_security_group_rule" "from-masters-kops-k8s-prodcrashed-live-egress-all-0to0-0-0-0-0--0" {
  cidr_blocks       = ["0.0.0.0/0"]
  from_port         = 0
  protocol          = "-1"
  security_group_id = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  to_port           = 0
  type              = "egress"
}

resource "aws_security_group_rule" "from-masters-kops-k8s-prodcrashed-live-egress-all-0to0-__--0" {
  from_port         = 0
  ipv6_cidr_blocks  = ["::/0"]
  protocol          = "-1"
  security_group_id = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  to_port           = 0
  type              = "egress"
}

resource "aws_security_group_rule" "from-masters-kops-k8s-prodcrashed-live-ingress-all-0to0-masters-kops-k8s-prodcrashed-live" {
  from_port                = 0
  protocol                 = "-1"
  security_group_id        = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  source_security_group_id = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  to_port                  = 0
  type                     = "ingress"
}

resource "aws_security_group_rule" "from-masters-kops-k8s-prodcrashed-live-ingress-all-0to0-nodes-kops-k8s-prodcrashed-live" {
  from_port                = 0
  protocol                 = "-1"
  security_group_id        = aws_security_group.nodes-kops-k8s-prodcrashed-live.id
  source_security_group_id = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  to_port                  = 0
  type                     = "ingress"
}

resource "aws_security_group_rule" "from-nodes-kops-k8s-prodcrashed-live-egress-all-0to0-0-0-0-0--0" {
  cidr_blocks       = ["0.0.0.0/0"]
  from_port         = 0
  protocol          = "-1"
  security_group_id = aws_security_group.nodes-kops-k8s-prodcrashed-live.id
  to_port           = 0
  type              = "egress"
}

resource "aws_security_group_rule" "from-nodes-kops-k8s-prodcrashed-live-egress-all-0to0-__--0" {
  from_port         = 0
  ipv6_cidr_blocks  = ["::/0"]
  protocol          = "-1"
  security_group_id = aws_security_group.nodes-kops-k8s-prodcrashed-live.id
  to_port           = 0
  type              = "egress"
}

resource "aws_security_group_rule" "from-nodes-kops-k8s-prodcrashed-live-ingress-all-0to0-nodes-kops-k8s-prodcrashed-live" {
  from_port                = 0
  protocol                 = "-1"
  security_group_id        = aws_security_group.nodes-kops-k8s-prodcrashed-live.id
  source_security_group_id = aws_security_group.nodes-kops-k8s-prodcrashed-live.id
  to_port                  = 0
  type                     = "ingress"
}

resource "aws_security_group_rule" "from-nodes-kops-k8s-prodcrashed-live-ingress-tcp-1to2379-masters-kops-k8s-prodcrashed-live" {
  from_port                = 1
  protocol                 = "tcp"
  security_group_id        = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  source_security_group_id = aws_security_group.nodes-kops-k8s-prodcrashed-live.id
  to_port                  = 2379
  type                     = "ingress"
}

resource "aws_security_group_rule" "from-nodes-kops-k8s-prodcrashed-live-ingress-tcp-2382to4000-masters-kops-k8s-prodcrashed-live" {
  from_port                = 2382
  protocol                 = "tcp"
  security_group_id        = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  source_security_group_id = aws_security_group.nodes-kops-k8s-prodcrashed-live.id
  to_port                  = 4000
  type                     = "ingress"
}

resource "aws_security_group_rule" "from-nodes-kops-k8s-prodcrashed-live-ingress-tcp-4003to65535-masters-kops-k8s-prodcrashed-live" {
  from_port                = 4003
  protocol                 = "tcp"
  security_group_id        = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  source_security_group_id = aws_security_group.nodes-kops-k8s-prodcrashed-live.id
  to_port                  = 65535
  type                     = "ingress"
}

resource "aws_security_group_rule" "from-nodes-kops-k8s-prodcrashed-live-ingress-udp-1to65535-masters-kops-k8s-prodcrashed-live" {
  from_port                = 1
  protocol                 = "udp"
  security_group_id        = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  source_security_group_id = aws_security_group.nodes-kops-k8s-prodcrashed-live.id
  to_port                  = 65535
  type                     = "ingress"
}

resource "aws_security_group_rule" "https-elb-to-master" {
  from_port                = 443
  protocol                 = "tcp"
  security_group_id        = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  source_security_group_id = aws_security_group.api-elb-kops-k8s-prodcrashed-live.id
  to_port                  = 443
  type                     = "ingress"
}

resource "aws_security_group_rule" "icmp-pmtu-api-elb-0-0-0-0--0" {
  cidr_blocks       = ["0.0.0.0/0"]
  from_port         = 3
  protocol          = "icmp"
  security_group_id = aws_security_group.api-elb-kops-k8s-prodcrashed-live.id
  to_port           = 4
  type              = "ingress"
}

resource "aws_security_group_rule" "icmp-pmtu-cp-to-elb" {
  from_port                = 3
  protocol                 = "icmp"
  security_group_id        = aws_security_group.api-elb-kops-k8s-prodcrashed-live.id
  source_security_group_id = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  to_port                  = 4
  type                     = "ingress"
}

resource "aws_security_group_rule" "icmp-pmtu-elb-to-cp" {
  from_port                = 3
  protocol                 = "icmp"
  security_group_id        = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  source_security_group_id = aws_security_group.api-elb-kops-k8s-prodcrashed-live.id
  to_port                  = 4
  type                     = "ingress"
}

resource "aws_security_group_rule" "icmpv6-pmtu-api-elb-__--0" {
  from_port         = -1
  ipv6_cidr_blocks  = ["::/0"]
  protocol          = "icmpv6"
  security_group_id = aws_security_group.api-elb-kops-k8s-prodcrashed-live.id
  to_port           = -1
  type              = "ingress"
}

resource "aws_security_group_rule" "kops-controller-elb-to-cp" {
  from_port                = 3988
  protocol                 = "tcp"
  security_group_id        = aws_security_group.masters-kops-k8s-prodcrashed-live.id
  source_security_group_id = aws_security_group.api-elb-kops-k8s-prodcrashed-live.id
  to_port                  = 3988
  type                     = "ingress"
}

resource "aws_security_group_rule" "node-to-elb" {
  from_port                = 0
  protocol                 = "-1"
  security_group_id        = aws_security_group.api-elb-kops-k8s-prodcrashed-live.id
  source_security_group_id = aws_security_group.nodes-kops-k8s-prodcrashed-live.id
  to_port                  = 0
  type                     = "ingress"
}

resource "aws_sqs_queue" "kops-k8s-prodcrashed-live-nth" {
  message_retention_seconds = 300
  name                      = "kops-k8s-prodcrashed-live-nth"
  policy                    = file("${path.module}/data/aws_sqs_queue_kops-k8s-prodcrashed-live-nth_policy")
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "kops-k8s-prodcrashed-live-nth"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_subnet" "ap-south-1a-kops-k8s-prodcrashed-live" {
  availability_zone                           = "ap-south-1a"
  cidr_block                                  = "172.20.0.0/16"
  enable_resource_name_dns_a_record_on_launch = true
  private_dns_hostname_type_on_launch         = "resource-name"
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "ap-south-1a.kops-k8s.prodcrashed.live"
    "SubnetType"                                      = "Public"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
    "kubernetes.io/role/elb"                          = "1"
    "kubernetes.io/role/internal-elb"                 = "1"
  }
  vpc_id = aws_vpc.kops-k8s-prodcrashed-live.id
}

resource "aws_vpc" "kops-k8s-prodcrashed-live" {
  assign_generated_ipv6_cidr_block = true
  cidr_block                       = "172.20.0.0/16"
  enable_dns_hostnames             = true
  enable_dns_support               = true
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_vpc_dhcp_options" "kops-k8s-prodcrashed-live" {
  domain_name         = "ap-south-1.compute.internal"
  domain_name_servers = ["AmazonProvidedDNS"]
  tags = {
    "KubernetesCluster"                               = "kops-k8s.prodcrashed.live"
    "Name"                                            = "kops-k8s.prodcrashed.live"
    "kubernetes.io/cluster/kops-k8s.prodcrashed.live" = "owned"
  }
}

resource "aws_vpc_dhcp_options_association" "kops-k8s-prodcrashed-live" {
  dhcp_options_id = aws_vpc_dhcp_options.kops-k8s-prodcrashed-live.id
  vpc_id          = aws_vpc.kops-k8s-prodcrashed-live.id
}

terraform {
  required_version = ">= 0.15.0"
  required_providers {
    aws = {
      "configuration_aliases" = [aws.files]
      "source"                = "hashicorp/aws"
      "version"               = ">= 5.0.0"
    }
  }
}
