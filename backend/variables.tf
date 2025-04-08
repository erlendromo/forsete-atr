# Credentials

variable "application_credential_id" {
  description = "OpenStack Application Credential ID"
  nullable    = false
  type        = string
}

variable "application_credential_secret" {
  description = "OpenStack Application Credential Secret"
  nullable    = false
  type        = string
  sensitive   = true
}

# Network

variable "network_name" {
  default     = "network-1"
  description = "Name of the network"
  type        = string
}

variable "admin_state_up" {
  default     = true
  description = "Admin state of the network"
  type        = bool
}

variable "subnet_name" {
  default     = "subnet-1"
  description = "Name of the subnet"
  type        = string
}

variable "subnet_cidr" {
  default     = "192.168.1.0/24"
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

# Routing

variable "router_name" {
  default     = "router-1"
  description = "Name of the router"
  type        = string
}

# Security group

variable "application_port" {
  default     = 8080
  description = "Application port"
  type        = number
}

# Key pair

variable "my_openstack_key_name" {
  default     = "my-openstack-key"
  description = "Name of the OpenStack keypair"
  type        = string
}

variable "my_openstack_key_public" {
  description = "Public key for OpenStack keypair"
  nullable    = false
  type        = string
  sensitive   = true
}

# VM

variable "vm_name" {
  default     = "vm-1"
  description = "Name of the VM"
  type        = string
}

variable "vm_image_id" {
  default     = "5bdb1498-831c-4de0-b7a0-8f63379c96ed"
  description = "ID of the image to use for the VM"
  type        = string
}

variable "vm_flavor_name" {
  default     = "de3.12c60r.a100-10g"
  description = "Name of the flavor to use for the VM"
  type        = string
}
