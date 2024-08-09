resource "aws_cloudwatch_log_group" "ec2_log_group" {
  name = "/aws/ec2/golang_react_instance"
  retention_in_days = 7
}


resource "aws_iam_policy" "ec2_cloudwatch_policy" {
  name = "ec2_cloudwatch_policy"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents",
          "logs:DescribeLogStreams",
          "logs:PutRetentionPolicy"
        ],
        Effect = "Allow",
        Resource = "*",
      },
    ],
  })
}



