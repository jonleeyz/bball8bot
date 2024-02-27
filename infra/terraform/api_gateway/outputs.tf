output "invoke_url" {
  value = aws_api_gateway_deployment.bball8bot_dev.invoke_url
}

output "resource_path" {
  value = aws_api_gateway_resource.bball8bot.path
}

output "resource_method" {
  value = aws_api_gateway_method.bball8bot.http_method
}
