resource "openstack_networking_port_v2" "main" {
  network_id         = var.network_id
  security_group_ids = var.security_group_ids

  fixed_ip {
    subnet_id = var.subnet_id
  }
}
