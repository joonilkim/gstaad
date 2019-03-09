resource "aws_iam_role" "ecs" {
  assume_role_policy = <<-JSON
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Principal": {
          "Service": "ecs.amazonaws.com"
        }
      }
    ]
  }
  JSON

  name = "${var.ns}-ecs-service-role"
}

resource "aws_iam_role_policy_attachment" "ecs" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceRole"
  role       = "${aws_iam_role.ecs.id}"
}


resource "aws_iam_role" "task" {
  assume_role_policy = <<-JSON
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Principal": {
          "Service": "ecs-tasks.amazonaws.com"
        }
      }
    ]
  }
  JSON

  name = "${var.ns}-ecs-task-role"
}

resource "aws_iam_role_policy" "task" {
  name = "${var.ns}-ecs-task-policy"
  policy = "${data.aws_iam_policy_document.task.json}"
  role = "${aws_iam_role.task.id}"
}
data "aws_iam_policy_document" "task" {
  statement {
    effect    = "Allow"
    actions   = [
      "cloudwatch:*",
      "events:*",
      "logs:*",
      "ecr:GetAuthorizationToken",
      "ecr:BatchCheckLayerAvailability",
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "cognito-idp:*",
    ]
    resources = ["*"]
  }
  statement {
    effect    = "Deny"
    actions   = [
      "cognito-idp:*IdentityProvider*",
      "cognito-idp:*UserPool*",
      "cognito-idp:*ResourceServer*",
    ]
    resources = ["*"]
  }
}


## ecs instance ##

resource "aws_iam_role" "ec2" {
  assume_role_policy = <<-JSON
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Principal": {
          "Service": "ec2.amazonaws.com"
        }
      }
    ]
  }
  JSON

  name = "${var.ns}-ecs-instance-role"
}

resource "aws_iam_role_policy_attachment" "ec2" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
  role       = "${aws_iam_role.ec2.id}"
}

output "ecs_instance_role" {
  value = "${aws_iam_role.ec2.name}"
}

output "ecs_task_role" {
  value = "${aws_iam_role.task.name}"
}

output "ecs_service_role" {
  value = "${aws_iam_role.ecs.name}"
}