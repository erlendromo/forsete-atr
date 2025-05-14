variable "network_id" {
  description = "ID of the network to port"
  type        = string
}

variable "security_group_ids" {
  description = "List of security group ids to port"
  type        = list(string)
}

variable "subnet_id" {
  description = "ID of the subnet to port"
  type        = string
}
