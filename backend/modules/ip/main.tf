resource "openstack_networking_floatingip_v2" "main" {
  pool = var.external_network_name
}

resource "openstack_networking_floatingip_associate_v2" "main" {
  floating_ip = openstack_networking_floatingip_v2.main.address
  port_id     = var.port_id

  depends_on = [
    openstack_networking_floatingip_v2.main,
  ]
}
