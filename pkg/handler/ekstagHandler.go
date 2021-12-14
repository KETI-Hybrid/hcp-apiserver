package handler

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

func ListTagsForResource(listTagsForResourceInput eks.ListTagsForResourceInput) (*eks.ListTagsForResourceOutput, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	eksSvc := eks.New(sess)

	input := &eks.ListTagsForResourceInput{
		ResourceArn: listTagsForResourceInput.ResourceArn,
	}
	out, err := eksSvc.ListTagsForResource(input)

	return out, err
}

func TagResource(input eks.TagResourceInput) (*eks.TagResourceOutput, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	eksSvc := eks.New(sess)

	input = eks.TagResourceInput{
		ResourceArn: input.ResourceArn,
		Tags:        input.Tags,
	}
	out, err := eksSvc.TagResource(&input)

	return out, err
}

func UntagResource(input eks.UntagResourceInput) (*eks.UntagResourceOutput, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	eksSvc := eks.New(sess)

	fmt.Println(input.TagKeys)
	input = eks.UntagResourceInput{
		ResourceArn: input.ResourceArn,
		TagKeys:     input.TagKeys,
	}
	out, err := eksSvc.UntagResource(&input)

	return out, err
}
