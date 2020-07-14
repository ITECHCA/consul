variable "allowed_inbound_cidrs" {
  type        = list(string)
  description = "List of CIDR blocks to permit inbound Consul access from"
}

  
variable "acl_bootstrap_bool" {
  type        = bool
  default     = true
  description = "Initial ACL Bootstrap configurations"
}


variable "consul_clients" {
  default     = "3"
  description = "number of Consul instances"
}

variable "consul_config" {
  description = "HCL Object with additional configuration overrides supplied to the consul servers.  This is converted to JSON before rendering via the template."
  default     = {}
}

variable "consul_cluster_version" {
  default     = "0.0.1"
  description = "Custom Version Tag for Upgrade Migrations"
}

variable "consul_servers" {
  default     = "5"
  description = "number of Consul instances"
}

variable "consul_version" {
  description = "Consul version"
}

variable "enable_connect" {
  type        = bool
  description = "Whether Consul Connect should be enabled on the cluster"
  default     = false
}

variable "instance_type" {
  default     = "t2.micro"
  description = "Instance type for Consul instances"
}

variable "key_name" {
  default     = "default"
  description = "SSH key name for Consul instances"
}

variable "name_prefix" {
  description = "prefix used in resource names"
}

variable "owner" {
  description = "value of owner tag on EC2 instances"
}

variable "public_ip" {
  type        = bool
  default     = false
  description = "should ec2 instance have public ip?"
}

# variable "vpc_id" {
#   description = "VPC ID"
# }

variable "vpc_az" {
  type        = list(string)
  description = "VPC Availability Zone"
}

variable "vpc_name" {
  description = "Name of the VPC"
}

variable "vpc_cidr" {
  description = "List of CIDR blocks for the VPC module"
}

variable "public_subnet_cidr" {
  type        = list(string)
  description = "CIDR Block for the Public Subnet, must be within VPC CIDR range"
}