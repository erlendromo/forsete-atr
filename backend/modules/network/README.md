# Network Module

This module provisions a private network in OpenStack along with a subnet, a router, and a router interface.
It also connects the private network to an external network for internet access.

---

## Variables

| Name                    | Description                                              | Type    | Required |
|-------------------------|----------------------------------------------------------|---------|----------|
| `external_network_name` | Name of the external network to connect the router to    | string  | Yes      |
| `network_name`          | Name of the internal network to create                   | string  | Yes      |
| `admin_state_up`        | Administrative state of the network (usually `true`)     | bool    | Yes      |
| `subnet_name`           | Name of the subnet                                       | string  | Yes      |
| `subnet_cidr`           | CIDR block for the subnet (e.g., `192.168.1.0/24`)       | string  | Yes      |
| `subnet_ip_version`     | IP version (typically `4` for IPv4)                      | number  | Yes      |
| `subnet_enable_dhcp`    | Whether to enable DHCP on the subnet                     | bool    | Yes      |
| `router_name`           | Name of the router to create                             | string  | Yes      |

---

## Outputs

| Name         | Description                         |
|--------------|-------------------------------------|
| `network_id` | ID of the created network           |
| `subnet_id`  | ID of the created subnet            |

---

## Example Usage

```hcl
module "network" {
  source = "../../modules/network"

  external_network_name = var.external_network_name
  network_name          = "My-Network"
  admin_state_up        = true
  subnet_name           = "My-Subnet"
  subnet_cidr           = "192.168.1.0/24"
  subnet_ip_version     = 4
  subnet_enable_dhcp    = true
  router_name           = "My-Router"
}
```
