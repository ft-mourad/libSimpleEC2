package SEC2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func indexResult(resp *ec2.DescribeInstancesOutput) ([]SimpleInstance, []string) {
	var instance SimpleInstance
	var instances []SimpleInstance
	var iids []string
	//looping through the results (list of instances)
	for idx, _ := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			//create a SimpleInstance object per returned existing EC2 instance
			instance.State = *inst.State.Name
			instance.Id = *inst.InstanceId
			//parse the list of tags to find the Owner and Name tags
			for i := 0; i < len(inst.Tags); i++ {
				if *inst.Tags[i].Key == "Name" {
					instance.Name = *inst.Tags[i].Value
				}
				if *inst.Tags[i].Key == "Owner" {
					instance.Owner = *inst.Tags[i].Value
				}
			}
			//add the Instance ID to an array of IDs
			iids = append(iids, instance.Id)
			//add the SimpleInstance Object to an array of SimpleInstances
			instances = append(instances, instance)
		}
	}
	return instances, iids
}

func SearchInstancesFromTag(svc *ec2.EC2, keyTag string, valueTag string) ([]SimpleInstance, []string) {

	var inputFilter *ec2.DescribeInstancesInput
	inputFilter = addTagFilter(keyTag, valueTag, inputFilter)
	//call
	resp, err := svc.DescribeInstances(inputFilter)
	if err != nil {
		panic(err)
	}
	instances, iids := indexResult(resp)

	fmt.Println(iids)
	return instances, iids
}

func addTagFilter(keyTag string, valueTag string, param *ec2.DescribeInstancesInput) *ec2.DescribeInstancesInput {
	param = &ec2.DescribeInstancesInput{
		DryRun: aws.Bool(false),
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:" + keyTag),
				Values: []*string{
					aws.String(valueTag),
				},
			},
		},
	}
	return param
}
