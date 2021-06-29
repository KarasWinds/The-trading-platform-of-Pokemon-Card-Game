terraform {
  backend "s3" {
    bucket = "minchi-terraform"
    key = "terraform/remote-state"
    region = "us-west-2"
  }
}