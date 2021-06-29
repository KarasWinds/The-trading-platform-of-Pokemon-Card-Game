# variable "AWS_ACCESS_KEY" {

# }

# variable "AWS_SECRET_KEY" {

# }

variable "AWS_REGION" {
  default = "us-west-2"
}

variable "INSTANCE_TYPE" {
  default = "t2.micro"
}

variable "AMI" {
  default = "ami-0721c9af7b9b75114"
}