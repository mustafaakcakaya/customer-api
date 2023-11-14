package main

import (
	customerApplication "CustomerAPI/cmd"
	"CustomerAPI/pkg/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	utils.PrintBanner()

	env := utils.GetGoEnv()

	fmt.Println("Customer API running on \"" + env + "\" environment.")

	if utils.GetGoPprofEnv() == "on" {
		go func() {
			if err := http.ListenAndServe("localhost:6060", nil); err != nil {
				log.Fatalf("Pprof failed: %v ", err)
			}
		}()
	}

	customerApplication.Execute(env)
}
