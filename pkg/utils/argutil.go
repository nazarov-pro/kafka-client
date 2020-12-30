package utils

import (
	"strings"
)

// ArgumentParser - Agument list parser generates map of arguments with their value(s)
func ArgumentParser(args []string) map[string]interface{} {
	argsMap := make(map[string]interface{})
	for i := 1; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				argsMap[arg[1:]] = args[i+1]
			} else if strings.Contains(strings.ToLower(arg[1:]), "enabled") {
				argsMap[arg[1:]] = true
			}
		}
	}
	return argsMap
}