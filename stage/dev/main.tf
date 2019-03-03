variable "ns" {
  default = "gstaad"
}

variable "domain" {
  default = "gstaad.co"
}

variable "region" {
  default = "ap-northeast-1"
}

locals {
  stage = "dev"
}

module "init" {
  source          = "../../infra/init"

  region          = "${var.region}"
  stage           = "${local.stage}"
  domain          = "${var.domain}"
  ns              = "${var.ns}"
}

module "post" {
  source          = "../../infra/post"

  region          = "${var.region}"
  stage           = "${local.stage}"
  ns              = "${var.ns}"
  service         = "${var.ns}-post"
  image_tag       = "develop"  # should be same as branch name

  vpc_id          = "${module.init.vpc_id}"
  pub_subnets     = "${module.init.pub_subnets}"
  priv_subnets    = "${module.init.priv_subnets}"

  secg_lb         = "${module.init.secg_lb}"
  secg_svc        = "${module.init.secg_svc}"
  secg_db         = "${module.init.secg_db}"

  lb_arn          = "${module.init.lb_arn}"
  lb_tg_arn       = "${module.init.lb_tg_arn}"

  ecs_service_role  = "${module.init.ecs_service_role}"
  ecs_task_role     = "${module.init.ecs_task_role}"
  ecs_instance_role = "${module.init.ecs_instance_role}"

  ecs_cluster     = "${module.init.ecs_cluster}"
  service_discovery_arn = "${module.init.service_discovery_arn}"
}

module "user" {
  source          = "../../infra/user"

  region          = "${var.region}"
  stage           = "${local.stage}"
  ns              = "${var.ns}"
  service         = "${var.ns}-user"
  image_tag       = "develop"  # should be same as branch name

  vpc_id          = "${module.init.vpc_id}"
  pub_subnets     = "${module.init.pub_subnets}"
  priv_subnets    = "${module.init.priv_subnets}"

  secg_lb         = "${module.init.secg_lb}"
  secg_svc        = "${module.init.secg_svc}"
  secg_db         = "${module.init.secg_db}"

  lb_arn          = "${module.init.lb_arn}"
  lb_tg_arn       = "${module.init.lb_tg_arn}"

  ecs_service_role  = "${module.init.ecs_service_role}"
  ecs_task_role     = "${module.init.ecs_task_role}"
  ecs_instance_role = "${module.init.ecs_instance_role}"

  ecs_cluster     = "${module.init.ecs_cluster}"
  service_discovery_arn = "${module.init.service_discovery_arn}"
}

output "lb_dns_name" {
  value = "${module.init.lb_dns_name}"
}
