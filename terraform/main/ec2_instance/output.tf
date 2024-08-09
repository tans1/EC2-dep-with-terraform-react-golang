output "instance_id" {
  value = aws_instance.golang_react_instance.id
}

output "region" {
  value = var.region
}