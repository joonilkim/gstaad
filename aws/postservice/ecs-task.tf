locals {
  ecr_uri = "${replace(aws_ecr_repository._.repository_url, "https://", "")}"
  gw_ecr = "${replace(aws_ecr_repository.gw.repository_url, "https://", "")}"
}

resource "aws_ecs_task_definition" "_" {
  family       = "${var.ns}-${var.service}"
  network_mode = "awsvpc"

  container_definitions = <<-JSON
  [
    {
      "name": "${var.service}",
      "image": "${local.ecr_uri}:latest",
      "memoryReservation": 128,
      "environment": [
        {
          "name": "APP_ENV",
          "value": "production"
        },
        {
          "name": "PORT",
          "value": "8080"
        },
        {
          "name": "POSTSERVICE",
          "value": "80"
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
      "name": "${var.service}-gw",
      "image": "${local.gw_ecr}:latest",
      "memoryReservation": 128,
      "portMappings": [
        {
          "containerPort": 80
        }
      ],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "${aws_cloudwatch_log_group._.name}",
          "awslogs-region": "${var.region}",
          "awslogs-stream-prefix": "gw"
        }
      }
    }
  ]
  JSON

  task_role_arn = "${data.aws_iam_role.task.arn}"
}
