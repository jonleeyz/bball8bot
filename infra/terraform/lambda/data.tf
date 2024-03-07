# Creates a zip file at the output path from the specified source file path.
data "archive_file" "zipped_binary_for_deploy" {
  type        = "zip"
  source_file = local.binary_path
  output_path = local.output_archive_path
}
