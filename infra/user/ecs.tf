data "aws_ecs_cluster" "_" {
  cluster_name = "${var.ecs_cluster}"
}

resource "aws_ecs_service" "_" {
  cluster                            = "${data.aws_ecs_cluster._.id}"
  deployment_minimum_healthy_percent = 0
  desired_count                      = 1
  name                               = "${var.service}"
  task_definition                    = "${aws_ecs_task_definition._.arn}"

  network_configuration {
    subnets         = ["${var.priv_subnets}"]
    security_groups = ["${var.secg_svc}"]
  }

  service_registries {
    registry_arn     = "${aws_service_discovery_service._.arn}"
  }

  lifecycle {
    # ecs update makes a new task definition
    ignore_changes = ["task_definition"]
  }
}