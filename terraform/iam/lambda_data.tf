data "aws_iam_policy_document" "lambda_policy" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    # TODO @jonlee: Tighten scope
    resources = ["arn:aws:logs:*:*"]
  }
  statement {
    effect = "Allow"
    actions = [
      "sqs:DeleteMessage",
      "sqs:ReceiveMessage",
      "sqs:GetQueueAttributes"
    ]

    # TODO @jonlee: Tighten scope
    resources = ["*"]
  }
  statement {
    effect = "Allow"
    actions = [
      "sqs:ListQueues"
    ]

    # TODO @jonlee: Tighten scope
    resources = ["*"]
  }
}

data "aws_iam_policy_document" "assume_lambda_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}
