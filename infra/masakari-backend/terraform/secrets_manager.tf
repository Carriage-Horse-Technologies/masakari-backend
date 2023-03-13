resource "aws_secretsmanager_secret" "openai_api_token" {
  name = "openai-api-token"
}

resource "aws_secretsmanager_secret_version" "openai_api_token" {
  secret_id = aws_secretsmanager_secret.openai_api_token.id
  secret_string = jsonencode({
    openai_api_token = var.OPENAI_API_TOKEN
  })
}
