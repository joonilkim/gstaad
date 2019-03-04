locals {
  ecr_uri = "${replace(aws_ecr_repository._.repository_url, "https://", "")}"
}

resource "aws_ecs_task_definition" "_" {
  family       = "${var.ns}-${var.service}"
  network_mode = "awsvpc"

  container_definitions = <<-JSON
  [
    {
      "name": "${var.ns}-${var.service}",
      "image": "${local.ecr_uri}:${var.image_tag}",
      "memory": 900,
      "portMappings": [
        {
          "containerPort": 8080
        }
      ],
      "environment": [
        {
          "name": "APP_ENV",
          "value": "production"
        }
      ],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "${aws_cloudwatch_log_group._.name}",
          "awslogs-region": "${var.region}",
          "awslogs-stream-prefix": "app"
        }
      }
    }
  ]
  JSON

  task_role_arn = "${data.aws_iam_role.task.arn}"
}