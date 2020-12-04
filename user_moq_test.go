// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package labdeploy

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"sync"
)

// Ensure, that IAMInterfaceMock does implement IAMInterface.
// If this is not the case, regenerate this file with moq.
var _ IAMInterface = &IAMInterfaceMock{}

// IAMInterfaceMock is a mock implementation of IAMInterface.
//
//     func TestSomethingThatUsesIAMInterface(t *testing.T) {
//
//         // make and configure a mocked IAMInterface
//         mockedIAMInterface := &IAMInterfaceMock{
//             AttachUserPolicyFunc: func(ctx context.Context, params *iam.AttachUserPolicyInput, optFns ...func(*iam.Options)) (*iam.AttachUserPolicyOutput, error) {
// 	               panic("mock out the AttachUserPolicy method")
//             },
//             CreateAccessKeyFunc: func(ctx context.Context, params *iam.CreateAccessKeyInput, optFns ...func(*iam.Options)) (*iam.CreateAccessKeyOutput, error) {
// 	               panic("mock out the CreateAccessKey method")
//             },
//             CreateLoginProfileFunc: func(ctx context.Context, params *iam.CreateLoginProfileInput, optFns ...func(*iam.Options)) (*iam.CreateLoginProfileOutput, error) {
// 	               panic("mock out the CreateLoginProfile method")
//             },
//             CreatePolicyFunc: func(ctx context.Context, params *iam.CreatePolicyInput, optFns ...func(*iam.Options)) (*iam.CreatePolicyOutput, error) {
// 	               panic("mock out the CreatePolicy method")
//             },
//             CreateUserFunc: func(ctx context.Context, params *iam.CreateUserInput, optFns ...func(*iam.Options)) (*iam.CreateUserOutput, error) {
// 	               panic("mock out the CreateUser method")
//             },
//             DeleteAccessKeyFunc: func(ctx context.Context, params *iam.DeleteAccessKeyInput, optFns ...func(*iam.Options)) (*iam.DeleteAccessKeyOutput, error) {
// 	               panic("mock out the DeleteAccessKey method")
//             },
//             DeleteLoginProfileFunc: func(ctx context.Context, params *iam.DeleteLoginProfileInput, optFns ...func(*iam.Options)) (*iam.DeleteLoginProfileOutput, error) {
// 	               panic("mock out the DeleteLoginProfile method")
//             },
//             DeletePolicyFunc: func(ctx context.Context, params *iam.DeletePolicyInput, optFns ...func(*iam.Options)) (*iam.DeletePolicyOutput, error) {
// 	               panic("mock out the DeletePolicy method")
//             },
//             DeleteUserFunc: func(ctx context.Context, params *iam.DeleteUserInput, optFns ...func(*iam.Options)) (*iam.DeleteUserOutput, error) {
// 	               panic("mock out the DeleteUser method")
//             },
//             DetachUserPolicyFunc: func(ctx context.Context, params *iam.DetachUserPolicyInput, optFns ...func(*iam.Options)) (*iam.DetachUserPolicyOutput, error) {
// 	               panic("mock out the DetachUserPolicy method")
//             },
//             ListAccessKeysFunc: func(ctx context.Context, params *iam.ListAccessKeysInput, optFns ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error) {
// 	               panic("mock out the ListAccessKeys method")
//             },
//             ListPoliciesFunc: func(ctx context.Context, params *iam.ListPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListPoliciesOutput, error) {
// 	               panic("mock out the ListPolicies method")
//             },
//             ListUsersFunc: func(ctx context.Context, params *iam.ListUsersInput, optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
// 	               panic("mock out the ListUsers method")
//             },
//         }
//
//         // use mockedIAMInterface in code that requires IAMInterface
//         // and then make assertions.
//
//     }
type IAMInterfaceMock struct {
	// AttachUserPolicyFunc mocks the AttachUserPolicy method.
	AttachUserPolicyFunc func(ctx context.Context, params *iam.AttachUserPolicyInput, optFns ...func(*iam.Options)) (*iam.AttachUserPolicyOutput, error)

	// CreateAccessKeyFunc mocks the CreateAccessKey method.
	CreateAccessKeyFunc func(ctx context.Context, params *iam.CreateAccessKeyInput, optFns ...func(*iam.Options)) (*iam.CreateAccessKeyOutput, error)

	// CreateLoginProfileFunc mocks the CreateLoginProfile method.
	CreateLoginProfileFunc func(ctx context.Context, params *iam.CreateLoginProfileInput, optFns ...func(*iam.Options)) (*iam.CreateLoginProfileOutput, error)

	// CreatePolicyFunc mocks the CreatePolicy method.
	CreatePolicyFunc func(ctx context.Context, params *iam.CreatePolicyInput, optFns ...func(*iam.Options)) (*iam.CreatePolicyOutput, error)

	// CreateUserFunc mocks the CreateUser method.
	CreateUserFunc func(ctx context.Context, params *iam.CreateUserInput, optFns ...func(*iam.Options)) (*iam.CreateUserOutput, error)

	// DeleteAccessKeyFunc mocks the DeleteAccessKey method.
	DeleteAccessKeyFunc func(ctx context.Context, params *iam.DeleteAccessKeyInput, optFns ...func(*iam.Options)) (*iam.DeleteAccessKeyOutput, error)

	// DeleteLoginProfileFunc mocks the DeleteLoginProfile method.
	DeleteLoginProfileFunc func(ctx context.Context, params *iam.DeleteLoginProfileInput, optFns ...func(*iam.Options)) (*iam.DeleteLoginProfileOutput, error)

	// DeletePolicyFunc mocks the DeletePolicy method.
	DeletePolicyFunc func(ctx context.Context, params *iam.DeletePolicyInput, optFns ...func(*iam.Options)) (*iam.DeletePolicyOutput, error)

	// DeleteUserFunc mocks the DeleteUser method.
	DeleteUserFunc func(ctx context.Context, params *iam.DeleteUserInput, optFns ...func(*iam.Options)) (*iam.DeleteUserOutput, error)

	// DetachUserPolicyFunc mocks the DetachUserPolicy method.
	DetachUserPolicyFunc func(ctx context.Context, params *iam.DetachUserPolicyInput, optFns ...func(*iam.Options)) (*iam.DetachUserPolicyOutput, error)

	// ListAccessKeysFunc mocks the ListAccessKeys method.
	ListAccessKeysFunc func(ctx context.Context, params *iam.ListAccessKeysInput, optFns ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error)

	// ListPoliciesFunc mocks the ListPolicies method.
	ListPoliciesFunc func(ctx context.Context, params *iam.ListPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListPoliciesOutput, error)

	// ListUsersFunc mocks the ListUsers method.
	ListUsersFunc func(ctx context.Context, params *iam.ListUsersInput, optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)

	// calls tracks calls to the methods.
	calls struct {
		// AttachUserPolicy holds details about calls to the AttachUserPolicy method.
		AttachUserPolicy []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.AttachUserPolicyInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
		// CreateAccessKey holds details about calls to the CreateAccessKey method.
		CreateAccessKey []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.CreateAccessKeyInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
		// CreateLoginProfile holds details about calls to the CreateLoginProfile method.
		CreateLoginProfile []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.CreateLoginProfileInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
		// CreatePolicy holds details about calls to the CreatePolicy method.
		CreatePolicy []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.CreatePolicyInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
		// CreateUser holds details about calls to the CreateUser method.
		CreateUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.CreateUserInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
		// DeleteAccessKey holds details about calls to the DeleteAccessKey method.
		DeleteAccessKey []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.DeleteAccessKeyInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
		// DeleteLoginProfile holds details about calls to the DeleteLoginProfile method.
		DeleteLoginProfile []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.DeleteLoginProfileInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
		// DeletePolicy holds details about calls to the DeletePolicy method.
		DeletePolicy []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.DeletePolicyInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
		// DeleteUser holds details about calls to the DeleteUser method.
		DeleteUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.DeleteUserInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
		// DetachUserPolicy holds details about calls to the DetachUserPolicy method.
		DetachUserPolicy []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.DetachUserPolicyInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
		// ListAccessKeys holds details about calls to the ListAccessKeys method.
		ListAccessKeys []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.ListAccessKeysInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
		// ListPolicies holds details about calls to the ListPolicies method.
		ListPolicies []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.ListPoliciesInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
		// ListUsers holds details about calls to the ListUsers method.
		ListUsers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *iam.ListUsersInput
			// OptFns is the optFns argument value.
			OptFns []func(*iam.Options)
		}
	}
	lockAttachUserPolicy   sync.RWMutex
	lockCreateAccessKey    sync.RWMutex
	lockCreateLoginProfile sync.RWMutex
	lockCreatePolicy       sync.RWMutex
	lockCreateUser         sync.RWMutex
	lockDeleteAccessKey    sync.RWMutex
	lockDeleteLoginProfile sync.RWMutex
	lockDeletePolicy       sync.RWMutex
	lockDeleteUser         sync.RWMutex
	lockDetachUserPolicy   sync.RWMutex
	lockListAccessKeys     sync.RWMutex
	lockListPolicies       sync.RWMutex
	lockListUsers          sync.RWMutex
}

