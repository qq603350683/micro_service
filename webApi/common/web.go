package common

import "html"

func XssFilter(str string) string {
	return html.EscapeString(str)
}
