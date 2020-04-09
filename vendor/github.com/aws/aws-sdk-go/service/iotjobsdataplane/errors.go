// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotjobsdataplane

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeCertificateValidationException for service response error code
	// "CertificateValidationException".
	//
	// The certificate is invalid.
	ErrCodeCertificateValidationException = "CertificateValidationException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The contents of the request were invalid. For example, this code is returned
	// when an UpdateJobExecution request contains invalid status details. The message
	// contains details about the error.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeInvalidStateTransitionException for service response error code
	// "InvalidStateTransitionException".
	//
	// An update attempted to change the job execution to a state that is invalid
	// because of the job execution's current state (for example, an attempt to
	// change a request in state SUCCESS to state IN_PROGRESS). In this case, the
	// body of the error message also contains the executionState field.
	ErrCodeInvalidStateTransitionException = "InvalidStateTransitionException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is temporarily unavailable.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeTerminalStateException for service response error code
	// "TerminalStateException".
	//
	// The job is in a terminal state.
	ErrCodeTerminalStateException = "TerminalStateException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The rate exceeds the limit.
	ErrCodeThrottlingException = "ThrottlingException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"CertificateValidationException":  newErrorCertificateValidationException,
	"InvalidRequestException":         newErrorInvalidRequestException,
	"InvalidStateTransitionException": newErrorInvalidStateTransitionException,
	"ResourceNotFoundException":       newErrorResourceNotFoundException,
	"ServiceUnavailableException":     newErrorServiceUnavailableException,
	"TerminalStateException":          newErrorTerminalStateException,
	"ThrottlingException":             newErrorThrottlingException,
}
