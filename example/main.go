package main

import (
	"os"

	"github.com/withmandala/go-log"
)

func main() {
	l := log.New(os.Stdout)
	l.Info("hello krythera")
}