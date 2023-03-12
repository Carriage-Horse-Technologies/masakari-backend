resource "aws_route_table" "AlloMasakariBackendPublicRouteTable" {
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.AlloMasakariBackendIGW.id
  }

  tags = {
    Name  = "prod-AlloMasakariBackendProgect-rtb-public"
    Group = "AlloMasakariBackend"
  }

  tags_all = {
    Name  = "prod-AlloMasakariBackendProgect-rtb-public"
    Group = "AlloMasakariBackend"
  }

  vpc_id = aws_vpc.AlloMasakariBackendVPC.id
}

resource "aws_route_table" "AlloMasakariBackendPrivateRouteTable" {
  tags = {
    Name  = "prod-AlloMasakariBackendProgect-rtb-private"
    Group = "AlloMasakariBackend"
  }

  tags_all = {
    Name  = "prod-AlloMasakariBackendProgect-rtb-private"
    Group = "AlloMasakariBackend"
  }

  vpc_id = aws_vpc.AlloMasakariBackendVPC.id
}
