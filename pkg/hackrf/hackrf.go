package hackrf

import (
	"io"
	"log"
	"os/exec"
	"strconv"
)

const MinSampleRate = 2000000

func Transmit(sampleRate int, freq int, reader io.Reader) error {
	cmd := exec.Command("hackrf_transfer", "-x", "47", "-a", "1", "-f", strconv.Itoa(freq), "-s", strconv.Itoa(sampleRate), "-t", "-")
	cmd.Stdin = reader
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("hackrf_transfer cmd %s failed:\nstdout: %s\nerr:%s", cmd, out, err)
	}
	return err
}
