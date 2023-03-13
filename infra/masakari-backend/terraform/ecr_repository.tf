resource "aws_ecr_repository" "AlloMasakariBackendRepository" {
  encryption_configuration {
    encryption_type = "KMS"
  }

  image_scanning_configuration {
    scan_on_push = "true"
  }

  image_tag_mutability = "MUTABLE"
  name                 = "allo-masakari-backend"

  tags = {
    Name  = "allo-masakari-backend-repository"
    Group = "AlloMasakariBackend"
  }

  tags_all = {
    Name  = "allo-masakari-backend-repository"
    Group = "AlloMasakariBackend"
  }
}
