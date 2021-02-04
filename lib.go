package lib

import "runtime/debug"

func Stack() string {
	return string(debug.Stack())
}
