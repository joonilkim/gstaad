module "init" {
  source          = "../../init"

  region          = "${var.region}"
  stage           = "${var.stage}"
  domain          = "${var.domain}"
  ns              = "${var.ns}"
}

module "post" {
  source          = "../../svc/post/tf"

  region          = "${var.region}"
  stage           = "${var.stage}"
  ns              = "${var.ns}"
  service         = "post"

  vpc_id          = "${module.init.vpc_id}"
  subnet_privs    = "${module.init.subnet_privs}"
  subnet_pubs     = "${module.init.subnet_pubs}"
  secg_lb_id      = "${module.init.secg_lb_id}"

  lb_arn          = "${module.init.lb}"
  lb_tg_http_arn  = "${module.init.lb_tg_http_arn}"
  path_patterns   = ["/post/*"]

  image_tag       = "${var.stage}"
}
