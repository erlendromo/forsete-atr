# Security Group Module

This module creates an OpenStack security group and optionally adds ingress rules for specified TCP ports using the `openstack_networking_secgroup_v2` and `openstack_networking_secgroup_rule_v2` resources.

---

## Variables

| Name               | Description                                      | Type         | Required |
|--------------------|--------------------------------------------------|--------------|----------|
| `name`             | Name of the security group                       | string       | Yes      |
| `description`      | Description of the security group                | string       | Yes      |
| `tcp_ingress_ports`| List of TCP ports to allow as ingress rules      | list(number) | Yes      |

---

## Outputs

| Name                | Description                          |
|---------------------|--------------------------------------|
| `security_group_id` | ID of the created security group     |

---

## Example Usage

```hcl
module "ssh_sg" {
  source = "../../modules/security_group"

  name              = "SSH"
  description       = "Allow SSH ingress"
  tcp_ingress_ports = [22]
}
```
