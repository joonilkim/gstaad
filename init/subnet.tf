resource "aws_subnet" "pub_a" {
  availability_zone = "${var.region}a"
  cidr_block        = "10.1.1.0/24"
  vpc_id            = "${aws_vpc._.id}"
}

resource "aws_subnet" "pub_c" {
  availability_zone = "${var.region}c"
  cidr_block        = "10.1.2.0/24"
  vpc_id            = "${aws_vpc._.id}"
}

resource "aws_subnet" "priv_a" {
  availability_zone = "${var.region}a"
  cidr_block        = "10.1.3.0/24"
  vpc_id            = "${aws_vpc._.id}"
}

resource "aws_subnet" "priv_c" {
  availability_zone = "${var.region}c"
  cidr_block        = "10.1.4.0/24"
  vpc_id            = "${aws_vpc._.id}"
}

resource "aws_route_table_association" "pub_a" {
  route_table_id = "${aws_route_table._.id}"
  subnet_id      = "${aws_subnet.pub_a.id}"
}

resource "aws_route_table_association" "pub_c" {
  route_table_id = "${aws_route_table._.id}"
  subnet_id      = "${aws_subnet.pub_c.id}"
}

output "pub_subnets" {
  value = ["${aws_subnet.pub_a.id}", "${aws_subnet.pub_c.id}"]
}

output "priv_subnets" {
  value = ["${aws_subnet.priv_a.id}", "${aws_subnet.priv_c.id}"]
}