// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package efsiface provides an interface to enable mocking the Amazon Elastic File System service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package efsiface

import (
	"github.com/aws/aws-sdk-go-v2/service/efs"
)

// EFSAPI provides an interface to enable mocking the
// efs.EFS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Elastic File System.
//    func myFunc(svc efsiface.EFSAPI) bool {
//        // Make svc.CreateFileSystem request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := efs.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockEFSClient struct {
//        efsiface.EFSAPI
//    }
//    func (m *mockEFSClient) CreateFileSystem(input *efs.CreateFileSystemInput) (*efs.UpdateFileSystemOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockEFSClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type EFSAPI interface {
	CreateFileSystemRequest(*efs.CreateFileSystemInput) efs.CreateFileSystemRequest

	CreateMountTargetRequest(*efs.CreateMountTargetInput) efs.CreateMountTargetRequest

	CreateTagsRequest(*efs.CreateTagsInput) efs.CreateTagsRequest

	DeleteFileSystemRequest(*efs.DeleteFileSystemInput) efs.DeleteFileSystemRequest

	DeleteMountTargetRequest(*efs.DeleteMountTargetInput) efs.DeleteMountTargetRequest

	DeleteTagsRequest(*efs.DeleteTagsInput) efs.DeleteTagsRequest

	DescribeFileSystemsRequest(*efs.DescribeFileSystemsInput) efs.DescribeFileSystemsRequest

	DescribeLifecycleConfigurationRequest(*efs.DescribeLifecycleConfigurationInput) efs.DescribeLifecycleConfigurationRequest

	DescribeMountTargetSecurityGroupsRequest(*efs.DescribeMountTargetSecurityGroupsInput) efs.DescribeMountTargetSecurityGroupsRequest

	DescribeMountTargetsRequest(*efs.DescribeMountTargetsInput) efs.DescribeMountTargetsRequest

	DescribeTagsRequest(*efs.DescribeTagsInput) efs.DescribeTagsRequest

	ModifyMountTargetSecurityGroupsRequest(*efs.ModifyMountTargetSecurityGroupsInput) efs.ModifyMountTargetSecurityGroupsRequest

	PutLifecycleConfigurationRequest(*efs.PutLifecycleConfigurationInput) efs.PutLifecycleConfigurationRequest

	UpdateFileSystemRequest(*efs.UpdateFileSystemInput) efs.UpdateFileSystemRequest
}

var _ EFSAPI = (*efs.EFS)(nil)
