package orm

import (
	"testing"
)

func setupBeforeAllTestCases(t *testing.T) {
	repository := GetRepository()
	repository.Initialize("file::memory:?cache=shared", SQLite)
	repository.Migrate()
}

func setupTestCase(t *testing.T) func(t *testing.T) {
	repository := GetRepository()
	return func(t *testing.T) {
		repository.DeleteAllData()
	}
}
