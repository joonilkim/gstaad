## ecs ##

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

  name = "${var.ns}-ecs"
}

resource "aws_iam_role_policy_attachment" "ecs" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceRole"
  role       = "${aws_iam_role.ecs.id}"
}


## ecs task ##

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

  name = "${var.ns}-task"
}

resource "aws_iam_role_policy" "task" {
  name = "${var.ns}_${var.service}_task"
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

  name = "${var.ns}-${var.service}-ec2"
}

resource "aws_iam_role_policy_attachment" "ec2" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
  role       = "${aws_iam_role.ec2.id}"
}