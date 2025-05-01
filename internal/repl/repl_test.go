package repl

// func TestParseInput(t *testing.T) {
// 	cases := []struct {
// 		input    string
// 		expected []string
// 	}{
// 		{
// 			input:    " hello   world  ",
// 			expected: []string{"hello", "world"},
// 		},
// 		{
// 			input:    " HELLO   world  ",
// 			expected: []string{"hello", "world"},
// 		},
// 	}

// 	for _, c := range cases {
// 		cmd, args := parseInput(c.input)
// 		actual := []string{cmd, args}

// 		if len(actual) != len(c.expected) {
// 			t.Errorf("mismatched lengths: got %d, want %d", len(actual), len(c.expected))
// 			continue
// 		}

// 		for i := range actual {
// 			if actual[i] != c.expected[i] {
// 				t.Errorf("word mismatch at index %d: got %s, want %s", i, actual[i], c.expected[i])
// 			}
// 		}
// 	}
// }
