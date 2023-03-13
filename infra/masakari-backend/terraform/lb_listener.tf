resource "aws_lb_listener" "AlloMasakariBackendBackendECSIbListernerHttp" {
  default_action {
    target_group_arn = aws_lb_target_group.AlloMasakariBackendTG.arn
    type             = "forward"
  }

  load_balancer_arn = aws_lb.AlloMasakariBackendALB.arn
  port              = "80"
  protocol          = "HTTP"

  tags = {
    Name  = "prod-AlloMasakariBackend-backend-ecs-lb-listerner-http"
    Group = "AlloMasakariBackend"
  }

  tags_all = {
    Name  = "prod-AlloMasakariBackend-backend-ecs-lb-listerner-http"
    Group = "AlloMasakariBackend"
  }
}

resource "aws_lb_listener" "AlloMasakariBackendBackendECSIbListernerHttps" {
  certificate_arn = var.CERTIFICATE_ARN

  default_action {
    target_group_arn = aws_lb_target_group.AlloMasakariBackendTG.arn
    type             = "forward"
  }

  load_balancer_arn = aws_lb.AlloMasakariBackendALB.arn
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"

  tags = {
    Name  = "prod-AlloMasakariBackend-backend-ecs-lb-listerner-https"
    Group = "AlloMasakariBackend"
  }

  tags_all = {
    Name  = "prod-AlloMasakariBackend-backend-ecs-lb-listerner-https"
    Group = "AlloMasakariBackend"
  }
}
