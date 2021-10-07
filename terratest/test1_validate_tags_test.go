package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Actual Terraform Test
func TestTerraformTags(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested
		TerraformDir: "../terraform",
		Vars: map[string]interface{}{
			"BUCKET_NAME": "test1bucketmdemartinis",
		},
	})

	// Clean up resources with "terraform destroy" at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// The Tags will be saved in a map for each resource. A map doesn't maintain order and Tags won't necessary be read in order, hence it's easier to compare maps with the assert module
	var mapbuckettags map[string]string = make(map[string]string)
	var mapec2tags map[string]string = make(map[string]string)

	// Creates the Session to connect to AWS. Region is hardcoded but AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY are read from Environment
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open session to AWS, %v", err)
	} else {
		// Create S3 service client
		svc := s3.New(sess)

		// Read S3 Bucket Tags for bucket BUCKET_NAME
		getS3Input := &s3.GetBucketTaggingInput{
			Bucket: aws.String("test1bucketmdemartinis"),
		}
		bucketTags, err := svc.GetBucketTagging(getS3Input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to get Bucket Tags, %v", err)
		} else {
			for _, b := range bucketTags.TagSet {
				mapbuckettags[aws.StringValue(b.Key)] = aws.StringValue(b.Value)
			}
			fmt.Println("\nS3 Bucket Tags:")
			fmt.Println(mapbuckettags)
		}

		// Create EC2 service client
		svcec2 := ec2.New(sess)

		// Read EC2 instances
		getEC2Input := &ec2.DescribeTagsInput{
			Filters: []*ec2.Filter{
				{
					Name: aws.String("resource-type"),
					Values: []*string{
						aws.String("instance"),
					},
				},
			},
		}

		ec2result, err := svcec2.DescribeTags(getEC2Input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to get EC2 Instance Tags, %v", err)
		} else {
			for _, b := range ec2result.Tags {
				mapec2tags[aws.StringValue(b.Key)] = aws.StringValue(b.Value)
			}
			fmt.Println("\nEC2 Instance Tags:")
			fmt.Println(mapec2tags)
			fmt.Println("")
		}
	}

	// Test if Tags from actual infra match with the expected ones. For that, we create a map with expected tags to compare

	var mapexpectedtags map[string]string = make(map[string]string)
	mapexpectedtags["Owner"] = "InfraTeam"
	mapexpectedtags["Name"] = "Flugel"

	assert.Equal(t, mapexpectedtags, mapbuckettags)
	assert.Equal(t, mapexpectedtags, mapec2tags)
}
