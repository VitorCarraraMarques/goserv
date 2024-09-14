package auth

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadEnvVars() {
	f, err := os.Open("web/server/auth/secrets.txt")
	if err != nil {
		log.Fatalf("[ERROR] Error loading Env Vars: %s\n", err.Error())
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		envvar := strings.Split(line, "=")
		name := envvar[0]
		val := envvar[1]
		err := os.Setenv(name, val)
		if err != nil {
			log.Fatalf("[ERROR] Error settings Env Vars: %s\n", err.Error())
		}
	}
	f.Close()
}
