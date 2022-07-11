// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	amqp "github.com/streadway/amqp"

	mock "github.com/stretchr/testify/mock"

	mongo "go.mongodb.org/mongo-driver/mongo"

	redis "github.com/go-redis/redis/v8"

	sql "database/sql"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// MongoDB provides a mock function with given fields: uri, database
func (_m *Config) MongoDB(uri string, database string) (*mongo.Client, *mongo.Database) {
	ret := _m.Called(uri, database)

	var r0 *mongo.Client
	if rf, ok := ret.Get(0).(func(string, string) *mongo.Client); ok {
		r0 = rf(uri, database)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Client)
		}
	}

	var r1 *mongo.Database
	if rf, ok := ret.Get(1).(func(string, string) *mongo.Database); ok {
		r1 = rf(uri, database)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*mongo.Database)
		}
	}

	return r0, r1
}

// Postgresql provides a mock function with given fields: dsn, SetMaxIdleConns, SetMaxOpenConns
func (_m *Config) Postgresql(dsn string, SetMaxIdleConns int, SetMaxOpenConns int) *sql.DB {
	ret := _m.Called(dsn, SetMaxIdleConns, SetMaxOpenConns)

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func(string, int, int) *sql.DB); ok {
		r0 = rf(dsn, SetMaxIdleConns, SetMaxOpenConns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	return r0
}

// RabbitMQ provides a mock function with given fields: addrs
func (_m *Config) RabbitMQ(addrs string) *amqp.Connection {
	ret := _m.Called(addrs)

	var r0 *amqp.Connection
	if rf, ok := ret.Get(0).(func(string) *amqp.Connection); ok {
		r0 = rf(addrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amqp.Connection)
		}
	}

	return r0
}

// Redis provides a mock function with given fields: address, password
func (_m *Config) Redis(address string, password string) redis.UniversalClient {
	ret := _m.Called(address, password)

	var r0 redis.UniversalClient
	if rf, ok := ret.Get(0).(func(string, string) redis.UniversalClient); ok {
		r0 = rf(address, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(redis.UniversalClient)
		}
	}

	return r0
}

type NewConfigT interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfig creates a new instance of Config. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfig(t NewConfigT) *Config {
	mock := &Config{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
