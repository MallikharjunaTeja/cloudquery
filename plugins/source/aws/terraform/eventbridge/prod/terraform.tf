terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "eventbridge"
    region = "us-east-1"
  }
}