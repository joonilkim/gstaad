resource "aws_cloudwatch_log_group" "_" {
  name = "${var.ns}/${var.service}"
  retention_in_days = 30
}
