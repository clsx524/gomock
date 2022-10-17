// Code generated by MockGen. DO NOT EDIT.
// Source: generics.go

// Package source is a generated GoMock package.
package source

import (
	reflect "reflect"

	other "github.com/golang/mock/mockgen/internal/tests/generics/other"
	gomock "github.com/golang/mock/gomock"
	generics "github.com/golang/mock/mockgen/internal/tests/generics"
)

// MockBar is a mock of Bar interface.
type MockBar[T any, R any] struct {
	ctrl     *gomock.Controller
	recorder *MockBarMockRecorder[T, R]
}

// MockBarMockRecorder is the mock recorder for MockBar.
type MockBarMockRecorder[T any, R any] struct {
	mock *MockBar[T, R]
}

// NewMockBar creates a new mock instance.
func NewMockBar[T any, R any](ctrl *gomock.Controller) *MockBar[T, R] {
	mock := &MockBar[T, R]{ctrl: ctrl}
	mock.recorder = &MockBarMockRecorder[T, R]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBar[T, R]) EXPECT() *MockBarMockRecorder[T, R] {
	return m.recorder
}

// Eight mocks base method.
func (m *MockBar[T, R]) Eight(arg0 T) other.Two[T, R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eight", arg0)
	ret0, _ := ret[0].(other.Two[T, R])
	return ret0
}

// Eight indicates an expected call of Eight.
func (mr *MockBarMockRecorder[T, R]) Eight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eight", reflect.TypeOf((*MockBar[T, R])(nil).Eight), arg0)
}

// Eighteen mocks base method.
func (m *MockBar[T, R]) Eighteen() (generics.Iface[*other.Five], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eighteen")
	ret0, _ := ret[0].(generics.Iface[*other.Five])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Eighteen indicates an expected call of Eighteen.
func (mr *MockBarMockRecorder[T, R]) Eighteen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eighteen", reflect.TypeOf((*MockBar[T, R])(nil).Eighteen))
}

// Eleven mocks base method.
func (m *MockBar[T, R]) Eleven() (*other.One[T], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eleven")
	ret0, _ := ret[0].(*other.One[T])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Eleven indicates an expected call of Eleven.
func (mr *MockBarMockRecorder[T, R]) Eleven() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eleven", reflect.TypeOf((*MockBar[T, R])(nil).Eleven))
}

// Fifteen mocks base method.
func (m *MockBar[T, R]) Fifteen() (generics.Iface[generics.StructType], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fifteen")
	ret0, _ := ret[0].(generics.Iface[generics.StructType])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fifteen indicates an expected call of Fifteen.
func (mr *MockBarMockRecorder[T, R]) Fifteen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fifteen", reflect.TypeOf((*MockBar[T, R])(nil).Fifteen))
}

// Five mocks base method.
func (m *MockBar[T, R]) Five(arg0 T) generics.Baz[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Five", arg0)
	ret0, _ := ret[0].(generics.Baz[T])
	return ret0
}

// Five indicates an expected call of Five.
func (mr *MockBarMockRecorder[T, R]) Five(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Five", reflect.TypeOf((*MockBar[T, R])(nil).Five), arg0)
}

// Four mocks base method.
func (m *MockBar[T, R]) Four(arg0 T) generics.Foo[T, R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Four", arg0)
	ret0, _ := ret[0].(generics.Foo[T, R])
	return ret0
}

// Four indicates an expected call of Four.
func (mr *MockBarMockRecorder[T, R]) Four(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Four", reflect.TypeOf((*MockBar[T, R])(nil).Four), arg0)
}

// Fourteen mocks base method.
func (m *MockBar[T, R]) Fourteen() (*generics.Foo[generics.StructType, generics.StructType2], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fourteen")
	ret0, _ := ret[0].(*generics.Foo[generics.StructType, generics.StructType2])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fourteen indicates an expected call of Fourteen.
func (mr *MockBarMockRecorder[T, R]) Fourteen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fourteen", reflect.TypeOf((*MockBar[T, R])(nil).Fourteen))
}

// Nine mocks base method.
func (m *MockBar[T, R]) Nine(arg0 generics.Iface[T]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Nine", arg0)
}

// Nine indicates an expected call of Nine.
func (mr *MockBarMockRecorder[T, R]) Nine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nine", reflect.TypeOf((*MockBar[T, R])(nil).Nine), arg0)
}

// Nineteen mocks base method.
func (m *MockBar[T, R]) Nineteen() generics.AliasType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nineteen")
	ret0, _ := ret[0].(generics.AliasType)
	return ret0
}

