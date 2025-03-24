output "vm-ip" {
  value       = openstack_networking_floatingip_v2.main.address
  description = "IP to VM"
}
