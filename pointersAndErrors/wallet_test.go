package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Depositing method works", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		assertCorrectMessage(t, wallet, Bitcoin(20))
	})

	t.Run("Withdrawing method works", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertCorrectMessage(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertCorrectMessage(t, wallet, startingBalance)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}