// Nineteen indicates an expected call of Nineteen.
func (mr *MockBarMockRecorder[T, R]) Nineteen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nineteen", reflect.TypeOf((*MockBar[T, R])(nil).Nineteen))
}

// One mocks base method.
func (m *MockBar[T, R]) One(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// One indicates an expected call of One.
func (mr *MockBarMockRecorder[T, R]) One(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockBar[T, R])(nil).One), arg0)
}

// Seven mocks base method.
func (m *MockBar[T, R]) Seven(arg0 T) other.One[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seven", arg0)
	ret0, _ := ret[0].(other.One[T])
	return ret0
}

// Seven indicates an expected call of Seven.
func (mr *MockBarMockRecorder[T, R]) Seven(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seven", reflect.TypeOf((*MockBar[T, R])(nil).Seven), arg0)
}

// Seventeen mocks base method.
func (m *MockBar[T, R]) Seventeen() (*generics.Foo[other.Three, other.Four], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seventeen")
	ret0, _ := ret[0].(*generics.Foo[other.Three, other.Four])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seventeen indicates an expected call of Seventeen.
func (mr *MockBarMockRecorder[T, R]) Seventeen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seventeen", reflect.TypeOf((*MockBar[T, R])(nil).Seventeen))
}

// Six mocks base method.
func (m *MockBar[T, R]) Six(arg0 T) *generics.Baz[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Six", arg0)
	ret0, _ := ret[0].(*generics.Baz[T])
	return ret0
}

// Six indicates an expected call of Six.
func (mr *MockBarMockRecorder[T, R]) Six(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Six", reflect.TypeOf((*MockBar[T, R])(nil).Six), arg0)
}

// Sixteen mocks base method.
func (m *MockBar[T, R]) Sixteen() (generics.Baz[other.Three], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sixteen")
	ret0, _ := ret[0].(generics.Baz[other.Three])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sixteen indicates an expected call of Sixteen.
func (mr *MockBarMockRecorder[T, R]) Sixteen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sixteen", reflect.TypeOf((*MockBar[T, R])(nil).Sixteen))
}

// Ten mocks base method.
func (m *MockBar[T, R]) Ten(arg0 *T) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ten", arg0)
}

// Ten indicates an expected call of Ten.
func (mr *MockBarMockRecorder[T, R]) Ten(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ten", reflect.TypeOf((*MockBar[T, R])(nil).Ten), arg0)
}

// Thirteen mocks base method.
func (m *MockBar[T, R]) Thirteen() (generics.Baz[generics.StructType], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Thirteen")
	ret0, _ := ret[0].(generics.Baz[generics.StructType])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Thirteen indicates an expected call of Thirteen.
func (mr *MockBarMockRecorder[T, R]) Thirteen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Thirteen", reflect.TypeOf((*MockBar[T, R])(nil).Thirteen))
}

// Three mocks base method.
func (m *MockBar[T, R]) Three(arg0 T) R {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Three", arg0)
	ret0, _ := ret[0].(R)
	return ret0
}

// Three indicates an expected call of Three.
func (mr *MockBarMockRecorder[T, R]) Three(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Three", reflect.TypeOf((*MockBar[T, R])(nil).Three), arg0)
}

// Twelve mocks base method.
func (m *MockBar[T, R]) Twelve() (*other.Two[T, R], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Twelve")
	ret0, _ := ret[0].(*other.Two[T, R])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Twelve indicates an expected call of Twelve.
func (mr *MockBarMockRecorder[T, R]) Twelve() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Twelve", reflect.TypeOf((*MockBar[T, R])(nil).Twelve))
}

// Twenty mocks base method.
func (m *MockBar[T, R]) Twenty(arg0 *other.One[T]) *other.Two[T, R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Twenty", arg0)
	ret0, _ := ret[0].(*other.Two[T, R])
	return ret0
}

// Twenty indicates an expected call of Twenty.
func (mr *MockBarMockRecorder[T, R]) Twenty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Twenty", reflect.TypeOf((*MockBar[T, R])(nil).Twenty), arg0)
}

// TwentyOne mocks base method.
func (m *MockBar[T, R]) TwentyOne(arg0 *string) *other.Two[*T, *R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TwentyOne", arg0)
	ret0, _ := ret[0].(*other.Two[*T, *R])
	return ret0
}

// TwentyOne indicates an expected call of TwentyOne.
func (mr *MockBarMockRecorder[T, R]) TwentyOne(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TwentyOne", reflect.TypeOf((*MockBar[T, R])(nil).TwentyOne), arg0)
}

