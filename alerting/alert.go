package alerting

import (
	"fmt"
	"strings"

	"github.com/gen2brain/beeep"
)

func Alert(op string, filename string) {
	alertTitle := fmt.Sprintf("%s to %s", strings.Title(strings.ToLower(op)), filename)
	err := beeep.Notify(alertTitle, filename, "")
	if err != nil {
		panic(err)
	}
}
