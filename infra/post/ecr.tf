resource "aws_ecr_repository" "_" {
  name = "${replace(var.service, "-", "/")}"
}
