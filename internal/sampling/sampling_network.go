package sampling

import (
	"bufio"
	"internal/utils"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"log"
	"path/filepath"
	"os"
	"strings"
	"strconv"
)

const (
	tcp4ConnectionFile = "/proc/net/tcp"
	tcp4ConnectionTitlePrefix = "sl"

	tcp4INodeIndex = 9
	tcp4UIDIndex = 7
	tcp4StatusIndex = 3
	tcp4RemoteAddressIndex = 2

	invalidAddress = "invalid address"
	invalidPort = -1

	tcpUnknown = "Unknown"
	tcpEstablished = "Established"
	tcpSynSent = "SynSent"
	tcpSynRecv = "SynRecv"
	tcpFinWait1 = "FinWait1"
	tcpFinWait2 = "FinWait2"
	tcpTimeWait = "TimeWait"
	tcpClose = "Close"
	tcpCloseWait = "CloseWait"
	tcpLastACK = "LastACK"
	tcpListen = "Listen"
	tcpClosing = "Closing"
	tcpNewSynRecv = "NewSynRecv"

	tcpCodeDefaultIndex = 0
)

var tcpCodes = []string {
	tcpUnknown,
	tcpEstablished,
	tcpSynSent,
	tcpSynRecv,
	tcpFinWait1,
	tcpFinWait2,
	tcpTimeWait,
	tcpClose,
	tcpCloseWait,
	tcpLastACK,
	tcpListen,
	tcpClosing,
	tcpNewSynRecv,
}

func isTCPConnectionData(line string) bool {
	line = strings.Trim(line, " ")

	return !strings.HasPrefix(line, tcp4ConnectionTitlePrefix)
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
	index := tcpCodeDefaultIndex
	val, err := strconv.ParseInt(code, 16, 32)

	if err != nil {
		log.Printf("Failed to convert %s to integer due to %+v\n", code, err)
	} else if int(val) < len(tcpCodes) {
		index = int(val)
	}

	return tcpCodes[index]
}

func getTCP4Connection(tcp4Data []string) snapshot.TCP4Connection {
	address, port := getTCP4AddressAndPort(tcp4Data[tcp4RemoteAddressIndex])

	return snapshot.TCP4Connection{
		RemoteAddress: address,
		RemotePort: port,
		INode: tcp4Data[tcp4INodeIndex],
		UID: tcp4Data[tcp4UIDIndex],
		State: getTCPState(tcp4Data[tcp4StatusIndex]),
	}
}

func getTCP4ConnectionData(line string) []string {
	tcp4Data := []string{}

	for _, val := range strings.Split(line, " ") {
		trimVal := strings.Trim(val, " ")

		if len(trimVal) > 0 {
			tcp4Data = append(tcp4Data, trimVal)
		}
	}

	return tcp4Data
}

func sampleTCPConnection() map[string]snapshot.TCP4Connection {
	nodeToTCP4 := map[string]snapshot.TCP4Connection {}
	file, err := os.Open(filepath.Clean(tcp4ConnectionFile))

	if err != nil {
		log.Printf("Failed to open file %s due to %+v\n", tcp4ConnectionFile, err)
		return nodeToTCP4
	}

	defer utils.CloseFile(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if isTCPConnectionData(line) {
			tcp4Data := getTCP4ConnectionData(line)
			tcp4Connection := getTCP4Connection(tcp4Data)
			nodeToTCP4[tcp4Connection.INode] = tcp4Connection
		}
	}

	return nodeToTCP4
}

func sampleNetwork() snapshot.Network {
	return snapshot.Network{
		INodeToTCP4: sampleTCPConnection(),
	}
}

