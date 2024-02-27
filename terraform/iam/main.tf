resource "aws_iam_policy" "lambda" {
  name   = "bball8botLambdaPolicy"
  path   = "/"
  policy = data.aws_iam_policy_document.lambda_policy.json
}

resource "aws_iam_role" "lambda" {
  name                = "bball8botLambdaRole"
  assume_role_policy  = data.aws_iam_policy_document.assume_lambda_role.json
  managed_policy_arns = [aws_iam_policy.lambda.arn]
}

resource "aws_iam_policy" "api_gateway" {
  name   = "bball8botAPIGatewayPolicy"
  path   = "/"
  policy = data.aws_iam_policy_document.api_gateway_policy.json
}

resource "aws_iam_role" "api_gateway" {
  name                = "bball8botAPIGatewayRole"
  assume_role_policy  = data.aws_iam_policy_document.assume_api_gateway_role.json
  managed_policy_arns = [aws_iam_policy.api_gateway.arn]
}
