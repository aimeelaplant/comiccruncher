// Code generated by MockGen. DO NOT EDIT.
// Source: comic/repositories.go

// Package mock_comic is a generated GoMock package.
package mock_comic

import (
	comic "github.com/aimeelaplant/comiccruncher/comic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPublisherRepository is a mock of PublisherRepository interface
type MockPublisherRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherRepositoryMockRecorder
}

// MockPublisherRepositoryMockRecorder is the mock recorder for MockPublisherRepository
type MockPublisherRepositoryMockRecorder struct {
	mock *MockPublisherRepository
}

// NewMockPublisherRepository creates a new mock instance
func NewMockPublisherRepository(ctrl *gomock.Controller) *MockPublisherRepository {
	mock := &MockPublisherRepository{ctrl: ctrl}
	mock.recorder = &MockPublisherRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublisherRepository) EXPECT() *MockPublisherRepositoryMockRecorder {
	return m.recorder
}

// FindBySlug mocks base method
func (m *MockPublisherRepository) FindBySlug(slug comic.PublisherSlug) (*comic.Publisher, error) {
	ret := m.ctrl.Call(m, "FindBySlug", slug)
	ret0, _ := ret[0].(*comic.Publisher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBySlug indicates an expected call of FindBySlug
func (mr *MockPublisherRepositoryMockRecorder) FindBySlug(slug interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBySlug", reflect.TypeOf((*MockPublisherRepository)(nil).FindBySlug), slug)
}

// MockIssueRepository is a mock of IssueRepository interface
type MockIssueRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIssueRepositoryMockRecorder
}

// MockIssueRepositoryMockRecorder is the mock recorder for MockIssueRepository
type MockIssueRepositoryMockRecorder struct {
	mock *MockIssueRepository
}

// NewMockIssueRepository creates a new mock instance
func NewMockIssueRepository(ctrl *gomock.Controller) *MockIssueRepository {
	mock := &MockIssueRepository{ctrl: ctrl}
	mock.recorder = &MockIssueRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIssueRepository) EXPECT() *MockIssueRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockIssueRepository) Create(issue *comic.Issue) error {
	ret := m.ctrl.Call(m, "Create", issue)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockIssueRepositoryMockRecorder) Create(issue interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIssueRepository)(nil).Create), issue)
}

// CreateAll mocks base method
func (m *MockIssueRepository) CreateAll(issues []*comic.Issue) error {
	ret := m.ctrl.Call(m, "CreateAll", issues)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAll indicates an expected call of CreateAll
func (mr *MockIssueRepositoryMockRecorder) CreateAll(issues interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAll", reflect.TypeOf((*MockIssueRepository)(nil).CreateAll), issues)
}

// Update mocks base method
func (m *MockIssueRepository) Update(issue *comic.Issue) error {
	ret := m.ctrl.Call(m, "Update", issue)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockIssueRepositoryMockRecorder) Update(issue interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIssueRepository)(nil).Update), issue)
}

// FindByVendorID mocks base method
func (m *MockIssueRepository) FindByVendorID(vendorID string) (*comic.Issue, error) {
	ret := m.ctrl.Call(m, "FindByVendorID", vendorID)
	ret0, _ := ret[0].(*comic.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByVendorID indicates an expected call of FindByVendorID
func (mr *MockIssueRepositoryMockRecorder) FindByVendorID(vendorID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByVendorID", reflect.TypeOf((*MockIssueRepository)(nil).FindByVendorID), vendorID)
}

// FindAll mocks base method
func (m *MockIssueRepository) FindAll(c comic.IssueCriteria) ([]*comic.Issue, error) {
	ret := m.ctrl.Call(m, "FindAll", c)
	ret0, _ := ret[0].([]*comic.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockIssueRepositoryMockRecorder) FindAll(c interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIssueRepository)(nil).FindAll), c)
}

// MockCharacterRepository is a mock of CharacterRepository interface
type MockCharacterRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCharacterRepositoryMockRecorder
}

// MockCharacterRepositoryMockRecorder is the mock recorder for MockCharacterRepository
type MockCharacterRepositoryMockRecorder struct {
	mock *MockCharacterRepository
}

// NewMockCharacterRepository creates a new mock instance
func NewMockCharacterRepository(ctrl *gomock.Controller) *MockCharacterRepository {
	mock := &MockCharacterRepository{ctrl: ctrl}
	mock.recorder = &MockCharacterRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharacterRepository) EXPECT() *MockCharacterRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCharacterRepository) Create(c *comic.Character) error {
	ret := m.ctrl.Call(m, "Create", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCharacterRepositoryMockRecorder) Create(c interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCharacterRepository)(nil).Create), c)
}

