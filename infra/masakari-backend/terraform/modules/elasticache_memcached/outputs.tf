output "security_group_id" {
  value = aws_security_group.this.id
}

output "cluster_address" {
  value = aws_elasticache_cluster.this.cluster_address
}

output "cluster_id" {
  value = aws_elasticache_cluster.this.cluster_id
}

output "replication_group_id" {
  value = aws_elasticache_cluster.this.id
}
