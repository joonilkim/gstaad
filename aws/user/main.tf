data "aws_caller_identity" "_" {}

variable "depends_on" {
    default = []
}

resource "null_resource" "depends_on" {
    triggers = {
        value = "${join(",", var.depends_on)}"
    } 
}

