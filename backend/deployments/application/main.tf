module "network" {
  source = "../../modules/network"

  external_network_name = var.external_network_name
  network_name          = "Application-Network"
  subnet_name           = "Application-Subnet"
  subnet_cidr           = "192.168.1.0/24"
  router_name           = "Application-Router"
}

module "ssh_sg" {
  source = "../../modules/security_group"

  name              = "SSH"
  description       = "Allow SSH ingress"
  tcp_ingress_ports = [22]
}

module "app_sg" {
  source = "../../modules/security_group"

  name              = "App"
  description       = "Allow application port ingress"
  tcp_ingress_ports = [8080]
}

module "keypair" {
  source = "../../modules/keypair"

  name       = var.keypair_name
  public_key = var.public_key
}

module "port" {
  source = "../../modules/port"

  network_id = module.network.network_id
  subnet_id  = module.network.subnet_id
  security_group_ids = [
    module.ssh_sg.security_group_id,
    module.app_sg.security_group_id
  ]

  depends_on = [
    module.network,
    module.ssh_sg,
    module.app_sg
  ]
}

module "instance" {
  source = "../../modules/instance"

  name         = "Application-Instance"
  image_id     = "5bdb1498-831c-4de0-b7a0-8f63379c96ed"
  flavor_name  = "de3.12c60r.a100-10g"
  keypair_name = var.keypair_name
  port_id      = module.port.networking_port_id

  depends_on = [
    module.port
  ]
}

module "ip" {
  source = "../../modules/ip"

  external_network_name = var.external_network_name
  port_id               = module.port.networking_port_id

  depends_on = [
    module.port
  ]
}
