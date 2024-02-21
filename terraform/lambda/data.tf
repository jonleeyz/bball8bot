data "aws_iam_policy_document" "bball8bot_lambda_policy" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = ["arn:aws:logs:*:*"]
  }
  statement {
    effect = "Allow"
    actions = [
      "sqs:DeleteMessage",
      "sqs:ReceiveMessage",
      "sqs:GetQueueAttributes"
    ]

    resources = ["*"]
  }
  statement {
    effect = "Allow"
    actions = [
      "sqs:ListQueues"
    ]

    resources = ["*"]
  }
}

data "aws_iam_policy_document" "assume_bball8bot_lambda_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

# Creates a zip file at the output path from the specified source file path.
data "archive_file" "zipped_binary_for_deploy" {
  type        = "zip"
  source_file = local.binary_path
  output_path = local.output_archive_path
}
