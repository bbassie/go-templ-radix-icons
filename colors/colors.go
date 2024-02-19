package colors

import "strings"

func ColorsOrFallback(colors ...string) string {
	if len(colors) == 0 {
		return "currentColor"
	}
	return strings.Join(colors, " ")
}
