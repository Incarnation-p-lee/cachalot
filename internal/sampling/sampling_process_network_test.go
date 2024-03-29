package sampling

import (
	"fmt"
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"
	"time"
)

func TestGetSocketFileINodeWithError(t *testing.T) {
	_, err := getSocketFileINode("invalid file")

	assert.IsNotNil(t, err, "invalid file will have err")

	_, err = getSocketFileINode("socket[30862686]")

	assert.IsNotNil(t, err, "no comma file will have err")
}

func TestGetSocketFileINode(t *testing.T) {
	iNode, err := getSocketFileINode("socket:[1243]")

	assert.IsNil(t, err, "valid file will have no err")
	assert.IsEqual(t, "1243", iNode, "valid file will have the same inode")
}

func TestSampleProcessNetworkStatInvalid(t *testing.T) {
	testStatChan := make(chan snapshot.NetworkStat)

	go sampleProcessNetworkStat(1000000, snapshot.Snapshot{}, testStatChan)

	stat := <-testStatChan

	assert.IsEqual(t, invalidConnectionCount, stat.TCP4Stat.ConnectionCount,
		"invalid pID should have invalid process network stat")
}

func getProcessFirstSocketFileINode(pID int) string {
	fdDir := fmt.Sprintf("/proc/%d/fd", pID)
	files, _ := ioutil.ReadDir(filepath.Clean(fdDir))
	socketINode := invalidINode

	for _, file := range files {
		filename := fmt.Sprintf("/proc/%d/fd/%s", pID, file.Name())
		targetFilename, _ := os.Readlink(filepath.Clean(filename))

		if iNode, err := getSocketFileINode(targetFilename); err == nil {
			socketINode = iNode
			break
		}
	}

	return socketINode
}

func getTestSnapShot(testINode string) snapshot.Snapshot {
	return snapshot.Snapshot{
		Network: snapshot.Network{
			INodeToTCP4: map[string]snapshot.TCPConnection{
				testINode: snapshot.TCPConnection{
					INode: testINode,
					State: "Established",
				},
			},
			INodeToTCP6: map[string]snapshot.TCPConnection{
				testINode: snapshot.TCPConnection{
					INode: testINode,
					State: "Established",
				},
			},
		},
	}
}

func TestSampleProcessNetworkStat(t *testing.T) {
	cmd := exec.Command("python3", "-m", "http.server", "9843")
	cmd.Start()

	time.Sleep(time.Duration(2) * time.Second)

	testPID := cmd.Process.Pid
	testStatChan := make(chan snapshot.NetworkStat)
	testINode := getProcessFirstSocketFileINode(testPID)
	testSnapshot := getTestSnapShot(testINode)

	go sampleProcessNetworkStat(testPID, testSnapshot, testStatChan)
	stat := <-testStatChan

	assert.IsTrue(t, invalidConnectionCount != stat.TCP4Stat.ConnectionCount,
		"valid pID should have valid process network stat tcp4")

	assert.IsTrue(t, invalidConnectionCount != stat.TCP6Stat.ConnectionCount,
		"valid pID should have valid process network stat tcp6")

	for _, count := range stat.TCP4Stat.ConnectionCountByState {
		assert.IsTrue(t, invalidConnectionCount != count,
			"valid pID should have valid process network count by state for tcp4")
	}

	for _, count := range stat.TCP6Stat.ConnectionCountByState {
		assert.IsTrue(t, invalidConnectionCount != count,
			"valid pID should have valid process network count by state for tcp6")
	}

	cmd.Process.Signal(syscall.SIGTERM)
}
