provider "aws" {
  region = "${var.region}"
}

provider "null" {
  version = "~> 1.0"
}

provider "external" {
  version = "~> 1.0"
}
