package account

import "sync"

type Account struct {
	balance  int64
	isClosed bool
	mux      sync.Mutex
}

func Open(amt int64) *Account {
	if amt < 0 {
		return nil
	}
	return &Account{balance: amt, isClosed: false}
}

func (act *Account) Balance() (int64, bool) {
	act.mux.Lock()
	defer act.mux.Unlock()

	return act.balance, !act.isClosed
}

func (act *Account) Close() (int64, bool) {
	act.mux.Lock()
	defer act.mux.Unlock()

	if act.isClosed {
		return 0, false
	}

	act.isClosed = true
	payout := act.balance
	act.balance = 0

	return payout, true
}

func (act *Account) Deposit(amt int64) (int64, bool) {
	act.mux.Lock()
	defer act.mux.Unlock()
	if act.isClosed {
		return 0, false
	}
	if amt < 0 && amt+act.balance < 0 {
		return 0, false
	}

	act.balance += amt

	return act.balance, true
}
