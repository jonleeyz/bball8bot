variable "aws_region" {
  description = "The AWS region that that resources will be provisioned in, should any of them need to belong to one"
  type        = string
}

variable "telegram_bot_token" {
  description = "The Telegram Bot API token for bball8bot"
  type        = string
  sensitive   = true
}
