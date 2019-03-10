package adb

import (
	"fmt"
	"os/exec"
)

// Manager enables ADB commands usage from main
type Manager struct {
	PhoneList []string

	MinicapChan  chan []byte
	StatusChange chan []string
	StatusStop   chan bool
}

func (manager *Manager) setPhoneList() {

}

// Start : EntryPoint
func (manager *Manager) Start() {

	fmt.Println("Starting ADB...")

	cmd := exec.Command("adb", "devices")
	startingResult, err := cmd.Output()
	if err != nil {
		fmt.Print("???")
	}
	fmt.Println(string(startingResult))
	manager.StatusChange = make(chan []string, 1)
	manager.StatusStop = make(chan bool, 1)
	manager.MinicapChan = make(chan []byte)
	go manager.setupMinicap("8425384f34503231")
	// go manager.StatusNotifier()
}
