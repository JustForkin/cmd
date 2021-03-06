package cmd

import (
	"os"
	"runtime"
	"testing"
)

func TestIsSystemdService(t *testing.T) {
	t.Parallel()

	if runtime.GOOS != "linux" {
		if IsSystemdService() {
			t.Error(`IsSystemdService()`)
		}
		return
	}

	if IsSystemdService() {
		t.Error(`IsSystemdService()`)
	}

	err := os.Setenv("JOURNAL_STREAM", "10:20")
	if err != nil {
		t.Fatal(err)
	}

	if !IsSystemdService() {
		t.Error(`!IsSystemdService()`)
	}
}
