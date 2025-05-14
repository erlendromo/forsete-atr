# Port Module

This module provisions a networking port in OpenStack using the `openstack_networking_port_v2` resource.
The port is created on a specified network and subnet, and can be associated with one or more security groups.

---

## Variables

| Name                 | Description                                      | Type           | Required |
|----------------------|--------------------------------------------------|----------------|----------|
| `network_id`         | ID of the network to attach the port to          | string         | Yes      |
| `subnet_id`          | ID of the subnet for the port's fixed IP         | string         | Yes      |
| `security_group_ids` | List of security group IDs to associate          | list(string)   | Yes      |

---

## Outputs

| Name                | Description                          |
|---------------------|--------------------------------------|
| `networking_port_id`| ID of the created networking port    |

---

## Example Usage

```hcl
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
```
