package SEC2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func createNewTags(svc *ec2.EC2, iid string, keyTag string, valueTag string) {
	fmt.Println("\ttag -- ", keyTag, " : ", valueTag, "\n")
	svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{
			aws.String(iid)},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(keyTag),
				Value: aws.String(valueTag),
			},
		},
	})
}
