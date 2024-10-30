package main

import (
	"flag"

	_ "time/tzdata"

	"github.com/aceberg/ForAuth/internal/web"
)

const dirPath = "/data/ForAuth"
const nodePath = ""

func main() {
	dirPtr := flag.String("d", dirPath, "Path to config dir")
	nodePtr := flag.String("n", nodePath, "Path to node modules")
	flag.Parse()

	web.Gui(*dirPtr, *nodePtr) // webgui.go
}
