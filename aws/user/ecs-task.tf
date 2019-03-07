locals {
  ecr_uri = "${replace(aws_ecr_repository._.repository_url, "https://", "")}"
  ecr_envoy_uri = "${replace(aws_ecr_repository.envoy.repository_url, "https://", "")}"
}

resource "aws_ecs_task_definition" "_" {
  family       = "${var.ns}-${var.service}"
  network_mode = "awsvpc"

  container_definitions = <<-JSON
  [
    {
      "name": "${var.service}",
      "image": "${local.ecr_uri}:${var.image_tag}",
      "memory": 600,
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
    },
    {
      "name": "${var.service}-envoy",
      "image": "${local.ecr_envoy_uri}:${var.image_tag}",
      "memory": 350,
      "portMappings": [
        {
          "containerPort": 10000
        }
      ],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "${aws_cloudwatch_log_group._.name}",
          "awslogs-region": "${var.region}",
          "awslogs-stream-prefix": "envoy"
        }
      }
    }
  ]
  JSON

  task_role_arn = "${data.aws_iam_role.task.arn}"
}