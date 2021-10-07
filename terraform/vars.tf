# AWS region where we will deploy resources
variable "AWS_REGION" {
  default = "us-east-1"
}

# Next, AMI and size for EC2 instances.
# This AMI corresponds to Ubuntu 20.04 LTS
variable "AMI_USEAST1" {
  default = "ami-036490d46656c4818"
}

variable "INSTANCE_TYPE" {
  default = "t2.micro"
}

# Bucket name. Random from pipeline

variable "BUCKET_NAME" {
  default = "test1bucketmdemartinis"
}

# General Tags for all resources

variable "TAGS" {
  type    = map(string)
  default = {
    Name  = "Flugel"
    Owner = "InfraTeam"
  }
}