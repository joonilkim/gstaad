variable "ns" {
  default = "gstaad"
}
variable "domain" {
  default = "gstaad.io"
}
variable "region" {
  default = "ap-northeast-1"
}

variable "bastion_key" {
  default = ""
}
variable "bastion_on" {
  default = false
}

locals {
  stage = "dev"
}

module "init" {
  source          = "../aws/init"

  region          = "${var.region}"
  stage           = "${local.stage}"
  domain          = "${var.domain}"
  ns              = "${var.ns}"

  bastion_on     = "${var.bastion_on}"
  bastion_key     = "${var.bastion_key}"
}

module "postservice" {
  source          = "../aws/postservice"

  region          = "${var.region}"
  stage           = "${local.stage}"
  ns              = "${var.ns}"
  service         = "postservice"

  vpc_id          = "${module.init.vpc_id}"
  pub_subnets     = "${module.init.pub_subnets}"
  priv_subnets    = "${module.init.priv_subnets}"

  secg_lb         = "${module.init.secg_lb}"
  secg_svc        = "${module.init.secg_svc}"
  secg_db         = "${module.init.secg_db}"

  lb_tg_arn       = "${module.init.postservice_tg}"

  ecs_service_role  = "${module.init.ecs_service_role}"
  ecs_task_role     = "${module.init.ecs_task_role}"
  ecs_instance_role = "${module.init.ecs_instance_role}"

  ecs_cluster          = "${module.init.ecs_cluster}"
  service_discovery_ns = "${module.init.service_discovery_ns_id}"
}

module "userservice" {
  source          = "../aws/userservice"

  region          = "${var.region}"
  stage           = "${local.stage}"
  ns              = "${var.ns}"
  service         = "userservice"

  vpc_id          = "${module.init.vpc_id}"
  pub_subnets     = "${module.init.pub_subnets}"
  priv_subnets    = "${module.init.priv_subnets}"

  secg_lb         = "${module.init.secg_lb}"
  secg_svc        = "${module.init.secg_svc}"
  secg_db         = "${module.init.secg_db}"

  lb_tg_arn       = "${module.init.userservice_tg}"

  cognito_userpool_id  = "${module.init.cognito_userpool_id}"
  cognito_client_id    = "${module.init.cognito_client_id}"

  ecs_service_role     = "${module.init.ecs_service_role}"
  ecs_task_role        = "${module.init.ecs_task_role}"
  ecs_instance_role    = "${module.init.ecs_instance_role}"

  ecs_cluster          = "${module.init.ecs_cluster}"
  service_discovery_ns = "${module.init.service_discovery_ns_id}"
}

output "lb_dns_name" {
  value = "${module.init.lb_dns_name}"
}