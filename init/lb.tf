resource "aws_alb" "_" {
  name    = "${var.ns}"

  subnets = [
    "${aws_subnet.priv_a.id}",
    "${aws_subnet.priv_c.id}",
  ]

  security_groups = [
    "${aws_security_group.lb.id}",
  ]
}

resource "aws_alb_listener" "_80" {
  load_balancer_arn = "${aws_alb._.arn}"
  port              = "80"
  protocol          = "HTTP"

  default_action {
    target_group_arn = "${aws_alb_target_group._.id}"
    type             = "forward"
  }
}

resource "aws_alb_target_group" "_" {
  name                  = "${var.ns}-http"
  port                  = 9000
  protocol              = "HTTP"
  vpc_id                = "${aws_vpc._.id}"
  target_type           = "ip"
  deregistration_delay  = 10

  health_check {
    healthy_threshold   = 2
    interval            = 10
    matcher             = "200,301"
    path                = "/ping"
    timeout             = 3
    unhealthy_threshold = 3
  }
}

output "lb_arn" {
  value = "${aws_alb._.arn}"
}

output "lb_tg_arn" {
  value = "${aws_alb_target_group._.arn}"
}

output "lb_dns_name" {
  value = "${aws_alb._.dns_name}"
}