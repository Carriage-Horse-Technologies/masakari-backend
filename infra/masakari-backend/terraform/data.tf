# account id
data "aws_caller_identity" "current" {}

data "aws_elasticache_replication_group" "AlloMasakariBackendElasticache" {
  replication_group_id = module.elasticache.replication_group_id
}
