package q

import (
	"fmt"
	"os"
)

func Boom(msg string) {
	fmt.Fprint(os.Stderr, msg)
	os.Exit(1)
}
