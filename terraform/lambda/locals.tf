locals {
  # Ref: https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html#golang-package-mac-linux
  # Note: golang executable binary has to be named bootstrap and output to this specific directory
  binary_path         = "bin/bootstrap"
  output_archive_path = "bin/main.zip"
}
