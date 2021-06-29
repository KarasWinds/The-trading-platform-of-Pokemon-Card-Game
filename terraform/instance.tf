resource "aws_instance" "terraform_test" {
  ami           = var.AMI
  instance_type = var.INSTANCE_TYPE

  # VPC subnet
  subnet_id = aws_subnet.main-public.id

  # security group
  vpc_security_group_ids = [aws_security_group.allow-5000.id]

  # user data
  user_data_base64 = base64encode(local.user_data)
}
