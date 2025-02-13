// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	chartRepoRepository "github.com/devtron-labs/devtron/pkg/chartRepo/repository"
	mock "github.com/stretchr/testify/mock"

	pg "github.com/go-pg/pg"
)

// ChartRepoRepository is an autogenerated mock type for the ChartRepoRepository type
type ChartRepoRepository struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *ChartRepoRepository) FindAll() ([]*chartRepoRepository.ChartRepo, error) {
	ret := _m.Called()

	var r0 []*chartRepoRepository.ChartRepo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*chartRepoRepository.ChartRepo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*chartRepoRepository.ChartRepo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*chartRepoRepository.ChartRepo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllWithDeploymentCount provides a mock function with given fields:
func (_m *ChartRepoRepository) FindAllWithDeploymentCount() ([]*chartRepoRepository.ChartRepoWithDeploymentCount, error) {
	ret := _m.Called()

	var r0 []*chartRepoRepository.ChartRepoWithDeploymentCount
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*chartRepoRepository.ChartRepoWithDeploymentCount, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*chartRepoRepository.ChartRepoWithDeploymentCount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*chartRepoRepository.ChartRepoWithDeploymentCount)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *ChartRepoRepository) FindById(id int) (*chartRepoRepository.ChartRepo, error) {
	ret := _m.Called(id)

	var r0 *chartRepoRepository.ChartRepo
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*chartRepoRepository.ChartRepo, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *chartRepoRepository.ChartRepo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chartRepoRepository.ChartRepo)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: name
func (_m *ChartRepoRepository) FindByName(name string) (*chartRepoRepository.ChartRepo, error) {
	ret := _m.Called(name)

	var r0 *chartRepoRepository.ChartRepo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*chartRepoRepository.ChartRepo, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *chartRepoRepository.ChartRepo); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chartRepoRepository.ChartRepo)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnection provides a mock function with given fields:
func (_m *ChartRepoRepository) GetConnection() *pg.DB {
	ret := _m.Called()

	var r0 *pg.DB
	if rf, ok := ret.Get(0).(func() *pg.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pg.DB)
		}
	}

	return r0
}

// GetDefault provides a mock function with given fields:
func (_m *ChartRepoRepository) GetDefault() (*chartRepoRepository.ChartRepo, error) {
	ret := _m.Called()

	var r0 *chartRepoRepository.ChartRepo
	var r1 error
	if rf, ok := ret.Get(0).(func() (*chartRepoRepository.ChartRepo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *chartRepoRepository.ChartRepo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chartRepoRepository.ChartRepo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkChartRepoDeleted provides a mock function with given fields: chartRepo, tx
func (_m *ChartRepoRepository) MarkChartRepoDeleted(chartRepo *chartRepoRepository.ChartRepo, tx *pg.Tx) error {
	ret := _m.Called(chartRepo, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*chartRepoRepository.ChartRepo, *pg.Tx) error); ok {
		r0 = rf(chartRepo, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: chartRepo, tx
func (_m *ChartRepoRepository) Save(chartRepo *chartRepoRepository.ChartRepo, tx *pg.Tx) error {
	ret := _m.Called(chartRepo, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*chartRepoRepository.ChartRepo, *pg.Tx) error); ok {
		r0 = rf(chartRepo, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: chartRepo, tx
func (_m *ChartRepoRepository) Update(chartRepo *chartRepoRepository.ChartRepo, tx *pg.Tx) error {
	ret := _m.Called(chartRepo, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*chartRepoRepository.ChartRepo, *pg.Tx) error); ok {
		r0 = rf(chartRepo, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewChartRepoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewChartRepoRepository creates a new instance of ChartRepoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChartRepoRepository(t mockConstructorTestingTNewChartRepoRepository) *ChartRepoRepository {
	mock := &ChartRepoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
