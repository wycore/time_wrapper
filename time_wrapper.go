package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var verbose bool
var days string
var daysList []string
var hourFrom int
var hourTo int
var timeZone string
var location *time.Location
var tail []string
var version string
var printHelp bool
var printVersion bool

func initArgs() {
	flag.BoolVar(&printHelp, "help", false, "print help and exit")
	flag.BoolVar(&printVersion, "V", false, "print version and exit")
	flag.BoolVar(&verbose, "v", false, "be verbose")
	flag.StringVar(&days, "days", "0,1,2,3,4,5,6", "days from Sun - Sat")
	flag.IntVar(&hourFrom, "from", 0, "active from, as an hour, 0-23 (default 0)")
	flag.IntVar(&hourTo, "to", 24, "active until, as an hour, 0-24")
	flag.StringVar(&timeZone, "timezone", "UTC", "timezone e.g. 'Europe/Berlin'")

	flag.Parse()

	if printHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if printVersion {
		fmt.Println("time_wrapper Version:", version)
		os.Exit(0)
	}

	if hourFrom < 0 || hourFrom > 23 {
		fmt.Println("Invalid hour for 'from':", hourFrom)
		os.Exit(3)
	}
	if hourTo < 0 || hourTo > 24 {
		fmt.Println("Invalid hour for 'to':", hourTo)
		os.Exit(3)
	}

	// this is the binary to run, if the time's right
	tail = flag.Args()
	if len(tail) < 1 {
		os.Exit(3)
	}

	var err error
	if location, err = time.LoadLocation(timeZone); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

}

func main() {
	initArgs()

	now := time.Now().In(location)
	nowDay := strconv.Itoa(int(now.Weekday()))
	nowHour := now.Hour()
	daysList := strings.Split(days, ",")

	// defaulting to not execute to make the next part easier
	willExecute := false

	for _, _day := range daysList {
		if string(_day) == nowDay {
			willExecute = true
		}
	}

	// don't run if it's too soon
	if nowHour < hourFrom {
		willExecute = false
	}

	// don't run if it's too late
	if nowHour >= hourTo {
		willExecute = false
	}

	if verbose {
		fmt.Println("days     : ", days)
		fmt.Println("from     : ", hourFrom)
		fmt.Println("to       : ", hourTo)
		fmt.Println("tail     : ", tail)
		fmt.Println("timezone : ", location)
		fmt.Println("now      : ", now)
		fmt.Println("nowD     : ", nowDay)
		fmt.Println("nowH     : ", nowHour)
		fmt.Println("run?     : ", willExecute)
	}

	if !willExecute {
		os.Exit(0)
	}

	// get the real path of the executable
	binary, lookErr := exec.LookPath(tail[0])
	if lookErr != nil {
		fmt.Println(lookErr)
		os.Exit(3)
	}
	env := os.Environ()

	// execute, incl. stdout, stderr and exit code \o/
	execErr := syscall.Exec(binary, tail, env)
	if execErr != nil {
		fmt.Println(lookErr)
		os.Exit(2)
	}
}
