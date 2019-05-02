// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilAppExists uses the AWS OpsWorks API operation
// DescribeApps to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *OpsWorks) WaitUntilAppExists(ctx context.Context, input *DescribeAppsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilAppExists",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(1 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    aws.FailureWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 400,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeAppsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeAppsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDeploymentSuccessful uses the AWS OpsWorks API operation
// DescribeDeployments to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *OpsWorks) WaitUntilDeploymentSuccessful(ctx context.Context, input *DescribeDeploymentsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDeploymentSuccessful",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "Deployments[].Status",
				Expected: "successful",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Deployments[].Status",
				Expected: "failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDeploymentsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDeploymentsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilInstanceOnline uses the AWS OpsWorks API operation
// DescribeInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *OpsWorks) WaitUntilInstanceOnline(ctx context.Context, input *DescribeInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilInstanceOnline",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "Instances[].Status",
				Expected: "online",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "setup_failed",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "shutting_down",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "start_failed",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stopped",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stopping",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "terminating",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "terminated",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stop_failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilInstanceRegistered uses the AWS OpsWorks API operation
// DescribeInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *OpsWorks) WaitUntilInstanceRegistered(ctx context.Context, input *DescribeInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilInstanceRegistered",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "Instances[].Status",
				Expected: "registered",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "setup_failed",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "shutting_down",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stopped",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stopping",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "terminating",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "terminated",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stop_failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilInstanceStopped uses the AWS OpsWorks API operation
// DescribeInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *OpsWorks) WaitUntilInstanceStopped(ctx context.Context, input *DescribeInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilInstanceStopped",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "Instances[].Status",
				Expected: "stopped",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "booting",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "pending",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "rebooting",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "requested",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "running_setup",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "setup_failed",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "start_failed",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stop_failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilInstanceTerminated uses the AWS OpsWorks API operation
// DescribeInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *OpsWorks) WaitUntilInstanceTerminated(ctx context.Context, input *DescribeInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilInstanceTerminated",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "Instances[].Status",
				Expected: "terminated",
			},
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "booting",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "online",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "pending",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "rebooting",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "requested",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "running_setup",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "setup_failed",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "start_failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
