resource "aws_elasticache_subnet_group" "this" {
  name        = "${var.cluster_name}-subnet"
  description = "Memcached Subnet Group"

  subnet_ids = var.subnet_ids
}
