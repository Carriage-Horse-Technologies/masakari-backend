resource "aws_elasticache_cluster" "this" {
  cluster_id                    = var.cluster_name
  engine                        = var.engine
  engine_version                = var.engine_version
  node_type                     = var.node_type
  parameter_group_name          = var.parameter_group_name
  num_cache_nodes         = var.num_cache_nodes
  port                          = var.port
  subnet_group_name             = aws_elasticache_subnet_group.this.name
  security_group_ids = [aws_security_group.this.id, ]
  maintenance_window = var.maintenance_window
  apply_immediately = var.apply_immediately
}
