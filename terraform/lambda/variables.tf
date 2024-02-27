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

variable "sqs_arn" {
  description = "The ARN of the SQS event queue that the Lambda will consume events from"
  type        = string
}

variable "is_sqs_to_lambda_integration_enabled" {
  description = "True if the SQS -> Lambda integration is enabled, and false otherwise"
  type        = bool
  default     = false
}

variable "sqs_to_lambda_batch_size" {
  description = "The largest number of records that the Lambda event handler will retrieve from the SQS event queue at time of invocation"
  type        = number
  default     = 1
}
