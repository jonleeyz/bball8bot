variable "queue_arn" {
  description = "The AWS resource name of the SQS event queue that the API Gateway will write to"
  type        = string
}

variable "queue_name" {
  description = "The name of the SQS event queue that the API Gateway will write to"
  type        = string
}

variable "aws_region" {
  description = "The AWS region that the API Gateway will be provisioned in"
  type        = string
}
