resource "aws_service_discovery_private_dns_namespace" "_" {
  name = "${var.domain}"
  vpc  = "${aws_vpc._.id}"
}

resource "aws_service_discovery_service" "_" {
    name = "${var.ns}"

    dns_config {
        namespace_id = "${aws_service_discovery_private_dns_namespace._.id}"

        dns_records {
            ttl  = 5
            type = "A"
        }
        routing_policy = "MULTIVALUE"
    }

    health_check_custom_config {
        # The number of 30-second intervals before it changes the health status of a service instance
        failure_threshold = 1
    }
}

output "service_discovery_arn" {
    value = "${aws_service_discovery_service._.arn}"
}