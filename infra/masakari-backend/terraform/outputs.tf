# account id
output "account_id" {
  value = data.aws_caller_identity.current.account_id
}

# network
# vpc
output "aws_vpc_AlloMasakariBackendVPC_id" {
  value = aws_vpc.AlloMasakariBackendVPC.id
}

# subnet
output "aws_subnet_AlloMasakariBackendPublicSubnet1_id" {
  value = aws_subnet.AlloMasakariBackendPublicSubnet1.id
}

output "aws_subnet_AlloMasakariBackendPrivateSubnet2_id" {
  value = aws_subnet.AlloMasakariBackendPrivateSubnet2.id
}

output "aws_subnet_AlloMasakariBackendPrivateSubnet1_id" {
  value = aws_subnet.AlloMasakariBackendPrivateSubnet1.id
}

output "aws_subnet_AlloMasakariBackendPublicSubnet2_id" {
  value = aws_subnet.AlloMasakariBackendPublicSubnet2.id
}

output "aws_subnet_AlloMasakariBackendPublicSubnet1_availability_zone" {
  value = aws_subnet.AlloMasakariBackendPublicSubnet1.availability_zone
}

output "aws_subnet_AlloMasakariBackendPrivateSubnet2_availability_zone" {
  value = aws_subnet.AlloMasakariBackendPrivateSubnet2.availability_zone
}

output "aws_subnet_AlloMasakariBackendPrivateSubnet1_availability_zone" {
  value = aws_subnet.AlloMasakariBackendPrivateSubnet1.availability_zone
}

output "aws_subnet_AlloMasakariBackendPublicSubnet2_availability_zone" {
  value = aws_subnet.AlloMasakariBackendPublicSubnet2.availability_zone
}

# igw
output "aws_internet_gateway_AlloMasakariBackendIGW_id" {
  value = aws_internet_gateway.AlloMasakariBackendIGW.id
}

# route_table
output "aws_route_table_AlloMasakariBackendPublicRouteTable_id" {
  value = aws_route_table.AlloMasakariBackendPublicRouteTable.id
}

output "aws_route_table_AlloMasakariBackendPrivateRouteTable_id" {
  value = aws_route_table.AlloMasakariBackendPrivateRouteTable.id
}

output "aws_route_table_association_AlloMasakariBackendPublicSubnet1_id" {
  value = aws_route_table_association.AlloMasakariBackendPublicSubnet1.id
}

output "aws_route_table_association_AlloMasakariBackendPublicSubnet2_id" {
  value = aws_route_table_association.AlloMasakariBackendPublicSubnet2.id
}

output "aws_route_table_association_AlloMasakariBackendPrivateSubnet1_id" {
  value = aws_route_table_association.AlloMasakariBackendPrivateSubnet1.id
}

# sg
output "aws_security_group_AlloMasakariBackendALBSecurityGroup_id" {
  value = aws_security_group.AlloMasakariBackendALBSecurityGroup.id
}

output "aws_security_group_AlloMasakariBackendPostgresSecurityGroup_id" {
  value = aws_security_group.AlloMasakariBackendPostgresSecurityGroup.id
}

output "aws_security_group_AlloMasakariBackendECSTaskSecurityGroup_id" {
  value = aws_security_group.AlloMasakariBackendECSTaskSecurityGroup.id
}
output "aws_route_table_association_AlloMasakariBackendPrivateSubnet2_id" {
  value = aws_route_table_association.AlloMasakariBackendPrivateSubnet2.id
}

# db
# rds
# output "aws_db_instance_AlloMasakariBackendPostgres_id" {
#   value = aws_db_instance.AlloMasakariBackendPostgres.id
# }

output "aws_db_subnet_group_AlloMasakariBackendDBSubnetGroup_id" {
  value = aws_db_subnet_group.AlloMasakariBackendDBSubnetGroup.id
}

# backend
# ecr
output "aws_ecr_repository_AlloMasakariBackendRepository_id" {
  value = aws_ecr_repository.AlloMasakariBackendRepository.id
}


# ecs
output "aws_ecs_cluster_AlloMasakariBackendECSCluster_id" {
  value = aws_ecs_cluster.AlloMasakariBackendECSCluster.id
}

output "aws_ecs_cluster_AlloMasakariBackendECSCluster_name" {
  value = aws_ecs_cluster.AlloMasakariBackendECSCluster.name
}

# alb
output "aws_lb_listener_AlloMasakariBackendBackendECSIbListernerHttp_id" {
  value = aws_lb_listener.AlloMasakariBackendBackendECSIbListernerHttp.id
}

output "aws_lb_listener_AlloMasakariBackendBackendECSIbListernerHttps_id" {
  value = aws_lb_listener.AlloMasakariBackendBackendECSIbListernerHttps.id
}

output "aws_lb_target_group_AlloMasakariBackendTG_id" {
  value = aws_lb_target_group.AlloMasakariBackendTG.id
}

output "aws_lb_AlloMasakariBackendALB_id" {
  value = aws_lb.AlloMasakariBackendALB.id
}
