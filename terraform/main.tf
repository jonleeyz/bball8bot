module "lambda" {
  source                = "./lambda"
  lambda_name           = "bball8bot_event_handler"
  handler_function_name = "main"
}