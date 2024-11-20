package q

import (
	"fmt"
	"os"
)

func Boom(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	os.Exit(1)
}
