resource "aws_lb" "AlloMasakariBackendALB" {
  desync_mitigation_mode     = "defensive"
  drop_invalid_header_fields = "false"
  enable_deletion_protection = "false"
  enable_http2               = "true"
  enable_waf_fail_open       = "false"
  idle_timeout               = "60"
  internal                   = "false"
  ip_address_type            = "ipv4"
  load_balancer_type         = "application"
  name                       = "prod-AlloMasakariBackendALB"
  preserve_host_header       = "false"
  security_groups            = [aws_security_group.AlloMasakariBackendALBSecurityGroup.id]

  subnet_mapping {
    subnet_id = aws_subnet.AlloMasakariBackendPublicSubnet1.id
  }

  subnet_mapping {
    subnet_id = aws_subnet.AlloMasakariBackendPublicSubnet2.id
  }

  tags = {
    Name  = "prod-AlloMasakariBackend-backend-ecs-alb"
    Group = "AlloMasakariBackend"
  }

  tags_all = {
    Name  = "prod-AlloMasakariBackend-backend-ecs-alb"
    Group = "AlloMasakariBackend"
  }
}
