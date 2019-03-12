resource "aws_subnet" "pub_c" {
  availability_zone = "${var.region}c"
  cidr_block        = "10.1.64.0/22"
  vpc_id            = "${aws_vpc._.id}"

  tags = {
    Name = "${var.ns}_pub"
  }
}


resource "aws_subnet" "db_c" {
  availability_zone = "${var.region}c"
  cidr_block        = "10.1.80.0/22"
  vpc_id            = "${aws_vpc._.id}"

  tags = {
    Name = "${var.ns}_db"
  }
}


resource "aws_subnet" "priv_c" {
  availability_zone = "${var.region}c"
  cidr_block        = "10.1.96.0/20"
  vpc_id            = "${aws_vpc._.id}"

  tags = {
    Name = "${var.ns}_priv"
  }
}

resource "aws_subnet" "pub_d" {
  availability_zone = "${var.region}d"
  cidr_block        = "10.1.0.0/22"
  vpc_id            = "${aws_vpc._.id}"

  tags = {
    Name = "${var.ns}_pub"
  }
}

resource "aws_subnet" "db_d" {
  availability_zone = "${var.region}d"
  cidr_block        = "10.1.16.0/22"
  vpc_id            = "${aws_vpc._.id}"

  tags = {
    Name = "${var.ns}_db"
  }
}

resource "aws_subnet" "priv_d" {
  availability_zone = "${var.region}d"
  cidr_block        = "10.1.32.0/20"
  vpc_id            = "${aws_vpc._.id}"

  tags = {
    Name = "${var.ns}_priv"
  }
}

resource "aws_route_table_association" "pub_c" {
  route_table_id = "${aws_route_table._.id}"
  subnet_id      = "${aws_subnet.pub_c.id}"
}

resource "aws_route_table_association" "priv_c" {
  route_table_id = "${aws_route_table._.id}"
  subnet_id      = "${aws_subnet.priv_c.id}"
}

resource "aws_route_table_association" "pub_d" {
  route_table_id = "${aws_route_table._.id}"
  subnet_id      = "${aws_subnet.pub_d.id}"
}

resource "aws_route_table_association" "priv_d" {
  route_table_id = "${aws_route_table._.id}"
  subnet_id      = "${aws_subnet.priv_d.id}"
}

resource "aws_network_acl" "priv" {
  vpc_id = "${aws_vpc._.id}"
  subnet_ids = [
    "${aws_subnet.priv_c.id}",
    "${aws_subnet.priv_d.id}",
  ]

  egress {
    protocol   = "-1"
    rule_no    = 100
    action     = "allow"
    cidr_block = "0.0.0.0/0"
    from_port  = 0
    to_port    = 0
  }

  ingress {
    protocol   = "tcp"
    rule_no    = 100
    action     = "allow"
    cidr_block = "0.0.0.0/0"
    from_port  = 32768
    to_port    = 65535
  }

  ingress {
    protocol   = "-1"
    rule_no    = 200
    action     = "allow"
    cidr_block = "10.1.0.0/16"
    from_port  = 0
    to_port    = 0
  }

  tags = {
    Name = "${var.ns}_priv"
  }
}

output "pub_subnets" {
  value = ["${aws_subnet.pub_d.id}", "${aws_subnet.pub_c.id}"]
}

output "db_subnets" {
  value = ["${aws_subnet.db_d.id}", "${aws_subnet.db_c.id}"]
}

output "priv_subnets" {
  value = ["${aws_subnet.priv_d.id}", "${aws_subnet.priv_c.id}"]
}