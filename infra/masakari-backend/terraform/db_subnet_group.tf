resource "aws_db_subnet_group" "AlloMasakariBackendDBSubnetGroup" {
  description = "Created from the RDS Management Console"
  name        = aws_vpc.AlloMasakariBackendVPC.id
  subnet_ids  = [aws_subnet.AlloMasakariBackendPublicSubnet1.id, aws_subnet.AlloMasakariBackendPublicSubnet2.id, aws_subnet.AlloMasakariBackendPrivateSubnet1.id, aws_subnet.AlloMasakariBackendPrivateSubnet2.id]
  tags = {
    Group = "AlloMasakariBackend"
  }

  tags_all = {
    Group = "AlloMasakariBackend"
  }
}
