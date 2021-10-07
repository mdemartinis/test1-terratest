# Test 1 for Flugel.it

This repository holds the files needed to run the **Test1** requested by Flugel.it to Marco de Martinis.

The requirements are as follows:

- Create Terraform code to create an S3 bucket and an EC2 instance. Both resources must be tagged with Name=Flugel, Owner=InfraTeam.
- Using Terratest, create the test automation for the Terraform code, validating that both resources are tagged properly.
- Setup Github Actions to run a pipeline to validate this code.
- Publish your code in a public GitHub repository, and share a Pull Request with your code. Do not merge into master until the PR is approved.
- Include documentation describing the steps to run and test the automation.

## About the commit history

The development has been realized in the repository [https://github.com/mdemartinis/GithubActions], including testing the CD workflow on merge with the main branch. For that reason, a new repository is presented with the final result of the development.
You are free to check the other repository to see the development process.

## How it Works

The solution is composed by Terraform, Go, and YML files, in order to deploy the infrastructure, test it and automate the entire process using GitHub Actions.

As you may notice, the branch *main* is empty. This is because all files were committed to the branch *pipeline* and are pending to merge with a Pull Request.

The GitHub Action workflow for CI runs on Pull Request creation. As first step, it lints the entire code base, to verify that everything has been written following the highest standards. Then it proceeds to prepare the runner to execute the test with Terratest.

---

### Running the solution locally

#### Requirements and Dependencies

To run the code locally, you must have a set of AWS credentials that are allowed to read and write the AWS S3 Bucket "terraform-state-mdemartinis", otherwise you will need to modifiy [backend.tf] to point to a new bucket or delete the ***backend*** definition to save the Terraform state locally.
Also, it's important to update `BUCKET_NAME` passed as a parameter in Terratest, as it could already exist an S3 bucket with that name and, in that case, the solution will fail. You can find this in [test1_validate_tags_test.go]:
```go
Vars: map[string]interface{}{
  "BUCKET_NAME": "test1bucketmdemartinis",
},
```

It is required to have Terraform and Go installed, with minimum versions as follows (based on versions used to develop the solution):
- Terraform: 1.0.7
- Go: 1.17.1

Additionally, if you want to clone this repository, you will need:
- Git: 2.33

#### Clone repository

You can clone this repository and select the branch *pipeline* with the following commands:
```bash
git clone https://github.com/mdemartinis/test1-terratest.git
cd ./test1-terratest/
git checkout pipeline
```

Although it's not needed to run the full test with Terratest, if you want to execute Terraform commands, you have to move to the `./terraform/` folder first, and then you will be able to initialize it using:
```bash
terraform init
```

After initializing Terraform, you are now able to run any of the following commands:
```bash
terraform validate
terraform plan
terraform apply
terraform destroy
```

---

If you want to run the full test, you have to move first to the `./terratest/` folder, initialize `go mod` and download all the dependencies. Then, execute the test. Use the following commands:
```bash
cd ./terratest/
go mod init <MODULE_NAME>
go mod tidy
go test
```

Replace `<MODULE_NAME>` with your desired module name, tipically, your repository name in the form of: **github.com/mdemartinis/test1-terratest**

### Running the GitHub Action workflow

#### Requirements

You can fork this repository or clone it and re-upload to your own repository.

In this case, running the CI pipeline should be easier than running the solution locally, as it takes care of the dependencies automatically. Though it's required to set up two Repository Secrets, to use later as Environment Variables in the runner. These secrets are the AWS credentials and must be called as `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`.

Also,  you'll need to take care of the [backend.tf] and `BUCKET_NAME` as mentioned above in **Running the solution locally**.


**Always remember to verify that the run has completed successfully and the cloud resources have been destroyed, to not incur in undesired billing.**


[//]: #

   [backend.tf]: <https://github.com/mdemartinis/test1-terratest/blob/pipeline/terraform/backend.tf>
   [test1_validate_tags_test.go]: <https://github.com/mdemartinis/test1-terratest/blob/pipeline/terratest/test1_validate_tags_test.go>
   [https://github.com/mdemartinis/GithubActions]: <https://github.com/mdemartinis/GithubActions>