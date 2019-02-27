terraform {
  backend "s3" {
    bucket = "dev.gstaad.co"
    key    = "init.tfstate"
    region = "ap-northeast-1"
  }
}
