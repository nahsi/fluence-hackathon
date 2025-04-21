terraform {
  required_providers {
    fluence = {
      source  = "registry.terraform.io/cloudlesslabs/fluence"
      version = "0.0.0"
    }
  }
}

provider "fluence" {
  api_key = "2a062304-e8fe-4a69-9669-2ff94126ec9e"
}

data "fluence_vm_instances" "example" {}

resource "fluence_ssh_key" "example" {
  name       = "examplekeey"
  public_key = file("/home/nahsi/.ssh/id_ed25519.pub")
}

# output "public_ip" {
#     value = data.fluence_vm_instances.example.instances[0].public_ip
# }

resource "fluence_vm" "example" {
  # How many VMs to spin up
  instances = 1

  # Mandatory VM configuration
  vm_configuration = {
    # human‑readable name
    name     = "hackathon-vm"

    # image URL
    os_image = "https://fluence-misc.s3.eu-west-1.amazonaws.com/nixos.qcow2"

    # open_ports is a list of { port, protocol } objects
    open_ports = [
      { port = 22,  protocol = "tcp" },
      { port = 80,  protocol = "tcp" },
      { port = 443, protocol = "tcp" },
    ]

    # ssh_keys is a list of SSH‐key fingerprints
    ssh_keys = [
      fluence_ssh_key.example.public_key
    ]
  }

  # Optional: if you need to specify datacenter/resources,
  # use the 'constraints' block’s nested 'additional_resources' and/or
  # 'datacenter' objects.  Otherwise you can omit it.
  # constraints = {
  #   additional_resources = {
  #     storage = [
  #       { type = "NVMe", megabytes = 25600 },
  #     ]
  #   }
  #   basic_configuration = "0x0000000000000000000000000000000000000000000000000000000000000005"
  # }
}

# output "vm_deal_id" {
#   description = "The Fluence deal ID of the newly created VM"
#   value       = fluence_vm.example.id
# }

# output "vm_public_ip" {
#   description = "The Fluence deal ID of the newly created VM"
# value       = fluence_vm.example.public_ip
# }

output "vm" {
value       = fluence_vm.example
}
