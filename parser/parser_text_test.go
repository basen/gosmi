package parser

import (
	"strings"
	"testing"
)

func TestParsePreservesDescriptionBlankLinesAndTrailingNewline(t *testing.T) {
	const mib = `TEST-MIB DEFINITIONS ::= BEGIN

testObject OBJECT-TYPE
    SYNTAX      INTEGER
    MAX-ACCESS  read-only
    STATUS      current
    DESCRIPTION
        "First line.

        Third line.
"
    ::= { iso 1 }

END
`

	module, err := Parse(strings.NewReader(mib))
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(module.Body.Nodes) != 1 || module.Body.Nodes[0].ObjectType == nil {
		t.Fatalf("expected one object type node")
	}

	want := "First line.\n\nThird line.\n"
	got := module.Body.Nodes[0].ObjectType.Description
	if got != want {
		t.Fatalf("description mismatch:\n got: %q\nwant: %q", got, want)
	}
}
