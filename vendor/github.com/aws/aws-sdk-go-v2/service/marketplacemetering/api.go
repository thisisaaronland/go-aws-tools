// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacemetering

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

const opBatchMeterUsage = "BatchMeterUsage"

// BatchMeterUsageRequest is a API request type for the BatchMeterUsage API operation.
type BatchMeterUsageRequest struct {
	*aws.Request
	Input *BatchMeterUsageInput
	Copy  func(*BatchMeterUsageInput) BatchMeterUsageRequest
}

// Send marshals and sends the BatchMeterUsage API request.
func (r BatchMeterUsageRequest) Send(ctx context.Context) (*BatchMeterUsageOutput, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*BatchMeterUsageOutput), nil
}

// BatchMeterUsageRequest returns a request value for making API operation for
// AWSMarketplace Metering.
//
// BatchMeterUsage is called from a SaaS application listed on the AWS Marketplace
// to post metering records for a set of customers.
//
// For identical requests, the API is idempotent; requests can be retried with
// the same records or a subset of the input records.
//
// Every request to BatchMeterUsage is for one product. If you need to meter
// usage for multiple products, you must make multiple calls to BatchMeterUsage.
//
// BatchMeterUsage can process up to 25 UsageRecords at a time.
//
//    // Example sending a request using the BatchMeterUsageRequest method.
//    req := client.BatchMeterUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/BatchMeterUsage
func (c *MarketplaceMetering) BatchMeterUsageRequest(input *BatchMeterUsageInput) BatchMeterUsageRequest {
	op := &aws.Operation{
		Name:       opBatchMeterUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchMeterUsageInput{}
	}

	output := &BatchMeterUsageOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return BatchMeterUsageRequest{Request: req, Input: input, Copy: c.BatchMeterUsageRequest}
}

const opMeterUsage = "MeterUsage"

// MeterUsageRequest is a API request type for the MeterUsage API operation.
type MeterUsageRequest struct {
	*aws.Request
	Input *MeterUsageInput
	Copy  func(*MeterUsageInput) MeterUsageRequest
}

// Send marshals and sends the MeterUsage API request.
func (r MeterUsageRequest) Send(ctx context.Context) (*MeterUsageOutput, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*MeterUsageOutput), nil
}

// MeterUsageRequest returns a request value for making API operation for
// AWSMarketplace Metering.
//
// API to emit metering records. For identical requests, the API is idempotent.
// It simply returns the metering record ID.
//
// MeterUsage is authenticated on the buyer's AWS account, generally when running
// from an EC2 instance on the AWS Marketplace.
//
//    // Example sending a request using the MeterUsageRequest method.
//    req := client.MeterUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/MeterUsage
func (c *MarketplaceMetering) MeterUsageRequest(input *MeterUsageInput) MeterUsageRequest {
	op := &aws.Operation{
		Name:       opMeterUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MeterUsageInput{}
	}

	output := &MeterUsageOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return MeterUsageRequest{Request: req, Input: input, Copy: c.MeterUsageRequest}
}

const opRegisterUsage = "RegisterUsage"

// RegisterUsageRequest is a API request type for the RegisterUsage API operation.
type RegisterUsageRequest struct {
	*aws.Request
	Input *RegisterUsageInput
	Copy  func(*RegisterUsageInput) RegisterUsageRequest
}

// Send marshals and sends the RegisterUsage API request.
func (r RegisterUsageRequest) Send(ctx context.Context) (*RegisterUsageOutput, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*RegisterUsageOutput), nil
}

