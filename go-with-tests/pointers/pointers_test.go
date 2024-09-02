package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but want one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := &Wallet{}
		wallet.Deposit(Bitcoin(10))
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		want := Bitcoin(10)
		assertBalance(t, *wallet, want)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := &Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))
		want := Bitcoin(5)
		assertNoError(t, err)
		assertBalance(t, *wallet, want)
	})
	t.Run("withdraw when insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := &Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, *wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
