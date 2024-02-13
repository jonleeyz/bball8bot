variable "queue_arn" {
  description = "The AWS resource name of the SQS event queue that the API Gateway will write to"
  type        = string
}

variable "queue_name" {
  description = "The name of the SQS event queue that the API Gateway will write to"
  type        = string
}
