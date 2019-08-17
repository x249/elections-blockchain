package identity_test

import (
	"reflect"
	"testing"

	"github.com/phr3nzy/elections-blockchain/identity"
)

func Test(t *testing.T) {
	t.Run("test byte conversion", func(t *testing.T) {
		got := identity.ConvertToByte("Hello")
		want := []byte("Hello")
		if reflect.DeepEqual(got, want) != true {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("test hashing", func(t *testing.T) {
		got := identity.Hash("Hello")
		want := "53427dce55a4d75552a20f0deb0c42e1545b1411a6a979d0077f3fe26c3eba15"
		if reflect.DeepEqual(got, want) != true {
			t.Logf("got %v want %v", got, want)
		}
	})
}
