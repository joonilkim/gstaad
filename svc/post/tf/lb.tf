data "aws_lb_listener" "http" {
  load_balancer_arn = "${var.lb_arn}"
  port = 80
}

data "aws_alb_target_group" "http" {
  arn = "${var.lb_tg_http_arn}"
}

resource "aws_lb_listener_rule" "_" {
  listener_arn = "${data.aws_lb_listener.http.arn}"
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = "${data.aws_alb_target_group.http.arn}"
  }

  condition {
    field  = "path-pattern"
    values = "${var.path_patterns}"
  }
}
