data "aws_iam_policy_document" "bball8bot_api_gateway_policy" {
  statement {
    effect = "Allow"
    actions = [
      "sqs:SendMessage"
    ]

    resources = [var.queue_arn]
  }
}

data "aws_iam_policy_document" "assume_bball8bot_api_gateway_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["apigateway.amazonaws.com"]
    }
  }
}
