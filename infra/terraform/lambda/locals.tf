locals {
  # Ref: https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html#golang-package-mac-linux
  # Note: golang executable binary has to be named bootstrap and output to this specific directory
  binary_path         = "bin/bootstrap"
  output_archive_path = "bin/main.zip"

  # Ref: https://docs.aws.amazon.com/systems-manager/latest/userguide/ps-integration-lambda-extensions.html#ps-integration-lambda-extensions-add
  # Note: Region-specific
  parameters_and_secrets_extension_layer_arn = "arn:aws:lambda:ap-southeast-1:044395824272:layer:AWS-Parameters-and-Secrets-Lambda-Extension:11"
}
