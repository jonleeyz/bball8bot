locals {
  # golang executable binary has to be named bootstrap before being zipped
  # Ref: https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html#golang-package-mac-linux
  binary_path         = "bin/bootstrap"
  output_archive_path = "bin/main.zip"
}
