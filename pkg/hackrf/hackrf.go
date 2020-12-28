package hackrf

import (
	"io"
	"log"
	"os/exec"
	"strconv"
	"sync"
)

const MinSampleRate = 2000000

type reader struct {
	mu     sync.Mutex
	cur    []byte
	pos    int
	repeat bool
}

func (r *reader) Read(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := 0; i < len(p); i++ {
		if len(r.cur) == 0 {
			p[i] = 0
		} else {
			p[i] = r.cur[r.pos]
			r.pos = (r.pos + 1)
			if !r.repeat && r.pos == len(r.cur) {
				return i + 1, io.EOF
			}
			r.pos = r.pos % len(r.cur)
		}
	}
	return len(p), nil
}

func (r *reader) setCur(cur []byte) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cur = cur
	r.pos = 0
}

func (r *reader) setRepeat(repeat bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.repeat = repeat
}

// HackRFTX implements transmitting using a HackRF.
type HackRFTX struct {
	r          *reader
	sampleRate int
	freq       int
}

func NewHackRFTX(sampleRate int, freq int) *HackRFTX {
	return &HackRFTX{&reader{}, sampleRate, freq}
}

func (h *HackRFTX) SetIQData(iq []byte) {
	h.r.setCur(iq)
}

func (h *HackRFTX) SetRepeat(repeat bool) {
	h.r.setRepeat(repeat)
}

func (h *HackRFTX) Start() error {
	cmd := exec.Command("hackrf_transfer", "-x", "47", "-a", "1", "-f", strconv.Itoa(h.freq), "-s", strconv.Itoa(h.sampleRate), "-t", "-")
	cmd.Stdin = h.r
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("hackrf_transfer cmd %s failed:\nstdout: %s\nerr:%s", cmd, out, err)
	}
	return err
}
