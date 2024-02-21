resource "aws_sqs_queue" "bball8bot_events" {
  name                       = var.queue_name
  delay_seconds              = 0
  visibility_timeout_seconds = 30
  max_message_size           = 1024
  message_retention_seconds  = 60
  receive_wait_time_seconds  = 2
  sqs_managed_sse_enabled    = true
}

##### Enables Lambda event handler to be triggered by SQS events

resource "aws_lambda_permission" "allow_sqs_event_to_trigger_lambda_event_handler" {
  statement_id = "AllowLambdaExecutionfromSQS"
  action       = "lambda:InvokeFuntion"

  principal     = "sqs.amazonaws.com"
  function_name = var.lambda_name
  source_arn    = aws_sqs_queue.bball8bot_events.arn
}

resource "aws_lambda_event_source_mapping" "sqs_event_queue_to_lambda_event_handler" {
  enabled    = var.is_sqs_to_lambda_integration_enabled
  batch_size = var.sqs_to_lambda_batch_size

  function_name    = var.lambda_name
  event_source_arn = aws_sqs_queue.bball8bot_events.arn
}