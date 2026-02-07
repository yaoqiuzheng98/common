package environment

import "os"

func GetEnvironment() Environment {
	env := os.Getenv("ENV")
	switch env {
	case Development.String(), Production.String(), Test.String():
		return Environment(env)
	}
	panic("Environment variable not set")
}
