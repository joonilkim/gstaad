data "aws_iam_role" "ecs" {
    name = "${var.ecs_service_role}"
}

data "aws_iam_role" "task" {
    name = "${var.ecs_task_role}"
}