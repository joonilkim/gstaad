resource "aws_ecr_repository" "_" {
  name = "${var.ns}/${var.service}"
}

resource "aws_ecr_repository" "gw" {
  name = "${var.ns}/${var.service}-gw"
}

