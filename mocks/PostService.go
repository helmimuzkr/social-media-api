// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	post "social-media-app/feature/post"
)

// PostService is an autogenerated mock type for the PostService type
type PostService struct {
	mock.Mock
}

// Create provides a mock function with given fields: token, newPost, fileHeader
func (_m *PostService) Create(token interface{}, newPost post.Core, fileHeader *multipart.FileHeader) error {
	ret := _m.Called(token, newPost, fileHeader)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, post.Core, *multipart.FileHeader) error); ok {
		r0 = rf(token, newPost, fileHeader)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: token, postID
func (_m *PostService) Delete(token interface{}, postID uint) error {
	ret := _m.Called(token, postID)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, uint) error); ok {
		r0 = rf(token, postID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *PostService) GetAll() ([]post.Core, error) {
	ret := _m.Called()

	var r0 []post.Core
	if rf, ok := ret.Get(0).(func() []post.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]post.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: postID
func (_m *PostService) GetByID(postID uint) (post.Core, error) {
	ret := _m.Called(postID)

	var r0 post.Core
	if rf, ok := ret.Get(0).(func(uint) post.Core); ok {
		r0 = rf(postID)
	} else {
		r0 = ret.Get(0).(post.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(postID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserID provides a mock function with given fields: userID
func (_m *PostService) GetByUserID(userID uint) ([]post.Core, error) {
	ret := _m.Called(userID)

	var r0 []post.Core
	if rf, ok := ret.Get(0).(func(uint) []post.Core); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]post.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MyPost provides a mock function with given fields: token
func (_m *PostService) MyPost(token interface{}) ([]post.Core, error) {
	ret := _m.Called(token)

	var r0 []post.Core
	if rf, ok := ret.Get(0).(func(interface{}) []post.Core); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]post.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: token, postID, updatePost, fileHeader
func (_m *PostService) Update(token interface{}, postID uint, updatePost post.Core, fileHeader *multipart.FileHeader) error {
	ret := _m.Called(token, postID, updatePost, fileHeader)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, uint, post.Core, *multipart.FileHeader) error); ok {
		r0 = rf(token, postID, updatePost, fileHeader)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPostService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostService creates a new instance of PostService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostService(t mockConstructorTestingTNewPostService) *PostService {
	mock := &PostService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
