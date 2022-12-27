package main

import (
	"github.com/rommms07/idream-erp/core/source/mysql"
)

func main() {
	if err := mysql.Connect(); err != nil {
		println(err.Error())
	}
}
