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
