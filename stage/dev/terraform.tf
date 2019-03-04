terraform {
  backend "s3" {
    bucket = "dev.gstaad.co"
    key    = "dev.tfstate"
    region = "ap-northeast-1"
  }
}
