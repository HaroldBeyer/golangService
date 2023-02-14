resource "aws_sns_topic" "example_topic" {
  name = "example-topic"
}

output "sns_topic_arn" {
  value = aws_sns_topic.example_topic.arn
}