// Two mocks base method.
func (m *MockBar[T, R]) Two(arg0 T) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Two", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Two indicates an expected call of Two.
func (mr *MockBarMockRecorder[T, R]) Two(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Two", reflect.TypeOf((*MockBar[T, R])(nil).Two), arg0)
}

// MockIface is a mock of Iface interface.
type MockIface[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockIfaceMockRecorder[T]
}

// MockIfaceMockRecorder is the mock recorder for MockIface.
type MockIfaceMockRecorder[T any] struct {
	mock *MockIface[T]
}

// NewMockIface creates a new mock instance.
func NewMockIface[T any](ctrl *gomock.Controller) *MockIface[T] {
	mock := &MockIface[T]{ctrl: ctrl}
	mock.recorder = &MockIfaceMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIface[T]) EXPECT() *MockIfaceMockRecorder[T] {
	return m.recorder
}

// MockEmbeddingIface is a mock of EmbeddingIface interface.
type MockEmbeddingIface struct {
	ctrl     *gomock.Controller
	recorder *MockEmbeddingIfaceMockRecorder
}

// MockEmbeddingIfaceMockRecorder is the mock recorder for MockEmbeddingIface.
type MockEmbeddingIfaceMockRecorder struct {
	mock *MockEmbeddingIface
}

// NewMockEmbeddingIface creates a new mock instance.
func NewMockEmbeddingIface(ctrl *gomock.Controller) *MockEmbeddingIface {
	mock := &MockEmbeddingIface{ctrl: ctrl}
	mock.recorder = &MockEmbeddingIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmbeddingIface) EXPECT() *MockEmbeddingIfaceMockRecorder {
	return m.recorder
}

// AddChan mocks base method.
func (m *MockEmbeddingIface) AddChan(arg0 chan generics.StructType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddChan", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddChan indicates an expected call of AddChan.
func (mr *MockEmbeddingIfaceMockRecorder) AddChan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChan", reflect.TypeOf((*MockEmbeddingIface)(nil).AddChan), arg0)
}

// DoR mocks base method.
func (m *MockEmbeddingIface) DoR(arg0 other.Five) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoR", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoR indicates an expected call of DoR.
func (mr *MockEmbeddingIfaceMockRecorder) DoR(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoR", reflect.TypeOf((*MockEmbeddingIface)(nil).DoR), arg0)
}

// DoT mocks base method.
func (m *MockEmbeddingIface) DoT(arg0 generics.StructType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoT", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoT indicates an expected call of DoT.
func (mr *MockEmbeddingIfaceMockRecorder) DoT(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoT", reflect.TypeOf((*MockEmbeddingIface)(nil).DoT), arg0)
}

// Eight mocks base method.
func (m *MockEmbeddingIface) Eight(arg0 other.Three) other.Two[other.Three, error] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eight", arg0)
	ret0, _ := ret[0].(other.Two[other.Three, error])
	return ret0
}

// Eight indicates an expected call of Eight.
func (mr *MockEmbeddingIfaceMockRecorder) Eight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eight", reflect.TypeOf((*MockEmbeddingIface)(nil).Eight), arg0)
}

// Eighteen mocks base method.
func (m *MockEmbeddingIface) Eighteen() (generics.Iface[*other.Five], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eighteen")
	ret0, _ := ret[0].(generics.Iface[*other.Five])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Eighteen indicates an expected call of Eighteen.
func (mr *MockEmbeddingIfaceMockRecorder) Eighteen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eighteen", reflect.TypeOf((*MockEmbeddingIface)(nil).Eighteen))
}

// Eleven mocks base method.
func (m *MockEmbeddingIface) Eleven() (*other.One[other.Three], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eleven")
	ret0, _ := ret[0].(*other.One[other.Three])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Eleven indicates an expected call of Eleven.
func (mr *MockEmbeddingIfaceMockRecorder) Eleven() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eleven", reflect.TypeOf((*MockEmbeddingIface)(nil).Eleven))
}

// Fifteen mocks base method.
func (m *MockEmbeddingIface) Fifteen() (generics.Iface[generics.StructType], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fifteen")
	ret0, _ := ret[0].(generics.Iface[generics.StructType])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fifteen indicates an expected call of Fifteen.
func (mr *MockEmbeddingIfaceMockRecorder) Fifteen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fifteen", reflect.TypeOf((*MockEmbeddingIface)(nil).Fifteen))
}

// Five mocks base method.
func (m *MockEmbeddingIface) Five(arg0 other.Three) generics.Baz[other.Three] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Five", arg0)
	ret0, _ := ret[0].(generics.Baz[other.Three])
	return ret0
}

// Five indicates an expected call of Five.
func (mr *MockEmbeddingIfaceMockRecorder) Five(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Five", reflect.TypeOf((*MockEmbeddingIface)(nil).Five), arg0)
}

