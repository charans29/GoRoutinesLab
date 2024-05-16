package main

import ("testing"
		"fmt")

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