// Update mocks base method
func (m *MockCharacterRepository) Update(c *comic.Character) error {
	ret := m.ctrl.Call(m, "Update", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockCharacterRepositoryMockRecorder) Update(c interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCharacterRepository)(nil).Update), c)
}

// FindBySlug mocks base method
func (m *MockCharacterRepository) FindBySlug(slug comic.CharacterSlug, includeIsDisabled bool) (*comic.Character, error) {
	ret := m.ctrl.Call(m, "FindBySlug", slug, includeIsDisabled)
	ret0, _ := ret[0].(*comic.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBySlug indicates an expected call of FindBySlug
func (mr *MockCharacterRepositoryMockRecorder) FindBySlug(slug, includeIsDisabled interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBySlug", reflect.TypeOf((*MockCharacterRepository)(nil).FindBySlug), slug, includeIsDisabled)
}

// FindAll mocks base method
func (m *MockCharacterRepository) FindAll(cr comic.CharacterCriteria) ([]*comic.Character, error) {
	ret := m.ctrl.Call(m, "FindAll", cr)
	ret0, _ := ret[0].([]*comic.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockCharacterRepositoryMockRecorder) FindAll(cr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCharacterRepository)(nil).FindAll), cr)
}

// UpdateAll mocks base method
func (m *MockCharacterRepository) UpdateAll(characters []*comic.Character) error {
	ret := m.ctrl.Call(m, "UpdateAll", characters)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAll indicates an expected call of UpdateAll
func (mr *MockCharacterRepositoryMockRecorder) UpdateAll(characters interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockCharacterRepository)(nil).UpdateAll), characters)
}

// Remove mocks base method
func (m *MockCharacterRepository) Remove(id comic.CharacterID) error {
	ret := m.ctrl.Call(m, "Remove", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockCharacterRepositoryMockRecorder) Remove(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCharacterRepository)(nil).Remove), id)
}

// Total mocks base method
func (m *MockCharacterRepository) Total(cr comic.CharacterCriteria) (int64, error) {
	ret := m.ctrl.Call(m, "Total", cr)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Total indicates an expected call of Total
func (mr *MockCharacterRepositoryMockRecorder) Total(cr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Total", reflect.TypeOf((*MockCharacterRepository)(nil).Total), cr)
}

// MockCharacterSourceRepository is a mock of CharacterSourceRepository interface
type MockCharacterSourceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCharacterSourceRepositoryMockRecorder
}

// MockCharacterSourceRepositoryMockRecorder is the mock recorder for MockCharacterSourceRepository
type MockCharacterSourceRepositoryMockRecorder struct {
	mock *MockCharacterSourceRepository
}

// NewMockCharacterSourceRepository creates a new mock instance
func NewMockCharacterSourceRepository(ctrl *gomock.Controller) *MockCharacterSourceRepository {
	mock := &MockCharacterSourceRepository{ctrl: ctrl}
	mock.recorder = &MockCharacterSourceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharacterSourceRepository) EXPECT() *MockCharacterSourceRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCharacterSourceRepository) Create(s *comic.CharacterSource) error {
	ret := m.ctrl.Call(m, "Create", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCharacterSourceRepositoryMockRecorder) Create(s interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCharacterSourceRepository)(nil).Create), s)
}

// FindAll mocks base method
func (m *MockCharacterSourceRepository) FindAll(criteria comic.CharacterSourceCriteria) ([]*comic.CharacterSource, error) {
	ret := m.ctrl.Call(m, "FindAll", criteria)
	ret0, _ := ret[0].([]*comic.CharacterSource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockCharacterSourceRepositoryMockRecorder) FindAll(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCharacterSourceRepository)(nil).FindAll), criteria)
}

// Remove mocks base method
func (m *MockCharacterSourceRepository) Remove(id comic.CharacterSourceID) error {
	ret := m.ctrl.Call(m, "Remove", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockCharacterSourceRepositoryMockRecorder) Remove(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCharacterSourceRepository)(nil).Remove), id)
}

// Raw mocks base method
func (m *MockCharacterSourceRepository) Raw(query string, params ...interface{}) error {
	varargs := []interface{}{query}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Raw", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Raw indicates an expected call of Raw
func (mr *MockCharacterSourceRepositoryMockRecorder) Raw(query interface{}, params ...interface{}) *gomock.Call {
	varargs := append([]interface{}{query}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockCharacterSourceRepository)(nil).Raw), varargs...)
}

