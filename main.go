package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type FocusSession struct {
	StartedWhen time.Time
	Duration time.Duration
	CurrentlyActive bool
	ID string
}



func NewFocusSession(duration time.Duration) *FocusSession {
    return &FocusSession{
        ID: uuid.NewString(),
        StartedWhen: time.Now(),
        Duration:  duration,
        CurrentlyActive:  true,
    }
}

func (fs *FocusSession) IsStillActive() bool {
    // First check: Is it marked as active?
    if !fs.CurrentlyActive {
        return false
    }
    
    // Second check: Has the time run out?
    endTime := fs.StartedWhen.Add(fs.Duration)
    return time.Now().Before(endTime)
}

func main() {
	// Create a new 30 seconds focus session
	session := NewFocusSession(30 * time.Second)

	fmt.Println("==starting a session==")
	fmt.Printf("Session ID: %s\n", session.ID)
	fmt.Printf("Started: %s\n", session.StartedWhen.Format(time.RFC1123))
	fmt.Printf("Duration: %s\n", session.Duration)
	fmt.Printf("Active: %t\n", session.CurrentlyActive)

	//checking session activity

	fmt.Println("\n=== Testing right after creation ===")
	fmt.Printf("Is still active: %t\n", session.IsStillActive())

	fmt.Println("\n=== Waiting 10 seconds ===")
	time.Sleep(10 * time.Second)
	fmt.Printf("Is still active: %t\n", session.IsStillActive())

	fmt.Println("\n=== Waiting another 8 seconds (just to see) ===")
	time.Sleep(8 * time.Second)
	fmt.Printf("Is still active: %t\n", session.IsStillActive())

	fmt.Println("\n=== Waiting another 20 seconds (should expire) ===")
	time.Sleep(20 * time.Second)
	fmt.Printf("Is still active: %t\n", session.IsStillActive())

	fmt.Println("\n=== Test complete ===")
}