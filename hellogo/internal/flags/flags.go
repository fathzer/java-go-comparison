package flags

import (
	"fmt"
	"strconv"
	"strings"
)

// parseIntFlag parses an integer command line flag with the given parameters.
// It handles both short and long form flags and performs validation.
// shortFlag: the short form of the flag (e.g., "pl" for -pl)
// longFlag: the long form of the flag (e.g., "piLoops" for --piLoops)
// defaultValue: the default value if the flag is not provided
// args: command line arguments
func ParseIntFlag(shortFlag, longFlag string, defaultValue int, args []string) (*int, error) {
	if err := checkIllegalUsage(shortFlag, longFlag, args); err != nil {
		return nil, err
	}
	var value *int = &defaultValue
	value, err := parseOneIntegerArgument("-"+shortFlag+"=", defaultValue, args)
	if err != nil {
		return nil, err
	}
	value, err = parseOneIntegerArgument("--"+longFlag+"=", *value, args)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func checkIllegalUsage(shortFlag, longFlag string, args []string) error {
	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "--"+shortFlag+"=") || strings.HasPrefix(arg, "-"+longFlag+"=") {
			return fmt.Errorf("error: %s is not supported. Use -%s or --%s instead", arg, shortFlag, longFlag)
		}
	}
	return nil
}

func parseOneIntegerArgument(prefix string, defaultValue int, args []string) (*int, error) {
	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, prefix) {
			i, err := strconv.Atoi(arg[len(prefix):])
			if err != nil {
				return nil, err
			}
			return &i, nil
		}
	}
	return &defaultValue, nil
}

func ParseStringFlag(shortFlag, longFlag string, defaultValue string, args []string) (*string, error) {
	if err := checkIllegalUsage(shortFlag, longFlag, args); err != nil {
		return nil, err
	}
	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "-"+shortFlag+"=") {
			result := arg[len("-"+shortFlag+"="):]
			return &result, nil
		}
		if strings.HasPrefix(arg, "--"+longFlag+"=") {
			result := arg[len("--"+longFlag+"="):]
			return &result, nil
		}
	}
	return &defaultValue, nil
}

// ParseBoolFlag checks if a boolean flag is present in the command line arguments.
// Boolean flags are specified without a value (e.g., -b or --blackPlaying).
// It returns true if either the short or long form of the flag is present, false otherwise.
// It returns an error if an illegal flag format is detected.
func ParseBoolFlag(shortFlag, longFlag string, args []string) (bool, error) {
	// Check for illegal usages like --shortFlag or -longFlag
	for _, arg := range args[1:] {
		if arg == "--"+shortFlag || arg == "-"+longFlag {
			return false, fmt.Errorf("invalid flag: %s is not supported. Use -%s or --%s", arg, shortFlag, longFlag)
		}
	}

	// Check for the presence of the flag
	for _, arg := range args[1:] {
		if arg == "-"+shortFlag || arg == "--"+longFlag {
			return true, nil
		}
	}
	return false, nil
}
