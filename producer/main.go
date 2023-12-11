package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"net"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

var (
	port      int
	delay     float64
	numEvents int
	events    []string
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info(
		fmt.Sprintf(
			"Sending %d events with a delay of %.2f seconds",
			numEvents,
			delay,
		),
	)

	// Read raw file
	_, fileraw, _, _ := runtime.Caller(0)
	fileraw = path.Join(path.Dir(fileraw), "raw.txt")

	raw, err := os.ReadFile(fileraw)
	if err != nil {
		panic(err.Error())
	}

	// Store raw data in memory as a list of strings
	for _, line := range strings.Split(string(raw), "\n") {
		events = append(events, line)
	}
	logger.Info("All events successfully loaded")

	// Connect on the given port with the UDP protocol
	server, err := net.ListenPacket("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	logger.Info(fmt.Sprintf("UDP server listening on port %d", port))

	// Initialize a new multi printer
	multi := pterm.DefaultMultiPrinter
	multi.Start()
	fmt.Printf("\n")

	// Infinite loop to handle incoming connection requests
	for {
		buffer := make([]byte, 50)
		_, address, err := server.ReadFrom(buffer)
		if err != nil {
			logger.Error(
				fmt.Sprintf("Error in reading from the incoming connection: %s\n", err.Error()),
			)
			continue
		}

		// Add a new progress bar
		pbname := fmt.Sprintf("%10s client", strings.Trim(string(buffer), "\u0000"))
		pb, _ := pterm.DefaultProgressbar.WithTotal(numEvents).WithWriter(multi.NewWriter()).Start(pbname)

		// Spawn a new goroutine for the new incoming connection
		go producer(server, address, pb)
	}
}

func producer(s net.PacketConn, a net.Addr, pb *pterm.ProgressbarPrinter) {

	// Send the number of events requested
	for i := 0; i < numEvents; i++ {
		s.WriteTo([]byte(events[rand.Intn(len(events))]), a)
		pb.Increment()

		time.Sleep(time.Duration(delay * float64(time.Second)))
	}

	// Tell client no more data will be sent
	s.WriteTo([]byte("EOF"), a)
}

func init() {
	// Configuration flags
	flag.IntVar(&port, "port", 1053, "UDP server port")

	// Events flag
	flag.IntVar(&numEvents, "events", 100, "the number of events to send")
	flag.Float64Var(&delay, "delay", 0.01, "time to wait before sending next event")

	flag.Parse()
}
