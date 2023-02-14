terraform {
  required_version = ">= 1.3.6"

  backend "s3" {
    bucket = "golang-test-haroldo"
    key    = "terraform/bucket/key"
    region = "us-east-1"
  }
}
