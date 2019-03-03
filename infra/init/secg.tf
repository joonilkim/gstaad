resource "aws_security_group" "lb" {
  name   = "${var.ns}-lb"
  vpc_id = "${aws_vpc._.id}"

  ingress = {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress = {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}


resource "aws_security_group" "svc" {
  name   = "${var.ns}-svc"
  vpc_id = "${aws_vpc._.id}"

  ingress = {
    from_port = 0
    to_port   = 65535
    protocol  = "tcp"

    security_groups = [
      "${aws_security_group.lb.id}",
      "${aws_security_group.bastion.id}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "db" {
  name   = "${var.ns}-db"
  vpc_id = "${aws_vpc._.id}"

  ingress = {
    from_port = 5432
    to_port   = 5432
    protocol  = "tcp"

    security_groups = [
      "${aws_security_group.svc.id}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "bastion" {
  name   = "${var.ns}-bastion"
  vpc_id = "${aws_vpc._.id}"

  ingress = {
    from_port = 22
    to_port   = 22
    protocol  = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

output "secg_lb" {
  value = "${aws_security_group.lb.id}"
}

output "secg_svc" {
  value = "${aws_security_group.svc.id}"
}

output "secg_db" {
  value = "${aws_security_group.db.id}"
}

output "secg_bastion" {
  value = "${aws_security_group.bastion.id}"
}