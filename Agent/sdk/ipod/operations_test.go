package ipod

import (
	"log/slog"
	"testing"

	"github.com/Synchro/common/log"
	"github.com/stretchr/testify/assert"
)

func TestThatIPodIsInRange(t *testing.T) {
	ipod := New(log.New(slog.LevelInfo), nil)
	inRange := ipod.IsInNetwork()
	assert.True(t, inRange, "ipod should be in range")
}

func TestThatIpodFileTransferWorks(t *testing.T) {
	ipod := New(log.New(slog.LevelInfo), nil)
	inRange := ipod.IsInNetwork()
	assert.True(t, inRange, "ipod should be in range")

	fm := make(map[string][]string)
	fm["dummyplaylist"] = []string{"/home/tars/dummy.txt"}

	_, err := ipod.SendFiles(fm)
	assert.Nil(t, err, "error should be nil")

}
