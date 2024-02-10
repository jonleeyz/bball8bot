resource "aws_lambda_function" "bball8bot_event_handler" {
  function_name = var.lambda_name
  handler       = var.handler_function_name

  role = aws_iam_role.bball8bot_lambda_role.arn

  # Ref: https://docs.aws.amazon.com/lambda/latest/dg/runtimes-provided.html
  runtime  = "provided.al2023"
  filename = local.output_archive_path
}

resource "aws_iam_policy" "bball8bot_lambda_policy" {
  name   = "bball8botLambdaPolicy"
  path   = "/"
  policy = data.aws_iam_policy_document.bball8bot_lambda_policy.json
}

resource "aws_iam_role" "bball8bot_lambda_role" {
  name                = "bball8botLambdaRole"
  assume_role_policy  = data.aws_iam_policy_document.assume_bball8bot_lambda_role.json
  managed_policy_arns = [aws_iam_policy.bball8bot_lambda_policy.arn]
}
