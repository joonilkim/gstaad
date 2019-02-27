resource "aws_ecs_task_definition" "_" {
  family = "${var.ns}-${var.service}"

  container_definitions = <<-JSON
  [
    {
      "image": "${replace(aws_ecr_repository._.repository_url, "https://", "")}:${var.image_tag}",
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "${var.ns}",
          "awslogs-region": "${var.region}",
          "awslogs-stream-prefix": "${var.service}"
        }
      },
      "memory": 300,
      "name": "${var.ns}-${var.service}",
      "portMappings": [
        {
          "containerPort": 9000,
          "protocol": "tcp"
        }
      ]
    }
  ]
  JSON

  task_role_arn = "${aws_iam_role.task.arn}"
}