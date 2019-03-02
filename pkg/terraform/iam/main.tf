variable "role" { 
    description = "role/aws-service-role/ecs.amazonaws.com/AWSServiceRoleForECS"
}

data "aws_caller_identity" "_" {}

output "arn" {
  value = "arn:aws:iam::${data.aws_caller_identity._.account_id}:${var.role_uri}"
}

