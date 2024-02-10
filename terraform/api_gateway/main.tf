resource "aws_iam_policy" "bball8bot_api_gateway_policy" {
  name   = "bball8botAPIGatewayPolicy"
  path   = "/"
  policy = data.aws_iam_policy_document.bball8bot_api_gateway_policy.json
}

resource "aws_iam_role" "bball8bot_api_gateway_role" {
  name                = "bball8botAPIGatewayRole"
  assume_role_policy  = data.aws_iam_policy_document.assume_bball8bot_api_gateway_role.json
  managed_policy_arns = [aws_iam_policy.bball8bot_api_gateway_policy.arn]
}