// Update mocks base method
func (m *MockCharacterSourceRepository) Update(s *comic.CharacterSource) error {
	ret := m.ctrl.Call(m, "Update", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockCharacterSourceRepositoryMockRecorder) Update(s interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCharacterSourceRepository)(nil).Update), s)
}

// MockCharacterSyncLogRepository is a mock of CharacterSyncLogRepository interface
type MockCharacterSyncLogRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCharacterSyncLogRepositoryMockRecorder
}

// MockCharacterSyncLogRepositoryMockRecorder is the mock recorder for MockCharacterSyncLogRepository
type MockCharacterSyncLogRepositoryMockRecorder struct {
	mock *MockCharacterSyncLogRepository
}

// NewMockCharacterSyncLogRepository creates a new mock instance
func NewMockCharacterSyncLogRepository(ctrl *gomock.Controller) *MockCharacterSyncLogRepository {
	mock := &MockCharacterSyncLogRepository{ctrl: ctrl}
	mock.recorder = &MockCharacterSyncLogRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharacterSyncLogRepository) EXPECT() *MockCharacterSyncLogRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCharacterSyncLogRepository) Create(s *comic.CharacterSyncLog) error {
	ret := m.ctrl.Call(m, "Create", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCharacterSyncLogRepositoryMockRecorder) Create(s interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCharacterSyncLogRepository)(nil).Create), s)
}

// FindAllByCharacterID mocks base method
func (m *MockCharacterSyncLogRepository) FindAllByCharacterID(characterID comic.CharacterID) ([]*comic.CharacterSyncLog, error) {
	ret := m.ctrl.Call(m, "FindAllByCharacterID", characterID)
	ret0, _ := ret[0].([]*comic.CharacterSyncLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllByCharacterID indicates an expected call of FindAllByCharacterID
func (mr *MockCharacterSyncLogRepositoryMockRecorder) FindAllByCharacterID(characterID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllByCharacterID", reflect.TypeOf((*MockCharacterSyncLogRepository)(nil).FindAllByCharacterID), characterID)
}

// Update mocks base method
func (m *MockCharacterSyncLogRepository) Update(s *comic.CharacterSyncLog) error {
	ret := m.ctrl.Call(m, "Update", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockCharacterSyncLogRepositoryMockRecorder) Update(s interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCharacterSyncLogRepository)(nil).Update), s)
}

// FindByID mocks base method
func (m *MockCharacterSyncLogRepository) FindByID(id comic.CharacterSyncLogID) (*comic.CharacterSyncLog, error) {
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*comic.CharacterSyncLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockCharacterSyncLogRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockCharacterSyncLogRepository)(nil).FindByID), id)
}

// LastSyncs mocks base method
func (m *MockCharacterSyncLogRepository) LastSyncs(id comic.CharacterID) ([]*comic.LastSync, error) {
	ret := m.ctrl.Call(m, "LastSyncs", id)
	ret0, _ := ret[0].([]*comic.LastSync)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastSyncs indicates an expected call of LastSyncs
func (mr *MockCharacterSyncLogRepositoryMockRecorder) LastSyncs(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastSyncs", reflect.TypeOf((*MockCharacterSyncLogRepository)(nil).LastSyncs), id)
}

// MockCharacterIssueRepository is a mock of CharacterIssueRepository interface
type MockCharacterIssueRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCharacterIssueRepositoryMockRecorder
}

// MockCharacterIssueRepositoryMockRecorder is the mock recorder for MockCharacterIssueRepository
type MockCharacterIssueRepositoryMockRecorder struct {
	mock *MockCharacterIssueRepository
}

// NewMockCharacterIssueRepository creates a new mock instance
func NewMockCharacterIssueRepository(ctrl *gomock.Controller) *MockCharacterIssueRepository {
	mock := &MockCharacterIssueRepository{ctrl: ctrl}
	mock.recorder = &MockCharacterIssueRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharacterIssueRepository) EXPECT() *MockCharacterIssueRepositoryMockRecorder {
	return m.recorder
}

