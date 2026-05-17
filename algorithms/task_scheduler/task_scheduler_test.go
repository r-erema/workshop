package taskscheduler_test

import (
	"testing"

	taskscheduler "github.com/r-erema/workshop/algorithms/task_scheduler"
	"github.com/stretchr/testify/assert"
)

func TestLeastInterval(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		tasks []byte
		n     int
		want  int
	}{
		{
			name:  "1 task",
			tasks: []byte{'A'},
			n:     22,
			want:  1,
		},
		{
			name:  "2 tasks with gap",
			tasks: []byte{'A', 'B', 'B'},
			n:     2,
			want:  4,
		},
		{
			name:  "many tasks",
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K'},
			n:     7,
			want:  18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, taskscheduler.LeastInterval(tt.tasks, tt.n))
		})
	}
}
