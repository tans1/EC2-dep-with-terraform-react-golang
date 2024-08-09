variable "vpc_name" {
  type        = string
  description = "the vpc name"

}

variable "azs" {
  type        = list(string)
  description = "availablity zones"
}