// CreateAll mocks base method
func (m *MockCharacterIssueRepository) CreateAll(cis []*comic.CharacterIssue) error {
	ret := m.ctrl.Call(m, "CreateAll", cis)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAll indicates an expected call of CreateAll
func (mr *MockCharacterIssueRepositoryMockRecorder) CreateAll(cis interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAll", reflect.TypeOf((*MockCharacterIssueRepository)(nil).CreateAll), cis)
}

// Create mocks base method
func (m *MockCharacterIssueRepository) Create(ci *comic.CharacterIssue) error {
	ret := m.ctrl.Call(m, "Create", ci)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCharacterIssueRepositoryMockRecorder) Create(ci interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCharacterIssueRepository)(nil).Create), ci)
}

// FindOneBy mocks base method
func (m *MockCharacterIssueRepository) FindOneBy(characterID comic.CharacterID, issueID comic.IssueID) (*comic.CharacterIssue, error) {
	ret := m.ctrl.Call(m, "FindOneBy", characterID, issueID)
	ret0, _ := ret[0].(*comic.CharacterIssue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneBy indicates an expected call of FindOneBy
func (mr *MockCharacterIssueRepositoryMockRecorder) FindOneBy(characterID, issueID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneBy", reflect.TypeOf((*MockCharacterIssueRepository)(nil).FindOneBy), characterID, issueID)
}

// InsertFast mocks base method
func (m *MockCharacterIssueRepository) InsertFast(issues []*comic.CharacterIssue) error {
	ret := m.ctrl.Call(m, "InsertFast", issues)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertFast indicates an expected call of InsertFast
func (mr *MockCharacterIssueRepositoryMockRecorder) InsertFast(issues interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertFast", reflect.TypeOf((*MockCharacterIssueRepository)(nil).InsertFast), issues)
}

// MockAppearancesByYearsRepository is a mock of AppearancesByYearsRepository interface
type MockAppearancesByYearsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAppearancesByYearsRepositoryMockRecorder
}

// MockAppearancesByYearsRepositoryMockRecorder is the mock recorder for MockAppearancesByYearsRepository
type MockAppearancesByYearsRepositoryMockRecorder struct {
	mock *MockAppearancesByYearsRepository
}

// NewMockAppearancesByYearsRepository creates a new mock instance
func NewMockAppearancesByYearsRepository(ctrl *gomock.Controller) *MockAppearancesByYearsRepository {
	mock := &MockAppearancesByYearsRepository{ctrl: ctrl}
	mock.recorder = &MockAppearancesByYearsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppearancesByYearsRepository) EXPECT() *MockAppearancesByYearsRepositoryMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockAppearancesByYearsRepository) List(slugs comic.CharacterSlug) (comic.AppearancesByYears, error) {
	ret := m.ctrl.Call(m, "List", slugs)
	ret0, _ := ret[0].(comic.AppearancesByYears)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockAppearancesByYearsRepositoryMockRecorder) List(slugs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAppearancesByYearsRepository)(nil).List), slugs)
}

// MockAppearancesByYearsMapRepository is a mock of AppearancesByYearsMapRepository interface
type MockAppearancesByYearsMapRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAppearancesByYearsMapRepositoryMockRecorder
}

// MockAppearancesByYearsMapRepositoryMockRecorder is the mock recorder for MockAppearancesByYearsMapRepository
type MockAppearancesByYearsMapRepositoryMockRecorder struct {
	mock *MockAppearancesByYearsMapRepository
}

// NewMockAppearancesByYearsMapRepository creates a new mock instance
func NewMockAppearancesByYearsMapRepository(ctrl *gomock.Controller) *MockAppearancesByYearsMapRepository {
	mock := &MockAppearancesByYearsMapRepository{ctrl: ctrl}
	mock.recorder = &MockAppearancesByYearsMapRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppearancesByYearsMapRepository) EXPECT() *MockAppearancesByYearsMapRepositoryMockRecorder {
	return m.recorder
}

// ListMap mocks base method
func (m *MockAppearancesByYearsMapRepository) ListMap(slugs ...comic.CharacterSlug) (map[comic.CharacterSlug][]comic.AppearancesByYears, error) {
	varargs := []interface{}{}
	for _, a := range slugs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMap", varargs...)
	ret0, _ := ret[0].(map[comic.CharacterSlug][]comic.AppearancesByYears)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMap indicates an expected call of ListMap
func (mr *MockAppearancesByYearsMapRepositoryMockRecorder) ListMap(slugs ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMap", reflect.TypeOf((*MockAppearancesByYearsMapRepository)(nil).ListMap), slugs...)
}

// MockStatsRepository is a mock of StatsRepository interface
type MockStatsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStatsRepositoryMockRecorder
}