// AttachUserPolicy calls AttachUserPolicyFunc.
func (mock *IAMInterfaceMock) AttachUserPolicy(ctx context.Context, params *iam.AttachUserPolicyInput, optFns ...func(*iam.Options)) (*iam.AttachUserPolicyOutput, error) {
	if mock.AttachUserPolicyFunc == nil {
		panic("IAMInterfaceMock.AttachUserPolicyFunc: method is nil but IAMInterface.AttachUserPolicy was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.AttachUserPolicyInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockAttachUserPolicy.Lock()
	mock.calls.AttachUserPolicy = append(mock.calls.AttachUserPolicy, callInfo)
	mock.lockAttachUserPolicy.Unlock()
	return mock.AttachUserPolicyFunc(ctx, params, optFns...)
}

// AttachUserPolicyCalls gets all the calls that were made to AttachUserPolicy.
// Check the length with:
//     len(mockedIAMInterface.AttachUserPolicyCalls())
func (mock *IAMInterfaceMock) AttachUserPolicyCalls() []struct {
	Ctx    context.Context
	Params *iam.AttachUserPolicyInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.AttachUserPolicyInput
		OptFns []func(*iam.Options)
	}
	mock.lockAttachUserPolicy.RLock()
	calls = mock.calls.AttachUserPolicy
	mock.lockAttachUserPolicy.RUnlock()
	return calls
}

// CreateAccessKey calls CreateAccessKeyFunc.
func (mock *IAMInterfaceMock) CreateAccessKey(ctx context.Context, params *iam.CreateAccessKeyInput, optFns ...func(*iam.Options)) (*iam.CreateAccessKeyOutput, error) {
	if mock.CreateAccessKeyFunc == nil {
		panic("IAMInterfaceMock.CreateAccessKeyFunc: method is nil but IAMInterface.CreateAccessKey was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.CreateAccessKeyInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockCreateAccessKey.Lock()
	mock.calls.CreateAccessKey = append(mock.calls.CreateAccessKey, callInfo)
	mock.lockCreateAccessKey.Unlock()
	return mock.CreateAccessKeyFunc(ctx, params, optFns...)
}

// CreateAccessKeyCalls gets all the calls that were made to CreateAccessKey.
// Check the length with:
//     len(mockedIAMInterface.CreateAccessKeyCalls())
func (mock *IAMInterfaceMock) CreateAccessKeyCalls() []struct {
	Ctx    context.Context
	Params *iam.CreateAccessKeyInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.CreateAccessKeyInput
		OptFns []func(*iam.Options)
	}
	mock.lockCreateAccessKey.RLock()
	calls = mock.calls.CreateAccessKey
	mock.lockCreateAccessKey.RUnlock()
	return calls
}

// CreateLoginProfile calls CreateLoginProfileFunc.
func (mock *IAMInterfaceMock) CreateLoginProfile(ctx context.Context, params *iam.CreateLoginProfileInput, optFns ...func(*iam.Options)) (*iam.CreateLoginProfileOutput, error) {
	if mock.CreateLoginProfileFunc == nil {
		panic("IAMInterfaceMock.CreateLoginProfileFunc: method is nil but IAMInterface.CreateLoginProfile was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.CreateLoginProfileInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockCreateLoginProfile.Lock()
	mock.calls.CreateLoginProfile = append(mock.calls.CreateLoginProfile, callInfo)
	mock.lockCreateLoginProfile.Unlock()
	return mock.CreateLoginProfileFunc(ctx, params, optFns...)
}

// CreateLoginProfileCalls gets all the calls that were made to CreateLoginProfile.
// Check the length with:
//     len(mockedIAMInterface.CreateLoginProfileCalls())
func (mock *IAMInterfaceMock) CreateLoginProfileCalls() []struct {
	Ctx    context.Context
	Params *iam.CreateLoginProfileInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.CreateLoginProfileInput
		OptFns []func(*iam.Options)
	}
	mock.lockCreateLoginProfile.RLock()
	calls = mock.calls.CreateLoginProfile
	mock.lockCreateLoginProfile.RUnlock()
	return calls
}

// CreatePolicy calls CreatePolicyFunc.
func (mock *IAMInterfaceMock) CreatePolicy(ctx context.Context, params *iam.CreatePolicyInput, optFns ...func(*iam.Options)) (*iam.CreatePolicyOutput, error) {
	if mock.CreatePolicyFunc == nil {
		panic("IAMInterfaceMock.CreatePolicyFunc: method is nil but IAMInterface.CreatePolicy was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.CreatePolicyInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockCreatePolicy.Lock()
	mock.calls.CreatePolicy = append(mock.calls.CreatePolicy, callInfo)
	mock.lockCreatePolicy.Unlock()
	return mock.CreatePolicyFunc(ctx, params, optFns...)
}

// CreatePolicyCalls gets all the calls that were made to CreatePolicy.
// Check the length with:
//     len(mockedIAMInterface.CreatePolicyCalls())
func (mock *IAMInterfaceMock) CreatePolicyCalls() []struct {
	Ctx    context.Context
	Params *iam.CreatePolicyInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.CreatePolicyInput
		OptFns []func(*iam.Options)
	}
	mock.lockCreatePolicy.RLock()
	calls = mock.calls.CreatePolicy
	mock.lockCreatePolicy.RUnlock()
	return calls
}

// CreateUser calls CreateUserFunc.
func (mock *IAMInterfaceMock) CreateUser(ctx context.Context, params *iam.CreateUserInput, optFns ...func(*iam.Options)) (*iam.CreateUserOutput, error) {
	if mock.CreateUserFunc == nil {
		panic("IAMInterfaceMock.CreateUserFunc: method is nil but IAMInterface.CreateUser was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.CreateUserInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockCreateUser.Lock()
	mock.calls.CreateUser = append(mock.calls.CreateUser, callInfo)
	mock.lockCreateUser.Unlock()
	return mock.CreateUserFunc(ctx, params, optFns...)
}

// CreateUserCalls gets all the calls that were made to CreateUser.
// Check the length with:
//     len(mockedIAMInterface.CreateUserCalls())
func (mock *IAMInterfaceMock) CreateUserCalls() []struct {
	Ctx    context.Context
	Params *iam.CreateUserInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.CreateUserInput
		OptFns []func(*iam.Options)
	}
	mock.lockCreateUser.RLock()
	calls = mock.calls.CreateUser
	mock.lockCreateUser.RUnlock()
	return calls
}

// DeleteAccessKey calls DeleteAccessKeyFunc.
func (mock *IAMInterfaceMock) DeleteAccessKey(ctx context.Context, params *iam.DeleteAccessKeyInput, optFns ...func(*iam.Options)) (*iam.DeleteAccessKeyOutput, error) {
	if mock.DeleteAccessKeyFunc == nil {
		panic("IAMInterfaceMock.DeleteAccessKeyFunc: method is nil but IAMInterface.DeleteAccessKey was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.DeleteAccessKeyInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockDeleteAccessKey.Lock()
	mock.calls.DeleteAccessKey = append(mock.calls.DeleteAccessKey, callInfo)
	mock.lockDeleteAccessKey.Unlock()
	return mock.DeleteAccessKeyFunc(ctx, params, optFns...)
}

// DeleteAccessKeyCalls gets all the calls that were made to DeleteAccessKey.
// Check the length with:
//     len(mockedIAMInterface.DeleteAccessKeyCalls())
func (mock *IAMInterfaceMock) DeleteAccessKeyCalls() []struct {
	Ctx    context.Context
	Params *iam.DeleteAccessKeyInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.DeleteAccessKeyInput
		OptFns []func(*iam.Options)
	}
	mock.lockDeleteAccessKey.RLock()
	calls = mock.calls.DeleteAccessKey
	mock.lockDeleteAccessKey.RUnlock()
	return calls
}

// DeleteLoginProfile calls DeleteLoginProfileFunc.
func (mock *IAMInterfaceMock) DeleteLoginProfile(ctx context.Context, params *iam.DeleteLoginProfileInput, optFns ...func(*iam.Options)) (*iam.DeleteLoginProfileOutput, error) {
	if mock.DeleteLoginProfileFunc == nil {
		panic("IAMInterfaceMock.DeleteLoginProfileFunc: method is nil but IAMInterface.DeleteLoginProfile was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.DeleteLoginProfileInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockDeleteLoginProfile.Lock()
	mock.calls.DeleteLoginProfile = append(mock.calls.DeleteLoginProfile, callInfo)
	mock.lockDeleteLoginProfile.Unlock()
	return mock.DeleteLoginProfileFunc(ctx, params, optFns...)
}

// DeleteLoginProfileCalls gets all the calls that were made to DeleteLoginProfile.
// Check the length with:
//     len(mockedIAMInterface.DeleteLoginProfileCalls())
func (mock *IAMInterfaceMock) DeleteLoginProfileCalls() []struct {
	Ctx    context.Context
	Params *iam.DeleteLoginProfileInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.DeleteLoginProfileInput
		OptFns []func(*iam.Options)
	}
	mock.lockDeleteLoginProfile.RLock()
	calls = mock.calls.DeleteLoginProfile
	mock.lockDeleteLoginProfile.RUnlock()
	return calls
}

// DeletePolicy calls DeletePolicyFunc.
func (mock *IAMInterfaceMock) DeletePolicy(ctx context.Context, params *iam.DeletePolicyInput, optFns ...func(*iam.Options)) (*iam.DeletePolicyOutput, error) {
	if mock.DeletePolicyFunc == nil {
		panic("IAMInterfaceMock.DeletePolicyFunc: method is nil but IAMInterface.DeletePolicy was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.DeletePolicyInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockDeletePolicy.Lock()
	mock.calls.DeletePolicy = append(mock.calls.DeletePolicy, callInfo)
	mock.lockDeletePolicy.Unlock()
	return mock.DeletePolicyFunc(ctx, params, optFns...)
}

// DeletePolicyCalls gets all the calls that were made to DeletePolicy.
// Check the length with:
//     len(mockedIAMInterface.DeletePolicyCalls())
func (mock *IAMInterfaceMock) DeletePolicyCalls() []struct {
	Ctx    context.Context
	Params *iam.DeletePolicyInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.DeletePolicyInput
		OptFns []func(*iam.Options)
	}
	mock.lockDeletePolicy.RLock()
	calls = mock.calls.DeletePolicy
	mock.lockDeletePolicy.RUnlock()
	return calls
}

// DeleteUser calls DeleteUserFunc.
func (mock *IAMInterfaceMock) DeleteUser(ctx context.Context, params *iam.DeleteUserInput, optFns ...func(*iam.Options)) (*iam.DeleteUserOutput, error) {
	if mock.DeleteUserFunc == nil {
		panic("IAMInterfaceMock.DeleteUserFunc: method is nil but IAMInterface.DeleteUser was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.DeleteUserInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockDeleteUser.Lock()
	mock.calls.DeleteUser = append(mock.calls.DeleteUser, callInfo)
	mock.lockDeleteUser.Unlock()
	return mock.DeleteUserFunc(ctx, params, optFns...)
}

// DeleteUserCalls gets all the calls that were made to DeleteUser.
// Check the length with:
//     len(mockedIAMInterface.DeleteUserCalls())
func (mock *IAMInterfaceMock) DeleteUserCalls() []struct {
	Ctx    context.Context
	Params *iam.DeleteUserInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.DeleteUserInput
		OptFns []func(*iam.Options)
	}
	mock.lockDeleteUser.RLock()
	calls = mock.calls.DeleteUser
	mock.lockDeleteUser.RUnlock()
	return calls
}

// DetachUserPolicy calls DetachUserPolicyFunc.
func (mock *IAMInterfaceMock) DetachUserPolicy(ctx context.Context, params *iam.DetachUserPolicyInput, optFns ...func(*iam.Options)) (*iam.DetachUserPolicyOutput, error) {
	if mock.DetachUserPolicyFunc == nil {
		panic("IAMInterfaceMock.DetachUserPolicyFunc: method is nil but IAMInterface.DetachUserPolicy was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.DetachUserPolicyInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockDetachUserPolicy.Lock()
	mock.calls.DetachUserPolicy = append(mock.calls.DetachUserPolicy, callInfo)
	mock.lockDetachUserPolicy.Unlock()
	return mock.DetachUserPolicyFunc(ctx, params, optFns...)
}

// DetachUserPolicyCalls gets all the calls that were made to DetachUserPolicy.
// Check the length with:
//     len(mockedIAMInterface.DetachUserPolicyCalls())
func (mock *IAMInterfaceMock) DetachUserPolicyCalls() []struct {
	Ctx    context.Context
	Params *iam.DetachUserPolicyInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.DetachUserPolicyInput
		OptFns []func(*iam.Options)
	}
	mock.lockDetachUserPolicy.RLock()
	calls = mock.calls.DetachUserPolicy
	mock.lockDetachUserPolicy.RUnlock()
	return calls
}

// ListAccessKeys calls ListAccessKeysFunc.
func (mock *IAMInterfaceMock) ListAccessKeys(ctx context.Context, params *iam.ListAccessKeysInput, optFns ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error) {
	if mock.ListAccessKeysFunc == nil {
		panic("IAMInterfaceMock.ListAccessKeysFunc: method is nil but IAMInterface.ListAccessKeys was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.ListAccessKeysInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockListAccessKeys.Lock()
	mock.calls.ListAccessKeys = append(mock.calls.ListAccessKeys, callInfo)
	mock.lockListAccessKeys.Unlock()
	return mock.ListAccessKeysFunc(ctx, params, optFns...)
}

// ListAccessKeysCalls gets all the calls that were made to ListAccessKeys.
// Check the length with:
//     len(mockedIAMInterface.ListAccessKeysCalls())
func (mock *IAMInterfaceMock) ListAccessKeysCalls() []struct {
	Ctx    context.Context
	Params *iam.ListAccessKeysInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.ListAccessKeysInput
		OptFns []func(*iam.Options)
	}
	mock.lockListAccessKeys.RLock()
	calls = mock.calls.ListAccessKeys
	mock.lockListAccessKeys.RUnlock()
	return calls
}

// ListPolicies calls ListPoliciesFunc.
func (mock *IAMInterfaceMock) ListPolicies(ctx context.Context, params *iam.ListPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListPoliciesOutput, error) {
	if mock.ListPoliciesFunc == nil {
		panic("IAMInterfaceMock.ListPoliciesFunc: method is nil but IAMInterface.ListPolicies was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.ListPoliciesInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockListPolicies.Lock()
	mock.calls.ListPolicies = append(mock.calls.ListPolicies, callInfo)
	mock.lockListPolicies.Unlock()
	return mock.ListPoliciesFunc(ctx, params, optFns...)
}

// ListPoliciesCalls gets all the calls that were made to ListPolicies.
// Check the length with:
//     len(mockedIAMInterface.ListPoliciesCalls())
func (mock *IAMInterfaceMock) ListPoliciesCalls() []struct {
	Ctx    context.Context
	Params *iam.ListPoliciesInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.ListPoliciesInput
		OptFns []func(*iam.Options)
	}
	mock.lockListPolicies.RLock()
	calls = mock.calls.ListPolicies
	mock.lockListPolicies.RUnlock()
	return calls
}

// ListUsers calls ListUsersFunc.
func (mock *IAMInterfaceMock) ListUsers(ctx context.Context, params *iam.ListUsersInput, optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	if mock.ListUsersFunc == nil {
		panic("IAMInterfaceMock.ListUsersFunc: method is nil but IAMInterface.ListUsers was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *iam.ListUsersInput
		OptFns []func(*iam.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockListUsers.Lock()
	mock.calls.ListUsers = append(mock.calls.ListUsers, callInfo)
	mock.lockListUsers.Unlock()
	return mock.ListUsersFunc(ctx, params, optFns...)
}

// ListUsersCalls gets all the calls that were made to ListUsers.
// Check the length with:
//     len(mockedIAMInterface.ListUsersCalls())
func (mock *IAMInterfaceMock) ListUsersCalls() []struct {
	Ctx    context.Context
	Params *iam.ListUsersInput
	OptFns []func(*iam.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *iam.ListUsersInput
		OptFns []func(*iam.Options)
	}
	mock.lockListUsers.RLock()
	calls = mock.calls.ListUsers
	mock.lockListUsers.RUnlock()
	return calls
}
