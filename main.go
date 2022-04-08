package main

import "github.com/lateralusd/laserver/cmd"

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
