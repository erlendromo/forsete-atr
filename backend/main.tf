# Create Network
resource "openstack_networking_network_v2" "private_net" {
  name           = "private_network"
  admin_state_up = true
}

# Create Subnet
resource "openstack_networking_subnet_v2" "private_subnet" {
  name        = "private_subnet"
  network_id  = openstack_networking_network_v2.private_net.id
  cidr        = "192.168.1.0/24"
  ip_version  = 4
  enable_dhcp = true
}

# Create Router
resource "openstack_networking_router_v2" "router" {
  name                = "router"
  external_network_id = "public_network_id"
}

# Router Interface
resource "openstack_networking_router_interface_v2" "router_interface" {
  router_id = openstack_networking_router_v2.router.id
  subnet_id = openstack_networking_subnet_v2.private_subnet.id
}

# Security Group
resource "openstack_networking_secgroup_v2" "security_group_1" {
  name        = "Security Group 1"
  description = "Allow SSH, HTTP, and HTTPS traffic"
}

# Allow SSH (Port 22)
resource "openstack_networking_secgroup_rule_v2" "ssh" {
  security_group_id = openstack_networking_secgroup_v2.security_group_1.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22
}

# Allow HTTP (Port 80)
resource "openstack_networking_secgroup_rule_v2" "http" {
  security_group_id = openstack_networking_secgroup_v2.security_group_1.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 80
  port_range_max    = 80
}

# Allow HTTPS (Port 443)
resource "openstack_networking_secgroup_rule_v2" "https" {
  security_group_id = openstack_networking_secgroup_v2.security_group_1.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 443
  port_range_max    = 443
}


# Block Storage
resource "openstack_blockstorage_volume_v3" "volume" {
  name        = "my-volume"
  size        = 10
  description = "A test volume"
}

# Compute Instance (VM)
resource "openstack_compute_instance_v2" "vm" {
  name            = "my-vm"
  image_name      = "Ubuntu-22.04"
  flavor_name     = "m1.small"
  key_pair        = "my-key"
  admin_pass      = var.vm_admin_pass
  security_groups = [openstack_networking_secgroup_v2.allow_all.name]

  network {
    uuid = openstack_networking_network_v2.private_net.id
  }

  block_device {
    uuid                  = openstack_blockstorage_volume_v3.volume.id
    source_type           = "volume"
    destination_type      = "volume"
    boot_index            = 0
    delete_on_termination = true
  }
}

# Floating IP
resource "openstack_networking_floatingip_v2" "floating_ip" {
  pool = "public"
}


# Associate Floating IP
resource "openstack_networking_floatingip_associate_v2" "floating_ip_assoc" {
  floating_ip = openstack_networking_floatingip_v2.floating_ip.address
  port_id     = openstack_compute_instance_v2.vm.network.0.uuid
}
