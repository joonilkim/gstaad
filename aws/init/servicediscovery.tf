resource "aws_service_discovery_private_dns_namespace" "_" {
  name = "${var.domain}"
  vpc  = "${aws_vpc._.id}"
}

output "service_discovery_ns_id" {
    value = "${aws_service_discovery_private_dns_namespace._.id}"
}