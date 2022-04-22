package nou

import (
	"fmt"
	"os"
	"strings"
)

func splitArgs(arg string) string {
	arr := strings.Split(arg, "/")
	return arr[len(arr)-1]
}

func usageMessage() {
	message := `Usage: module [OPTION] ...
Try 'module -help' for more information.
`
	arg := splitArgs(os.Args[0])
	message = strings.ReplaceAll(message, "module", arg)
	fmt.Print(message)
}

func DispUsage(usageMes string) {
	if len(os.Args) < 2 {
		if usageMes == "" {
			usageMessage()
		} else {
			fmt.Println(usageMes)
		}
		os.Exit(0)
	}
}
