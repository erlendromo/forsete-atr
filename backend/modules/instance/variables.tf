variable "name" {
  description = "Name of the instance"
  type        = string
}

variable "image_id" {
  description = "ID of the instance image"
  type        = string
}

variable "flavor_name" {
  description = "Name of the instance flavor"
  type        = string
}

variable "keypair_name" {
  description = "Name of the keypair"
  type        = string
}

variable "port_id" {
  description = "ID of the networking port"
  type        = string
}
