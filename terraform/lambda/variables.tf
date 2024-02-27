variable "lambda_name" {
  description = "The name of the Lambda function"
  type        = string
}

variable "handler_function_name" {
  description = "The name of the golang handler function"
  type        = string
}

variable "lambda_iam_role_arn" {
  description = "The ARN of the IAM role the Lambda will assume during execution"
  type        = string
}
