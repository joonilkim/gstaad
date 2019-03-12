resource "aws_cognito_user_pool" "_" {
  name = "${var.ns}"

  username_attributes = ["email"]

  admin_create_user_config {
    allow_admin_create_user_only = false
  }

  password_policy {
    minimum_length    = 8
    require_uppercase = true
    require_lowercase = true
    require_numbers   = true
    require_symbols   = true
  }

  schema {
    name                = "email"
    attribute_data_type = "String"
    mutable             = false
    required            = true
  }

  lifecycle {
    ignore_changes = ["schema"]
  }
}

resource "aws_cognito_user_pool_client" "_" {
  name                   = "${var.ns}"
  user_pool_id           = "${aws_cognito_user_pool._.id}"
  generate_secret        = false
  refresh_token_validity = 30
  explicit_auth_flows    = ["ADMIN_NO_SRP_AUTH", "USER_PASSWORD_AUTH"]
}

output "cognito_userpool_id" {
  value = "${aws_cognito_user_pool._.id}"
}

output "cognito_client_id" {
  value = "${aws_cognito_user_pool_client._.id}"
}
