// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/awslabs/ecs-task-kite/lib/ecsclient (interfaces: AugmentedTask,AugmentedContainer)

package mock_ecsclient

import (
	ecsclient "github.com/awslabs/ecs-task-kite/lib/ecsclient"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	ecs "github.com/aws/aws-sdk-go/service/ecs"
	gomock "github.com/golang/mock/gomock"
)

// Mock of AugmentedTask interface
type MockAugmentedTask struct {
	ctrl     *gomock.Controller
	recorder *_MockAugmentedTaskRecorder
}

// Recorder for MockAugmentedTask (not exported)
type _MockAugmentedTaskRecorder struct {
	mock *MockAugmentedTask
}

func NewMockAugmentedTask(ctrl *gomock.Controller) *MockAugmentedTask {
	mock := &MockAugmentedTask{ctrl: ctrl}
	mock.recorder = &_MockAugmentedTaskRecorder{mock}
	return mock
}

func (_m *MockAugmentedTask) EXPECT() *_MockAugmentedTaskRecorder {
	return _m.recorder
}

func (_m *MockAugmentedTask) Container(_param0 string) ecsclient.AugmentedContainer {
	ret := _m.ctrl.Call(_m, "Container", _param0)
	ret0, _ := ret[0].(ecsclient.AugmentedContainer)
	return ret0
}

func (_mr *_MockAugmentedTaskRecorder) Container(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Container", arg0)
}

func (_m *MockAugmentedTask) EC2Instance() *ec2.Instance {
	ret := _m.ctrl.Call(_m, "EC2Instance")
	ret0, _ := ret[0].(*ec2.Instance)
	return ret0
}

func (_mr *_MockAugmentedTaskRecorder) EC2Instance() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EC2Instance")
}

func (_m *MockAugmentedTask) ECSTask() *ecs.Task {
	ret := _m.ctrl.Call(_m, "ECSTask")
	ret0, _ := ret[0].(*ecs.Task)
	return ret0
}

func (_mr *_MockAugmentedTaskRecorder) ECSTask() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ECSTask")
}

func (_m *MockAugmentedTask) PrivateIP() string {
	ret := _m.ctrl.Call(_m, "PrivateIP")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockAugmentedTaskRecorder) PrivateIP() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PrivateIP")
}

func (_m *MockAugmentedTask) PublicIP() string {
	ret := _m.ctrl.Call(_m, "PublicIP")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockAugmentedTaskRecorder) PublicIP() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PublicIP")
}

// Mock of AugmentedContainer interface
type MockAugmentedContainer struct {
	ctrl     *gomock.Controller
	recorder *_MockAugmentedContainerRecorder
}

// Recorder for MockAugmentedContainer (not exported)
type _MockAugmentedContainerRecorder struct {
	mock *MockAugmentedContainer
}

func NewMockAugmentedContainer(ctrl *gomock.Controller) *MockAugmentedContainer {
	mock := &MockAugmentedContainer{ctrl: ctrl}
	mock.recorder = &_MockAugmentedContainerRecorder{mock}
	return mock
}

func (_m *MockAugmentedContainer) EXPECT() *_MockAugmentedContainerRecorder {
	return _m.recorder
}

func (_m *MockAugmentedContainer) ContainerPorts(_param0 string) []uint16 {
	ret := _m.ctrl.Call(_m, "ContainerPorts", _param0)
	ret0, _ := ret[0].([]uint16)
	return ret0
}

func (_mr *_MockAugmentedContainerRecorder) ContainerPorts(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerPorts", arg0)
}

func (_m *MockAugmentedContainer) ECSContainer() *ecs.Container {
	ret := _m.ctrl.Call(_m, "ECSContainer")
	ret0, _ := ret[0].(*ecs.Container)
	return ret0
}

func (_mr *_MockAugmentedContainerRecorder) ECSContainer() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ECSContainer")
}

func (_m *MockAugmentedContainer) ResolvePort(_param0 uint16) uint16 {
	ret := _m.ctrl.Call(_m, "ResolvePort", _param0)
	ret0, _ := ret[0].(uint16)
	return ret0
}

func (_mr *_MockAugmentedContainerRecorder) ResolvePort(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResolvePort", arg0)
}

func (_m *MockAugmentedContainer) Running() bool {
	ret := _m.ctrl.Call(_m, "Running")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockAugmentedContainerRecorder) Running() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Running")
}
