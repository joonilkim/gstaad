resource "aws_security_group" "ec2" {
  name   = "${var.ns}-${var.service}"
  vpc_id = "${var.vpc_id}"

  ingress = {
    from_port = 0
    to_port   = 65535
    protocol  = "tcp"

    security_groups = [
      "${var.secg_lb_id}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
