package SEC2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func ListInstances(svc *ec2.EC2) ([]SimpleInstance, []string) {
	var inputFilter *ec2.DescribeInstancesInput
	//call DescribeInstances
	resp, err := svc.DescribeInstances(inputFilter)
	if err != nil {
		panic(err)
	}

	//fmt.Println(resp)
	instances, iids := indexResult(resp)
	// fmt.Println(iids)
	return instances, iids
}

func StartInstance(svc *ec2.EC2, iid string) {
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(iid),
		},
		DryRun: aws.Bool(false),
	}
	_, err := svc.StartInstances(input)
	if err != nil {
		fmt.Println(err)
	}

}

func StopInstance(svc *ec2.EC2, iid string) {
	fmt.Println("stopping - ", iid)
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(iid),
		},
		DryRun: aws.Bool(false),
	}
	_, err := svc.StopInstances(input)
	if err != nil {
		fmt.Println(err)
	}
}

func terminateInstance() {

}

func TagInstance(svc *ec2.EC2, iid, keyTag, valueTag string) {
	createNewTags(svc, iid, keyTag, valueTag)
}
