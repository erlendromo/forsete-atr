# Frontend Application Deployment

This Terraform configuration deploys a frontend application infrastructure on OpenStack using reusable modules.
It provisions networking, security groups, a key pair, a compute instance, and a floating IP.

---

## Architecture

This configuration provisions a frontend application environment in OpenStack, including:

- A private network with a router connected to the external network.
- Security groups for SSH and application access.
- A virtual machine deployed with a specific image, flavor, and key pair.
- A floating IP assigned to the instance for external connectivity.

---

## Modules Used

| Module     | Description                                           |
|------------|-------------------------------------------------------|
| `network`  | Creates network, subnet, router, and router interface |
| `ssh_sg`   | Creates a security group allowing SSH access          |
| `app_sg`   | Creates a security group for the app port (3000)      |
| `keypair`  | Registers an OpenStack key pair for SSH               |
| `port`     | Provisions a network port with security groups        |
| `instance` | Launches the compute instance                         |
| `ip`       | Allocates and associates a floating IP                |

---

## Variables

| Name                    | Description                            | Type   | Required |
|-------------------------|----------------------------------------|--------|----------|
| `keypair_name`          | Name of the SSH keypair                | string | Yes      |
| `public_key`            | Public key for the keypair             | string | Yes      |
| `external_network_name` | Name of the external OpenStack network | string | Yes      |

---

## Outputs

| Name  | Description                  |
|-------|------------------------------|
| `ip`  | Public floating IP of the VM |

---

## Example Usage

```hcl
module "frontend" {
  source = "./deployments/frontend"

  keypair_name           = var.keypair_name
  public_key             = var.public_key
  external_network_name  = var.external_network_name
}
```
