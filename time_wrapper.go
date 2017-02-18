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

func main() {
	// debugging help
	// args := os.Args
	// fmt.Println(args)

	helpPtrShort := flag.Bool("h", false, "show help")
	helpPtrLong := flag.Bool("help", false, "show help")
	verbosePtrShort := flag.Bool("v", false, "be verbose")
	verbosePtrLong := flag.Bool("verbose", false, "be verbose")
	daysPtr := flag.String("days", "0,1,2,3,4,5,6", "days from Sun - Sat")
	fromPtr := flag.Int("from", 0, "active from, as an hour, 0-23 (default 0)")
	toPtr := flag.Int("to", 24, "active until, as an hour, 0-24")

	flag.Parse()

	showHelp := *helpPtrShort || *helpPtrLong
	if showHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	isVerbose := *verbosePtrShort || *verbosePtrLong
	days := strings.Split(*daysPtr, ",")
	hourFrom := *fromPtr
	hourTo := *toPtr

	if hourFrom < 0 || hourFrom > 23 {
		fmt.Println("Invalid hour for 'from':", hourFrom)
		os.Exit(3)
	}
	if hourTo < 0 || hourTo > 24 {
		fmt.Println("Invalid hour for 'to':", hourTo)
		os.Exit(3)
	}

	// this is the binary to run, if the time's right
	tail := flag.Args()
	if len(tail) < 1 {
		os.Exit(3)
	}

	now := time.Now().UTC()
	nowDay := strconv.Itoa(int(now.Weekday()))
	nowHour := now.Hour()

	// defaulting to not execute to make the next part easier
	willExecute := false

	for _, _day := range days {
		if _day == nowDay {
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

	if isVerbose {
		fmt.Println("days : ", days)
		fmt.Println("from : ", hourFrom)
		fmt.Println("to   : ", hourTo)
		fmt.Println("tail : ", tail)
		fmt.Println("now  : ", now)
		fmt.Println("nowD : ", nowDay)
		fmt.Println("nowH : ", nowHour)
		fmt.Println("run? : ", willExecute)
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
