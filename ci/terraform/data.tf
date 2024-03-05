data "aws_iam_policy_document" "terraform_state_management_policy" {
  statement {
    effect  = "Allow"
    actions = ["s3:ListBucket"]

    resources = [local.state_store_bucket_arn]
  }
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject",
    ]

    resources = ["${local.state_store_bucket_arn}/${local.workspace_bucket_key}"]
  }
  statement {
    effect = "Allow"
    actions = [
      "dynamodb:DescribeTable",
      "dynamodb:GetItem",
      "dynamodb:PutItem",
      "dynamodb:DeleteItem"
    ]

    resources = [local.state_lock_table_arn]
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
      variable = "${local.github_oidc_provider_url}:sub"
      values = [
        "repo:jonleeyz/bball8bot:*",
      ]
    }

    principals {
      type        = "Federated"
      identifiers = ["arn:aws:iam::574182556674:oidc-provider/${local.github_oidc_provider_url}"]
    }
  }
}
