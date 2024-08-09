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

variable "region" {
  type        = string
  description = "aws instance region"
}

variable "security_group_id" {
  type = string
  
}

variable "private_subnets" {
  type = list(string)
}

variable "public_subnets" {
  type = list(string)
}