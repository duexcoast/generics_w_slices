package main

import "testing"

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Ryan",
			Sum:  100,
		},
		{
			From: "Dana",
			To:   "Chris",
			Sum:  25,
		},
	}
	AssertEqual(t, BalanceFor(transactions, "Ryan"), 100)
	AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	AssertEqual(t, BalanceFor(transactions, "Dana"), -25)
}