// Four mocks base method.
func (m *MockEmbeddingIface) Four(arg0 other.Three) generics.Foo[other.Three, error] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Four", arg0)
	ret0, _ := ret[0].(generics.Foo[other.Three, error])
	return ret0
}

// Four indicates an expected call of Four.
func (mr *MockEmbeddingIfaceMockRecorder) Four(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Four", reflect.TypeOf((*MockEmbeddingIface)(nil).Four), arg0)
}

// Fourteen mocks base method.
func (m *MockEmbeddingIface) Fourteen() (*generics.Foo[generics.StructType, generics.StructType2], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fourteen")
	ret0, _ := ret[0].(*generics.Foo[generics.StructType, generics.StructType2])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fourteen indicates an expected call of Fourteen.
func (mr *MockEmbeddingIfaceMockRecorder) Fourteen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fourteen", reflect.TypeOf((*MockEmbeddingIface)(nil).Fourteen))
}

// GetMap mocks base method.
func (m *MockEmbeddingIface) GetMap() (map[bool]generics.StructType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMap")
	ret0, _ := ret[0].(map[bool]generics.StructType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMap indicates an expected call of GetMap.
func (mr *MockEmbeddingIfaceMockRecorder) GetMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMap", reflect.TypeOf((*MockEmbeddingIface)(nil).GetMap))
}

// GetThem mocks base method.
func (m *MockEmbeddingIface) GetThem() ([]generics.StructType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThem")
	ret0, _ := ret[0].([]generics.StructType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThem indicates an expected call of GetThem.
func (mr *MockEmbeddingIfaceMockRecorder) GetThem() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThem", reflect.TypeOf((*MockEmbeddingIface)(nil).GetThem))
}

// GetThemMapped mocks base method.
func (m *MockEmbeddingIface) GetThemMapped() ([]map[int64]*generics.StructType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThemMapped")
	ret0, _ := ret[0].([]map[int64]*generics.StructType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThemMapped indicates an expected call of GetThemMapped.
func (mr *MockEmbeddingIfaceMockRecorder) GetThemMapped() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThemMapped", reflect.TypeOf((*MockEmbeddingIface)(nil).GetThemMapped))
}

// GetThemPtr mocks base method.
func (m *MockEmbeddingIface) GetThemPtr() ([]*generics.StructType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThemPtr")
	ret0, _ := ret[0].([]*generics.StructType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThemPtr indicates an expected call of GetThemPtr.
func (mr *MockEmbeddingIfaceMockRecorder) GetThemPtr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThemPtr", reflect.TypeOf((*MockEmbeddingIface)(nil).GetThemPtr))
}

// LocalFunc mocks base method.
func (m *MockEmbeddingIface) LocalFunc() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalFunc")
	ret0, _ := ret[0].(error)
	return ret0
}

// LocalFunc indicates an expected call of LocalFunc.
func (mr *MockEmbeddingIfaceMockRecorder) LocalFunc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalFunc", reflect.TypeOf((*MockEmbeddingIface)(nil).LocalFunc))
}

// MakeThem mocks base method.
func (m *MockEmbeddingIface) MakeThem(arg0 ...generics.StructType) (other.Five, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MakeThem", varargs...)
	ret0, _ := ret[0].(other.Five)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeThem indicates an expected call of MakeThem.
func (mr *MockEmbeddingIfaceMockRecorder) MakeThem(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeThem", reflect.TypeOf((*MockEmbeddingIface)(nil).MakeThem), arg0...)
}

// Nine mocks base method.
func (m *MockEmbeddingIface) Nine(arg0 generics.Iface[other.Three]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Nine", arg0)
}

// Nine indicates an expected call of Nine.
func (mr *MockEmbeddingIfaceMockRecorder) Nine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nine", reflect.TypeOf((*MockEmbeddingIface)(nil).Nine), arg0)
}

// Nineteen mocks base method.
func (m *MockEmbeddingIface) Nineteen() generics.AliasType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nineteen")
	ret0, _ := ret[0].(generics.AliasType)
	return ret0
}

// Nineteen indicates an expected call of Nineteen.
func (mr *MockEmbeddingIfaceMockRecorder) Nineteen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nineteen", reflect.TypeOf((*MockEmbeddingIface)(nil).Nineteen))
}

// One mocks base method.
func (m *MockEmbeddingIface) One(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// One indicates an expected call of One.
func (mr *MockEmbeddingIfaceMockRecorder) One(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockEmbeddingIface)(nil).One), arg0)
}

