// Code generated by mockery v1.0.0
package validating

import context "context"
import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import validating "github.com/paalka/kubewebhook/pkg/webhook/validating"

// Validator is an autogenerated mock type for the Validator type
type Validator struct {
	mock.Mock
}

// Validate provides a mock function with given fields: _a0, _a1
func (_m *Validator) Validate(_a0 context.Context, _a1 v1.Object) (bool, validating.ValidatorResult, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, v1.Object) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 validating.ValidatorResult
	if rf, ok := ret.Get(1).(func(context.Context, v1.Object) validating.ValidatorResult); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(validating.ValidatorResult)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, v1.Object) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
