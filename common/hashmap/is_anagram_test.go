package hashmap

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	target := isAnagram("aacc", "ccac")
	if target != false {
		t.Errorf("isAnagram(aacc,ccac) = %v; want false", target)
	}
}
