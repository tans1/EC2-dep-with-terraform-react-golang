module "vpc" {
  source   = "./vpc"
  azs      = var.azs
  vpc_name = var.vpc_name
}

module "ec2_instance" {
  source        = "./ec2_instance"
  instance_ami  = var.instance_ami
  instance_type = var.instance_type
  key_name      = var.key_name
  region        = var.region
  security_group_id = module.security-group.security_group_id
  private_subnets = module.vpc.private_subnets
  public_subnets = module.vpc.public_subnets
}

module "loadbalancer" {
  source              = "./loadbalancer"
  ssl_policy          = var.ssl_policy
  ssl_certificate_arn = var.ssl_certificate_arn
  main_domain = var.main_domain
  api_domain = var.api_domain 
  instance_id = module.ec2_instance.instance_id
  security_group_id = module.security-group.security_group_id
  public_subnets = module.vpc.public_subnets
  vpc_id = module.vpc.vpc_id
}

module "security-group" {
  source              = "./security_group"
  security_group_name = var.security_group_name
  vpc_id = module.vpc.vpc_id
  vpc_cidr_block = module.vpc.vpc_cidr_block
  
}
