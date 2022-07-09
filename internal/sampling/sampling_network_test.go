package sampling

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"testing"
)

func TestSampleNetwork(t *testing.T) {
	network := sampleNetwork()

	assert.IsNotNil(t, network.INodeToTCP4, "inode to tcp4 maps should not be nil")
	assert.IsTrue(t, len(network.INodeToTCP4) > 0, "inode to tcp4 maps should not be empty")

	for _, tcpConnect := range network.INodeToTCP4 {
		assert.IsTrue(t, len(tcpConnect.INode) > 0, "inode of tcp connection cannot be empty")
	}
}

func TestGetTCPStateUnknown(t *testing.T) {
	assert.IsEqual(t, tcpUnknown, getTCPState("ie"), "invalid hex will have unknown state")
	assert.IsEqual(t, tcpUnknown, getTCPState("a0"), "out of range hex will have unknown state")
}

func TestGetTCP4AddressAndPortInvalid(t *testing.T) {
	address, port := getTCP4AddressAndPort("0.0.0.0:123:345")

	assert.IsEqual(t, invalidAddress, address, "invalid format will have invalid address")
	assert.IsEqual(t, invalidPort, port, "invalid format will have invalid port")

	address, port = getTCP4AddressAndPort("0.0.0.0:du5")

	assert.IsFalse(t, invalidAddress == address, "should not be invalid address")
	assert.IsEqual(t, invalidPort, port, "invalid hex will have invalid port")
}

