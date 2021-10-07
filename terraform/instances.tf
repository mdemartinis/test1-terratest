resource "aws_instance" "test1instance" {
  #ts:skip=AC-AW-IS-IN-M-0144 need to skip this rule
  #ts:skip=AWS.AI.LM.HIGH.0070 need to skip this rule
  ami           = var.AMI_USEAST1
  instance_type = var.INSTANCE_TYPE
  tags          = var.TAGS

  metadata_options {
    http_endpoint = "enabled"
    http_tokens = "required"
  }
}