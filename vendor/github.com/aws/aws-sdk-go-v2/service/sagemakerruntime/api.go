// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakerruntime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

const opInvokeEndpoint = "InvokeEndpoint"

// InvokeEndpointRequest is a API request type for the InvokeEndpoint API operation.
type InvokeEndpointRequest struct {
	*aws.Request
	Input *InvokeEndpointInput
	Copy  func(*InvokeEndpointInput) InvokeEndpointRequest
}

// Send marshals and sends the InvokeEndpoint API request.
func (r InvokeEndpointRequest) Send(ctx context.Context) (*InvokeEndpointOutput, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*InvokeEndpointOutput), nil
}

// InvokeEndpointRequest returns a request value for making API operation for
// Amazon SageMaker Runtime.
//
// After you deploy a model into production using Amazon SageMaker hosting services,
// your client applications use this API to get inferences from the model hosted
// at the specified endpoint.
//
// For an overview of Amazon SageMaker, see How It Works (http://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html).
//
// Amazon SageMaker strips all POST headers except those supported by the API.
// Amazon SageMaker might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax.
//
// Cals to InvokeEndpoint are authenticated by using AWS Signature Version 4.
// For information, see Authenticating Requests (AWS Signature Version 4) (http://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html)
// in the Amazon S3 API Reference.
//
// Endpoints are scoped to an individual account, and are not public. The URL
// does not contain the account ID, but Amazon SageMaker determines the account
// ID from the authentication token that is supplied by the caller.
//
//    // Example sending a request using the InvokeEndpointRequest method.
//    req := client.InvokeEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.sagemaker-2017-05-13/InvokeEndpoint
func (c *SageMakerRuntime) InvokeEndpointRequest(input *InvokeEndpointInput) InvokeEndpointRequest {
	op := &aws.Operation{
		Name:       opInvokeEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/endpoints/{EndpointName}/invocations",
	}

	if input == nil {
		input = &InvokeEndpointInput{}
	}

	output := &InvokeEndpointOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return InvokeEndpointRequest{Request: req, Input: input, Copy: c.InvokeEndpointRequest}
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.sagemaker-2017-05-13/InvokeEndpointInput
type InvokeEndpointInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The desired MIME type of the inference in the response.
	Accept *string `location:"header" locationName:"Accept" type:"string"`

	// Provides input data, in the format specified in the ContentType request header.
	// Amazon SageMaker passes all of the data in the body to the model.
	//
	// For information about the format of the request body, see Common Data Formats—Inference
	// (http://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html).
	//
	// Body is a required field
	Body []byte `type:"blob" required:"true"`

	// The MIME type of the input data in the request body.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	CustomAttributes *string `location:"header" locationName:"X-Amzn-SageMaker-Custom-Attributes" type:"string"`

	// The name of the endpoint that you specified when you created the endpoint
	// using the CreateEndpoint (http://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpoint.html)
	// API.
	//
	// EndpointName is a required field
	EndpointName *string `location:"uri" locationName:"EndpointName" type:"string" required:"true"`
}

// String returns the string representation
func (s InvokeEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InvokeEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InvokeEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InvokeEndpointInput"}

	if s.Body == nil {
		invalidParams.Add(aws.NewErrParamRequired("Body"))
	}

	if s.EndpointName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s InvokeEndpointInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Accept != nil {
		v := *s.Accept

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Accept", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CustomAttributes != nil {
		v := *s.CustomAttributes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amzn-SageMaker-Custom-Attributes", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndpointName != nil {
		v := *s.EndpointName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EndpointName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Body", protocol.BytesStream(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.sagemaker-2017-05-13/InvokeEndpointOutput
type InvokeEndpointOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	responseMetadata aws.Response

	// Includes the inference provided by the model.
	//
	// For information about the format of the response body, see Common Data Formats—Inference
	// (http://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html).
	//
	// Body is a required field
	Body []byte `type:"blob" required:"true"`

	// The MIME type of the inference returned in the response body.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	CustomAttributes *string `location:"header" locationName:"X-Amzn-SageMaker-Custom-Attributes" type:"string"`

	// Identifies the production variant that was invoked.
	InvokedProductionVariant *string `location:"header" locationName:"x-Amzn-Invoked-Production-Variant" type:"string"`
}

// String returns the string representation
func (s InvokeEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InvokeEndpointOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s InvokeEndpointOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s InvokeEndpointOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CustomAttributes != nil {
		v := *s.CustomAttributes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amzn-SageMaker-Custom-Attributes", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InvokedProductionVariant != nil {
		v := *s.InvokedProductionVariant

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-Amzn-Invoked-Production-Variant", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Body", protocol.BytesStream(v), metadata)
	}
	return nil
}
