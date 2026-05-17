package groupanagrams_test

import (
	"sort"
	"testing"

	groupanagrams "github.com/r-erema/workshop/algorithms/group_anagrams"
	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "Case 0",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want:  [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := groupanagrams.GroupAnagrams(tt.input)

			sort.Slice(result, func(i, j int) bool {
				return len(result[i]) > len(result[j])
			})

			for i := range result {
				sort.Strings(result[i])
			}

			assert.Equal(t, tt.want, result)
		})
	}
}
