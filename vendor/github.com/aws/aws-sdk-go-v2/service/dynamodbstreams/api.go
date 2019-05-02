// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodbstreams

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const opDescribeStream = "DescribeStream"

// DescribeStreamRequest is a API request type for the DescribeStream API operation.
type DescribeStreamRequest struct {
	*aws.Request
	Input *DescribeStreamInput
	Copy  func(*DescribeStreamInput) DescribeStreamRequest
}

// Send marshals and sends the DescribeStream API request.
func (r DescribeStreamRequest) Send(ctx context.Context) (*DescribeStreamOutput, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*DescribeStreamOutput), nil
}

// DescribeStreamRequest returns a request value for making API operation for
// Amazon DynamoDB Streams.
//
// Returns information about a stream, including the current status of the stream,
// its Amazon Resource Name (ARN), the composition of its shards, and its corresponding
// DynamoDB table.
//
// You can call DescribeStream at a maximum rate of 10 times per second.
//
// Each shard in the stream has a SequenceNumberRange associated with it. If
// the SequenceNumberRange has a StartingSequenceNumber but no EndingSequenceNumber,
// then the shard is still open (able to receive more stream records). If both
// StartingSequenceNumber and EndingSequenceNumber are present, then that shard
// is closed and can no longer receive more data.
//
//    // Example sending a request using the DescribeStreamRequest method.
//    req := client.DescribeStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/DescribeStream
func (c *DynamoDBStreams) DescribeStreamRequest(input *DescribeStreamInput) DescribeStreamRequest {
	op := &aws.Operation{
		Name:       opDescribeStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStreamInput{}
	}

	output := &DescribeStreamOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return DescribeStreamRequest{Request: req, Input: input, Copy: c.DescribeStreamRequest}
}

const opGetRecords = "GetRecords"

// GetRecordsRequest is a API request type for the GetRecords API operation.
type GetRecordsRequest struct {
	*aws.Request
	Input *GetRecordsInput
	Copy  func(*GetRecordsInput) GetRecordsRequest
}

// Send marshals and sends the GetRecords API request.
func (r GetRecordsRequest) Send(ctx context.Context) (*GetRecordsOutput, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetRecordsOutput), nil
}

// GetRecordsRequest returns a request value for making API operation for
// Amazon DynamoDB Streams.
//
// Retrieves the stream records from a given shard.
//
// Specify a shard iterator using the ShardIterator parameter. The shard iterator
// specifies the position in the shard from which you want to start reading
// stream records sequentially. If there are no stream records available in
// the portion of the shard that the iterator points to, GetRecords returns
// an empty list. Note that it might take multiple calls to get to a portion
// of the shard that contains stream records.
//
// GetRecords can retrieve a maximum of 1 MB of data or 1000 stream records,
// whichever comes first.
//
//    // Example sending a request using the GetRecordsRequest method.
//    req := client.GetRecordsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/GetRecords
func (c *DynamoDBStreams) GetRecordsRequest(input *GetRecordsInput) GetRecordsRequest {
	op := &aws.Operation{
		Name:       opGetRecords,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRecordsInput{}
	}

	output := &GetRecordsOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetRecordsRequest{Request: req, Input: input, Copy: c.GetRecordsRequest}
}

const opGetShardIterator = "GetShardIterator"

// GetShardIteratorRequest is a API request type for the GetShardIterator API operation.
type GetShardIteratorRequest struct {
	*aws.Request
	Input *GetShardIteratorInput
	Copy  func(*GetShardIteratorInput) GetShardIteratorRequest
}

// Send marshals and sends the GetShardIterator API request.
func (r GetShardIteratorRequest) Send(ctx context.Context) (*GetShardIteratorOutput, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetShardIteratorOutput), nil
}

// GetShardIteratorRequest returns a request value for making API operation for
// Amazon DynamoDB Streams.
//
// Returns a shard iterator. A shard iterator provides information about how
// to retrieve the stream records from within a shard. Use the shard iterator
// in a subsequent GetRecords request to read the stream records from the shard.
//
// A shard iterator expires 15 minutes after it is returned to the requester.
//
//    // Example sending a request using the GetShardIteratorRequest method.
//    req := client.GetShardIteratorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/GetShardIterator
func (c *DynamoDBStreams) GetShardIteratorRequest(input *GetShardIteratorInput) GetShardIteratorRequest {
	op := &aws.Operation{
		Name:       opGetShardIterator,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetShardIteratorInput{}
	}

	output := &GetShardIteratorOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetShardIteratorRequest{Request: req, Input: input, Copy: c.GetShardIteratorRequest}
}

