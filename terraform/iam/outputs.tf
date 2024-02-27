output "lambda_iam_role_arn" {
  value = aws_iam_role.lambda.arn
}

output "api_gateway_iam_role_arn" {
  value = aws_iam_role.api_gateway.arn
}
