// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rdsiface provides an interface to enable mocking the Amazon Relational Database Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rdsiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

// RDSAPI provides an interface to enable mocking the
// rds.RDS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Relational Database Service.
//    func myFunc(svc rdsiface.RDSAPI) bool {
//        // Make svc.AddRoleToDBCluster request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := rds.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRDSClient struct {
//        rdsiface.RDSAPI
//    }
//    func (m *mockRDSClient) AddRoleToDBCluster(input *rds.AddRoleToDBClusterInput) (*rds.AddRoleToDBClusterOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRDSClient{}
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
type RDSAPI interface {
	AddRoleToDBClusterRequest(*rds.AddRoleToDBClusterInput) rds.AddRoleToDBClusterRequest

	AddSourceIdentifierToSubscriptionRequest(*rds.AddSourceIdentifierToSubscriptionInput) rds.AddSourceIdentifierToSubscriptionRequest

	AddTagsToResourceRequest(*rds.AddTagsToResourceInput) rds.AddTagsToResourceRequest

	ApplyPendingMaintenanceActionRequest(*rds.ApplyPendingMaintenanceActionInput) rds.ApplyPendingMaintenanceActionRequest

	AuthorizeDBSecurityGroupIngressRequest(*rds.AuthorizeDBSecurityGroupIngressInput) rds.AuthorizeDBSecurityGroupIngressRequest

	BacktrackDBClusterRequest(*rds.BacktrackDBClusterInput) rds.BacktrackDBClusterRequest

	CopyDBClusterParameterGroupRequest(*rds.CopyDBClusterParameterGroupInput) rds.CopyDBClusterParameterGroupRequest

	CopyDBClusterSnapshotRequest(*rds.CopyDBClusterSnapshotInput) rds.CopyDBClusterSnapshotRequest

	CopyDBParameterGroupRequest(*rds.CopyDBParameterGroupInput) rds.CopyDBParameterGroupRequest

	CopyDBSnapshotRequest(*rds.CopyDBSnapshotInput) rds.CopyDBSnapshotRequest

	CopyOptionGroupRequest(*rds.CopyOptionGroupInput) rds.CopyOptionGroupRequest

	CreateDBClusterRequest(*rds.CreateDBClusterInput) rds.CreateDBClusterRequest

	CreateDBClusterParameterGroupRequest(*rds.CreateDBClusterParameterGroupInput) rds.CreateDBClusterParameterGroupRequest

	CreateDBClusterSnapshotRequest(*rds.CreateDBClusterSnapshotInput) rds.CreateDBClusterSnapshotRequest

	CreateDBInstanceRequest(*rds.CreateDBInstanceInput) rds.CreateDBInstanceRequest

	CreateDBInstanceReadReplicaRequest(*rds.CreateDBInstanceReadReplicaInput) rds.CreateDBInstanceReadReplicaRequest

	CreateDBParameterGroupRequest(*rds.CreateDBParameterGroupInput) rds.CreateDBParameterGroupRequest

	CreateDBSecurityGroupRequest(*rds.CreateDBSecurityGroupInput) rds.CreateDBSecurityGroupRequest

	CreateDBSnapshotRequest(*rds.CreateDBSnapshotInput) rds.CreateDBSnapshotRequest

	CreateDBSubnetGroupRequest(*rds.CreateDBSubnetGroupInput) rds.CreateDBSubnetGroupRequest

	CreateEventSubscriptionRequest(*rds.CreateEventSubscriptionInput) rds.CreateEventSubscriptionRequest

	CreateOptionGroupRequest(*rds.CreateOptionGroupInput) rds.CreateOptionGroupRequest

	DeleteDBClusterRequest(*rds.DeleteDBClusterInput) rds.DeleteDBClusterRequest

	DeleteDBClusterParameterGroupRequest(*rds.DeleteDBClusterParameterGroupInput) rds.DeleteDBClusterParameterGroupRequest

	DeleteDBClusterSnapshotRequest(*rds.DeleteDBClusterSnapshotInput) rds.DeleteDBClusterSnapshotRequest

	DeleteDBInstanceRequest(*rds.DeleteDBInstanceInput) rds.DeleteDBInstanceRequest

	DeleteDBParameterGroupRequest(*rds.DeleteDBParameterGroupInput) rds.DeleteDBParameterGroupRequest

	DeleteDBSecurityGroupRequest(*rds.DeleteDBSecurityGroupInput) rds.DeleteDBSecurityGroupRequest

	DeleteDBSnapshotRequest(*rds.DeleteDBSnapshotInput) rds.DeleteDBSnapshotRequest

	DeleteDBSubnetGroupRequest(*rds.DeleteDBSubnetGroupInput) rds.DeleteDBSubnetGroupRequest

	DeleteEventSubscriptionRequest(*rds.DeleteEventSubscriptionInput) rds.DeleteEventSubscriptionRequest

	DeleteOptionGroupRequest(*rds.DeleteOptionGroupInput) rds.DeleteOptionGroupRequest

	DescribeAccountAttributesRequest(*rds.DescribeAccountAttributesInput) rds.DescribeAccountAttributesRequest

	DescribeCertificatesRequest(*rds.DescribeCertificatesInput) rds.DescribeCertificatesRequest

	DescribeDBClusterBacktracksRequest(*rds.DescribeDBClusterBacktracksInput) rds.DescribeDBClusterBacktracksRequest

	DescribeDBClusterParameterGroupsRequest(*rds.DescribeDBClusterParameterGroupsInput) rds.DescribeDBClusterParameterGroupsRequest

	DescribeDBClusterParametersRequest(*rds.DescribeDBClusterParametersInput) rds.DescribeDBClusterParametersRequest

	DescribeDBClusterSnapshotAttributesRequest(*rds.DescribeDBClusterSnapshotAttributesInput) rds.DescribeDBClusterSnapshotAttributesRequest

	DescribeDBClusterSnapshotsRequest(*rds.DescribeDBClusterSnapshotsInput) rds.DescribeDBClusterSnapshotsRequest

	DescribeDBClustersRequest(*rds.DescribeDBClustersInput) rds.DescribeDBClustersRequest

	DescribeDBEngineVersionsRequest(*rds.DescribeDBEngineVersionsInput) rds.DescribeDBEngineVersionsRequest

	DescribeDBInstancesRequest(*rds.DescribeDBInstancesInput) rds.DescribeDBInstancesRequest

	DescribeDBLogFilesRequest(*rds.DescribeDBLogFilesInput) rds.DescribeDBLogFilesRequest

	DescribeDBParameterGroupsRequest(*rds.DescribeDBParameterGroupsInput) rds.DescribeDBParameterGroupsRequest

	DescribeDBParametersRequest(*rds.DescribeDBParametersInput) rds.DescribeDBParametersRequest

	DescribeDBSecurityGroupsRequest(*rds.DescribeDBSecurityGroupsInput) rds.DescribeDBSecurityGroupsRequest

	DescribeDBSnapshotAttributesRequest(*rds.DescribeDBSnapshotAttributesInput) rds.DescribeDBSnapshotAttributesRequest

	DescribeDBSnapshotsRequest(*rds.DescribeDBSnapshotsInput) rds.DescribeDBSnapshotsRequest

	DescribeDBSubnetGroupsRequest(*rds.DescribeDBSubnetGroupsInput) rds.DescribeDBSubnetGroupsRequest

	DescribeEngineDefaultClusterParametersRequest(*rds.DescribeEngineDefaultClusterParametersInput) rds.DescribeEngineDefaultClusterParametersRequest

	DescribeEngineDefaultParametersRequest(*rds.DescribeEngineDefaultParametersInput) rds.DescribeEngineDefaultParametersRequest

	DescribeEventCategoriesRequest(*rds.DescribeEventCategoriesInput) rds.DescribeEventCategoriesRequest

	DescribeEventSubscriptionsRequest(*rds.DescribeEventSubscriptionsInput) rds.DescribeEventSubscriptionsRequest

	DescribeEventsRequest(*rds.DescribeEventsInput) rds.DescribeEventsRequest

	DescribeOptionGroupOptionsRequest(*rds.DescribeOptionGroupOptionsInput) rds.DescribeOptionGroupOptionsRequest

	DescribeOptionGroupsRequest(*rds.DescribeOptionGroupsInput) rds.DescribeOptionGroupsRequest

	DescribeOrderableDBInstanceOptionsRequest(*rds.DescribeOrderableDBInstanceOptionsInput) rds.DescribeOrderableDBInstanceOptionsRequest

	DescribePendingMaintenanceActionsRequest(*rds.DescribePendingMaintenanceActionsInput) rds.DescribePendingMaintenanceActionsRequest

	DescribeReservedDBInstancesRequest(*rds.DescribeReservedDBInstancesInput) rds.DescribeReservedDBInstancesRequest

	DescribeReservedDBInstancesOfferingsRequest(*rds.DescribeReservedDBInstancesOfferingsInput) rds.DescribeReservedDBInstancesOfferingsRequest

	DescribeSourceRegionsRequest(*rds.DescribeSourceRegionsInput) rds.DescribeSourceRegionsRequest

	DescribeValidDBInstanceModificationsRequest(*rds.DescribeValidDBInstanceModificationsInput) rds.DescribeValidDBInstanceModificationsRequest

	DownloadDBLogFilePortionRequest(*rds.DownloadDBLogFilePortionInput) rds.DownloadDBLogFilePortionRequest

	FailoverDBClusterRequest(*rds.FailoverDBClusterInput) rds.FailoverDBClusterRequest

	ListTagsForResourceRequest(*rds.ListTagsForResourceInput) rds.ListTagsForResourceRequest

	ModifyDBClusterRequest(*rds.ModifyDBClusterInput) rds.ModifyDBClusterRequest

	ModifyDBClusterParameterGroupRequest(*rds.ModifyDBClusterParameterGroupInput) rds.ModifyDBClusterParameterGroupRequest

	ModifyDBClusterSnapshotAttributeRequest(*rds.ModifyDBClusterSnapshotAttributeInput) rds.ModifyDBClusterSnapshotAttributeRequest

	ModifyDBInstanceRequest(*rds.ModifyDBInstanceInput) rds.ModifyDBInstanceRequest

	ModifyDBParameterGroupRequest(*rds.ModifyDBParameterGroupInput) rds.ModifyDBParameterGroupRequest

	ModifyDBSnapshotRequest(*rds.ModifyDBSnapshotInput) rds.ModifyDBSnapshotRequest

	ModifyDBSnapshotAttributeRequest(*rds.ModifyDBSnapshotAttributeInput) rds.ModifyDBSnapshotAttributeRequest

	ModifyDBSubnetGroupRequest(*rds.ModifyDBSubnetGroupInput) rds.ModifyDBSubnetGroupRequest

	ModifyEventSubscriptionRequest(*rds.ModifyEventSubscriptionInput) rds.ModifyEventSubscriptionRequest

	ModifyOptionGroupRequest(*rds.ModifyOptionGroupInput) rds.ModifyOptionGroupRequest

	PromoteReadReplicaRequest(*rds.PromoteReadReplicaInput) rds.PromoteReadReplicaRequest

	PromoteReadReplicaDBClusterRequest(*rds.PromoteReadReplicaDBClusterInput) rds.PromoteReadReplicaDBClusterRequest

	PurchaseReservedDBInstancesOfferingRequest(*rds.PurchaseReservedDBInstancesOfferingInput) rds.PurchaseReservedDBInstancesOfferingRequest

	RebootDBInstanceRequest(*rds.RebootDBInstanceInput) rds.RebootDBInstanceRequest

	RemoveRoleFromDBClusterRequest(*rds.RemoveRoleFromDBClusterInput) rds.RemoveRoleFromDBClusterRequest

	RemoveSourceIdentifierFromSubscriptionRequest(*rds.RemoveSourceIdentifierFromSubscriptionInput) rds.RemoveSourceIdentifierFromSubscriptionRequest

	RemoveTagsFromResourceRequest(*rds.RemoveTagsFromResourceInput) rds.RemoveTagsFromResourceRequest

	ResetDBClusterParameterGroupRequest(*rds.ResetDBClusterParameterGroupInput) rds.ResetDBClusterParameterGroupRequest

	ResetDBParameterGroupRequest(*rds.ResetDBParameterGroupInput) rds.ResetDBParameterGroupRequest

	RestoreDBClusterFromS3Request(*rds.RestoreDBClusterFromS3Input) rds.RestoreDBClusterFromS3Request

	RestoreDBClusterFromSnapshotRequest(*rds.RestoreDBClusterFromSnapshotInput) rds.RestoreDBClusterFromSnapshotRequest

	RestoreDBClusterToPointInTimeRequest(*rds.RestoreDBClusterToPointInTimeInput) rds.RestoreDBClusterToPointInTimeRequest

	RestoreDBInstanceFromDBSnapshotRequest(*rds.RestoreDBInstanceFromDBSnapshotInput) rds.RestoreDBInstanceFromDBSnapshotRequest

	RestoreDBInstanceFromS3Request(*rds.RestoreDBInstanceFromS3Input) rds.RestoreDBInstanceFromS3Request

	RestoreDBInstanceToPointInTimeRequest(*rds.RestoreDBInstanceToPointInTimeInput) rds.RestoreDBInstanceToPointInTimeRequest

	RevokeDBSecurityGroupIngressRequest(*rds.RevokeDBSecurityGroupIngressInput) rds.RevokeDBSecurityGroupIngressRequest

	StartDBInstanceRequest(*rds.StartDBInstanceInput) rds.StartDBInstanceRequest

	StopDBInstanceRequest(*rds.StopDBInstanceInput) rds.StopDBInstanceRequest

	WaitUntilDBInstanceAvailable(*rds.DescribeDBInstancesInput) error
	WaitUntilDBInstanceAvailableWithContext(aws.Context, *rds.DescribeDBInstancesInput, ...aws.WaiterOption) error

	WaitUntilDBInstanceDeleted(*rds.DescribeDBInstancesInput) error
	WaitUntilDBInstanceDeletedWithContext(aws.Context, *rds.DescribeDBInstancesInput, ...aws.WaiterOption) error

	WaitUntilDBSnapshotAvailable(*rds.DescribeDBSnapshotsInput) error
	WaitUntilDBSnapshotAvailableWithContext(aws.Context, *rds.DescribeDBSnapshotsInput, ...aws.WaiterOption) error

	WaitUntilDBSnapshotDeleted(*rds.DescribeDBSnapshotsInput) error
	WaitUntilDBSnapshotDeletedWithContext(aws.Context, *rds.DescribeDBSnapshotsInput, ...aws.WaiterOption) error
}

var _ RDSAPI = (*rds.RDS)(nil)
