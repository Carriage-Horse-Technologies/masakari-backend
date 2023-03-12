resource "aws_security_group" "AlloMasakariBackendALBSecurityGroup" {
  description = "prod-AlloMasakariBackendALBSG"

  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = "0"
    protocol    = "-1"
    self        = "false"
    to_port     = "0"
  }

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = "443"
    protocol    = "tcp"
    self        = "false"
    to_port     = "443"
  }

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = "80"
    protocol    = "tcp"
    self        = "false"
    to_port     = "80"
  }

  name = "prod-AlloMasakariBackendALBSecurityGroup"

  tags = {
    Name  = "prod-AlloMasakariBackend-alb-security-group"
    Group = "AlloMasakariBackend"
  }

  tags_all = {
    Name  = "prod-AlloMasakariBackend-alb-security-group"
    Group = "AlloMasakariBackend"
  }

  vpc_id = aws_vpc.AlloMasakariBackendVPC.id
}

resource "aws_security_group" "AlloMasakariBackendPostgresSecurityGroup" {
  description = "Created by RDS management console"

  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = "0"
    protocol    = "-1"
    self        = "false"
    to_port     = "0"
  }

  ingress {
    cidr_blocks = [
      aws_subnet.AlloMasakariBackendPublicSubnet1.cidr_block,
      aws_subnet.AlloMasakariBackendPublicSubnet2.cidr_block
    ]
    from_port = "5432"
    protocol  = "tcp"
    self      = "false"
    to_port   = "5432"
  }

  name = "prod-AlloMasakariBackendPostgresSecurityGroup"

  tags = {
    Name  = "prod-AlloMasakariBackend-postgres-security-group"
    Group = "AlloMasakariBackend"
  }

  tags_all = {
    Name  = "prod-AlloMasakariBackend-postgres-security-group"
    Group = "AlloMasakariBackend"
  }

  vpc_id = aws_vpc.AlloMasakariBackendVPC.id
}

resource "aws_security_group" "AlloMasakariBackendECSTaskSecurityGroup" {
  description = "prod-AlloMasakariBackendECSTaskSecurityGroup"

  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = "0"
    protocol    = "-1"
    self        = "false"
    to_port     = "0"
  }

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = "80"
    protocol    = "tcp"
    self        = "false"
    to_port     = "8000"
  }

  name = "prod-AlloMasakariBackendECSTaskSecurityGroup"

  tags = {
    Name  = "prod-AlloMasakariBackend-ecs-task-security-group"
    Group = "AlloMasakariBackend"
  }

  tags_all = {
    Name  = "prod-AlloMasakariBackend-ecs-task-security-group"
    Group = "AlloMasakariBackend"
  }

  vpc_id = aws_vpc.AlloMasakariBackendVPC.id
}
