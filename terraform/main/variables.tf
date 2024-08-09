variable "region" {
  type        = string
  description = "region where to apply the resources"
}

variable "profile" {
  type        = string
  description = "the users profile to apply those resources"
}


######### VPC ############
variable "vpc_name" {
  type        = string
  description = "vpc name"

}
variable "azs" {
  description = "availablity zones"
  type        = list(string)
}


############### EC2 Instance ##########
variable "instance_ami" {
  type        = string
  description = "aws ec2 instance ami"
}

variable "instance_type" {
  type        = string
  description = "aws ec2 instance type"
}

variable "key_name" {
  type        = string
  description = "ec2 isntance key-pair name"
}


############### loadbalancer #########################33
variable "ssl_policy" {
  type        = string
  description = "ssl termination policy"
}

variable "ssl_certificate_arn" {
  type        = string
  description = "ssl certificate arn"
}


################# Security group ###########################
variable "security_group_name" {
  type = string
}



############## Domain names #################
variable "main_domain" {
  type = string
  description = "example.com"
}

variable "api_domain" {
  type = string
  description = "api.example.com"
}