resource "aws_alb" "_" {
  name    = "${var.ns}-post"

  subnets = [
    "${aws_subnet.pub_a.id}",
    "${aws_subnet.pub_c.id}",
  ]

  security_groups = [
    "${aws_security_group.lb.id}",
  ]
}

resource "aws_alb_listener" "_" {
  load_balancer_arn = "${aws_alb._.arn}"
  port              = "80"
  protocol          = "HTTP"

  default_action {
    target_group_arn = "${aws_alb_target_group.http.id}"
    type             = "forward"
  }
}

resource "aws_alb_target_group" "http" {
  deregistration_delay = 10

  health_check {
    healthy_threshold   = 2
    interval            = 60
    matcher             = "200,301"
    path                = "/ping"
    timeout             = 10
    unhealthy_threshold = 5
  }

  name     = "${var.ns}-http"
  port     = 80
  protocol = "HTTP"

  vpc_id = "${aws_vpc._.id}"
}