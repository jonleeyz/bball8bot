resource "aws_lambda_function" "bball8bot_event_handler" {
  function_name = var.lambda_name
  handler       = var.handler_function_name
  role          = var.lambda_iam_role_arn

  # Ref: https://docs.aws.amazon.com/lambda/latest/dg/runtimes-provided.html
  runtime          = "provided.al2023"
  filename         = local.output_archive_path
  source_code_hash = data.archive_file.zipped_binary_for_deploy.output_base64sha256
}

##### Enables Lambda event handler to be triggered by SQS events

resource "aws_lambda_permission" "allow_sqs_event_to_trigger_lambda_event_handler" {
  statement_id = "AllowLambdaExecutionfromSQS"
  action       = "lambda:InvokeFuntion"

  principal     = "sqs.amazonaws.com"
  function_name = var.lambda_name
  source_arn    = var.sqs_arn
}

resource "aws_lambda_event_source_mapping" "sqs_event_queue_to_lambda_event_handler" {
  enabled    = var.is_sqs_to_lambda_integration_enabled
  batch_size = var.sqs_to_lambda_batch_size

  function_name    = var.lambda_name
  event_source_arn = var.sqs_arn
}
