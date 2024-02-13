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

resource "aws_api_gateway_rest_api" "bball8bot" {
  name = "bball8bot"
}

resource "aws_api_gateway_resource" "bball8bot" {
  rest_api_id = aws_api_gateway_rest_api.bball8bot.id
  parent_id   = aws_api_gateway_rest_api.bball8bot.root_resource_id

  path_part = "bot"
  # TODO @jonlee: change later to more appropriate path
}

resource "aws_api_gateway_method" "bball8bot" {
  rest_api_id = aws_api_gateway_rest_api.bball8bot.id
  resource_id = aws_api_gateway_resource.bball8bot.id

  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "bball8bot" {
  rest_api_id = aws_api_gateway_rest_api.bball8bot.id
  resource_id = aws_api_gateway_resource.bball8bot.id

  http_method             = aws_api_gateway_method.bball8bot.http_method
  integration_http_method = "POST"
  type                    = "AWS"
  uri                     = "arn:aws:apigateway:ap-southeast-1:sqs:path/${var.queue_name}"

  # TODO @jonlee: Understand why this is needed
  request_parameters = {
    "integration.request.header.Content-Type" = "'application/x-www-form-urlencoded'"
  }

  # TODO @jonlee: Understand why this is needed
  request_templates = {
    "application/json" = <<-EOF
      Action=SendMessage&MessageBody={
       "method": "$context.httpMethod",
       "body-json" : $input.json('$'),
       "queryParams": {
         #foreach($param in $input.params().querystring.keySet())
         "$param": "$util.escapeJavaScript($input.params().querystring.get($param))" #if($foreach.hasNext),#end
       #end
       },
       "pathParams": {
         #foreach($param in $input.params().path.keySet())
         "$param": "$util.escapeJavaScript($input.params().path.get($param))" #if($foreach.hasNext),#end
         #end
       }
      }"
      EOF
  }
}

resource "aws_api_gateway_method_response" "bball8bot_200" {
  rest_api_id = aws_api_gateway_rest_api.bball8bot.id
  resource_id = aws_api_gateway_resource.bball8bot.id

  http_method = aws_api_gateway_method.bball8bot.http_method
  status_code = 200

  response_models = {
    "application/json" : "Empty"
  }
}

resource "aws_api_gateway_integration_response" "bball8bot_200" {
  rest_api_id = aws_api_gateway_rest_api.bball8bot.id
  resource_id = aws_api_gateway_resource.bball8bot.id

  http_method = aws_api_gateway_method.bball8bot.http_method
  status_code = aws_api_gateway_method_response.bball8bot_200.status_code

  depends_on = [aws_api_gateway_integration.bball8bot]
}

resource "aws_api_gateway_deployment" "bball8bot_dev" {
  rest_api_id = aws_api_gateway_rest_api.bball8bot.id
  stage_name  = "dev"

  depends_on = [aws_api_gateway_integration.bball8bot]

  variables = {
    deployed_at = timestamp()
  }
}
