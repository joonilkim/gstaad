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
    target_group_arn = "${aws_alb_target_group.gw.id}"
    type             = "forward"
  }
}

resource "aws_alb_target_group" "gw" {
  name                  = "${var.ns}-gw"
  port                  = 8080 
  protocol              = "HTTP"
  vpc_id                = "${aws_vpc._.id}"
  target_type           = "ip"
  deregistration_delay  = 10

  health_check {
    healthy_threshold   = 2
    interval            = 10
    matcher             = 200
    path                = "/ping"
    timeout             = 3
    unhealthy_threshold = 3
  }

  lifecycle {
    create_before_destroy = true
  }
}

output "lb_arn" {
  value = "${aws_alb_listener._80.load_balancer_arn}"
}

output "gw_tg_arn" {
  value = "${aws_alb_target_group.gw.arn}"
}

output "lb_dns_name" {
  value = "${aws_alb._.dns_name}"
}