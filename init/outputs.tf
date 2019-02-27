output "lb" {
  value = "${aws_alb._.arn}"
}

output "lb_tg_http_arn" {
  value = "${aws_alb_target_group.http.arn}"
}