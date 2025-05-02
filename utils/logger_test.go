package utils

import "testing"

// TestHasToLogDebug calls hasToLog
// checking if its need to be logged
func TestHasToLogDebug(t *testing.T) {
	Flags.Debug = true

	if !hasToLog("info") {
		t.Errorf("Should log")
	}
	if !hasToLog("debug") {
		t.Errorf("Should log")
	}
	if !hasToLog("warn") {
		t.Errorf("Should log")
	}
	if !hasToLog("error") {
		t.Errorf("Should log")
	}
}

// TestHasToLogNotDebug calls hasToLog
// checking if its need to be logged
func TestHasToLogNotDebug(t *testing.T) {
	Flags.Debug = false

	if !hasToLog("info") {
		t.Errorf("Should not log")
	}
	if hasToLog("debug") {
		t.Errorf("Should log")
	}
	if !hasToLog("warn") {
		t.Errorf("Should log")
	}
	if !hasToLog("error") {
		t.Errorf("Should log")
	}

}
