module "elasticache" {
  source = "./modules/elasticache"

  ## network
  vpc_id         = aws_vpc.AlloMasakariBackendVPC.id
  vpc_cidr_block = aws_vpc.AlloMasakariBackendVPC.cidr_block
  subnet_ids = [
    aws_subnet.AlloMasakariBackendPrivateSubnet1.id,
    aws_subnet.AlloMasakariBackendPrivateSubnet2.id,
  ]

  ## base for redis(cluster mode disabled)
  cluster_name          = "allo-masakari-backend-redis"
  node_type             = "cache.t3.micro"
  engine_version        = "5.0.6"
  family                = "redis5.0"
  number_cache_clusters = 2

  ## not available for t1/t2 instances
  automatic_failover_enabled = true

  ## automatic Failover must also be enabled
  multi_az_enabled = true
}
