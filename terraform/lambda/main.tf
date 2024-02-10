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
