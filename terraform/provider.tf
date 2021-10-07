# These variables are used in the provider, in order to connect to AWS properly.
# ACCESS and SECRET keys will be declared but not populated.
# They will be fulfilled later with .tfvars file to keep the secrets. 
# The provider "aws" requires access key and a secret key. If not directly
# provided, these variables are read ENVIRONMENT, as AWS_ACCESS_KEY_ID and
# AWS_SECRET_ACCESS_KEY

provider "aws" {
  region     = var.AWS_REGION
}