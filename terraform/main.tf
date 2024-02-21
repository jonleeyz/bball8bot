module "lambda" {
  source                = "./lambda"
  lambda_name           = "bball8bot_event_handler"
  handler_function_name = "main"
}

module "sqs" {
  source      = "./sqs"
  queue_name  = "bball8bot_event_queue"
  is_fifo     = false
  lambda_name = module.lambda.function_name

  is_sqs_to_lambda_integration_enabled = false
  sqs_to_lambda_batch_size             = 1
}

module "api_gateway" {
  source     = "./api_gateway"
  queue_arn  = module.sqs.queue_arn
  queue_name = module.sqs.queue_name
  aws_region = var.aws_region
}
