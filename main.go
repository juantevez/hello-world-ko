package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("I'm running on OS %s and arch %s", runtime.GOOS, runtime.GOARCH)
}

/*
export KO_DOCKER_REPO=gcr.io/YOUR_PROJECT/my-app
ko build github.com/juantevez/hello-world-ko
*/