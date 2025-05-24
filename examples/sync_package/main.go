package main

// The sync package provides several benefits for concurrent programming in Go:
// 1. Thread Safety: Mutexes and RWMutexes prevent data races and ensure safe
//    concurrent access to shared resources.
// 2. Coordination: WaitGroups help coordinate multiple goroutines by waiting
//    for a collection of goroutines to finish.
// 3. One-time Initialization: sync.Once ensures a function is executed exactly
//    once, even in concurrent scenarios.
// 4. Atomic Operations: sync/atomic provides atomic operations for primitive
//    types without explicit locking.
// 5. Concurrent Maps: sync.Map provides a thread-safe map implementation
//    optimized for cases where items are written once but read many times.
// 6. Memory Synchronization: The sync primitives ensure proper memory ordering
//    and visibility between goroutines.
// Without proper synchronization using the sync package, concurrent programs
// can suffer from race conditions, data corruption, and unpredictable behavior.

import (
	"fmt"
	"sync"
	"time"
)

// BankAccount represents a simple bank account with mutex protection
type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

// Deposit adds money to the account
func (b *BankAccount) Deposit(amount int) {
	b.mutex.Lock()         // Lock before modifying
	defer b.mutex.Unlock() // Unlock when done

	// Simulate some processing time
	time.Sleep(100 * time.Millisecond)
	b.balance += amount
	fmt.Printf("Deposited %d, new balance: %d\n", amount, b.balance)
}

// Withdraw removes money from the account
func (b *BankAccount) Withdraw(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.balance >= amount {
		time.Sleep(100 * time.Millisecond)
		b.balance -= amount
		fmt.Printf("Withdrawn %d, new balance: %d\n", amount, b.balance)
	}
}

func main() {

	fmt.Println("Original balance: 1000")
	account := &BankAccount{balance: 1000}

	// Start multiple goroutines to modify the account
	for i := 0; i < 5; i++ {
		go func() {
			account.Deposit(100)
			account.Withdraw(50)
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("Final balance: %d\n", account.balance)
}
