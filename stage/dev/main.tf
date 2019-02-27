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

  lb_arn          = "${module.init.lb}"
  lb_tg_http_arn  = "${module.init.lb_tg_http_arn}"
  path_patterns   = ["/post/*"]
  image_tag       = "${var.stage}"
}
