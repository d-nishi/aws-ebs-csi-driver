// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bertinatto/ebs-csi-driver/pkg/cloud (interfaces: EC2)

// Package mocks is a generated GoMock package.
package mocks

import (
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEC2 is a mock of EC2 interface
type MockEC2 struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MockRecorder
}

// MockEC2MockRecorder is the mock recorder for MockEC2
type MockEC2MockRecorder struct {
	mock *MockEC2
}

// NewMockEC2 creates a new mock instance
func NewMockEC2(ctrl *gomock.Controller) *MockEC2 {
	mock := &MockEC2{ctrl: ctrl}
	mock.recorder = &MockEC2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEC2) EXPECT() *MockEC2MockRecorder {
	return m.recorder
}

// AttachVolume mocks base method
func (m *MockEC2) AttachVolume(arg0 *ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error) {
	ret := m.ctrl.Call(m, "AttachVolume", arg0)
	ret0, _ := ret[0].(*ec2.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachVolume indicates an expected call of AttachVolume
func (mr *MockEC2MockRecorder) AttachVolume(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachVolume", reflect.TypeOf((*MockEC2)(nil).AttachVolume), arg0)
}

// CreateVolume mocks base method
func (m *MockEC2) CreateVolume(arg0 *ec2.CreateVolumeInput) (*ec2.Volume, error) {
	ret := m.ctrl.Call(m, "CreateVolume", arg0)
	ret0, _ := ret[0].(*ec2.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolume indicates an expected call of CreateVolume
func (mr *MockEC2MockRecorder) CreateVolume(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockEC2)(nil).CreateVolume), arg0)
}

// DeleteVolume mocks base method
func (m *MockEC2) DeleteVolume(arg0 *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	ret := m.ctrl.Call(m, "DeleteVolume", arg0)
	ret0, _ := ret[0].(*ec2.DeleteVolumeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVolume indicates an expected call of DeleteVolume
func (mr *MockEC2MockRecorder) DeleteVolume(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockEC2)(nil).DeleteVolume), arg0)
}

// DescribeInstances mocks base method
func (m *MockEC2) DescribeInstances(arg0 *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	ret := m.ctrl.Call(m, "DescribeInstances", arg0)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstances indicates an expected call of DescribeInstances
func (mr *MockEC2MockRecorder) DescribeInstances(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstances", reflect.TypeOf((*MockEC2)(nil).DescribeInstances), arg0)
}

// DescribeVolumes mocks base method
func (m *MockEC2) DescribeVolumes(arg0 *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	ret := m.ctrl.Call(m, "DescribeVolumes", arg0)
	ret0, _ := ret[0].(*ec2.DescribeVolumesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVolumes indicates an expected call of DescribeVolumes
func (mr *MockEC2MockRecorder) DescribeVolumes(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVolumes", reflect.TypeOf((*MockEC2)(nil).DescribeVolumes), arg0)
}

// DetachVolume mocks base method
func (m *MockEC2) DetachVolume(arg0 *ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error) {
	ret := m.ctrl.Call(m, "DetachVolume", arg0)
	ret0, _ := ret[0].(*ec2.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachVolume indicates an expected call of DetachVolume
func (mr *MockEC2MockRecorder) DetachVolume(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachVolume", reflect.TypeOf((*MockEC2)(nil).DetachVolume), arg0)
}
