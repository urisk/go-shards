package server

import (
	"fmt"
	"go-shards/utils"
)

// Init <function>
// is used to initialize server and all the corresponding services such as DB, Utils, Workers
func Init() {
	// utils
	utils.InitEnvVars()

	r := NewRouter()
	err := r.Run(":6969")
	if err != nil {
		fmt.Println(err)
		return
	}
}
