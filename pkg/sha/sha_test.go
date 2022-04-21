package sha

import "testing"

func TestHash(t *testing.T) {
	Digest([]byte("abc"))
}
