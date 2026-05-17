package meetingrooms_test

import (
	"testing"

	meetingrooms "github.com/r-erema/workshop/algorithms/meeting_rooms"
	"github.com/stretchr/testify/assert"
)

func TestMeetingRooms(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		meetings [][2]int
		want     bool
	}{
		{
			name:     "can not attend all meetings",
			meetings: [][2]int{{0, 30}, {5, 10}, {15, 20}},
			want:     false,
		},
		{
			name:     "can attend all meetings",
			meetings: [][2]int{{5, 8}, {9, 15}},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, meetingrooms.CanAttendMeetings(tt.meetings))
		})
	}
}
