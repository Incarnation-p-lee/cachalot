package sampling

import (
	"fmt"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"io/ioutil"
	"log"
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
			snapshot.TcpEstablished: invalidConnectionCount,
			snapshot.TcpSynSent:     invalidConnectionCount,
			snapshot.TcpSynRecv:     invalidConnectionCount,
			snapshot.TcpFinWait1:    invalidConnectionCount,
			snapshot.TcpFinWait2:    invalidConnectionCount,
			snapshot.TcpTimeWait:    invalidConnectionCount,
			snapshot.TcpClose:       invalidConnectionCount,
			snapshot.TcpCloseWait:   invalidConnectionCount,
			snapshot.TcpLastACK:     invalidConnectionCount,
			snapshot.TcpListen:      invalidConnectionCount,
			snapshot.TcpClosing:     invalidConnectionCount,
			snapshot.TcpNewSynRecv:  invalidConnectionCount,
		},
	},
	TCP6Stat: snapshot.ProcessTCPStat{
		ConnectionCount: invalidConnectionCount,
		ConnectionCountByState: map[string]int{
			snapshot.TcpEstablished: invalidConnectionCount,
			snapshot.TcpSynSent:     invalidConnectionCount,
			snapshot.TcpSynRecv:     invalidConnectionCount,
			snapshot.TcpFinWait1:    invalidConnectionCount,
			snapshot.TcpFinWait2:    invalidConnectionCount,
			snapshot.TcpTimeWait:    invalidConnectionCount,
			snapshot.TcpClose:       invalidConnectionCount,
			snapshot.TcpCloseWait:   invalidConnectionCount,
			snapshot.TcpLastACK:     invalidConnectionCount,
			snapshot.TcpListen:      invalidConnectionCount,
			snapshot.TcpClosing:     invalidConnectionCount,
			snapshot.TcpNewSynRecv:  invalidConnectionCount,
		},
	},
}

var emptyProcessNetworkStat = snapshot.ProcessNetworkStat{
	TCP4Stat: snapshot.ProcessTCPStat{
		ConnectionCount: 0,
		ConnectionCountByState: map[string]int{
			snapshot.TcpEstablished: 0,
			snapshot.TcpSynSent:     0,
			snapshot.TcpSynRecv:     0,
			snapshot.TcpFinWait1:    0,
			snapshot.TcpFinWait2:    0,
			snapshot.TcpTimeWait:    0,
			snapshot.TcpClose:       0,
			snapshot.TcpCloseWait:   0,
			snapshot.TcpLastACK:     0,
			snapshot.TcpListen:      0,
			snapshot.TcpClosing:     0,
			snapshot.TcpNewSynRecv:  0,
		},
	},
	TCP6Stat: snapshot.ProcessTCPStat{
		ConnectionCount: 0,
		ConnectionCountByState: map[string]int{
			snapshot.TcpEstablished: 0,
			snapshot.TcpSynSent:     0,
			snapshot.TcpSynRecv:     0,
			snapshot.TcpFinWait1:    0,
			snapshot.TcpFinWait2:    0,
			snapshot.TcpTimeWait:    0,
			snapshot.TcpClose:       0,
			snapshot.TcpCloseWait:   0,
			snapshot.TcpLastACK:     0,
			snapshot.TcpListen:      0,
			snapshot.TcpClosing:     0,
			snapshot.TcpNewSynRecv:  0,
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
		log.Printf("Failed to read dir from %s due to %+v\n", fdDir, err)
		networkStatChan <- invalidProcessNetworkStat
		return
	}

	tcp4Count, tcp6Count, networkStat := 0, 0, emptyProcessNetworkStat

	for _, file := range files {
		filename := fmt.Sprintf("/proc/%d/fd/%s", pID, file.Name())
		targetFilename, err := os.Readlink(filepath.Clean(filename))

		if err != nil {
			log.Printf("Failed to read link from file %s due to %+v\n", filename, err)
		} else if iNode, err := getSocketFileINode(targetFilename); err == nil {
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

	networkStat.TCP4Stat.ConnectionCount = tcp4Count
	networkStat.TCP6Stat.ConnectionCount = tcp6Count

	networkStatChan <- networkStat
}
