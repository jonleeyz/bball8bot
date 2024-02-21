variable "queue_name" {
  description = "The name of the SQS queue"
  type        = string
}

variable "is_fifo" {
  description = "True if the SQS queue is FIFO, and false otherwise"
  type        = bool
}