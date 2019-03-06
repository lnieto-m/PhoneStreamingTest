package adb

import (
	"log"
	"os/exec"
	"regexp"
	"time"
)

// StatusNotifier starts the
func (manager *Manager) StatusNotifier() {
	regxp := regexp.MustCompile(`(?m)(.+)\sdevice\s`)
	ticker := time.NewTicker(2 * time.Second)

	for range ticker.C {
		select {
		case <-manager.StatusStop:
			log.Print("Notifier stopped\n")
			return
		default:
			var listToCompare []string
			cmd := exec.Command("adb", "devices")
			deviceList, err := cmd.Output()
			if err != nil {
				log.Printf("error %v\n", err)
			}
			for _, match := range regxp.FindAllStringSubmatch(string(deviceList), -1) {
				listToCompare = append(listToCompare, string(match[1]))
			}
			// log.Printf("tmp: %v\n", tmp)
			manager.StatusChange <- listToCompare
		}
	}
}
