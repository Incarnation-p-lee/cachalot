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

	invalidAddress = "invalid address"
	invalidPort    = -1

	tcpCodeDefaultIndex = 0
)

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

func sampleNetwork() snapshot.Network {
	return snapshot.Network{
		INodeToTCP4: sampleTCP4Connection(),
		INodeToTCP6: sampleTCP6Connection(),
	}
}
