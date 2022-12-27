package mysql_test

import (
	"testing"

	"github.com/rommms07/idream-erp/core/source/mysql"
	"github.com/stretchr/testify/assert"
)

func Test_shouldConnectTheDatabaseInstanceToALocalDb(t *testing.T) {
	if err := mysql.Connect(); err != nil {
		assert.Nil(t, err, err.Error())
	}
}