// RegisterUsageRequest returns a request value for making API operation for
// AWSMarketplace Metering.
//
// Paid container software products sold through AWS Marketplace must integrate
// with the AWS Marketplace Metering Service and call the RegisterUsage operation
// for software entitlement and metering. Calling RegisterUsage from containers
// running outside of ECS is not currently supported. Free and BYOL products
// for ECS aren't required to call RegisterUsage, but you may choose to do so
// if you would like to receive usage data in your seller reports. The sections
// below explain the behavior of RegisterUsage. RegisterUsage performs two primary
// functions: metering and entitlement.
//
//    * Entitlement: RegisterUsage allows you to verify that the customer running
//    your paid software is subscribed to your product on AWS Marketplace, enabling
//    you to guard against unauthorized use. Your container image that integrates
//    with RegisterUsage is only required to guard against unauthorized use
//    at container startup, as such a CustomerNotSubscribedException/PlatformNotSupportedException
//    will only be thrown on the initial call to RegisterUsage. Subsequent calls
//    from the same Amazon ECS task instance (e.g. task-id) will not throw a
//    CustomerNotSubscribedException, even if the customer unsubscribes while
//    the Amazon ECS task is still running.
//
//    * Metering: RegisterUsage meters software use per ECS task, per hour,
//    with usage prorated to the second. A minimum of 1 minute of usage applies
//    to tasks that are short lived. For example, if a customer has a 10 node
//    ECS cluster and creates an ECS service configured as a Daemon Set, then
//    ECS will launch a task on all 10 cluster nodes and the customer will be
//    charged: (10 * hourly_rate). Metering for software use is automatically
//    handled by the AWS Marketplace Metering Control Plane -- your software
//    is not required to perform any metering specific actions, other than call
//    RegisterUsage once for metering of software use to commence. The AWS Marketplace
//    Metering Control Plane will also continue to bill customers for running
//    ECS tasks, regardless of the customers subscription state, removing the
//    need for your software to perform entitlement checks at runtime.
//
//    // Example sending a request using the RegisterUsageRequest method.
//    req := client.RegisterUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/RegisterUsage
func (c *MarketplaceMetering) RegisterUsageRequest(input *RegisterUsageInput) RegisterUsageRequest {
	op := &aws.Operation{
		Name:       opRegisterUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterUsageInput{}
	}

	output := &RegisterUsageOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return RegisterUsageRequest{Request: req, Input: input, Copy: c.RegisterUsageRequest}
}

const opResolveCustomer = "ResolveCustomer"

// ResolveCustomerRequest is a API request type for the ResolveCustomer API operation.
type ResolveCustomerRequest struct {
	*aws.Request
	Input *ResolveCustomerInput
	Copy  func(*ResolveCustomerInput) ResolveCustomerRequest
}

// Send marshals and sends the ResolveCustomer API request.
func (r ResolveCustomerRequest) Send(ctx context.Context) (*ResolveCustomerOutput, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*ResolveCustomerOutput), nil
}

// ResolveCustomerRequest returns a request value for making API operation for
// AWSMarketplace Metering.
//
// ResolveCustomer is called by a SaaS application during the registration process.
// When a buyer visits your website during the registration process, the buyer
// submits a registration token through their browser. The registration token
// is resolved through this API to obtain a CustomerIdentifier and product code.
//
//    // Example sending a request using the ResolveCustomerRequest method.
//    req := client.ResolveCustomerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/ResolveCustomer
func (c *MarketplaceMetering) ResolveCustomerRequest(input *ResolveCustomerInput) ResolveCustomerRequest {
	op := &aws.Operation{
		Name:       opResolveCustomer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResolveCustomerInput{}
	}

	output := &ResolveCustomerOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return ResolveCustomerRequest{Request: req, Input: input, Copy: c.ResolveCustomerRequest}
}

