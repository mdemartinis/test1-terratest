resource "aws_s3_bucket" "test1bucket" {
  #ts:skip=AWS.S3Bucket.EncryptionandKeyManagement.HIGH.0405 Need to skip this rule
  #ts:skip=AWS.S3Bucket.LM.MEDIUM.0078 Need to skip this rule
  bucket = var.BUCKET_NAME
  acl    = "private"
  tags   = var.TAGS

  versioning {
    enabled = true
  }
}