package connect

import (
	"fmt"
	"strings"
	"testing"
)

func TestLiveConnectionToEthereumNode(t *testing.T) {
		// Add the `testing.Short()` line.
		testing.Short()

	const nodeURLSepolia string = "https://sepolia.infura.io/v3/bfdf85f781ae493aaf2773a6b495c36a"
	input := nodeURLSepolia
	conn, _ := LiveConnectionToEthereumNode(input)
	if conn == nil {
		t.Errorf("Function returns wrong data: \nExpected: to have a memory adress, \ngot %v", conn)
	}

	if !strings.HasPrefix(fmt.Sprintf("%v", conn), "0x") {
		t.Errorf("Expected conn to start with 0x, got %v", conn)
	}

}
