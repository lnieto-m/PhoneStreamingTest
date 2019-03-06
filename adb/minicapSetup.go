package adb

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func (manager *Manager) setupMinicap(serial string) {

	// Install minicap on Phone and launch it

	log.Println("installing minicap")
	os.Setenv("ANDROID_SERIAL", serial)
	log.Println(os.Getenv("ANDROID_SERIAL"))
	cmd := exec.Command("echo", "$ANDROID_SERIAL")
	cmd.Env = os.Environ()
	outt, _ := cmd.Output()
	log.Println(string(outt))
	log.Println("launching...")
	cmd = exec.Command(strings.Join([]string{"bash -c 'screen -S ", serial, " -d -m ./run.sh autosize'"}, ""))
	cmd.Dir = "/Volumes/SAMSUNG/Library/minicap"
	out, _ := cmd.Output()
	log.Println(string(out))

	// Create local forward to connect to the socket

	log.Println("tcp:1313 <- localabstract:minicap")
	cmd = exec.Command("adb forward tcp:1313 localabstract:minicap")
	cmd.Run()

	// Start getting frames from minicap and push them to the websocket

	log.Println("Connecting to localhost:1313 ...")
	conn, err := net.Dial("tcp", "localhost:1313")
	if err != nil {
		log.Println(err)
		return
	}

	first := 0

	log.Println("Reading from localhost:1313 ...")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if first == 0 {
			first++
			_ = scanner.Bytes()
		} else {
			log.Println("urmom")
			manager.MinicapChan <- scanner.Bytes()
		}
	}

}
