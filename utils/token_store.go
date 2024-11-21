package utils

import (
	"sync"
	"time"
)

var (
	tokenMutex    sync.Mutex
	revokedTokens = make(map[string]time.Time)
)

// AddTokenToRevokedList adds a token and its expiration time to the revoked list
func AddTokenToRevokedList(token string, exp time.Time) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()

	// Log the addition of a token to the revoked list
	log.Info("Adding token to revoked list", "token", token, "expiration", exp)

	revokedTokens[token] = exp
}

// IsTokenRevoked checks if a token is revoked
func IsTokenRevoked(token string) bool {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()

	exp, exists := revokedTokens[token]
	if !exists {
		// Log when the token is not found in the revoked list
		log.Info("Token not revoked", "token", token)
		return false
	}

	// Check if the token has expired
	if time.Now().After(exp) {
		// Log cleanup of expired token
		log.Info("Cleaning up expired token", "token", token, "expiration", exp)
		delete(revokedTokens, token)
		return false
	}

	// Log when the token is still revoked
	log.Info("Token is revoked", "token", token, "expiration", exp)
	return true
}
