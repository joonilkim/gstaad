resource "aws_instance" "bastion" {
  count           = "${var.bastion_on}"
  ami             = "${var.bastion_ami}"
  instance_type   = "t2.micro"
  subnet_id       = "${aws_subnet.pub_d.id}"
  key_name        = "${var.bastion_key}"
  security_groups = ["${aws_security_group.bastion.id}"]
  associate_public_ip_address = true

  tags {
      Name = "Bastion"
  }
}

output "bastion_ip" {
  value = "${aws_instance.bastion.*.public_ip}"
}