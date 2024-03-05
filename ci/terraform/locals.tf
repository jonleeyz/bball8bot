locals {
  github_oidc_provider_url = "token.actions.githubusercontent.com"
  state_store_bucket_arn = "arn:aws:s3:::jl-terraform-remote-state-store"
  infra_workspace_bucket_key     = "bball8bot/infra/terraform.tfstate"
  ci_workspace_bucket_key     = "bball8bot/ci/terraform.tfstate"
  state_lock_table_arn = "arn:aws:dynamodb:ap-southeast-1:574182556674:table/terraform_state_lock"
}
