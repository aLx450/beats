// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFpgaImageAttributeRequest
type DescribeFpgaImageAttributeInput struct {
	_ struct{} `type:"structure"`

	// The AFI attribute.
	//
	// Attribute is a required field
	Attribute FpgaImageAttributeName `type:"string" required:"true" enum:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the AFI.
	//
	// FpgaImageId is a required field
	FpgaImageId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeFpgaImageAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFpgaImageAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFpgaImageAttributeInput"}
	if len(s.Attribute) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Attribute"))
	}

	if s.FpgaImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FpgaImageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFpgaImageAttributeResult
type DescribeFpgaImageAttributeOutput struct {
	_ struct{} `type:"structure"`

	// Information about the attribute.
	FpgaImageAttribute *FpgaImageAttribute `locationName:"fpgaImageAttribute" type:"structure"`
}

// String returns the string representation
func (s DescribeFpgaImageAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFpgaImageAttribute = "DescribeFpgaImageAttribute"

// DescribeFpgaImageAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified attribute of the specified Amazon FPGA Image (AFI).
//
//    // Example sending a request using DescribeFpgaImageAttributeRequest.
//    req := client.DescribeFpgaImageAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFpgaImageAttribute
func (c *Client) DescribeFpgaImageAttributeRequest(input *DescribeFpgaImageAttributeInput) DescribeFpgaImageAttributeRequest {
	op := &aws.Operation{
		Name:       opDescribeFpgaImageAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeFpgaImageAttributeInput{}
	}

	req := c.newRequest(op, input, &DescribeFpgaImageAttributeOutput{})
	return DescribeFpgaImageAttributeRequest{Request: req, Input: input, Copy: c.DescribeFpgaImageAttributeRequest}
}

// DescribeFpgaImageAttributeRequest is the request type for the
// DescribeFpgaImageAttribute API operation.
type DescribeFpgaImageAttributeRequest struct {
	*aws.Request
	Input *DescribeFpgaImageAttributeInput
	Copy  func(*DescribeFpgaImageAttributeInput) DescribeFpgaImageAttributeRequest
}

// Send marshals and sends the DescribeFpgaImageAttribute API request.
func (r DescribeFpgaImageAttributeRequest) Send(ctx context.Context) (*DescribeFpgaImageAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFpgaImageAttributeResponse{
		DescribeFpgaImageAttributeOutput: r.Request.Data.(*DescribeFpgaImageAttributeOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFpgaImageAttributeResponse is the response type for the
// DescribeFpgaImageAttribute API operation.
type DescribeFpgaImageAttributeResponse struct {
	*DescribeFpgaImageAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFpgaImageAttribute request.
func (r *DescribeFpgaImageAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}