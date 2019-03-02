output "all" {
  value = [
    "${aws_ecs_cluster._.name}",
    "${aws_alb._.arn}",
    "${aws_alb_target_group._.arn}",
    "${aws_alb._.dns_name}",
    "${aws_iam_role.ec2.arn}",
    "${aws_iam_role.ecs.arn}",
    "${aws_iam_role.task.arn}",
    "${aws_security_group.lb.id}",
    "${aws_security_group.svc.id}",
    "${aws_security_group.db.id}",
    "${aws_service_discovery_service._.arn}",
    "${aws_subnet.pub_a.id}",
    "${aws_subnet.pub_c.id}",
    "${aws_subnet.priv_a.id}",
    "${aws_subnet.priv_c.id}",
    "${aws_vpc._.id}",
  ]
}
