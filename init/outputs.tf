
output "vpc_id" {
  value = "${aws_vpc._.id}"
}

output "subnet_privs" {
  value = ["${aws_subnet.priv_a.id}", "${aws_subnet.priv_c.id}"]
}

output "subnet_pubs" {
  value = ["${aws_subnet.pub_a.id}", "${aws_subnet.pub_c.id}"]
}

output "secg_lb_id" {
  value = "${aws_security_group.lb.id}"
}

output "lb" {
  value = "${aws_alb._.arn}"
}

output "lb_tg_http_arn" {
  value = "${aws_alb_target_group.http.arn}"
}
