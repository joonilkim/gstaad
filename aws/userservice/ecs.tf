data "aws_ecs_cluster" "_" {
  cluster_name = "${var.ecs_cluster}"
}

resource "aws_ecs_service" "_" {
  cluster                            = "${data.aws_ecs_cluster._.id}"
  launch_type                        = "FARGATE"
  deployment_minimum_healthy_percent = 50
  desired_count                      = 1
  name                               = "${var.service}"
  task_definition                    = "${aws_ecs_task_definition._.arn}"
  health_check_grace_period_seconds  = 20

  load_balancer {
    container_name   = "${var.service}-gw"
    container_port   = 80
    target_group_arn = "${var.lb_tg_arn}"
  }

  network_configuration {
    subnets          = ["${var.priv_subnets}"]
    security_groups  = ["${var.secg_svc}"]
    assign_public_ip = "true"
  }

  service_registries {
    registry_arn     = "${aws_service_discovery_service._.arn}"
  }

  lifecycle {
    # ecs update makes a new task definition
    ignore_changes = ["task_definition"]
  }
}