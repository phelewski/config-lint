resource "aws_iam_role" "role1" {
    name = "role1"
    assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": }
}
EOF
}

