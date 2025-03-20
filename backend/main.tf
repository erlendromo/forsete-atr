resource "openstack_networking_network_v2" "main" {
  name           = "network1"
  admin_state_up = true
}

resource "openstack_networking_subnet_v2" "main" {
  network_id = openstack_networking_network_v2.main.id

  name        = "subnet1"
  cidr        = "192.168.1.0/24"
  ip_version  = 4
  enable_dhcp = true

  depends_on = [openstack_networking_network_v2.main]
}

resource "openstack_networking_router_v2" "main" {
  name                = "router1"
  external_network_id = "730cb16e-a460-4a87-8c73-50a2cb2293f9"

  depends_on = [
    openstack_networking_network_v2.main
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
  name        = "SSH-Security-Group"
  description = "Allow SSH"

}

resource "openstack_networking_secgroup_rule_v2" "ssh" {
  security_group_id = openstack_networking_secgroup_v2.ssh.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22

  depends_on = [openstack_networking_secgroup_v2.ssh]
}

resource "openstack_compute_keypair_v2" "main" {
  name       = "my-openstack-key"
  public_key = file("~/.ssh/openstack_key.pub")
}

resource "openstack_networking_port_v2" "main" {
  network_id         = openstack_networking_network_v2.main.id
  security_group_ids = [openstack_networking_secgroup_v2.ssh.id]

  fixed_ip {
    subnet_id = openstack_networking_subnet_v2.main.id
  }

  depends_on = [
    openstack_networking_network_v2.main,
    openstack_networking_subnet_v2.main
  ]
}

resource "openstack_compute_instance_v2" "main" {
  name        = "forsete-vm"
  image_id    = "5bdb1498-831c-4de0-b7a0-8f63379c96ed"
  flavor_name = "de3.12c60r.a100-10g"
  key_pair    = openstack_compute_keypair_v2.main.name
  admin_pass  = var.vm_admin_pass

  network {
    port = openstack_networking_port_v2.main.id
  }

  depends_on = [
    openstack_compute_keypair_v2.main,
    openstack_networking_secgroup_v2.ssh,
    openstack_networking_network_v2.main,
    openstack_networking_port_v2.main
  ]
}

resource "openstack_networking_floatingip_v2" "main" {
  pool = "ntnu-internal"
}

resource "openstack_networking_floatingip_associate_v2" "main" {
  floating_ip = openstack_networking_floatingip_v2.main.address
  port_id     = openstack_networking_port_v2.main.id

  depends_on = [
    openstack_networking_floatingip_v2.main,
    openstack_networking_port_v2.main
  ]
}
