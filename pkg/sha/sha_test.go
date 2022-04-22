package sha

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	digest := Digest([]byte("abc"))
	fmt.Printf("Digest:\n%x", digest)
}
