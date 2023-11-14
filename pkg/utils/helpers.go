package utils

import (
	"fmt"
	"os"
)

func GetSwagHostEnv() string {
	e, ok := os.LookupEnv("SWAG_URL")
	if !ok {
		e = defaultSwagUrl
	}
	return e
}
func PrintBanner() {
	fmt.Println(asciiArt)
}

func GetGoEnv() string {
	e, ok := os.LookupEnv("GO_ENV")
	if !ok {
		e = defaultGoEnv
	}
	return e
}
func GetGoPprofEnv() string {
	pprof, ok := os.LookupEnv("GO_PPROF")
	if !ok {
		pprof = "off"
	}
	return pprof
}

const defaultSwagUrl = "localhost:1953"
const defaultGoEnv = "local-qa"
const asciiArt = "\n██╗░░██╗███████╗██████╗░░██████╗██╗░█████╗░██████╗░████████╗░█████╗░██╗░░██╗\n██║░░██║██╔════╝██╔══██╗██╔════╝██║██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗██║░██╔╝\n███████║█████╗░░██████╔╝╚█████╗░██║██║░░██║██████╔╝░░░██║░░░███████║█████═╝░\n██╔══██║██╔══╝░░██╔═══╝░░╚═══██╗██║██║░░██║██╔══██╗░░░██║░░░██╔══██║██╔═██╗░\n██║░░██║███████╗██║░░░░░██████╔╝██║╚█████╔╝██║░░██║░░░██║░░░██║░░██║██║░╚██╗\n╚═╝░░╚═╝╚══════╝╚═╝░░░░░╚═════╝░╚═╝░╚════╝░╚═╝░░╚═╝░░░╚═╝░░░╚═╝░░╚═╝╚═╝░░╚═╝"
