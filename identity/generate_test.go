package identity_test

import (
	"reflect"
	"testing"

	"github.com/phr3nzy/elections-blockchain/identity"
)

func Test(t *testing.T) {
	t.Run("generate blockchain id and compare", func(t *testing.T) {
		payload := identity.VoterInfo{Name: "Osama", NationalID: "0000-0000-0000-0000"}
		got := identity.GenerateBlockchainIdentity(payload)
		want := "09d1854b8830fc47ef48aee7055fdfdd36864e8a79d1a7e1d76ce5e35595f7e3"
		if reflect.DeepEqual(got, want) != true {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
