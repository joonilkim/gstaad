resource "aws_subnet" "pub_c" {
  availability_zone = "${var.region}c"
  cidr_block        = "10.1.64.0/22"
  vpc_id            = "${aws_vpc._.id}"

  tags = {
    Name = "gstaad_pub"
  }
}


resource "aws_subnet" "db_c" {
  availability_zone = "${var.region}c"
  cidr_block        = "10.1.80.0/22"
  vpc_id            = "${aws_vpc._.id}"

  tags = {
    Name = "gstaad_db"
  }
}


resource "aws_subnet" "priv_c" {
  availability_zone = "${var.region}c"
  cidr_block        = "10.1.96.0/20"
  vpc_id            = "${aws_vpc._.id}"

  tags = {
    Name = "gstaad_priv"
  }
}

resource "aws_subnet" "pub_d" {
  availability_zone = "${var.region}d"
  cidr_block        = "10.1.0.0/22"
  vpc_id            = "${aws_vpc._.id}"

  tags = {
    Name = "gstaad_pub"
  }
}

resource "aws_subnet" "db_d" {
  availability_zone = "${var.region}d"
  cidr_block        = "10.1.16.0/22"
  vpc_id            = "${aws_vpc._.id}"

  tags = {
    Name = "gstaad_db"
  }
}

resource "aws_subnet" "priv_d" {
  availability_zone = "${var.region}d"
  cidr_block        = "10.1.32.0/20"
  vpc_id            = "${aws_vpc._.id}"

  tags = {
    Name = "gstaad_priv"
  }
}


resource "aws_route_table_association" "pub_c" {
  route_table_id = "${aws_route_table._.id}"
  subnet_id      = "${aws_subnet.pub_c.id}"
}

resource "aws_route_table_association" "pub_d" {
  route_table_id = "${aws_route_table._.id}"
  subnet_id      = "${aws_subnet.pub_d.id}"
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