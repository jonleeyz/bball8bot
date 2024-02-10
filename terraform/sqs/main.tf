resource "aws_sqs_queue" "bball8bot_event_queue" {
  name                       = var.queue_name
  delay_seconds              = 0
  visibility_timeout_seconds = 30
  max_message_size           = 1024
  message_retention_seconds  = 60
  receive_wait_time_seconds  = 2
  sqs_managed_sse_enabled    = true
}
