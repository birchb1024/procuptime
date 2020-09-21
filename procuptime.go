package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type rebooted struct {
	BootEpochTime  int64  `json:"booted"`
	UpTimeSecs     int64  `json:"uptime"`
	BootedDateTime string `json:"datebooted"`
	UpTimeDuration string `json:"uptimeduration"`
}

func main() {

	var jso bool
	flag.BoolVar(&jso, "j", false, "Print help")
	var help bool
	flag.BoolVar(&help, "h", false, "JSON output")

	flag.Parse()
	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tok := strings.Split(line, " ")

	upduration, err := time.ParseDuration(tok[0] + "s")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	now := time.Now()

	if jso {
		var p rebooted
		p.BootEpochTime = now.Add(-upduration).Unix()
		p.UpTimeSecs = int64(upduration.Seconds())
		p.BootedDateTime = now.Add(-upduration).Format(time.RFC3339)
		p.UpTimeDuration = upduration.String()
		pj, _ := json.Marshal(p)
		fmt.Println(string(pj))
	} else {
		fmt.Println(now.Add(-upduration).Unix(),
			int(upduration.Seconds()),
			now.Add(-upduration).Format(time.RFC3339),
			upduration) // 195013
	}

}
