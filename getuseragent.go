package psychoclient

import (
	"math/rand"
	"time"
)

// GetChromeUserAgent : Get Latest Chrome Useragent
func GetChromeUserAgent() (useragent string) {
	chromeUseragents := GetChromeUserAgentList()
	rand.Seed(time.Now().UnixNano())
	return chromeUseragents[rand.Intn(len(chromeUseragents))]
}

// GetTestUserAgent : Get Latest Test Useragent
func GetTestUserAgent() (useragent string) {
	testUserAgents := GetTestUserAgentList()
	rand.Seed(time.Now().UnixNano())
	return testUserAgents[rand.Intn(len(testUserAgents))]
}
