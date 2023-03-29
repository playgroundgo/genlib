package poll_test

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/playgroundgo/genlib/poll"
)

func TestPollWhile(t *testing.T) {
	testPollWhile(t, false)
	testPollWhile(t, true)
}

func testPollWhile(t *testing.T, expectErr bool) {
	i := 0
	nums := make([]int, 0, 5)
	expected := []int{1, 2, 3, 4, 5}
	expectedErr := errors.New("test error")

	err := poll.PollWhile(context.Background(), 10*time.Millisecond, func() (bool, error) {
		i += 1
		if i <= 5 {
			nums = append(nums, i)
			return true, nil
		}
		if expectErr {
			return false, expectedErr
		}
		return false, nil
	})

	if expectErr {
		if !errors.Is(err, expectedErr) {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
		return
	}

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(nums, expected) {
		t.Fatalf("expected the nums slice to be %v, got %v", expected, nums)
	}
}
