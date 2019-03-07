resource "aws_ecr_repository" "_" {
  name = "${var.ns}/${var.service}"
}
