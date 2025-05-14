variable "external_network_name" {
  description = "Name of the external network"
}

variable "network_name" {
  description = "Name of the network"
  type        = string
}

variable "admin_state_up" {
  default     = true
  description = "Admin state of the network"
  type        = bool
}

variable "subnet_name" {
  description = "Name of the subnet"
  type        = string
}

variable "subnet_cidr" {
  description = "CIDR of the subnet"
  type        = string
}
variable "subnet_ip_version" {
  default     = 4
  description = "IP version of the subnet"
  type        = number
}

variable "subnet_enable_dhcp" {
  default     = true
  description = "Enable DHCP for the subnet"
  type        = bool
}

variable "router_name" {
  description = "Name of the router"
  type        = string
}
