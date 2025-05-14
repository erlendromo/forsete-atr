data "openstack_networking_network_v2" "ntnu_external_network" {
  name = var.external_network_name
}

resource "openstack_networking_network_v2" "main" {
  name           = var.network_name
  admin_state_up = var.admin_state_up
}

resource "openstack_networking_subnet_v2" "main" {
  network_id = openstack_networking_network_v2.main.id

  name        = var.subnet_name
  cidr        = var.subnet_cidr
  ip_version  = var.subnet_ip_version
  enable_dhcp = var.subnet_enable_dhcp

  depends_on = [
    openstack_networking_network_v2.main
  ]
}

resource "openstack_networking_router_v2" "main" {
  name                = var.router_name
  external_network_id = data.openstack_networking_network_v2.ntnu_external_network.id

  depends_on = [
    data.openstack_networking_network_v2.ntnu_external_network
  ]
}

resource "openstack_networking_router_interface_v2" "main" {
  router_id = openstack_networking_router_v2.main.id
  subnet_id = openstack_networking_subnet_v2.main.id

  depends_on = [
    openstack_networking_router_v2.main,
    openstack_networking_subnet_v2.main
  ]
}
