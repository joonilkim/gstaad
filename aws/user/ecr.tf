resource "aws_ecr_repository" "_" {
  name = "${var.ns}/${var.service}"
}

resource "aws_ecr_repository" "envoy" {
  name = "${var.ns}/${var.service}/envoy"
}

