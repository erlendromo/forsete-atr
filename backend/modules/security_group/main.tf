resource "openstack_networking_secgroup_v2" "main" {
  name        = var.name
  description = var.description
}

resource "openstack_networking_secgroup_rule_v2" "main" {
  count             = length(var.tcp_ingress_ports)
  security_group_id = openstack_networking_secgroup_v2.main.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = var.tcp_ingress_ports[count.index]
  port_range_max    = var.tcp_ingress_ports[count.index]

  depends_on = [
    openstack_networking_secgroup_v2.main
  ]
}
