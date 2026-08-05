package main

import "fmt"

func main() {

	// Log Counter

	logCounts := make(map[string]int)

	logs := []string{"INFO", "ERROR", "WARN", "INFO", "ERROR", "INFO", "DEBUG"}

	for _, log := range logs {
		logCounts[log]++
	}

	for name, count := range logCounts {
		fmt.Printf("%s appears %d times\n", name, count)
	}

	// Session Cache and Lookup

	sessions := map[string]string{"user1": "token_abc", "user2": "token_bcd"}

	checkSession(sessions, "user1")
	
	delete(sessions, "user1")

	checkSession(sessions, "user1")

	checkSession(sessions, "user4")
}


func checkSession(cache map[string]string, userID string) {
	token, ok := cache[userID]

	if !ok {
		fmt.Println("[MISSING] No active session\n")
	} else {
		fmt.Printf("[FOUND] User: %s with Token: %s\n", userID, token)
	}
}
