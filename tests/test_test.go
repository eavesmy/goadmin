package tests

import (
	"testing"

	"github.com/eavesmy/goadmin/modules/config"
	"github.com/eavesmy/goadmin/tests/frameworks/gin"
)

func TestBlackBoxTestSuitOfBuiltInTables(t *testing.T) {
	BlackBoxTestSuitOfBuiltInTables(t, gin.NewHandler, config.DatabaseList{
		"default": {
			Host:       "127.0.0.1",
			Port:       "3306",
			User:       "root",
			Pwd:        "root",
			Name:       "go-admin-test",
			MaxIdleCon: 50,
			MaxOpenCon: 150,
			Driver:     config.DriverMysql,
		},
	})
}
