package argparse

import (
	// "fmt"
	// "os"
	"testing"
)

func TestParserContext(t *testing.T) {
	pc := NewParserContext([]string{"a", "a", "a", "a", "a"})
	pc.Next()
}
