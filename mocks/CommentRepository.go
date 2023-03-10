// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	comment "social-media-app/feature/comment"

	mock "github.com/stretchr/testify/mock"
)

// CommentRepository is an autogenerated mock type for the CommentRepository type
type CommentRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: userID, postID, _a2
func (_m *CommentRepository) Add(userID uint, postID uint, _a2 string) error {
	ret := _m.Called(userID, postID, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, string) error); ok {
		r0 = rf(userID, postID, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: postID
func (_m *CommentRepository) GetAll(postID uint) ([]comment.Core, error) {
	ret := _m.Called(postID)

	var r0 []comment.Core
	if rf, ok := ret.Get(0).(func(uint) []comment.Core); ok {
		r0 = rf(postID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comment.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(postID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentRepository creates a new instance of CommentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentRepository(t mockConstructorTestingTNewCommentRepository) *CommentRepository {
	mock := &CommentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
