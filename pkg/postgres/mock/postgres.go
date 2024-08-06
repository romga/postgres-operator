// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/postgres/postgres.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	"github.com/movetokube/postgres-operator/pkg/postgres"
)

// MockPG is a mock of PG interface
type MockPG struct {
	ctrl     *gomock.Controller
	recorder *MockPGMockRecorder
}

// MockPGMockRecorder is the mock recorder for MockPG
type MockPGMockRecorder struct {
	mock *MockPG
}

// NewMockPG creates a new mock instance
func NewMockPG(ctrl *gomock.Controller) *MockPG {
	mock := &MockPG{ctrl: ctrl}
	mock.recorder = &MockPGMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPG) EXPECT() *MockPGMockRecorder {
	return m.recorder
}

// CreateDB mocks base method
func (m *MockPG) CreateDB(dbname, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDB", dbname, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDB indicates an expected call of CreateDB
func (mr *MockPGMockRecorder) CreateDB(dbname, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDB", reflect.TypeOf((*MockPG)(nil).CreateDB), dbname, username)
}

// CreateSchema mocks base method
func (m *MockPG) CreateSchema(db, role, schema string, logger logr.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchema", db, role, schema, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSchema indicates an expected call of CreateSchema
func (mr *MockPGMockRecorder) CreateSchema(db, role, schema, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchema", reflect.TypeOf((*MockPG)(nil).CreateSchema), db, role, schema, logger)
}

// CreateExtension mocks base method
func (m *MockPG) CreateExtension(db, extension string, logger logr.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExtension", db, extension, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExtension indicates an expected call of CreateExtension
func (mr *MockPGMockRecorder) CreateExtension(db, extension, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExtension", reflect.TypeOf((*MockPG)(nil).CreateExtension), db, extension, logger)
}

// CreateGroupRole mocks base method
func (m *MockPG) CreateGroupRole(role string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroupRole", role)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGroupRole indicates an expected call of CreateGroupRole
func (mr *MockPGMockRecorder) CreateGroupRole(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroupRole", reflect.TypeOf((*MockPG)(nil).CreateGroupRole), role)
}

// CreateUserRole mocks base method
func (m *MockPG) CreateUserRole(role, password string,iamAuthentication bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserRole", role, password, false)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserRole indicates an expected call of CreateUserRole
func (mr *MockPGMockRecorder) CreateUserRole(role, password interface{}, iamAuthentication bool) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserRole", reflect.TypeOf((*MockPG)(nil).CreateUserRole), role, password, false)
}

// UpdatePassword mocks base method
func (m *MockPG) UpdatePassword(role, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", role, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockPGMockRecorder) UpdatePassword(role, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockPG)(nil).UpdatePassword), role, password)
}

// GrantRole mocks base method
func (m *MockPG) GrantRole(role, grantee string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantRole", role, grantee)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantRole indicates an expected call of GrantRole
func (mr *MockPGMockRecorder) GrantRole(role, grantee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantRole", reflect.TypeOf((*MockPG)(nil).GrantRole), role, grantee)
}

// SetSchemaPrivileges mocks base method
func (m *MockPG) SetSchemaPrivileges(privileges postgres.PostgresSchemaPrivileges, logger logr.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSchemaPrivileges", privileges.DB, privileges.Creator, privileges.Role, privileges.Schema, privileges.Privs, privileges.CreateSchema, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSchemaPrivileges indicates an expected call of SetSchemaPrivileges
func (mr *MockPGMockRecorder) SetSchemaPrivileges(db, creator, role, schema, privs, createSchema, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSchemaPrivileges", reflect.TypeOf((*MockPG)(nil).SetSchemaPrivileges), db, creator, role, schema, privs, createSchema, logger)
}

// RevokeRole mocks base method
func (m *MockPG) RevokeRole(role, revoked string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeRole", role, revoked)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeRole indicates an expected call of RevokeRole
func (mr *MockPGMockRecorder) RevokeRole(role, revoked interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeRole", reflect.TypeOf((*MockPG)(nil).RevokeRole), role, revoked)
}

// AlterDefaultLoginRole mocks base method
func (m *MockPG) AlterDefaultLoginRole(role, setRole string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlterDefaultLoginRole", role, setRole)
	ret0, _ := ret[0].(error)
	return ret0
}

// AlterDefaultLoginRole indicates an expected call of AlterDefaultLoginRole
func (mr *MockPGMockRecorder) AlterDefaultLoginRole(role, setRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlterDefaultLoginRole", reflect.TypeOf((*MockPG)(nil).AlterDefaultLoginRole), role, setRole)
}

// DropDatabase mocks base method
func (m *MockPG) DropDatabase(db string, logger logr.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropDatabase", db, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropDatabase indicates an expected call of DropDatabase
func (mr *MockPGMockRecorder) DropDatabase(db, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropDatabase", reflect.TypeOf((*MockPG)(nil).DropDatabase), db, logger)
}

// DropRole mocks base method
func (m *MockPG) DropRole(role, newOwner, database string, logger logr.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropRole", role, newOwner, database, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropRole indicates an expected call of DropRole
func (mr *MockPGMockRecorder) DropRole(role, newOwner, database, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropRole", reflect.TypeOf((*MockPG)(nil).DropRole), role, newOwner, database, logger)
}

// GetUser mocks base method
func (m *MockPG) GetUser() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUser indicates an expected call of GetUser
func (mr *MockPGMockRecorder) GetUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockPG)(nil).GetUser))
}

// GetDefaultDatabase mocks base method
func (m *MockPG) GetDefaultDatabase() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultDatabase")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDefaultDatabase indicates an expected call of GetDefaultDatabase
func (mr *MockPGMockRecorder) GetDefaultDatabase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultDatabase", reflect.TypeOf((*MockPG)(nil).GetUser))
}
