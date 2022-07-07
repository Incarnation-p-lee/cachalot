package sampling

import (
	"os"
	"log"
	"internal/utils"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"strings"
}

const (
	tcp4ConnectionFile = "/proc/net/tcp"
	tcp4ConnectionTitlePrefix = "sl"

	tcp4INodeIndex = 9
	tcp4UIDIndex = 7
	tcp4StatusIndex = 3
	tcp4RemoteAddressIndex = 2
)

func isTCPConnectionData(line string) bool {
	line = strings.Trim(line, " ")

	return !strings.HasPrefix(line, tcp4ConnectionTitlePrefix)
}

func getTCP4AddressAndPort(addressAndPort string) (address string, port int) {
}

func getTCP4Status(statusCode string) string {

}

func getTCP4Connection(tcp4Data []string) snapshot.TCP4Connection {
	address, port := getTCP4AddressAndPort(tcp4Data[tcp4RemoteAddressIndex])

	return snapshot.TCP4Connection{
		RemoteAddress: address,
		RemotePort: port,
		INode: tcp4Data[tcp4INodeIndex],
		UID: tcp4Data[tcp4UIDIndex],
		Status: getTCP4Status(tcp4Data[tcp4StatusIndex]),
	}
}

func sampleTCPConnection() map[string]snapshot.TCP4Connection {
	nodeToTCP4 := map[string]snapshot.TCP4Connection {}
	file, err := os.Open(filepath.Clean(tcp4ConnectionFile))

	if err != nil {
		log.Printf("Failed to open file %s due to %+v\n", filename, err)
		return nodeToTCP4
	}

	defer utils.CloseFile(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if isTCPConnectionData(line) {
			tcp4Data := strings.Split(strings.Trim(line, " "), " ")
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

