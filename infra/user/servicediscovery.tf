resource "aws_service_discovery_service" "_" {
    name = "${var.service}"

    dns_config {
        namespace_id = "${var.service_discovery_ns}"

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