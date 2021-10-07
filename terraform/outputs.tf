output "instance-tags" {
  value = aws_instance.test1instance.tags
}

output "bucket-tags" {
  value = aws_s3_bucket.test1bucket.tags
}