package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s but wanted %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()

		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got.Error() != want {
			t.Errorf("got %s but wanted %s", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()

		if got != nil {
			t.Fatal("got an error but didn't want one.")
		}
	}
	t.Run("deposit 10 bitcoin", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10.00)
		assertBalance(t, wallet, Bitcoin(10.00))

	})

	t.Run("withdraw 8 bitcoin", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(8))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(12.00))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20.00)}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "insufficient funds")
		assertBalance(t, wallet, Bitcoin(20.00))
	})
}
