package mustache

import (
	"errors"
	"strings"
	"testing"
)

// TestDeeplyNestedSectionsDoNotOverflow ensures that a template consisting of a
// large number of nested section openers is rejected with an error instead of
// exhausting the goroutine stack and aborting the process.
func TestDeeplyNestedSectionsDoNotOverflow(t *testing.T) {
	tmpl := strings.Repeat("{{#a}}", maxNestingDepth+1)
	_, err := ParseString(tmpl)
	if err == nil {
		t.Fatal("expected an error for excessively nested sections, got nil")
	}
	var perr ParseError
	if !errors.As(err, &perr) || perr.Code != ErrNestingTooDeep {
		t.Fatalf("expected ErrNestingTooDeep, got %v", err)
	}
}

// TestReasonablyNestedSectionsStillParse ensures the depth guard does not
// reject templates with ordinary (well-formed, closed) nesting.
func TestReasonablyNestedSectionsStillParse(t *testing.T) {
	const depth = 50
	tmpl := strings.Repeat("{{#a}}", depth) + "x" + strings.Repeat("{{/a}}", depth)
	if _, err := ParseString(tmpl); err != nil {
		t.Fatalf("valid nested template failed to parse: %v", err)
	}
}
