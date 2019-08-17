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
		want := "47b1094306daf89edf0759839db82c29a1df2e07aaf8f0d27e93588b985d9231"
		if reflect.DeepEqual(got, want) != true {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
