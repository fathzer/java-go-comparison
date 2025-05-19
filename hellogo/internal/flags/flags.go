package flags

import (
	"flag"
	"fmt"
	"strings"
)

func ParseLoopsFlag(args []string) (*int, error) {
	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "-piLoops") && !strings.HasPrefix(arg, "--") {
			return nil, fmt.Errorf("Error: -piLoops is not supported. Use -pl or --piLoops instead.")
		}
		if strings.HasPrefix(arg, "--pl") {
			return nil, fmt.Errorf("Error: --pl is not supported. Use -pl or --piLoops instead.")
		}
	}
	var newArgs []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "--piLoops") {
			arg = strings.Replace(arg, "--piLoops", "-pl", 1)
		}
		newArgs = append(newArgs, arg)
	}
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	loops := fs.Int("pl", 2000, "number of Pi calculations to perform")
	err := fs.Parse(newArgs[1:])
	if err != nil {
		return nil, err
	}
	return loops, nil
}
