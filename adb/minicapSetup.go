package adb

import (
	"bufio"
	"encoding/binary"
	"io"
	"log"
	"net"
	"os/exec"
)

func (manager *Manager) setupMinicap(serial string) {

	// Launch minicap script

	log.Printf("Launching minicap on device %v\n", serial)

	cmd := exec.Command("./scripts/minicapSetup.sh", serial)
	err := cmd.Run()
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	log.Println("Connecting to localhost:1313 ...")
	conn, err := net.Dial("tcp", "localhost:1313")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Reading from localhost:1313 ...")

	var pid, rw, rh, vw, vh uint32
	var version uint8
	var unused uint8
	var orientation uint8
	binRead := func(data interface{}) error {
		if err != nil {
			return err
		}
		err = binary.Read(conn, binary.LittleEndian, data)
		return err
	}
	binRead(&version)
	binRead(&unused)
	binRead(&pid)
	binRead(&rw)
	binRead(&rh)
	binRead(&vw)
	binRead(&vh)
	binRead(&orientation)
	binRead(&unused)
	if err != nil {
		log.Printf("Error Banner: %v\n", err)
		return
	}
	bufrd := bufio.NewReader(conn) // Do not put it into for loop
	for {
		var size uint32
		if err = binRead(&size); err != nil {
			break
		}
		lr := &io.LimitedReader{
			R: bufrd,
			N: int64(size),
		}
		log.Printf("size frame %v\n", size)
		if size > 1000000000 {
			continue
		}
		// var im image.Image
		// im, err = jpeg.Decode(lr)
		buffer := make([]byte, int64(size))
		lr.Read(buffer)
		// im, _, err = image.Decode(lr)
		if err != nil {
			break
		}
		log.Println("Frame")
		manager.MinicapChan <- buffer

	}
	conn.Close()
}
