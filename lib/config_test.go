package nudger_test

# borrowed from github.com/codegangsta/gin

import (
	"github.com/sgrimee/nudger/lib"
	"testing"
)

func Test_LoadConfig(t *testing.T) {
	config, err := nudger.LoadConfig("test_fixtures/config.json")

	expect(t, err, nil)
	expect(t, config.ItemsDir, "~/Downloads")
	expect(t, config.NudgeCmd, "ls -lah")
}

func Test_LoadConfig_WithNonExistantFile(t *testing.T) {
	_, err := nudger.LoadConfig("im/not/here.json")

	refute(t, err, nil)
	expect(t, err.Error(), "Unable to read configuration file im/not/here.json")
}

func Test_LoadConfig_WithMalformedFile(t *testing.T) {
	_, err := nudger.LoadConfig("test_fixtures/bad_config.json")

	refute(t, err, nil)
	expect(t, err.Error(), "Unable to parse configuration file test_fixtures/bad_config.json")
}
