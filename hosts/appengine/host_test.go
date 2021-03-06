package gaehost

import (
	"testing"
)

const noPanic = "The code did not panic"

func TestGaeBotHost_GetHTTPClient_nilRequest(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(noPanic)
		}
	}()

	GaeBotHost{}.GetHTTPClient(nil)
}

func TestGaeBotHost_GetBotCoreStores_nil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(noPanic)
		}
	}()

	GaeBotHost{}.GetBotCoreStores("", nil, nil)
}
