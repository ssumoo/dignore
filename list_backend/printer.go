package list_backend

import (
	"fmt"

	"github.com/fatih/color"
)

func printMatches(
	res MatchResult,
	filter PrintFilter,
	explain bool,
) {
	if !shouldPrint(res.Mode, filter) {
		return
	}
	printString := getPrintString(res.Path, res, explain)
	if res.Mode == matchExclude {
		color.Red(printString)
	} else {
		color.Green(printString)
	}
}

func shouldPrint(
	mode matchMode,
	filter PrintFilter,
) bool {
	if filter == PrintAll {
		return true
	}
	if filter == PrintExclude && mode == matchExclude {
		return true
	}
	if filter == PrintInclude && (mode == matchInclude || mode == matchNone) {
		return true
	}
	return false
}

func getPrintString(
	path string,
	res MatchResult,
	explain bool,
) string {
	if explain {
		finalRes := "include"
		if res.Mode == matchExclude {
			finalRes = "exclude"
		}
		return fmt.Sprintf("[%s] %s [%s]", finalRes, path, res.Line)
	}
	return path
}
