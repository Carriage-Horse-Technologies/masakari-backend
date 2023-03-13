resource "aws_internet_gateway" "AlloMasakariBackendIGW" {
  tags = {
    Name  = "prod-AlloMasakariBackendProgect-igw"
    Group = "AlloMasakariBackend"
  }

  tags_all = {
    Name  = "prod-AlloMasakariBackendProgect-igw"
    Group = "AlloMasakariBackend"
  }

  vpc_id = aws_vpc.AlloMasakariBackendVPC.id
}
