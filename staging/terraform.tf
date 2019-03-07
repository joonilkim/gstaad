terraform {
  backend "s3" {
    bucket = "gstaad"
    key    = "staging.tfstate"
    region = "ap-northeast-1"
  }
}
