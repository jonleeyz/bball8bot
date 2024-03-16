variable "queue_name" {
  description = "The name of the SQS event queue that the API Gateway will write to"
  type        = string
}

variable "aws_region" {
  description = "The AWS region that the API Gateway will be provisioned in"
  type        = string
}

variable "api_gateway_iam_role_arn" {
  description = "The ARN of the IAM role the API Gateway resource will assume during invocation"
  type        = string
}
