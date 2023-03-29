package poll_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/playgroundgo/genlib/poll"
)

func TestPollWhile(t *testing.T) {
	i := 0
	nums := make([]int, 0, 5)
	expected := []int{1, 2, 3, 4, 5}
	poll.PollWhile(context.Background(), 100*time.Millisecond, func() bool {
		i += 1
		if i <= 5 {
			nums = append(nums, i)
			return true
		}
		return false
	})

	if !reflect.DeepEqual(nums, expected) {
		t.Fatalf("expected the nums slice to be %v, got %v", expected, nums)
	}
}
