package list_backend

import (
	"fmt"

	"github.com/fatih/color"
)

func printMatches(
	path string,
	res matchResult,
	filter PrintFilter,
	explain bool,
) {
	if !shouldPrint(res.mode, filter) {
		return
	}
	printString := getPrintString(path, res, explain)
	if res.mode == matchExclude {
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
	res matchResult,
	explain bool,
) string {
	if explain {
		return fmt.Sprintf("%s [%s]", path, res.line)
	}
	return path
}
