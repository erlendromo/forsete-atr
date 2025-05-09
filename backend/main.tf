data "openstack_networking_network_v2" "ntnu_global" {
  name = "ntnu-global"
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
  external_network_id = data.openstack_networking_network_v2.ntnu_global.id

  depends_on = [
    data.openstack_networking_network_v2.ntnu_global
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

resource "openstack_networking_secgroup_v2" "ssh" {
  name        = "My-Security-Group"
  description = "Allow SSH port"
}

resource "openstack_networking_secgroup_rule_v2" "ssh" {
  security_group_id = openstack_networking_secgroup_v2.ssh.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22

  depends_on = [
    openstack_networking_secgroup_v2.ssh
  ]
}

resource "openstack_networking_secgroup_v2" "application" {
  name        = "Application-Security-Group"
  description = "Allow application port"
}

resource "openstack_networking_secgroup_rule_v2" "application" {
  security_group_id = openstack_networking_secgroup_v2.application.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = var.application_port
  port_range_max    = var.application_port

  depends_on = [
    openstack_networking_secgroup_v2.application
  ]
}

resource "openstack_networking_port_v2" "main" {
  network_id = openstack_networking_network_v2.main.id
  security_group_ids = [
    openstack_networking_secgroup_v2.ssh.id,
    openstack_networking_secgroup_v2.application.id
  ]

  fixed_ip {
    subnet_id = openstack_networking_subnet_v2.main.id
  }

  depends_on = [
    openstack_networking_network_v2.main,
    openstack_networking_secgroup_v2.ssh,
    openstack_networking_secgroup_v2.application,
    openstack_networking_subnet_v2.main
  ]
}

resource "openstack_compute_keypair_v2" "main" {
  name       = var.my_openstack_key_name
  public_key = var.my_openstack_key_public
}

resource "openstack_compute_instance_v2" "main" {
  name        = var.vm_name
  image_id    = var.vm_image_id
  flavor_name = var.vm_flavor_name
  key_pair    = openstack_compute_keypair_v2.main.name

  network {
    port = openstack_networking_port_v2.main.id
  }

  depends_on = [
    openstack_compute_keypair_v2.main,
    openstack_networking_port_v2.main
  ]
}

resource "openstack_networking_floatingip_v2" "main" {
  pool = data.openstack_networking_network_v2.ntnu_global.name

  depends_on = [
    data.openstack_networking_network_v2.ntnu_global
  ]
}

resource "openstack_networking_floatingip_associate_v2" "main" {
  floating_ip = openstack_networking_floatingip_v2.main.address
  port_id     = openstack_networking_port_v2.main.id

  depends_on = [
    openstack_networking_floatingip_v2.main,
    openstack_networking_port_v2.main
  ]
}
