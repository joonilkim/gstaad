resource "aws_cloudwatch_log_group" "_" {
  name = "${replace(var.service, "-", "/")}"
}
