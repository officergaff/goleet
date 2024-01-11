package algos

import (
	"testing"

	"github.com/xinxin001/goleet/common/utils"
)

func TestBase(t *testing.T) {
	got := canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)
	want := true
	utils.TestError(t, got, want)
}
