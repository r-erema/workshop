package timebasedkeyvaluestore_test

import (
	"testing"

	timebasedkeyvaluestore "github.com/r-erema/workshop/algorithms/time_based_key_value_store"
	"github.com/stretchr/testify/assert"
)

func TestTimeBasedKeyValueStore(t *testing.T) {
	t.Parallel()

	timeMap := timebasedkeyvaluestore.Constructor()
	timeMap.Set("foo", "bar", 1)
	assert.Equal(t, "bar", timeMap.Get("foo", 1))
	assert.Equal(t, "bar", timeMap.Get("foo", 3))
	timeMap.Set("foo", "bar2", 4)
	assert.Equal(t, "bar2", timeMap.Get("foo", 4))
	assert.Equal(t, "bar2", timeMap.Get("foo", 5))
}
