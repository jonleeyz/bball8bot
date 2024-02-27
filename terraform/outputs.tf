output "queue_arn" {
  value = module.sqs.arn
}

output "lambda_name" {
  value = module.lambda.function_name
}

output "api_invoke_url" {
  value = module.api_gateway.invoke_url
}

output "api_resource_path" {
  value = module.api_gateway.resource_path
}

output "api_resource_method" {
  value = module.api_gateway.resource_method
}
