output "queue_arn" {
  value = aws_sqs_queue.bball8bot_event_queue.arn
}

output "queue_name" {
  value = aws_sqs_queue.bball8bot_event_queue.name
}