// A BatchMeterUsageRequest contains UsageRecords, which indicate quantities
// of usage within your application.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/BatchMeterUsageRequest
type BatchMeterUsageInput struct {
	_ struct{} `type:"structure"`

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of
	// a new product.
	//
	// ProductCode is a required field
	ProductCode *string `min:"1" type:"string" required:"true"`

	// The set of UsageRecords to submit. BatchMeterUsage accepts up to 25 UsageRecords
	// at a time.
	//
	// UsageRecords is a required field
	UsageRecords []UsageRecord `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchMeterUsageInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchMeterUsageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchMeterUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchMeterUsageInput"}

	if s.ProductCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductCode"))
	}
	if s.ProductCode != nil && len(*s.ProductCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductCode", 1))
	}

	if s.UsageRecords == nil {
		invalidParams.Add(aws.NewErrParamRequired("UsageRecords"))
	}
	if s.UsageRecords != nil {
		for i, v := range s.UsageRecords {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UsageRecords", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the UsageRecords processed by BatchMeterUsage and any records that
// have failed due to transient error.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/BatchMeterUsageResult
type BatchMeterUsageOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// Contains all UsageRecords processed by BatchMeterUsage. These records were
	// either honored by AWS Marketplace Metering Service or were invalid.
	Results []UsageRecordResult `type:"list"`

	// Contains all UsageRecords that were not processed by BatchMeterUsage. This
	// is a list of UsageRecords. You can retry the failed request by making another
	// BatchMeterUsage call with this list as input in the BatchMeterUsageRequest.
	UnprocessedRecords []UsageRecord `type:"list"`
}

// String returns the string representation
func (s BatchMeterUsageOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchMeterUsageOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s BatchMeterUsageOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/MeterUsageRequest
type MeterUsageInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the permissions required for the action, but does
	// not make the request. If you have the permissions, the request returns DryRunOperation;
	// otherwise, it returns UnauthorizedException. Defaults to false if not specified.
	DryRun *bool `type:"boolean"`

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of
	// a new product.
	//
	// ProductCode is a required field
	ProductCode *string `min:"1" type:"string" required:"true"`

	// Timestamp of the hour, recorded in UTC. The seconds and milliseconds portions
	// of the timestamp will be ignored.
	//
	// Timestamp is a required field
	Timestamp *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// It will be one of the fcp dimension name provided during the publishing of
	// the product.
	//
	// UsageDimension is a required field
	UsageDimension *string `min:"1" type:"string" required:"true"`

	// Consumption value for the hour. Defaults to 0 if not specified.
	UsageQuantity *int64 `type:"integer"`
}

// String returns the string representation
func (s MeterUsageInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MeterUsageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MeterUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MeterUsageInput"}

	if s.ProductCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductCode"))
	}
	if s.ProductCode != nil && len(*s.ProductCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductCode", 1))
	}

	if s.Timestamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("Timestamp"))
	}

	if s.UsageDimension == nil {
		invalidParams.Add(aws.NewErrParamRequired("UsageDimension"))
	}
	if s.UsageDimension != nil && len(*s.UsageDimension) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UsageDimension", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/MeterUsageResult
type MeterUsageOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// Metering record id.
	MeteringRecordId *string `type:"string"`
}

// String returns the string representation
func (s MeterUsageOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MeterUsageOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s MeterUsageOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/RegisterUsageRequest
type RegisterUsageInput struct {
	_ struct{} `type:"structure"`

	// (Optional) To scope down the registration to a specific running software
	// instance and guard against replay attacks.
	Nonce *string `type:"string"`

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of
	// a new product.
	//
	// ProductCode is a required field
	ProductCode *string `min:"1" type:"string" required:"true"`

	// Public Key Version provided by AWS Marketplace
	//
	// PublicKeyVersion is a required field
	PublicKeyVersion *int64 `min:"1" type:"integer" required:"true"`
}

// String returns the string representation
func (s RegisterUsageInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RegisterUsageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterUsageInput"}

	if s.ProductCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductCode"))
	}
	if s.ProductCode != nil && len(*s.ProductCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductCode", 1))
	}

	if s.PublicKeyVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("PublicKeyVersion"))
	}
	if s.PublicKeyVersion != nil && *s.PublicKeyVersion < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PublicKeyVersion", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/RegisterUsageResult
type RegisterUsageOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// (Optional) Only included when public key version has expired
	PublicKeyRotationTimestamp *time.Time `type:"timestamp" timestampFormat:"unix"`

	// JWT Token
	Signature *string `type:"string"`
}

// String returns the string representation
func (s RegisterUsageOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RegisterUsageOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s RegisterUsageOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Contains input to the ResolveCustomer operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/ResolveCustomerRequest
type ResolveCustomerInput struct {
	_ struct{} `type:"structure"`

	// When a buyer visits your website during the registration process, the buyer
	// submits a registration token through the browser. The registration token
	// is resolved to obtain a CustomerIdentifier and product code.
	//
	// RegistrationToken is a required field
	RegistrationToken *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResolveCustomerInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResolveCustomerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResolveCustomerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResolveCustomerInput"}

	if s.RegistrationToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistrationToken"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of the ResolveCustomer operation. Contains the CustomerIdentifier
// and product code.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/ResolveCustomerResult
type ResolveCustomerOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The CustomerIdentifier is used to identify an individual customer in your
	// application. Calls to BatchMeterUsage require CustomerIdentifiers for each
	// UsageRecord.
	CustomerIdentifier *string `min:"1" type:"string"`

	// The product code is returned to confirm that the buyer is registering for
	// your product. Subsequent BatchMeterUsage calls should be made using this
	// product code.
	ProductCode *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ResolveCustomerOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResolveCustomerOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s ResolveCustomerOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// A UsageRecord indicates a quantity of usage for a given product, customer,
// dimension and time.
//
// Multiple requests with the same UsageRecords as input will be deduplicated
// to prevent double charges.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/UsageRecord
type UsageRecord struct {
	_ struct{} `type:"structure"`

	// The CustomerIdentifier is obtained through the ResolveCustomer operation
	// and represents an individual buyer in your application.
	//
	// CustomerIdentifier is a required field
	CustomerIdentifier *string `min:"1" type:"string" required:"true"`

	// During the process of registering a product on AWS Marketplace, up to eight
	// dimensions are specified. These represent different units of value in your
	// application.
	//
	// Dimension is a required field
	Dimension *string `min:"1" type:"string" required:"true"`

	// The quantity of usage consumed by the customer for the given dimension and
	// time. Defaults to 0 if not specified.
	Quantity *int64 `type:"integer"`

	// Timestamp of the hour, recorded in UTC. The seconds and milliseconds portions
	// of the timestamp will be ignored.
	//
	// Your application can meter usage for up to one hour in the past.
	//
	// Timestamp is a required field
	Timestamp *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s UsageRecord) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UsageRecord) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UsageRecord) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UsageRecord"}

	if s.CustomerIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomerIdentifier"))
	}
	if s.CustomerIdentifier != nil && len(*s.CustomerIdentifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomerIdentifier", 1))
	}

	if s.Dimension == nil {
		invalidParams.Add(aws.NewErrParamRequired("Dimension"))
	}
	if s.Dimension != nil && len(*s.Dimension) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Dimension", 1))
	}

	if s.Timestamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("Timestamp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A UsageRecordResult indicates the status of a given UsageRecord processed
// by BatchMeterUsage.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/UsageRecordResult
type UsageRecordResult struct {
	_ struct{} `type:"structure"`

	// The MeteringRecordId is a unique identifier for this metering event.
	MeteringRecordId *string `type:"string"`

	// The UsageRecordResult Status indicates the status of an individual UsageRecord
	// processed by BatchMeterUsage.
	//
	//    * Success- The UsageRecord was accepted and honored by BatchMeterUsage.
	//
	//    * CustomerNotSubscribed- The CustomerIdentifier specified is not subscribed
	//    to your product. The UsageRecord was not honored. Future UsageRecords
	//    for this customer will fail until the customer subscribes to your product.
	//
	//    * DuplicateRecord- Indicates that the UsageRecord was invalid and not
	//    honored. A previously metered UsageRecord had the same customer, dimension,
	//    and time, but a different quantity.
	Status UsageRecordResultStatus `type:"string" enum:"true"`

	// The UsageRecord that was part of the BatchMeterUsage request.
	UsageRecord *UsageRecord `type:"structure"`
}

// String returns the string representation
func (s UsageRecordResult) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UsageRecordResult) GoString() string {
	return s.String()
}

type UsageRecordResultStatus string

// Enum values for UsageRecordResultStatus
const (
	UsageRecordResultStatusSuccess               UsageRecordResultStatus = "Success"
	UsageRecordResultStatusCustomerNotSubscribed UsageRecordResultStatus = "CustomerNotSubscribed"
	UsageRecordResultStatusDuplicateRecord       UsageRecordResultStatus = "DuplicateRecord"
)

func (enum UsageRecordResultStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UsageRecordResultStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
