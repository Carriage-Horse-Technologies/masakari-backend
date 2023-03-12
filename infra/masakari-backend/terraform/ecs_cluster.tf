resource "aws_ecs_cluster" "AlloMasakariBackendECSCluster" {
  configuration {
    execute_command_configuration {
      logging = "DEFAULT"
    }
  }

  name = "allo-masakari-backend-cluster"

  setting {
    name  = "containerInsights"
    value = "disabled"
  }

  tags = {
    "ecs:cluster:createdFrom" = "ecs-console-v2"
    Name                      = "AlloMasakariBackend-ecs-cluster"
    Group                     = "AlloMasakariBackend"
  }

  tags_all = {
    "ecs:cluster:createdFrom" = "ecs-console-v2"
    Name                      = "AlloMasakariBackend-ecs-cluster"
    Group                     = "AlloMasakariBackend"
  }
}
