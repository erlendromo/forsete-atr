variable "name" {
  description = "Security group name"
  type        = string
}

variable "description" {
  description = "Description for the security group"
  type        = string
}

variable "tcp_ingress_ports" {
  description = "List of ports to allow for ingress"
  type        = list(number)
}