// Seven mocks base method.
func (m *MockEmbeddingIface) Seven(arg0 other.Three) other.One[other.Three] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seven", arg0)
	ret0, _ := ret[0].(other.One[other.Three])
	return ret0
}

// Seven indicates an expected call of Seven.
func (mr *MockEmbeddingIfaceMockRecorder) Seven(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seven", reflect.TypeOf((*MockEmbeddingIface)(nil).Seven), arg0)
}

// Seventeen mocks base method.
func (m *MockEmbeddingIface) Seventeen() (*generics.Foo[other.Three, other.Four], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seventeen")
	ret0, _ := ret[0].(*generics.Foo[other.Three, other.Four])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seventeen indicates an expected call of Seventeen.
func (mr *MockEmbeddingIfaceMockRecorder) Seventeen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seventeen", reflect.TypeOf((*MockEmbeddingIface)(nil).Seventeen))
}

// Six mocks base method.
func (m *MockEmbeddingIface) Six(arg0 other.Three) *generics.Baz[other.Three] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Six", arg0)
	ret0, _ := ret[0].(*generics.Baz[other.Three])
	return ret0
}

// Six indicates an expected call of Six.
func (mr *MockEmbeddingIfaceMockRecorder) Six(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Six", reflect.TypeOf((*MockEmbeddingIface)(nil).Six), arg0)
}

// Sixteen mocks base method.
func (m *MockEmbeddingIface) Sixteen() (generics.Baz[other.Three], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sixteen")
	ret0, _ := ret[0].(generics.Baz[other.Three])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sixteen indicates an expected call of Sixteen.
func (mr *MockEmbeddingIfaceMockRecorder) Sixteen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sixteen", reflect.TypeOf((*MockEmbeddingIface)(nil).Sixteen))
}

// Ten mocks base method.
func (m *MockEmbeddingIface) Ten(arg0 *other.Three) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ten", arg0)
}

// Ten indicates an expected call of Ten.
func (mr *MockEmbeddingIfaceMockRecorder) Ten(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ten", reflect.TypeOf((*MockEmbeddingIface)(nil).Ten), arg0)
}

// Thirteen mocks base method.
func (m *MockEmbeddingIface) Thirteen() (generics.Baz[generics.StructType], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Thirteen")
	ret0, _ := ret[0].(generics.Baz[generics.StructType])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Thirteen indicates an expected call of Thirteen.
func (mr *MockEmbeddingIfaceMockRecorder) Thirteen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Thirteen", reflect.TypeOf((*MockEmbeddingIface)(nil).Thirteen))
}

// Three mocks base method.
func (m *MockEmbeddingIface) Three(arg0 other.Three) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Three", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Three indicates an expected call of Three.
func (mr *MockEmbeddingIfaceMockRecorder) Three(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Three", reflect.TypeOf((*MockEmbeddingIface)(nil).Three), arg0)
}

// Twelve mocks base method.
func (m *MockEmbeddingIface) Twelve() (*other.Two[other.Three, error], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Twelve")
	ret0, _ := ret[0].(*other.Two[other.Three, error])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Twelve indicates an expected call of Twelve.
func (mr *MockEmbeddingIfaceMockRecorder) Twelve() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Twelve", reflect.TypeOf((*MockEmbeddingIface)(nil).Twelve))
}

// Twenty mocks base method.
func (m *MockEmbeddingIface) Twenty(arg0 *other.One[other.Three]) *other.Two[other.Three, error] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Twenty", arg0)
	ret0, _ := ret[0].(*other.Two[other.Three, error])
	return ret0
}

// Twenty indicates an expected call of Twenty.
func (mr *MockEmbeddingIfaceMockRecorder) Twenty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Twenty", reflect.TypeOf((*MockEmbeddingIface)(nil).Twenty), arg0)
}

// TwentyOne mocks base method.
func (m *MockEmbeddingIface) TwentyOne(arg0 *string) *other.Two[*other.Three, *error] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TwentyOne", arg0)
	ret0, _ := ret[0].(*other.Two[*other.Three, *error])
	return ret0
}

// TwentyOne indicates an expected call of TwentyOne.
func (mr *MockEmbeddingIfaceMockRecorder) TwentyOne(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TwentyOne", reflect.TypeOf((*MockEmbeddingIface)(nil).TwentyOne), arg0)
}

// Two mocks base method.
func (m *MockEmbeddingIface) Two(arg0 other.Three) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Two", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Two indicates an expected call of Two.
func (mr *MockEmbeddingIfaceMockRecorder) Two(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Two", reflect.TypeOf((*MockEmbeddingIface)(nil).Two), arg0)
}
