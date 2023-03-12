resource "aws_ecs_cluster_capacity_providers" "AlloMasakariBackendECSCapacityProider" {
  cluster_name = aws_ecs_cluster.AlloMasakariBackendECSCluster.name

  capacity_providers = ["FARGATE", "FARGATE_SPOT"]
}
