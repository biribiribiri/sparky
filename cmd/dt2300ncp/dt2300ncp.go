// Binary dt2300ncp is a simple tool for sending commands to a Dogtra 2300NCP
// collar using HackRF.
package main

import (
	"flag"
	"log"
	"time"

	dt "github.com/biribiribiri/sparky/pkg/dt2300ncp"
	"github.com/biribiribiri/sparky/pkg/hackrf"
)

var cmdFlag = flag.String("cmd", "", "command, one of {nick, continuous, vibrate, pairing}")
var repeatFlag = flag.Bool("repeat", false, "if true, repeat the command until terminated")
var intensityFlag = flag.Int("intensity", 1, "intensity, between 0 and 127 inclusive")
var durationFlag = flag.Duration("duration", time.Duration(0), "duration of the command")

func parseCmd() dt.Cmd {
	switch *cmdFlag {
	case "nick":
		return dt.NickCmd
	case "continuous":
		return dt.ContinuousCmd
	case "vibrate":
		return dt.VibrateCmd
	case "pairing":
		return dt.PairingCmd
	default:
		log.Fatalf("invalid --cmd flag %q, expected {nick, continuous, vibrate, pairing}", *cmdFlag)
		return dt.VibrateCmd
	}
}

func main() {
	flag.Parse()
	hrf := hackrf.NewHackRFTX(hackrf.MinSampleRate, dt.Freq)
	hrf.SetIQData(dt.GenerateIQ(hackrf.MinSampleRate, parseCmd(), dt.CollarID1, *intensityFlag, *durationFlag))
	hrf.SetRepeat(*repeatFlag)
	hrf.Start()
}
