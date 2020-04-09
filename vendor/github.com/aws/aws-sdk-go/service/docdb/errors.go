// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

const (

	// ErrCodeAuthorizationNotFoundFault for service response error code
	// "AuthorizationNotFound".
	//
	// The specified CIDR IP or Amazon EC2 security group isn't authorized for the
	// specified security group.
	//
	// Amazon DocumentDB also might not be authorized to perform necessary actions
	// on your behalf using IAM.
	ErrCodeAuthorizationNotFoundFault = "AuthorizationNotFound"

	// ErrCodeCertificateNotFoundFault for service response error code
	// "CertificateNotFound".
	//
	// CertificateIdentifier doesn't refer to an existing certificate.
	ErrCodeCertificateNotFoundFault = "CertificateNotFound"

	// ErrCodeDBClusterAlreadyExistsFault for service response error code
	// "DBClusterAlreadyExistsFault".
	//
	// You already have a cluster with the given identifier.
	ErrCodeDBClusterAlreadyExistsFault = "DBClusterAlreadyExistsFault"

	// ErrCodeDBClusterNotFoundFault for service response error code
	// "DBClusterNotFoundFault".
	//
	// DBClusterIdentifier doesn't refer to an existing cluster.
	ErrCodeDBClusterNotFoundFault = "DBClusterNotFoundFault"

	// ErrCodeDBClusterParameterGroupNotFoundFault for service response error code
	// "DBClusterParameterGroupNotFound".
	//
	// DBClusterParameterGroupName doesn't refer to an existing cluster parameter
	// group.
	ErrCodeDBClusterParameterGroupNotFoundFault = "DBClusterParameterGroupNotFound"

	// ErrCodeDBClusterQuotaExceededFault for service response error code
	// "DBClusterQuotaExceededFault".
	//
	// The cluster can't be created because you have reached the maximum allowed
	// quota of clusters.
	ErrCodeDBClusterQuotaExceededFault = "DBClusterQuotaExceededFault"

	// ErrCodeDBClusterSnapshotAlreadyExistsFault for service response error code
	// "DBClusterSnapshotAlreadyExistsFault".
	//
	// You already have a cluster snapshot with the given identifier.
	ErrCodeDBClusterSnapshotAlreadyExistsFault = "DBClusterSnapshotAlreadyExistsFault"

	// ErrCodeDBClusterSnapshotNotFoundFault for service response error code
	// "DBClusterSnapshotNotFoundFault".
	//
	// DBClusterSnapshotIdentifier doesn't refer to an existing cluster snapshot.
	ErrCodeDBClusterSnapshotNotFoundFault = "DBClusterSnapshotNotFoundFault"

	// ErrCodeDBInstanceAlreadyExistsFault for service response error code
	// "DBInstanceAlreadyExists".
	//
	// You already have a instance with the given identifier.
	ErrCodeDBInstanceAlreadyExistsFault = "DBInstanceAlreadyExists"

	// ErrCodeDBInstanceNotFoundFault for service response error code
	// "DBInstanceNotFound".
	//
	// DBInstanceIdentifier doesn't refer to an existing instance.
	ErrCodeDBInstanceNotFoundFault = "DBInstanceNotFound"

	// ErrCodeDBParameterGroupAlreadyExistsFault for service response error code
	// "DBParameterGroupAlreadyExists".
	//
	// A parameter group with the same name already exists.
	ErrCodeDBParameterGroupAlreadyExistsFault = "DBParameterGroupAlreadyExists"

	// ErrCodeDBParameterGroupNotFoundFault for service response error code
	// "DBParameterGroupNotFound".
	//
	// DBParameterGroupName doesn't refer to an existing parameter group.
	ErrCodeDBParameterGroupNotFoundFault = "DBParameterGroupNotFound"

	// ErrCodeDBParameterGroupQuotaExceededFault for service response error code
	// "DBParameterGroupQuotaExceeded".
	//
	// This request would cause you to exceed the allowed number of parameter groups.
	ErrCodeDBParameterGroupQuotaExceededFault = "DBParameterGroupQuotaExceeded"

	// ErrCodeDBSecurityGroupNotFoundFault for service response error code
	// "DBSecurityGroupNotFound".
	//
	// DBSecurityGroupName doesn't refer to an existing security group.
	ErrCodeDBSecurityGroupNotFoundFault = "DBSecurityGroupNotFound"

	// ErrCodeDBSnapshotAlreadyExistsFault for service response error code
	// "DBSnapshotAlreadyExists".
	//
	// DBSnapshotIdentifier is already being used by an existing snapshot.
	ErrCodeDBSnapshotAlreadyExistsFault = "DBSnapshotAlreadyExists"

	// ErrCodeDBSnapshotNotFoundFault for service response error code
	// "DBSnapshotNotFound".
	//
	// DBSnapshotIdentifier doesn't refer to an existing snapshot.
	ErrCodeDBSnapshotNotFoundFault = "DBSnapshotNotFound"

	// ErrCodeDBSubnetGroupAlreadyExistsFault for service response error code
	// "DBSubnetGroupAlreadyExists".
	//
	// DBSubnetGroupName is already being used by an existing subnet group.
	ErrCodeDBSubnetGroupAlreadyExistsFault = "DBSubnetGroupAlreadyExists"

	// ErrCodeDBSubnetGroupDoesNotCoverEnoughAZs for service response error code
	// "DBSubnetGroupDoesNotCoverEnoughAZs".
	//
	// Subnets in the subnet group should cover at least two Availability Zones
	// unless there is only one Availability Zone.
	ErrCodeDBSubnetGroupDoesNotCoverEnoughAZs = "DBSubnetGroupDoesNotCoverEnoughAZs"

	// ErrCodeDBSubnetGroupNotFoundFault for service response error code
	// "DBSubnetGroupNotFoundFault".
	//
	// DBSubnetGroupName doesn't refer to an existing subnet group.
	ErrCodeDBSubnetGroupNotFoundFault = "DBSubnetGroupNotFoundFault"

	// ErrCodeDBSubnetGroupQuotaExceededFault for service response error code
	// "DBSubnetGroupQuotaExceeded".
	//
	// The request would cause you to exceed the allowed number of subnet groups.
	ErrCodeDBSubnetGroupQuotaExceededFault = "DBSubnetGroupQuotaExceeded"

	// ErrCodeDBSubnetQuotaExceededFault for service response error code
	// "DBSubnetQuotaExceededFault".
	//
	// The request would cause you to exceed the allowed number of subnets in a
	// subnet group.
	ErrCodeDBSubnetQuotaExceededFault = "DBSubnetQuotaExceededFault"

	// ErrCodeDBUpgradeDependencyFailureFault for service response error code
	// "DBUpgradeDependencyFailure".
	//
	// The upgrade failed because a resource that the depends on can't be modified.
	ErrCodeDBUpgradeDependencyFailureFault = "DBUpgradeDependencyFailure"

	// ErrCodeInstanceQuotaExceededFault for service response error code
	// "InstanceQuotaExceeded".
	//
	// The request would cause you to exceed the allowed number of instances.
	ErrCodeInstanceQuotaExceededFault = "InstanceQuotaExceeded"

	// ErrCodeInsufficientDBClusterCapacityFault for service response error code
	// "InsufficientDBClusterCapacityFault".
	//
	// The cluster doesn't have enough capacity for the current operation.
	ErrCodeInsufficientDBClusterCapacityFault = "InsufficientDBClusterCapacityFault"

	// ErrCodeInsufficientDBInstanceCapacityFault for service response error code
	// "InsufficientDBInstanceCapacity".
	//
	// The specified instance class isn't available in the specified Availability
	// Zone.
	ErrCodeInsufficientDBInstanceCapacityFault = "InsufficientDBInstanceCapacity"

	// ErrCodeInsufficientStorageClusterCapacityFault for service response error code
	// "InsufficientStorageClusterCapacity".
	//
	// There is not enough storage available for the current action. You might be
	// able to resolve this error by updating your subnet group to use different
	// Availability Zones that have more storage available.
	ErrCodeInsufficientStorageClusterCapacityFault = "InsufficientStorageClusterCapacity"

	// ErrCodeInvalidDBClusterSnapshotStateFault for service response error code
	// "InvalidDBClusterSnapshotStateFault".
	//
	// The provided value isn't a valid cluster snapshot state.
	ErrCodeInvalidDBClusterSnapshotStateFault = "InvalidDBClusterSnapshotStateFault"

	// ErrCodeInvalidDBClusterStateFault for service response error code
	// "InvalidDBClusterStateFault".
	//
	// The cluster isn't in a valid state.
	ErrCodeInvalidDBClusterStateFault = "InvalidDBClusterStateFault"

	// ErrCodeInvalidDBInstanceStateFault for service response error code
	// "InvalidDBInstanceState".
	//
	// The specified instance isn't in the available state.
	ErrCodeInvalidDBInstanceStateFault = "InvalidDBInstanceState"

	// ErrCodeInvalidDBParameterGroupStateFault for service response error code
	// "InvalidDBParameterGroupState".
	//
	// The parameter group is in use, or it is in a state that is not valid. If
	// you are trying to delete the parameter group, you can't delete it when the
	// parameter group is in this state.
	ErrCodeInvalidDBParameterGroupStateFault = "InvalidDBParameterGroupState"

	// ErrCodeInvalidDBSecurityGroupStateFault for service response error code
	// "InvalidDBSecurityGroupState".
	//
	// The state of the security group doesn't allow deletion.
	ErrCodeInvalidDBSecurityGroupStateFault = "InvalidDBSecurityGroupState"

	// ErrCodeInvalidDBSnapshotStateFault for service response error code
	// "InvalidDBSnapshotState".
	//
	// The state of the snapshot doesn't allow deletion.
	ErrCodeInvalidDBSnapshotStateFault = "InvalidDBSnapshotState"

	// ErrCodeInvalidDBSubnetGroupStateFault for service response error code
	// "InvalidDBSubnetGroupStateFault".
	//
	// The subnet group can't be deleted because it's in use.
	ErrCodeInvalidDBSubnetGroupStateFault = "InvalidDBSubnetGroupStateFault"

	// ErrCodeInvalidDBSubnetStateFault for service response error code
	// "InvalidDBSubnetStateFault".
	//
	// The subnet isn't in the available state.
	ErrCodeInvalidDBSubnetStateFault = "InvalidDBSubnetStateFault"

	// ErrCodeInvalidRestoreFault for service response error code
	// "InvalidRestoreFault".
	//
	// You cannot restore from a virtual private cloud (VPC) backup to a non-VPC
	// DB instance.
	ErrCodeInvalidRestoreFault = "InvalidRestoreFault"

	// ErrCodeInvalidSubnet for service response error code
	// "InvalidSubnet".
	//
	// The requested subnet is not valid, or multiple subnets were requested that
	// are not all in a common virtual private cloud (VPC).
	ErrCodeInvalidSubnet = "InvalidSubnet"

	// ErrCodeInvalidVPCNetworkStateFault for service response error code
	// "InvalidVPCNetworkStateFault".
	//
	// The subnet group doesn't cover all Availability Zones after it is created
	// because of changes that were made.
	ErrCodeInvalidVPCNetworkStateFault = "InvalidVPCNetworkStateFault"

	// ErrCodeKMSKeyNotAccessibleFault for service response error code
	// "KMSKeyNotAccessibleFault".
	//
	// An error occurred when accessing an AWS KMS key.
	ErrCodeKMSKeyNotAccessibleFault = "KMSKeyNotAccessibleFault"

	// ErrCodeResourceNotFoundFault for service response error code
	// "ResourceNotFoundFault".
	//
	// The specified resource ID was not found.
	ErrCodeResourceNotFoundFault = "ResourceNotFoundFault"

	// ErrCodeSharedSnapshotQuotaExceededFault for service response error code
	// "SharedSnapshotQuotaExceeded".
	//
	// You have exceeded the maximum number of accounts that you can share a manual
	// DB snapshot with.
	ErrCodeSharedSnapshotQuotaExceededFault = "SharedSnapshotQuotaExceeded"

	// ErrCodeSnapshotQuotaExceededFault for service response error code
	// "SnapshotQuotaExceeded".
	//
	// The request would cause you to exceed the allowed number of snapshots.
	ErrCodeSnapshotQuotaExceededFault = "SnapshotQuotaExceeded"

	// ErrCodeStorageQuotaExceededFault for service response error code
	// "StorageQuotaExceeded".
	//
	// The request would cause you to exceed the allowed amount of storage available
	// across all instances.
	ErrCodeStorageQuotaExceededFault = "StorageQuotaExceeded"

	// ErrCodeStorageTypeNotSupportedFault for service response error code
	// "StorageTypeNotSupported".
	//
	// Storage of the specified StorageType can't be associated with the DB instance.
	ErrCodeStorageTypeNotSupportedFault = "StorageTypeNotSupported"

	// ErrCodeSubnetAlreadyInUse for service response error code
	// "SubnetAlreadyInUse".
	//
	// The subnet is already in use in the Availability Zone.
	ErrCodeSubnetAlreadyInUse = "SubnetAlreadyInUse"
)
