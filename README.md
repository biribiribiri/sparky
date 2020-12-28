[![GoDoc](https://pkg.go.dev/github.com/biribiribiri/sparky?status.svg)](http://pkg.go.dev/github.com/biribiribiri/sparky)

# sparky
Golang library for controlling the Dogtra 2300NCP shock collar using the HackRF SDR.

DISCLAIMER: This project is not affiliated with or endorsed by Dogtra.

# Pairing

To pair your collar to the software remote:
1. Turn off the collar.
2. Hold a magnet to the red dot on the collar until the green LED on the collar blinks rapidly.
3. While the green LED on the collar is blinking rapidly, send a pairing command using the software remote.

	```
	./dt2300ncp --cmd=pairing --duration=5s
	```

4. The green LED should stop blinking rapidly. You should now be able to send
   other commands to the collar, e.g. vibrate.

	```
	./dt2300ncp --cmd=vibrate --duration=5s
	```

# CLI Examples
```
# Maximum intensity nick.
./dt2300ncp --cmd=nick --intensity=127

# Continuous stimulation for 1 second.
./dt2300ncp --cmd=continuous --intensity=30 --duration=1s

# Vibrate for 5 seconds.
./dt2300ncp --cmd=vibrate --duration=5s
```