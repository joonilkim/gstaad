resource "aws_vpc" "_" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = "gstaad"
  }
}

resource "aws_internet_gateway" "_" {
  vpc_id = "${aws_vpc._.id}"

  tags {
    Name = "gstaad"
  }
}

resource "aws_route_table" "_" {
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway._.id}"
  }

  vpc_id = "${aws_vpc._.id}"
}

output "vpc_id" {
  value = "${aws_vpc._.id}"
}

