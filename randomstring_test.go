package randomstring

import (
	"fmt"
	"testing"
)

// Uncomment this function to see that rand.NewSource(1) is the same as the unseeded behavior before Go 1.20
func init() {
	Seed()
}

func TestString(t *testing.T) {
	fmt.Printf("%x\n", String(10))
}

func TestStringNoAlloc(t *testing.T) {
	b := make([]byte, 10)
	StringNoAlloc(b)
	fmt.Printf("%x\n", string(b))
}

func TestHumanFriendlyString(t *testing.T) {
	fmt.Printf("%s\n", HumanFriendlyString(7))
	fmt.Printf("%s\n", HumanFriendlyString(20))
}

func TestCookieFriendlyString(t *testing.T) {
	fmt.Printf("%s\n", CookieFriendlyString(20))
}

func TestCookieFriendlyStringNoAlloc(t *testing.T) {
	b := make([]byte, 20)
	CookieFriendlyStringNoAlloc(b)
	fmt.Printf("%s\n", string(b))
}

func TestCookieFriendlyBytes(t *testing.T) {
	fmt.Printf("%s\n", CookieFriendlyBytes(20))
}

func TestEnglishFrequencyString(t *testing.T) {
	fmt.Printf("%s\n", EnglishFrequencyString(20))
}

func TestHumanFriendlyEnglishString(t *testing.T) {
	fmt.Printf("%s\n", HumanFriendlyEnglishString(7))
	fmt.Printf("%s\n", HumanFriendlyEnglishString(20))
}
