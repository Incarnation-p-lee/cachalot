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
	socketFilePrefix       = "'socket:["
)

var invalidProcessNetworkStat = snapshot.ProcessNetworkStat{
	TCP4Stat: snapshot.ProcessTCP4Stat{
		ConnectionCount: invalidConnectionCount,
		ConnectionCountByState: map[string]int{
			tcpEstablished: invalidConnectionCount,
			tcpSynSent:     invalidConnectionCount,
			tcpSynRecv:     invalidConnectionCount,
			tcpFinWait1:    invalidConnectionCount,
			tcpFinWait2:    invalidConnectionCount,
			tcpTimeWait:    invalidConnectionCount,
			tcpClose:       invalidConnectionCount,
			tcpCloseWait:   invalidConnectionCount,
			tcpLastACK:     invalidConnectionCount,
			tcpListen:      invalidConnectionCount,
			tcpClosing:     invalidConnectionCount,
			tcpNewSynRecv:  invalidConnectionCount,
		},
	},
}

var emptyProcessNetworkStat = snapshot.ProcessNetworkStat{
	TCP4Stat: snapshot.ProcessTCP4Stat{
		ConnectionCount: 0,
		ConnectionCountByState: map[string]int{
			tcpEstablished: 0,
			tcpSynSent:     0,
			tcpSynRecv:     0,
			tcpFinWait1:    0,
			tcpFinWait2:    0,
			tcpTimeWait:    0,
			tcpClose:       0,
			tcpCloseWait:   0,
			tcpLastACK:     0,
			tcpListen:      0,
			tcpClosing:     0,
			tcpNewSynRecv:  0,
		},
	},
}

func getSocketFileINode(targetFile string) (iNode string, err error) {
	iNode, err = invalidINode, nil

	if !strings.HasPrefix(targetFile, socketFilePrefix) { // 'socket:[30862686]'
		err = fmt.Errorf("target file should be started with %s", socketFilePrefix)
	} else {
		targetFile = strings.Trim(targetFile, "'")        // socket:[30862686]
		fileData := strings.Split(targetFile, ":")        // [socket [30862686]]
		iNode = strings.Trim(fileData[1], "[]")           // 30862686
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

	count, networkStat := 0, emptyProcessNetworkStat

	for _, file := range files {
		filename := file.Name()
		targetFilename, err := os.Readlink(filename)

		if err != nil {
			log.Printf("Failed to read link from file %s due to %+v\n", filename, err)
		} else if iNode, err := getSocketFileINode(targetFilename); err == nil {
			iNodeToTCP4 := spshot.Network.INodeToTCP4

			if tcp4, has := iNodeToTCP4[iNode]; has {
				networkStat.TCP4Stat.ConnectionCountByState[tcp4.State]++
				count++
			}
		}
	}

	networkStat.TCP4Stat.ConnectionCount = count
	networkStatChan <- networkStat
}
