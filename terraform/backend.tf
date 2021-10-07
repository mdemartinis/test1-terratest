terraform {
  backend "s3" {
    # In backend, it's not allowed to use variables, hence, we have to hard code these values
    bucket = "terraform-state-mdemartinis"
    key = "terraform/test1"
    region = "us-east-1"
  }
}