resource "aws_lambda_function" "bball8bot_event_handler" {
  function_name = var.lambda_name
  handler       = var.handler_function_name
  role          = var.lambda_iam_role_arn

  # Ref: https://docs.aws.amazon.com/lambda/latest/dg/runtimes-provided.html
  runtime  = "provided.al2023"
  filename = local.output_archive_path
}
