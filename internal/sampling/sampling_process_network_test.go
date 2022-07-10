package sampling

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"os/exec"
	"testing"
)

func TestGetSocketFileINodeWithError(t *testing.T) {
	_, err := getSocketFileINode("invalid file")

	assert.IsNotNil(t, err, "invalid file will have err")

	_, err = getSocketFileINode("'socket[30862686]'")

	assert.IsNotNil(t, err, "no comma file will have err")
}

func TestSampleProcessNetworkStatInvalid(t *testing.T) {
	testStatChan := make(chan snapshot.ProcessNetworkStat)

	go sampleProcessNetworkStat(1000000, snapshot.Snapshot{}, testStatChan)

	stat := <-testStatChan

	assert.IsEqual(t, invalidConnectionCount, stat.TCP4Stat.ConnectionCount,
		"invalid pID should have invalid process network stat")
}

func TestSampleProcessNetworkStat(t *testing.T) {
	cmd := exec.Command("sleep", "54321")
	cmd.Start()

	testPID := cmd.Process.Pid
	testStatChan := make(chan snapshot.ProcessNetworkStat)

	go sampleProcessNetworkStat(testPID, snapshot.Snapshot{}, testStatChan)

	stat := <-testStatChan

	assert.IsTrue(t, invalidConnectionCount != stat.TCP4Stat.ConnectionCount,
		"valid pID should have valid process network stat")

	for _, count := range stat.TCP4Stat.ConnectionCountByState {
		assert.IsTrue(t, invalidConnectionCount != count,
			"valid pID should have valid process network count by state")
	}
}
