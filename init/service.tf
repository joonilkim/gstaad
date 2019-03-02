resource "aws_lb_listener_rule" "user" {
  listener_arn = "${aws_alb_listener._80.arn}"
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = "${aws_alb_target_group._.arn}"
  }

  condition {
    field  = "path-pattern"
    values = ["/users/*"]
  }
}

resource "aws_lb_listener_rule" "post" {
  listener_arn = "${aws_alb_listener._80.arn}"
  priority     = 102

  action {
    type             = "forward"
    target_group_arn = "${aws_alb_target_group._.arn}"
  }

  condition {
    field  = "path-pattern"
    values = ["/posts/*"]
  }
}