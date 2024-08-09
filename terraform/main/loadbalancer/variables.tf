variable "ssl_policy" {
  type        = string
  description = "ssl termination policy"
}

variable "ssl_certificate_arn" {
  type        = string
  description = "ssl certificate arn"
}

variable "main_domain" {
  type = string
}

variable "api_domain" {
  type = string
}

variable "instance_id" {
  type = string
}

variable "security_group_id" {
  type = string
  
}

variable "public_subnets" {
  type = list(string)
}

variable "vpc_id" {
  type = string
}