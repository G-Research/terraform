package diff

import (
	"log"
	"os"
	"strings"
)

type DiffLevel string

// EnvDiff is the environment variable that determines the level of diff details.
const EnvDiff = "TF_DIFF"

// These are the supported diff levels.
const (
	AllLevel        = DiffLevel("ALL")
	HideValuesLevel = DiffLevel("HIDE_VALUES")
	RootLevel       = DiffLevel("ROOT")
)

// ValidLevels is the list of supported diff levels.
var ValidLevels = []DiffLevel{AllLevel, HideValuesLevel, RootLevel}

var level DiffLevel

// CurrentDiffLevel returns the current diff level based on the environment var.
func CurrentDiffLevel() DiffLevel {
	if level == "" {
		envLevel := os.Getenv(EnvDiff)
		level = AllLevel

		if isValidDiffLevel(envLevel) {
			level = DiffLevel(strings.ToUpper(envLevel))
		} else if envLevel != "" {
			log.Printf("[WARN] Invalid diff level: %q. Defaulting to level: %s. Valid levels are: %+v",
				envLevel, level, ValidLevels)
		}
	}

	return level
}

func isValidDiffLevel(level string) bool {
	for _, l := range ValidLevels {
		if strings.ToUpper(level) == string(l) {
			return true
		}
	}

	return false
}
