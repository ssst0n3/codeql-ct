package main

import (
	"codeql-ct/config"
	"codeql-ct/router"
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func main() {
	r := router.InitRouter()
	awesome_error.CheckFatal(r.Run(fmt.Sprintf(":%s", config.LocalListenPort)))
}
