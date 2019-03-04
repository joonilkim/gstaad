variable "region" {}
variable "stage" {}

variable "ns" {}
variable "service" {}

variable "vpc_id" {}
variable "pub_subnets" {
  type = "list"
}
variable "priv_subnets" {
  type = "list"
}

variable "secg" {}

variable "lb_arn" {}
variable "lb_tg_arn" {}

variable "ecs_service_role" {}
variable "ecs_task_role" {}
variable "ecs_instance_role" {}

variable "ecs_cluster" {}
variable "image_tag" {}