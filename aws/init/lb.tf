resource "aws_alb" "_" {
  name            = "${var.ns}"

  subnets = [
    "${aws_subnet.pub_d.id}",
    "${aws_subnet.pub_c.id}",
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
    target_group_arn = "${aws_alb_target_group.userservice.id}"
    type             = "forward"
  }
}

resource "aws_alb_target_group" "userservice" {
  name                  = "${var.ns}-userservice"
  port                  = 80
  protocol              = "HTTP"
  vpc_id                = "${aws_vpc._.id}"
  target_type           = "ip"
  deregistration_delay  = 10

  health_check {
    healthy_threshold   = 2
    interval            = 10
    matcher             = 200
    path                = "/health"
    timeout             = 3
    unhealthy_threshold = 3
  }

  lifecycle {
    create_before_destroy = true
  }

  depends_on = ["aws_alb._"]
}

resource "aws_lb_listener_rule" "postservice" {
  listener_arn = "${aws_alb_listener._80.arn}"
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = "${aws_alb_target_group.postservice.arn}"
  }

  condition {
    field  = "path-pattern"
    values = ["/posts*"]
  }
}

resource "aws_alb_target_group" "postservice" {
  name                  = "${var.ns}-postservice"
  port                  = 80
  protocol              = "HTTP"
  vpc_id                = "${aws_vpc._.id}"
  target_type           = "ip"
  deregistration_delay  = 10

  health_check {
    healthy_threshold   = 2
    interval            = 10
    matcher             = 200
    path                = "/health"
    timeout             = 3
    unhealthy_threshold = 3
  }

  lifecycle {
    create_before_destroy = true
  }

  depends_on = ["aws_alb._"]
}


output "lb_arn" {
  value = "${aws_alb_listener._80.load_balancer_arn}"
}

output "userservice_tg" {
  value = "${aws_alb_target_group.userservice.arn}"
}

output "postservice_tg" {
  value = "${aws_alb_target_group.postservice.arn}"
}

output "lb_dns_name" {
  value = "${aws_alb._.dns_name}"
}