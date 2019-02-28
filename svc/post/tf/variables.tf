variable "region" {
  description = "A AWS region"
}

variable "stage" {
  description = "A stage name"
}

variable "ns" {
  default = "example"
}

variable "service" {
  default = "post"
}

variable "path_patterns" {
  default = ["/post/*"]
}

variable "vpc_id" {
  description = "vpc_id"
}

variable "subnet_privs" {
  description = "private subnet ids"
  type = "list"
}

variable "subnet_pubs" {
  description = "public subnet ids"
  type = "list"
}

variable "secg_lb_id" {
  description = "LB security group id"
}

variable "lb_arn" {
  description = "A LB arn"
}

variable "lb_tg_http_arn" {
  description = "A LB target group arn"
}

variable "image_tag" {
  description = "A image tag"
}