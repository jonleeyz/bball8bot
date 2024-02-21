output "queue_arn" {
  value = aws_sqs_queue.bball8bot_events.arn
}

output "queue_name" {
  value = aws_sqs_queue.bball8bot_events.name
}
