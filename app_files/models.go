package models

import (
	"github.com/r0fls/reinhardt/src/model"
)

func main() {
	model.Connect("postgres", "postgres", "localhost", "235711", "test")
}
