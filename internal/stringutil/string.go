package stringutil

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

const randCharMap = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890"

var once sync.Once

// Generate the random seed at program start.
func init() {
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})
}

// HasAnyPrefix checks if any of the given `prefixes` are in the `s` string.
func HasAnyPrefix(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// HasAnyiPrefix checks if any of the given `prefixes` are in the `s` string.
// Case insensitive.
func HasAnyiPrefix(s string, prefixes ...string) bool {
	s = strings.ToLower(s)
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}

// EqualsIAny checks if the string `s` is equal to any of the strings `strs`.
// Case insensitive and trims the strings.
func EqualsIAny(s string, strs ...string) bool {
	s = strings.TrimSpace(strings.ToLower(s))
	for _, str := range strs {
		if s == strings.TrimSpace(strings.ToLower(str)) {
			return true
		}
	}
	return false
}

// Empty returns true if the string is empty or nil.
func Empty(s *string) bool {
	if s == nil {
		return true
	}
	if *s == "" {
		return true
	}
	return false
}

// RandString generates a random string of the length `n` with characters A-Z, a-z, and 0-9.
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = randCharMap[rand.Intn(len(randCharMap))]
	}
	rand.Int63()
	return string(b)
}

// AnyFunc checks if the `s` string is in any of the `strs` slices that satisfy the given `f` func
func AnyFunc(s string, strs []string, f func(s, substr string) bool) bool {
	for idx := range strs {
		// blank strings
		if s == "" && strs[idx] == "" {
			return true
		}
		if s != "" && f(strs[idx], s) {
			return true
		}
	}
	return false
}

// MustAtoi converts an `s` string using `strconv.Atoi` but panics if there's an error.
func MustAtoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}
