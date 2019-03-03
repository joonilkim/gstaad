variable "ns" {
  default = "gstaad"
  description = "A namespace"
}

variable "domain" {
  default = "gstaad.co"
}

variable "region" {
  default = "ap-northeast-1"
}

variable "stage" {
  default = "dev"
}

variable "bastion_ami" {
  default = "ami-07ad4b1c3af1ea214"
}
variable "bastion_key" {
  default = ""
}
variable "bastion_on" {
  default = false
}