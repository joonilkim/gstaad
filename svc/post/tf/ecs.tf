resource "aws_ecs_cluster" "_" {
  name = "${var.ns}"
}

resource "aws_ecs_service" "_" {
  cluster                            = "${aws_ecs_cluster._.id}"
  deployment_minimum_healthy_percent = 0
  desired_count                      = 1
  iam_role                           = "${aws_iam_role.ecs.arn}"
  name                               = "${var.ns}-${var.service}"

  load_balancer {
    container_name   = "${var.ns}-${var.service}"
    container_port   = "9000"
    target_group_arn = "${data.aws_alb_target_group.http.arn}"
  }

  task_definition = "${aws_ecs_task_definition._.arn}"
}