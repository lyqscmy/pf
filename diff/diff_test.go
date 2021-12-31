package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	testDo(t, "", "", nil)
	testDo(t, "abc", "abc", nil)
	testDo(t, "abc", "abd", nil, "-abc", "+abd")
	testDo(t, "abc\nabc", "abc\nabd", nil, "-abc", "+abd")
	testDo(t, "abc\nabc", "abc\nabc", nil)
}

func testDo(
	t *testing.T,
	input string,
	output string,
	expectedError error,
	expectedDiffs ...string,
) {
	diff, err := Diff([]byte(input), []byte(output), "")
	for _, expectedDiff := range expectedDiffs {
		assert.Contains(t, string(diff), expectedDiff, "diff does not contain %s", expectedDiff)
	}
	assert.Equal(t, expectedError, err)
}