// MockStatsRepositoryMockRecorder is the mock recorder for MockStatsRepository
type MockStatsRepositoryMockRecorder struct {
	mock *MockStatsRepository
}

// NewMockStatsRepository creates a new mock instance
func NewMockStatsRepository(ctrl *gomock.Controller) *MockStatsRepository {
	mock := &MockStatsRepository{ctrl: ctrl}
	mock.recorder = &MockStatsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStatsRepository) EXPECT() *MockStatsRepositoryMockRecorder {
	return m.recorder
}

// Stats mocks base method
func (m *MockStatsRepository) Stats() (comic.Stats, error) {
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(comic.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (mr *MockStatsRepositoryMockRecorder) Stats() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockStatsRepository)(nil).Stats))
}

// MockPopularRepository is a mock of PopularRepository interface
type MockPopularRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPopularRepositoryMockRecorder
}

// MockPopularRepositoryMockRecorder is the mock recorder for MockPopularRepository
type MockPopularRepositoryMockRecorder struct {
	mock *MockPopularRepository
}

// NewMockPopularRepository creates a new mock instance
func NewMockPopularRepository(ctrl *gomock.Controller) *MockPopularRepository {
	mock := &MockPopularRepository{ctrl: ctrl}
	mock.recorder = &MockPopularRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPopularRepository) EXPECT() *MockPopularRepositoryMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockPopularRepository) All(cr comic.PopularCriteria) ([]*comic.RankedCharacter, error) {
	ret := m.ctrl.Call(m, "All", cr)
	ret0, _ := ret[0].([]*comic.RankedCharacter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All
func (mr *MockPopularRepositoryMockRecorder) All(cr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockPopularRepository)(nil).All), cr)
}

