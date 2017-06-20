package SEC2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type SimpleInstance struct {
	Id    string
	Name  string
	State string
	Owner string
}

func init() {

}

func EC2_init(region string) *ec2.EC2 {
	//create session
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	//initialize session
	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
	return svc
}
