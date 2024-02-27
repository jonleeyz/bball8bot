output "queue_arn" {
  value = module.sqs.arn
}

output "lambda_name" {
  value = module.lambda.function_name
}

output "lambda_invoke_arn" {
  value = module.lambda.invoke_arn
}
