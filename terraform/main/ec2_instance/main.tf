########### SSH Key Generation ##############
resource "tls_private_key" "rsa_key" {
  algorithm = "RSA"
  rsa_bits  = 4096
}


resource "aws_key_pair" "generated_key" {
  key_name   = var.key_name
  public_key = tls_private_key.rsa_key.public_key_openssh

  # Incase to use for SSH
  provisioner "local-exec" {
    command = <<-EOT
		echo '${tls_private_key.rsa_key.private_key_pem}' > ./'${var.key_name}'.pem
		chmod 400 ./'${var.key_name}'.pem
    
	EOT
  }
}

####################### Instance #########################
resource "aws_instance" "golang_react_instance" {
  ami           = var.instance_ami
  instance_type = var.instance_type

  # associate_public_ip_address=true
  key_name = var.key_name

  vpc_security_group_ids = [var.security_group_id]

  iam_instance_profile = aws_iam_instance_profile.ec2_instance_profile.name

  # subnet_id        = var.public_subnets[0]
  subnet_id        = var.private_subnets[0]

  user_data_base64 = base64encode(templatefile("${path.module}/user_data.sh",{
    instance_name = "golang_react_instance"
  }))
}


