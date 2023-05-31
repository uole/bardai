package main

import (
	"flag"
	"fmt"
	"git.nspix.com/golang/kos"
	"github.com/uole/bardai"
	"github.com/uole/bardai/version"
	"os"
)

func main() {
	var (
		err error
	)
	flag.Parse()
	svr := kos.Init(
		kos.WithName("github.com/uole/bardai", version.Version),
		kos.WithServer(bardai.New()),
	)
	if err = svr.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
