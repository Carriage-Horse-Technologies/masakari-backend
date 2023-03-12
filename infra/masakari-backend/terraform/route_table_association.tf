resource "aws_route_table_association" "AlloMasakariBackendPublicSubnet1" {
  route_table_id = aws_route_table.AlloMasakariBackendPublicRouteTable.id
  subnet_id      = aws_subnet.AlloMasakariBackendPublicSubnet1.id
}

resource "aws_route_table_association" "AlloMasakariBackendPublicSubnet2" {
  route_table_id = aws_route_table.AlloMasakariBackendPublicRouteTable.id
  subnet_id      = aws_subnet.AlloMasakariBackendPublicSubnet2.id
}

resource "aws_route_table_association" "AlloMasakariBackendPrivateSubnet1" {
  route_table_id = aws_route_table.AlloMasakariBackendPrivateRouteTable.id
  subnet_id      = aws_subnet.AlloMasakariBackendPrivateSubnet1.id
}

resource "aws_route_table_association" "AlloMasakariBackendPrivateSubnet2" {
  route_table_id = aws_route_table.AlloMasakariBackendPrivateRouteTable.id
  subnet_id      = aws_subnet.AlloMasakariBackendPrivateSubnet2.id
}
