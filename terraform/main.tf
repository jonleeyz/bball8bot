module "lambda" {
  source                = "./lambda"
  lambda_name           = "bball8bot_event_handler"
  handler_function_name = "main"
}

module "sqs" {
  source     = "./sqs"
  queue_name = "bball8bot_event_queue"
  is_fifo    = false
}
