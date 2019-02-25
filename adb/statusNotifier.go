package adb

import (
	"log"
	"os/exec"
	"regexp"
	"time"
)

func (manager *Manager) findAndUpdate(listToCompare []string) {
	notFound := 0

	log.Printf("compare %v %v\n", manager.PhoneList, listToCompare)

globalList:
	for index := range listToCompare {
		for _, v := range manager.PhoneList {
			if v == listToCompare[index] {
				continue globalList
			}
		}
		log.Print("sending...\n")
		notFound = 1
		manager.StatusChange <- listToCompare[index]
	}

	if notFound == 1 {
		manager.PhoneList = listToCompare
	}
	log.Print("end reached\n")
}

func (manager *Manager) statusNotifier() {
	regxp := regexp.MustCompile(`(?m)(.+)\sdevice\s`)
	ticker := time.NewTicker(2 * time.Second)

	for range ticker.C {
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
		manager.findAndUpdate(listToCompare)
	}
}