// DC mocks base method
func (m *MockPopularRepository) DC(cr comic.PopularCriteria) ([]*comic.RankedCharacter, error) {
	ret := m.ctrl.Call(m, "DC", cr)
	ret0, _ := ret[0].([]*comic.RankedCharacter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DC indicates an expected call of DC
func (mr *MockPopularRepositoryMockRecorder) DC(cr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DC", reflect.TypeOf((*MockPopularRepository)(nil).DC), cr)
}

// Marvel mocks base method
func (m *MockPopularRepository) Marvel(cr comic.PopularCriteria) ([]*comic.RankedCharacter, error) {
	ret := m.ctrl.Call(m, "Marvel", cr)
	ret0, _ := ret[0].([]*comic.RankedCharacter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marvel indicates an expected call of Marvel
func (mr *MockPopularRepositoryMockRecorder) Marvel(cr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marvel", reflect.TypeOf((*MockPopularRepository)(nil).Marvel), cr)
}

// FindOneByDC mocks base method
func (m *MockPopularRepository) FindOneByDC(id comic.CharacterID) (*comic.RankedCharacter, error) {
	ret := m.ctrl.Call(m, "FindOneByDC", id)
	ret0, _ := ret[0].(*comic.RankedCharacter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByDC indicates an expected call of FindOneByDC
func (mr *MockPopularRepositoryMockRecorder) FindOneByDC(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByDC", reflect.TypeOf((*MockPopularRepository)(nil).FindOneByDC), id)
}

// FindOneByMarvel mocks base method
func (m *MockPopularRepository) FindOneByMarvel(id comic.CharacterID) (*comic.RankedCharacter, error) {
	ret := m.ctrl.Call(m, "FindOneByMarvel", id)
	ret0, _ := ret[0].(*comic.RankedCharacter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByMarvel indicates an expected call of FindOneByMarvel
func (mr *MockPopularRepositoryMockRecorder) FindOneByMarvel(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByMarvel", reflect.TypeOf((*MockPopularRepository)(nil).FindOneByMarvel), id)
}

// FindOneByAll mocks base method
func (m *MockPopularRepository) FindOneByAll(id comic.CharacterID) (*comic.RankedCharacter, error) {
	ret := m.ctrl.Call(m, "FindOneByAll", id)
	ret0, _ := ret[0].(*comic.RankedCharacter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByAll indicates an expected call of FindOneByAll
func (mr *MockPopularRepositoryMockRecorder) FindOneByAll(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByAll", reflect.TypeOf((*MockPopularRepository)(nil).FindOneByAll), id)
}

// MarvelTrending mocks base method
func (m *MockPopularRepository) MarvelTrending(limit, offset int) ([]*comic.RankedCharacter, error) {
	ret := m.ctrl.Call(m, "MarvelTrending", limit, offset)
	ret0, _ := ret[0].([]*comic.RankedCharacter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarvelTrending indicates an expected call of MarvelTrending
func (mr *MockPopularRepositoryMockRecorder) MarvelTrending(limit, offset interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarvelTrending", reflect.TypeOf((*MockPopularRepository)(nil).MarvelTrending), limit, offset)
}

// DCTrending mocks base method
func (m *MockPopularRepository) DCTrending(limit, offset int) ([]*comic.RankedCharacter, error) {
	ret := m.ctrl.Call(m, "DCTrending", limit, offset)
	ret0, _ := ret[0].([]*comic.RankedCharacter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DCTrending indicates an expected call of DCTrending
func (mr *MockPopularRepositoryMockRecorder) DCTrending(limit, offset interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DCTrending", reflect.TypeOf((*MockPopularRepository)(nil).DCTrending), limit, offset)
}

// MockPopularRefresher is a mock of PopularRefresher interface
type MockPopularRefresher struct {
	ctrl     *gomock.Controller
	recorder *MockPopularRefresherMockRecorder
}

// MockPopularRefresherMockRecorder is the mock recorder for MockPopularRefresher
type MockPopularRefresherMockRecorder struct {
	mock *MockPopularRefresher
}

// NewMockPopularRefresher creates a new mock instance
func NewMockPopularRefresher(ctrl *gomock.Controller) *MockPopularRefresher {
	mock := &MockPopularRefresher{ctrl: ctrl}
	mock.recorder = &MockPopularRefresherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPopularRefresher) EXPECT() *MockPopularRefresherMockRecorder {
	return m.recorder
}

// Refresh mocks base method
func (m *MockPopularRefresher) Refresh(view comic.MaterializedView) error {
	ret := m.ctrl.Call(m, "Refresh", view)
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh
func (mr *MockPopularRefresherMockRecorder) Refresh(view interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockPopularRefresher)(nil).Refresh), view)
}

// RefreshAll mocks base method
func (m *MockPopularRefresher) RefreshAll() error {
	ret := m.ctrl.Call(m, "RefreshAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshAll indicates an expected call of RefreshAll
func (mr *MockPopularRefresherMockRecorder) RefreshAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshAll", reflect.TypeOf((*MockPopularRefresher)(nil).RefreshAll))
}

// MockCharacterThumbRepository is a mock of CharacterThumbRepository interface
type MockCharacterThumbRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCharacterThumbRepositoryMockRecorder
}

// MockCharacterThumbRepositoryMockRecorder is the mock recorder for MockCharacterThumbRepository
type MockCharacterThumbRepositoryMockRecorder struct {
	mock *MockCharacterThumbRepository
}

// NewMockCharacterThumbRepository creates a new mock instance
func NewMockCharacterThumbRepository(ctrl *gomock.Controller) *MockCharacterThumbRepository {
	mock := &MockCharacterThumbRepository{ctrl: ctrl}
	mock.recorder = &MockCharacterThumbRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharacterThumbRepository) EXPECT() *MockCharacterThumbRepositoryMockRecorder {
	return m.recorder
}

// AllThumbnails mocks base method
func (m *MockCharacterThumbRepository) AllThumbnails(slugs ...comic.CharacterSlug) (map[comic.CharacterSlug]*comic.CharacterThumbnails, error) {
	varargs := []interface{}{}
	for _, a := range slugs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AllThumbnails", varargs...)
	ret0, _ := ret[0].(map[comic.CharacterSlug]*comic.CharacterThumbnails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllThumbnails indicates an expected call of AllThumbnails
func (mr *MockCharacterThumbRepositoryMockRecorder) AllThumbnails(slugs ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllThumbnails", reflect.TypeOf((*MockCharacterThumbRepository)(nil).AllThumbnails), slugs...)
}

// Thumbnails mocks base method
func (m *MockCharacterThumbRepository) Thumbnails(slug comic.CharacterSlug) (*comic.CharacterThumbnails, error) {
	ret := m.ctrl.Call(m, "Thumbnails", slug)
	ret0, _ := ret[0].(*comic.CharacterThumbnails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Thumbnails indicates an expected call of Thumbnails
func (mr *MockCharacterThumbRepositoryMockRecorder) Thumbnails(slug interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Thumbnails", reflect.TypeOf((*MockCharacterThumbRepository)(nil).Thumbnails), slug)
}
