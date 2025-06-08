package hello

import "testing"

func TestInc(t *testing.T) {
    Inc()
    if count != 1 {
        t.Fatal("expected count = 1, got", count)
    }
}
