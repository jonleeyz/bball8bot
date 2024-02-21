output "queue_arn" {
  value = module.sqs.queue_arn
}

output "lambda_name" {
  value = module.lambda.function_name
}

output "lambda_invoke_arn" {
  value = module.lambda.invoke_arn
}