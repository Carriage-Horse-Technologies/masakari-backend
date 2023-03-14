variable "vpc_id" {
  type = string
}

variable "vpc_cidr_block" {
  type = string
}

variable "cluster_name" {
  type    = string
  default = "allo-masakari-backend-memcached"
}

variable "num_cache_nodes" {
  type    = number
  default = 2
}

variable "engine" {
  type    = string
  default = "memcached"
}

variable "engine_version" {
  type    = string
  default = "1.6.6"
}

variable "node_type" {
  type    = string
  default = "cache.t3.micro"
}

variable "family" {
  type    = string
  default = "memcached"
}

variable "port" {
  type    = number
  default = 11211
}

variable "parameter_group_name" {
  type    = string
  default = "default.memcached1.6"
}

variable "automatic_failover_enabled" {
  type    = bool
  default = true
}

variable "multi_az_enabled" {
  type    = bool
  default = true
}

variable "snapshot_window" {
  type    = string
  default = "17:00-18:00"
}

variable "snapshot_retention_limit" {
  type    = number
  default = 2
}

variable "maintenance_window" {
  type    = string
  default = "Sat:18:00-Sat:19:00"
}

variable "transit_encryption_enabled" {
  type    = bool
  default = false
}

variable "at_rest_encryption_enabled" {
  type    = bool
  default = false
}

variable "apply_immediately" {
  type    = bool
  default = true
}

variable "subnet_ids" {
  type = list(string)
}
