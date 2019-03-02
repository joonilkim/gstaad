resource "aws_ecs_cluster" "_" {
  name = "${var.ns}"
}

output "ecs_cluster" {
  value = "${aws_ecs_cluster._.name}"
}
