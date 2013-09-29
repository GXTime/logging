package logging

import (
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	SetLevel(INFO)
	Debug("%d, %s", 1, "OK")
	Info("%d, %s", 2, "OK")
	Warning("%d, %s", 3, "OK")
	Error("%d, %s", 4, "OK")
}
