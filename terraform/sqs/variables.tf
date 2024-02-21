variable "queue_name" {
  description = "The name of the SQS queue"
  type        = string
}

variable "is_fifo" {
  description = "True if the SQS queue is FIFO, and false otherwise"
  type        = bool
}

variable "is_sqs_to_lambda_integration_enabled" {
  description = "True if the SQS -> Lambda integration is enabled, and false otherwise"
  type        = bool
  default     = false
}

variable "lambda_name" {
  description = "The name of the Lambda function"
  type        = string
}

variable "sqs_to_lambda_batch_size" {
  description = "The largest number of records that the Lambda event handler will retrieve from the SQS event queue at time of invocation"
  type        = number
  default     = 1
}
