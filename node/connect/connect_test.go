package connect

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLiveConnectionToEthereumNode(t *testing.T) {
	// Add the `testing.Short()` line.
	testing.Short()

	const nodeURLSepolia string = "https://sepolia.infura.io/v3/bfdf85f781ae493aaf2773a6b495c36a"
	input := nodeURLSepolia
	app := NewApp(logrus.New())
	app.LiveConnectionToEthereumNode(input)
	if app.conn == nil {
		t.Errorf("Function returns wrong data: \nExpected: to have a memory adress, \ngot %v", app.conn)
	}

	// Так сравнивать не нужно, переводить адрес объекта в строку и брать префикс такое себе
	// Уж лучше тогда с nil сравнить сам указатель на объект
	if !strings.HasPrefix(fmt.Sprintf("%v", app.conn), "&{0x") {
		t.Errorf("Expected conn to start with 0x, got %v", *app.conn)
	}
}
