module "iam" {
  source = "./iam"
}

module "lambda" {
  source                = "./lambda"
  lambda_name           = "bball8bot_event_handler"
  handler_function_name = "main"
  lambda_iam_role_arn   = module.iam.lambda_iam_role_arn

  is_sqs_to_lambda_integration_enabled = false
  sqs_to_lambda_batch_size             = 1
  sqs_arn = module.sqs.queue_arn
}

module "sqs" {
  source      = "./sqs"
  queue_name  = "bball8bot_event_queue"
  is_fifo     = false
}

module "api_gateway" {
  source     = "./api_gateway"
  queue_arn  = module.sqs.queue_arn
  queue_name = module.sqs.queue_name
  aws_region = var.aws_region
}

module "secrets_manager" {
  source = "./secrets_manager"
  telegram_bot_token = var.telegram_bot_token
}
