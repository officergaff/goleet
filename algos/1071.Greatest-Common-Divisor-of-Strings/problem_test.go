package algos

import (
	"testing"

	"github.com/xinxin001/goleet/common/utils"
)

func TestBase(t *testing.T) {
	got := gcdOfStrings("ABCABC", "ABC")
	want := "ABC"
	if want != got {
		utils.TestError(t, got, want)
	}
}

func TestCase1(t *testing.T) {
	got := gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX")
	want := "TAUXX"
	utils.TestError(t, got, want)
}

func TestCase2(t *testing.T) {
	got := gcdOfStrings("AAAAAAAAA", "AAACCC")
	want := ""
	utils.TestError(t, got, want)
}
