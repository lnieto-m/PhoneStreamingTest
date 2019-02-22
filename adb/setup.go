package adb

import (
	"fmt"
	"os/exec"
)

// Manager enables ADB commands usage from main
type Manager struct {
	PhoneList []string
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
}
