package sampling

import (
	"fmt"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	invalidConnectionCount = -1
	invalidINode           = "-1"
	socketFilePrefix       = "socket:["
)

var invalidProcessNetworkStat = snapshot.ProcessNetworkStat{
	TCP4Stat: snapshot.ProcessTCPStat{
		ConnectionCount: invalidConnectionCount,
		ConnectionCountByState: map[string]int{
			snapshot.TCPEstablished: invalidConnectionCount,
			snapshot.TCPSynSent:     invalidConnectionCount,
			snapshot.TCPSynRecv:     invalidConnectionCount,
			snapshot.TCPFinWait1:    invalidConnectionCount,
			snapshot.TCPFinWait2:    invalidConnectionCount,
			snapshot.TCPTimeWait:    invalidConnectionCount,
			snapshot.TCPClose:       invalidConnectionCount,
			snapshot.TCPCloseWait:   invalidConnectionCount,
			snapshot.TCPLastACK:     invalidConnectionCount,
			snapshot.TCPListen:      invalidConnectionCount,
			snapshot.TCPClosing:     invalidConnectionCount,
			snapshot.TCPNewSynRecv:  invalidConnectionCount,
		},
	},
	TCP6Stat: snapshot.ProcessTCPStat{
		ConnectionCount: invalidConnectionCount,
		ConnectionCountByState: map[string]int{
			snapshot.TCPEstablished: invalidConnectionCount,
			snapshot.TCPSynSent:     invalidConnectionCount,
			snapshot.TCPSynRecv:     invalidConnectionCount,
			snapshot.TCPFinWait1:    invalidConnectionCount,
			snapshot.TCPFinWait2:    invalidConnectionCount,
			snapshot.TCPTimeWait:    invalidConnectionCount,
			snapshot.TCPClose:       invalidConnectionCount,
			snapshot.TCPCloseWait:   invalidConnectionCount,
			snapshot.TCPLastACK:     invalidConnectionCount,
			snapshot.TCPListen:      invalidConnectionCount,
			snapshot.TCPClosing:     invalidConnectionCount,
			snapshot.TCPNewSynRecv:  invalidConnectionCount,
		},
	},
}

var emptyProcessNetworkStat = snapshot.ProcessNetworkStat{
	TCP4Stat: snapshot.ProcessTCPStat{
		ConnectionCount: 0,
		ConnectionCountByState: map[string]int{
			snapshot.TCPEstablished: 0,
			snapshot.TCPSynSent:     0,
			snapshot.TCPSynRecv:     0,
			snapshot.TCPFinWait1:    0,
			snapshot.TCPFinWait2:    0,
			snapshot.TCPTimeWait:    0,
			snapshot.TCPClose:       0,
			snapshot.TCPCloseWait:   0,
			snapshot.TCPLastACK:     0,
			snapshot.TCPListen:      0,
			snapshot.TCPClosing:     0,
			snapshot.TCPNewSynRecv:  0,
		},
	},
	TCP6Stat: snapshot.ProcessTCPStat{
		ConnectionCount: 0,
		ConnectionCountByState: map[string]int{
			snapshot.TCPEstablished: 0,
			snapshot.TCPSynSent:     0,
			snapshot.TCPSynRecv:     0,
			snapshot.TCPFinWait1:    0,
			snapshot.TCPFinWait2:    0,
			snapshot.TCPTimeWait:    0,
			snapshot.TCPClose:       0,
			snapshot.TCPCloseWait:   0,
			snapshot.TCPLastACK:     0,
			snapshot.TCPListen:      0,
			snapshot.TCPClosing:     0,
			snapshot.TCPNewSynRecv:  0,
		},
	},
}

func getSocketFileINode(targetFile string) (iNode string, err error) {
	iNode, err = invalidINode, nil

	if !strings.HasPrefix(targetFile, socketFilePrefix) { // socket:[30862686]
		err = fmt.Errorf("target file should be started with %s", socketFilePrefix)
	} else {
		fileData := strings.Split(targetFile, ":") // [socket [30862686]]
		iNode = strings.Trim(fileData[1], "[]")    // 30862686
	}

	return iNode, err
}

func sampleProcessNetworkStat(pID int, spshot snapshot.Snapshot,
	networkStatChan chan<- snapshot.ProcessNetworkStat) {
	fdDir := fmt.Sprintf("/proc/%d/fd", pID)
	files, err := ioutil.ReadDir(filepath.Clean(fdDir))

	if err != nil {
		networkStatChan <- invalidProcessNetworkStat
		return
	}

	tcp4Count, tcp6Count, networkStat := 0, 0, emptyProcessNetworkStat

	for _, file := range files {
		filename := fmt.Sprintf("/proc/%d/fd/%s", pID, file.Name())
		targetFilename, err := os.Readlink(filepath.Clean(filename))

		if err == nil {
			if iNode, err := getSocketFileINode(targetFilename); err == nil {
				iNodeToTCP4, iNodeToTCP6 := spshot.Network.INodeToTCP4, spshot.Network.INodeToTCP6

				if tcp4, has := iNodeToTCP4[iNode]; has {
					networkStat.TCP4Stat.ConnectionCountByState[tcp4.State]++
					tcp4Count++
				}

				if tcp6, has := iNodeToTCP6[iNode]; has {
					networkStat.TCP6Stat.ConnectionCountByState[tcp6.State]++
					tcp6Count++
				}
			}
		}
	}

	networkStat.TCP4Stat.ConnectionCount = tcp4Count
	networkStat.TCP6Stat.ConnectionCount = tcp6Count

	networkStatChan <- networkStat
}
