package designaddandsearchwordsdatastructure_test

import (
	"testing"

	designaddandsearchwordsdatastructure "github.com/r-erema/workshop/algorithms/design_add_and_search_words_data_structure"
	"github.com/stretchr/testify/assert"
)

func TestWordDictionary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		scenarioFn func(t *testing.T)
	}{
		{
			name: "scenario 1",
			scenarioFn: func(t *testing.T) {
				t.Helper()

				dict := designaddandsearchwordsdatastructure.Constructor()
				dict.AddWord("bad")
				dict.AddWord("dad")
				dict.AddWord("mad")
				assert.False(t, dict.Search("pad"))
				assert.True(t, dict.Search("bad"))
				assert.True(t, dict.Search(".ad"))
				assert.True(t, dict.Search("..d"))
				assert.True(t, dict.Search("b.."))
			},
		},
		{
			name: "scenario 2",
			scenarioFn: func(t *testing.T) {
				t.Helper()

				dict := designaddandsearchwordsdatastructure.Constructor()
				dict.AddWord("a")
				dict.AddWord("ab")
				assert.True(t, dict.Search("a"))
				assert.True(t, dict.Search("a."))
			},
		},
		{
			name: "scenario 3",
			scenarioFn: func(t *testing.T) {
				t.Helper()

				dict := designaddandsearchwordsdatastructure.Constructor()
				dict.AddWord("at")
				dict.AddWord("and")
				dict.AddWord("an")
				dict.AddWord("add")
				assert.False(t, dict.Search("a"))
				assert.False(t, dict.Search(".at"))

				dict.AddWord("bat")

				assert.True(t, dict.Search(".at"))
			},
		},
		{
			name: "scenario 4",
			scenarioFn: func(t *testing.T) {
				t.Helper()

				dict := designaddandsearchwordsdatastructure.Constructor()
				assert.False(t, dict.Search("a"))
			},
		},
		{
			name: "scenario 5",
			scenarioFn: func(t *testing.T) {
				t.Helper()

				dict := designaddandsearchwordsdatastructure.Constructor()
				dict.AddWord("a")
				dict.AddWord("ab")
				assert.True(t, dict.Search("."))
				assert.True(t, dict.Search(".."))
			},
		},
		{
			name: "scenario 6",
			scenarioFn: func(t *testing.T) {
				t.Helper()

				dict := designaddandsearchwordsdatastructure.Constructor()
				dict.AddWord("at")
				dict.AddWord("and")
				dict.AddWord("an")
				dict.AddWord("add")
				dict.AddWord("bat")
				assert.False(t, dict.Search("b."))
				assert.True(t, dict.Search(".."))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.scenarioFn(t)
		})
	}
}
