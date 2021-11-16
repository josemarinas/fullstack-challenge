package main

import (
	"server/cmd"
)

//go:generate sqlboiler --wipe psql && go run github.com/99designs/gqlgen

func main() {
	cmd.Execute()
}
