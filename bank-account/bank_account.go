package account

import (
    "sync"
    "sync/atomic"
    //"fmt"
)

type Account struct {
    Remainder int64
    IsClose bool
    Mutex *sync.Mutex
}


// Open(initialDeposit int64) *Account
// If Open is given a negative initial deposit, it must return nil.
// If any Account method is called on an closed account, it must not modify
// the account and must return ok = false.

func Open(initialDeposit int64) *Account {
    if initialDeposit < 0 {return nil}
    acc := Account{Remainder: initialDeposit,
                   IsClose: false,
                   Mutex: &sync.Mutex{}}
    return &acc
}

// (Account) Close() (payout int64, ok bool)
func (acc *Account) Close() (payout int64, ok bool) {
   acc.Mutex.Lock()
   if acc.IsClose {
       acc.Mutex.Unlock()
       return 0, false
   }
   acc.IsClose = true
   payout = acc.Remainder
   acc.Mutex.Unlock()
   return payout, true
}
// (Account) Balance() (balance int64, ok bool)
func (acc *Account) Balance() (balance int64, ok bool){
    acc.Mutex.Lock()
    if acc.IsClose {
        acc.Mutex.Unlock()
        return -1, false
    }
    acc.Mutex.Unlock()
    return acc.Remainder, true
}

// (Account) Deposit(amount int64) (newBalance int64, ok bool)
// Deposit must handle a negative amount as a withdrawal. Withdrawals must
// not succeed if they result in a negative balance.
func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool){
    acc.Mutex.Lock()
    if acc.IsClose || acc.Remainder + amount < 0{
        acc.Mutex.Unlock()
        return -1, false
    }
    acc.Mutex.Unlock()

    atomic.AddInt64(&(acc.Remainder), amount)
    return acc.Remainder, true
}

