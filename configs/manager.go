package configs

import (
	"fmt"
	"os"
)

var Manager manager

type manager struct {
	HostCredentials *hostCredentials
}

type hostCredentials struct {
	PORT string
}

func (m *manager) Setup() {

	defaultPort := "5000"
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	m.HostCredentials = &hostCredentials{PORT: fmt.Sprintf(":%s", port)}

}
