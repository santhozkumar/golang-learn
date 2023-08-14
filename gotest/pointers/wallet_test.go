package fintech

import (
	"testing"
)


func TestWallet(t *testing.T){
    t.Run("Deposit", func(t *testing.T){

        wallet := Wallet{balance: Bitcoin(100)}
        wallet.Deposit(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(110))
    })

    t.Run("Withdraw", func(t *testing.T){

        wallet := Wallet{balance: Bitcoin(100)}
        err := wallet.Withdraw(Bitcoin(10))

        assertNoError(t, err)
        assertBalance(t, wallet, Bitcoin(90))
    })

    t.Run("Withdraw Insufficient Funds", func(t *testing.T){
        startingBalance := Bitcoin(20)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(Bitcoin(100))

        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, startingBalance)
    })
}


func TestString(t *testing.T){
    got := Bitcoin(10).String()
    want := "10 BT"

    assert(t, got, want)
}


func assert(t testing.TB, got string, want string){
    if got != want {
        t.Errorf("expected %s but got %s",want, got)
    }

    
}
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin){
    t.Helper()

    got := wallet.Balance()
    if got != want {
        t.Errorf("expected %s but got %s",want, got)
    }
}

func assertError(t testing.TB, got error, want error){
    t.Helper()

    if got == nil {
        t.Fatal("didn't get an error but wanted one")
    }


    if got != want {
        t.Errorf("expected %q but got %q",want, got)
    }
}



func assertNoError(t testing.TB, got error){
    t.Helper()

    if got != nil {
        t.Fatal("shouldn't return error but got error")
    }

}