const opListStreams = "ListStreams"

// ListStreamsRequest is a API request type for the ListStreams API operation.
type ListStreamsRequest struct {
	*aws.Request
	Input *ListStreamsInput
	Copy  func(*ListStreamsInput) ListStreamsRequest
}

// Send marshals and sends the ListStreams API request.
func (r ListStreamsRequest) Send(ctx context.Context) (*ListStreamsOutput, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*ListStreamsOutput), nil
}

// ListStreamsRequest returns a request value for making API operation for
// Amazon DynamoDB Streams.
//
// Returns an array of stream ARNs associated with the current account and endpoint.
// If the TableName parameter is present, then ListStreams will return only
// the streams ARNs for that table.
//
// You can call ListStreams at a maximum rate of 5 times per second.
//
//    // Example sending a request using the ListStreamsRequest method.
//    req := client.ListStreamsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/ListStreams
func (c *DynamoDBStreams) ListStreamsRequest(input *ListStreamsInput) ListStreamsRequest {
	op := &aws.Operation{
		Name:       opListStreams,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListStreamsInput{}
	}

	output := &ListStreamsOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return ListStreamsRequest{Request: req, Input: input, Copy: c.ListStreamsRequest}
}

// Represents the input of a DescribeStream operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/DescribeStreamInput
type DescribeStreamInput struct {
	_ struct{} `type:"structure"`

	// The shard ID of the first item that this operation will evaluate. Use the
	// value that was returned for LastEvaluatedShardId in the previous operation.
	ExclusiveStartShardId *string `min:"28" type:"string"`

	// The maximum number of shard objects to return. The upper limit is 100.
	Limit *int64 `min:"1" type:"integer"`

	// The Amazon Resource Name (ARN) for the stream.
	//
	// StreamArn is a required field
	StreamArn *string `min:"37" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStreamInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeStreamInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStreamInput"}
	if s.ExclusiveStartShardId != nil && len(*s.ExclusiveStartShardId) < 28 {
		invalidParams.Add(aws.NewErrParamMinLen("ExclusiveStartShardId", 28))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.StreamArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamArn"))
	}
	if s.StreamArn != nil && len(*s.StreamArn) < 37 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamArn", 37))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a DescribeStream operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/DescribeStreamOutput
type DescribeStreamOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// A complete description of the stream, including its creation date and time,
	// the DynamoDB table associated with the stream, the shard IDs within the stream,
	// and the beginning and ending sequence numbers of stream records within the
	// shards.
	StreamDescription *StreamDescription `type:"structure"`
}

// String returns the string representation
func (s DescribeStreamOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeStreamOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s DescribeStreamOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Represents the input of a GetRecords operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/GetRecordsInput
type GetRecordsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of records to return from the shard. The upper limit is
	// 1000.
	Limit *int64 `min:"1" type:"integer"`

	// A shard iterator that was retrieved from a previous GetShardIterator operation.
	// This iterator can be used to access the stream records in this shard.
	//
	// ShardIterator is a required field
	ShardIterator *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRecordsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRecordsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRecordsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRecordsInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.ShardIterator == nil {
		invalidParams.Add(aws.NewErrParamRequired("ShardIterator"))
	}
	if s.ShardIterator != nil && len(*s.ShardIterator) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ShardIterator", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetRecords operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/GetRecordsOutput
type GetRecordsOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The next position in the shard from which to start sequentially reading stream
	// records. If set to null, the shard has been closed and the requested iterator
	// will not return any more data.
	NextShardIterator *string `min:"1" type:"string"`

	// The stream records from the shard, which were retrieved using the shard iterator.
	Records []Record `type:"list"`
}

// String returns the string representation
func (s GetRecordsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRecordsOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetRecordsOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Represents the input of a GetShardIterator operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/GetShardIteratorInput
type GetShardIteratorInput struct {
	_ struct{} `type:"structure"`

	// The sequence number of a stream record in the shard from which to start reading.
	SequenceNumber *string `min:"21" type:"string"`

	// The identifier of the shard. The iterator will be returned for this shard
	// ID.
	//
	// ShardId is a required field
	ShardId *string `min:"28" type:"string" required:"true"`

	// Determines how the shard iterator is used to start reading stream records
	// from the shard:
	//
	//    * AT_SEQUENCE_NUMBER - Start reading exactly from the position denoted
	//    by a specific sequence number.
	//
	//    * AFTER_SEQUENCE_NUMBER - Start reading right after the position denoted
	//    by a specific sequence number.
	//
	//    * TRIM_HORIZON - Start reading at the last (untrimmed) stream record,
	//    which is the oldest record in the shard. In DynamoDB Streams, there is
	//    a 24 hour limit on data retention. Stream records whose age exceeds this
	//    limit are subject to removal (trimming) from the stream.
	//
	//    * LATEST - Start reading just after the most recent stream record in the
	//    shard, so that you always read the most recent data in the shard.
	//
	// ShardIteratorType is a required field
	ShardIteratorType ShardIteratorType `type:"string" required:"true" enum:"true"`

	// The Amazon Resource Name (ARN) for the stream.
	//
	// StreamArn is a required field
	StreamArn *string `min:"37" type:"string" required:"true"`
}

// String returns the string representation
func (s GetShardIteratorInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetShardIteratorInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetShardIteratorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetShardIteratorInput"}
	if s.SequenceNumber != nil && len(*s.SequenceNumber) < 21 {
		invalidParams.Add(aws.NewErrParamMinLen("SequenceNumber", 21))
	}

	if s.ShardId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ShardId"))
	}
	if s.ShardId != nil && len(*s.ShardId) < 28 {
		invalidParams.Add(aws.NewErrParamMinLen("ShardId", 28))
	}
	if len(s.ShardIteratorType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ShardIteratorType"))
	}

	if s.StreamArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamArn"))
	}
	if s.StreamArn != nil && len(*s.StreamArn) < 37 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamArn", 37))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetShardIterator operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/GetShardIteratorOutput
type GetShardIteratorOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The position in the shard from which to start reading stream records sequentially.
	// A shard iterator specifies this position using the sequence number of a stream
	// record in a shard.
	ShardIterator *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetShardIteratorOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetShardIteratorOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetShardIteratorOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Contains details about the type of identity that made the request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/Identity
type Identity struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the entity that made the call. For Time To Live,
	// the principalId is "dynamodb.amazonaws.com".
	PrincipalId *string `type:"string"`

	// The type of the identity. For Time To Live, the type is "Service".
	Type *string `type:"string"`
}

// String returns the string representation
func (s Identity) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Identity) GoString() string {
	return s.String()
}

// Represents the input of a ListStreams operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/ListStreamsInput
type ListStreamsInput struct {
	_ struct{} `type:"structure"`

	// The ARN (Amazon Resource Name) of the first item that this operation will
	// evaluate. Use the value that was returned for LastEvaluatedStreamArn in the
	// previous operation.
	ExclusiveStartStreamArn *string `min:"37" type:"string"`

	// The maximum number of streams to return. The upper limit is 100.
	Limit *int64 `min:"1" type:"integer"`

	// If this parameter is provided, then only the streams associated with this
	// table name are returned.
	TableName *string `min:"3" type:"string"`
}

// String returns the string representation
func (s ListStreamsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListStreamsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStreamsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStreamsInput"}
	if s.ExclusiveStartStreamArn != nil && len(*s.ExclusiveStartStreamArn) < 37 {
		invalidParams.Add(aws.NewErrParamMinLen("ExclusiveStartStreamArn", 37))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.TableName != nil && len(*s.TableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a ListStreams operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/ListStreamsOutput
type ListStreamsOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The stream ARN of the item where the operation stopped, inclusive of the
	// previous result set. Use this value to start a new operation, excluding this
	// value in the new request.
	//
	// If LastEvaluatedStreamArn is empty, then the "last page" of results has been
	// processed and there is no more data to be retrieved.
	//
	// If LastEvaluatedStreamArn is not empty, it does not necessarily mean that
	// there is more data in the result set. The only way to know when you have
	// reached the end of the result set is when LastEvaluatedStreamArn is empty.
	LastEvaluatedStreamArn *string `min:"37" type:"string"`

	// A list of stream descriptors associated with the current account and endpoint.
	Streams []Stream `type:"list"`
}

// String returns the string representation
func (s ListStreamsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListStreamsOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s ListStreamsOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// A description of a unique event within a stream.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/Record
type Record struct {
	_ struct{} `type:"structure"`

	// The region in which the GetRecords request was received.
	AwsRegion *string `locationName:"awsRegion" type:"string"`

	// The main body of the stream record, containing all of the DynamoDB-specific
	// fields.
	Dynamodb *StreamRecord `locationName:"dynamodb" type:"structure"`

	// A globally unique identifier for the event that was recorded in this stream
	// record.
	EventID *string `locationName:"eventID" type:"string"`

	// The type of data modification that was performed on the DynamoDB table:
	//
	//    * INSERT - a new item was added to the table.
	//
	//    * MODIFY - one or more of an existing item's attributes were modified.
	//
	//    * REMOVE - the item was deleted from the table
	EventName OperationType `locationName:"eventName" type:"string" enum:"true"`

	// The AWS service from which the stream record originated. For DynamoDB Streams,
	// this is aws:dynamodb.
	EventSource *string `locationName:"eventSource" type:"string"`

	// The version number of the stream record format. This number is updated whenever
	// the structure of Record is modified.
	//
	// Client applications must not assume that eventVersion will remain at a particular
	// value, as this number is subject to change at any time. In general, eventVersion
	// will only increase as the low-level DynamoDB Streams API evolves.
	EventVersion *string `locationName:"eventVersion" type:"string"`

	// Items that are deleted by the Time to Live process after expiration have
	// the following fields:
	//
	//    * Records[].userIdentity.type
	//
	// "Service"
	//
	//    * Records[].userIdentity.principalId
	//
	// "dynamodb.amazonaws.com"
	UserIdentity *Identity `locationName:"userIdentity" type:"structure"`
}

// String returns the string representation
func (s Record) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Record) GoString() string {
	return s.String()
}

// The beginning and ending sequence numbers for the stream records contained
// within a shard.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/SequenceNumberRange
type SequenceNumberRange struct {
	_ struct{} `type:"structure"`

	// The last sequence number.
	EndingSequenceNumber *string `min:"21" type:"string"`

	// The first sequence number.
	StartingSequenceNumber *string `min:"21" type:"string"`
}

// String returns the string representation
func (s SequenceNumberRange) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SequenceNumberRange) GoString() string {
	return s.String()
}

// A uniquely identified group of stream records within a stream.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/Shard
type Shard struct {
	_ struct{} `type:"structure"`

	// The shard ID of the current shard's parent.
	ParentShardId *string `min:"28" type:"string"`

	// The range of possible sequence numbers for the shard.
	SequenceNumberRange *SequenceNumberRange `type:"structure"`

	// The system-generated identifier for this shard.
	ShardId *string `min:"28" type:"string"`
}

// String returns the string representation
func (s Shard) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Shard) GoString() string {
	return s.String()
}

// Represents all of the data describing a particular stream.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/Stream
type Stream struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the stream.
	StreamArn *string `min:"37" type:"string"`

	// A timestamp, in ISO 8601 format, for this stream.
	//
	// Note that LatestStreamLabel is not a unique identifier for the stream, because
	// it is possible that a stream from another table might have the same timestamp.
	// However, the combination of the following three elements is guaranteed to
	// be unique:
	//
	//    * the AWS customer ID.
	//
	//    * the table name
	//
	//    * the StreamLabel
	StreamLabel *string `type:"string"`

	// The DynamoDB table with which the stream is associated.
	TableName *string `min:"3" type:"string"`
}

// String returns the string representation
func (s Stream) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Stream) GoString() string {
	return s.String()
}

// Represents all of the data describing a particular stream.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/StreamDescription
type StreamDescription struct {
	_ struct{} `type:"structure"`

	// The date and time when the request to create this stream was issued.
	CreationRequestDateTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The key attribute(s) of the stream's DynamoDB table.
	KeySchema []dynamodb.KeySchemaElement `min:"1" type:"list"`

	// The shard ID of the item where the operation stopped, inclusive of the previous
	// result set. Use this value to start a new operation, excluding this value
	// in the new request.
	//
	// If LastEvaluatedShardId is empty, then the "last page" of results has been
	// processed and there is currently no more data to be retrieved.
	//
	// If LastEvaluatedShardId is not empty, it does not necessarily mean that there
	// is more data in the result set. The only way to know when you have reached
	// the end of the result set is when LastEvaluatedShardId is empty.
	LastEvaluatedShardId *string `min:"28" type:"string"`

	// The shards that comprise the stream.
	Shards []Shard `type:"list"`

	// The Amazon Resource Name (ARN) for the stream.
	StreamArn *string `min:"37" type:"string"`

	// A timestamp, in ISO 8601 format, for this stream.
	//
	// Note that LatestStreamLabel is not a unique identifier for the stream, because
	// it is possible that a stream from another table might have the same timestamp.
	// However, the combination of the following three elements is guaranteed to
	// be unique:
	//
	//    * the AWS customer ID.
	//
	//    * the table name
	//
	//    * the StreamLabel
	StreamLabel *string `type:"string"`

	// Indicates the current status of the stream:
	//
	//    * ENABLING - Streams is currently being enabled on the DynamoDB table.
	//
	//    * ENABLED - the stream is enabled.
	//
	//    * DISABLING - Streams is currently being disabled on the DynamoDB table.
	//
	//    * DISABLED - the stream is disabled.
	StreamStatus StreamStatus `type:"string" enum:"true"`

	// Indicates the format of the records within this stream:
	//
	//    * KEYS_ONLY - only the key attributes of items that were modified in the
	//    DynamoDB table.
	//
	//    * NEW_IMAGE - entire items from the table, as they appeared after they
	//    were modified.
	//
	//    * OLD_IMAGE - entire items from the table, as they appeared before they
	//    were modified.
	//
	//    * NEW_AND_OLD_IMAGES - both the new and the old images of the items from
	//    the table.
	StreamViewType StreamViewType `type:"string" enum:"true"`

	// The DynamoDB table with which the stream is associated.
	TableName *string `min:"3" type:"string"`
}

// String returns the string representation
func (s StreamDescription) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StreamDescription) GoString() string {
	return s.String()
}

// A description of a single data modification that was performed on an item
// in a DynamoDB table.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/StreamRecord
type StreamRecord struct {
	_ struct{} `type:"structure"`

	// The approximate date and time when the stream record was created, in UNIX
	// epoch time (http://www.epochconverter.com/) format.
	ApproximateCreationDateTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The primary key attribute(s) for the DynamoDB item that was modified.
	Keys map[string]dynamodb.AttributeValue `type:"map"`

	// The item in the DynamoDB table as it appeared after it was modified.
	NewImage map[string]dynamodb.AttributeValue `type:"map"`

	// The item in the DynamoDB table as it appeared before it was modified.
	OldImage map[string]dynamodb.AttributeValue `type:"map"`

	// The sequence number of the stream record.
	SequenceNumber *string `min:"21" type:"string"`

	// The size of the stream record, in bytes.
	SizeBytes *int64 `min:"1" type:"long"`

	// The type of data from the modified DynamoDB item that was captured in this
	// stream record:
	//
	//    * KEYS_ONLY - only the key attributes of the modified item.
	//
	//    * NEW_IMAGE - the entire item, as it appeared after it was modified.
	//
	//    * OLD_IMAGE - the entire item, as it appeared before it was modified.
	//
	//    * NEW_AND_OLD_IMAGES - both the new and the old item images of the item.
	StreamViewType StreamViewType `type:"string" enum:"true"`
}

// String returns the string representation
func (s StreamRecord) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StreamRecord) GoString() string {
	return s.String()
}

type KeyType string

// Enum values for KeyType
const (
	KeyTypeHash  KeyType = "HASH"
	KeyTypeRange KeyType = "RANGE"
)

func (enum KeyType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KeyType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OperationType string

// Enum values for OperationType
const (
	OperationTypeInsert OperationType = "INSERT"
	OperationTypeModify OperationType = "MODIFY"
	OperationTypeRemove OperationType = "REMOVE"
)

func (enum OperationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OperationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ShardIteratorType string

// Enum values for ShardIteratorType
const (
	ShardIteratorTypeTrimHorizon         ShardIteratorType = "TRIM_HORIZON"
	ShardIteratorTypeLatest              ShardIteratorType = "LATEST"
	ShardIteratorTypeAtSequenceNumber    ShardIteratorType = "AT_SEQUENCE_NUMBER"
	ShardIteratorTypeAfterSequenceNumber ShardIteratorType = "AFTER_SEQUENCE_NUMBER"
)

func (enum ShardIteratorType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ShardIteratorType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StreamStatus string

// Enum values for StreamStatus
const (
	StreamStatusEnabling  StreamStatus = "ENABLING"
	StreamStatusEnabled   StreamStatus = "ENABLED"
	StreamStatusDisabling StreamStatus = "DISABLING"
	StreamStatusDisabled  StreamStatus = "DISABLED"
)

func (enum StreamStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StreamStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StreamViewType string

// Enum values for StreamViewType
const (
	StreamViewTypeNewImage        StreamViewType = "NEW_IMAGE"
	StreamViewTypeOldImage        StreamViewType = "OLD_IMAGE"
	StreamViewTypeNewAndOldImages StreamViewType = "NEW_AND_OLD_IMAGES"
	StreamViewTypeKeysOnly        StreamViewType = "KEYS_ONLY"
)

func (enum StreamViewType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StreamViewType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
