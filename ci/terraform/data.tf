data "aws_iam_policy_document" "terraform_state_management_policy" {
  statement {
    effect  = "Allow"
    actions = ["s3:ListBucket"]

    // change
    resources = [aws_s3_bucket.terraform_state_store.arn]
  }
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject",
    ]

    // change
    resources = ["${aws_s3_bucket.terraform_state_store.arn}/${local.workspace_bucket_key}"]
  }
  statement {
    effect = "Allow"
    actions = [
      "dynamodb:DescribeTable",
      "dynamodb:GetItem",
      "dynamodb:PutItem",
      "dynamodb:DeleteItem"
    ]

    resources = [
      // change
      aws_dynamodb_table.terraform_state_lock.arn
    ]
  }
}

data "aws_iam_policy_document" "workspace_infra_policy" {
  statement {
    effect = "Allow"
    actions = [
      // change
      "s3:*",
      "dynamodb:*",
      "iam:*"
    ]

    resources = ["*"]
  }
}

data "aws_iam_policy_document" "assume_account_wide_terraform_support_role" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]

    condition {
      test     = "StringLike"
      // change
      variable = "${local.github_oidc_provider_url}:sub"
      values = [
        "repo:jonleeyz/terraform-backend:*",
        // change
      ]
    }

    principals {
      type        = "Federated"
      // change
      identifiers = ["arn:aws:iam::574182556674:oidc-provider/${local.github_oidc_provider_url}"]
    }
  }
}
