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
  cidr_block        = "10.1.128.0/24"
  vpc_id            = "${aws_vpc._.id}"
}

resource "aws_subnet" "priv_c" {
  availability_zone = "${var.region}c"
  cidr_block        = "10.1.129.0/24"
  vpc_id            = "${aws_vpc._.id}"
}
