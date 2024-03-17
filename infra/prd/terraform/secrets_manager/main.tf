resource "aws_secretsmanager_secret" "bot_token" {
  name        = "telegram_bot_token"
  description = "The Telegram Bot API token for bball8bot"
}

resource "aws_secretsmanager_secret_version" "bot_token" {
  secret_id     = aws_secretsmanager_secret.bot_token.id
  secret_string = var.telegram_bot_token
}
