package sampling

import (
	"bufio"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"internal/utils"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	tcp4ConnectionFile = "/proc/net/tcp"
	tcp6ConnectionFile = "/proc/net/tcp6"

	tcpConnectionTitlePrefix = "sl"

	tcpINodeIndex         = 9
	tcpUIDIndex           = 7
	tcpStatusIndex        = 3
	tcpRemoteAddressIndex = 2

	invalidAddress         = "invalid address"
	invalidPort            = -1
	invalidConnectionCount = -1

	tcpCodeDefaultIndex = 0
)

func getInvalidTCPConnectionCountByState() map[string]int {
	return map[string]int{
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
	}
}

func getInvalidTCPStat() snapshot.TCPStat {
	return snapshot.TCPStat{
		ConnectionCount:        invalidConnectionCount,
		ConnectionCountByState: getInvalidTCPConnectionCountByState(),
	}
}

func getInvalidNetworkStat() snapshot.NetworkStat {
	return snapshot.NetworkStat{
		TCP4Stat: getInvalidTCPStat(),
		TCP6Stat: getInvalidTCPStat(),
	}
}

func getEmptyTCPConnectionCountByState() map[string]int {
	return map[string]int{
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
	}
}

func getEmptyTCPStat() snapshot.TCPStat {
	return snapshot.TCPStat{
		ConnectionCount:        0,
		ConnectionCountByState: getEmptyTCPConnectionCountByState(),
	}
}

func getEmptyNetworkStat() snapshot.NetworkStat {
	return snapshot.NetworkStat{
		TCP4Stat: getEmptyTCPStat(),
		TCP6Stat: getEmptyTCPStat(),
	}
}

func isTCPConnectionData(line string) bool {
	line = strings.Trim(line, " ")

	return !strings.HasPrefix(line, tcpConnectionTitlePrefix)
}

func getTCP4AddressAndPort(addressAndPort string) (address string, port int) {
	address, port = invalidAddress, invalidPort
	data := strings.Split(addressAndPort, ":")

	if len(data) == 2 {
		address = data[0]

		if val, err := strconv.ParseInt(data[1], 16, 32); err != nil {
			log.Printf("Failed to convert %s to integer due to %+v\n", data[1], err)
		} else {
			port = int(val)
		}
	}

	return address, port
}

func getTCPState(code string) string {
	val, err := strconv.ParseInt(code, 16, 32)
	index, states := tcpCodeDefaultIndex, snapshot.GetTCPStates()

	if err != nil {
		log.Printf("Failed to convert %s to integer due to %+v\n", code, err)
	} else if int(val) < len(states) {
		index = int(val)
	}

	return states[index]
}

func getTCPConnection(tcpData []string) snapshot.TCPConnection {
	address, port := getTCP4AddressAndPort(tcpData[tcpRemoteAddressIndex])

	return snapshot.TCPConnection{
		RemoteAddress: address,
		RemotePort:    port,
		INode:         tcpData[tcpINodeIndex],
		UID:           tcpData[tcpUIDIndex],
		State:         getTCPState(tcpData[tcpStatusIndex]),
	}
}

func getTCPConnectionData(line string) []string {
	tcpData := []string{}

	for _, val := range strings.Split(line, " ") {
		trimVal := strings.Trim(val, " ")

		if len(trimVal) > 0 {
			tcpData = append(tcpData, trimVal)
		}
	}

	return tcpData
}

func sampleTCP4Connection() map[string]snapshot.TCPConnection {
	return sampleTCPConnection(tcp4ConnectionFile)
}

func sampleTCP6Connection() map[string]snapshot.TCPConnection {
	return sampleTCPConnection(tcp6ConnectionFile)
}

func sampleTCPConnection(connectionFile string) map[string]snapshot.TCPConnection {
	nodeToTCP := map[string]snapshot.TCPConnection{}
	file, err := os.Open(filepath.Clean(connectionFile))

	if err != nil {
		log.Printf("Failed to open file %s due to %+v\n", connectionFile, err)
		return nodeToTCP
	}

	defer utils.CloseFile(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if isTCPConnectionData(line) {
			tcpData := getTCPConnectionData(line)
			tcpConnection := getTCPConnection(tcpData)
			nodeToTCP[tcpConnection.INode] = tcpConnection
		}
	}

	return nodeToTCP
}

func calculateTCP4Stat(inodeToTCP4 map[string]snapshot.TCPConnection) snapshot.TCPStat {
	tcp4Stat, tcp4Count := getEmptyTCPStat(), 0

	for _, tcp4 := range inodeToTCP4 {
		tcp4Stat.ConnectionCountByState[tcp4.State]++
		tcp4Count++
	}

	tcp4Stat.ConnectionCount = tcp4Count

	return tcp4Stat
}

func calculateTCP6Stat(inodeToTCP6 map[string]snapshot.TCPConnection) snapshot.TCPStat {
	tcp6Stat, tcp6Count := getEmptyTCPStat(), 0

	for _, tcp6 := range inodeToTCP6 {
		tcp6Stat.ConnectionCountByState[tcp6.State]++
		tcp6Count++
	}

	tcp6Stat.ConnectionCount = tcp6Count

	return tcp6Stat
}

func sampleNetwork() snapshot.Network {
	inodeToTCP4, inodeToTCP6 := sampleTCP4Connection(), sampleTCP6Connection()
	networkStat := getEmptyNetworkStat()

	networkStat.TCP4Stat = calculateTCP4Stat(inodeToTCP4)
	networkStat.TCP6Stat = calculateTCP6Stat(inodeToTCP6)

	return snapshot.Network{
		INodeToTCP4: inodeToTCP4,
		INodeToTCP6: inodeToTCP6,
		NetworkStat: networkStat,
	}
}
