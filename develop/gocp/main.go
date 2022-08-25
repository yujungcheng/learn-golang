package main

import (
	"flag"
	"fmt"
	buffer "gocp/bufferAlign"
	progress "gocp/progressBar"
	limiter "gocp/ioLimiter"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var BUFFER_SIZE int
var IO_LIMIT int
var DIRECTIO bool
var SHOW_PROGRESS bool
var SHOW_HELP bool
var WRITE_COUNT int64

func getTime() string {
	timeLayout := "2006-01-02T15:04:05 -"
	now := time.Now()
	return now.Format(timeLayout)
}

func printLog(s ...string) {
	ts := getTime()
	msg := strings.Join(s, " ")
	fmt.Println(ts, msg)
}

func parseOptions() {
	buffer_size := flag.Int("buffersize", 128, "set buffer size in KB, default 128KB.")
	io_limit := flag.Int("iolimit", 0, "limit IO bandwidth in MB.")
	directio := flag.Bool("directio", false, "use direct IO mode.")
	show_progress := flag.Bool("progress", false, "show progress.")
	show_help := flag.Bool("help", false, "show command help info.")

	flag.Parse()

	BUFFER_SIZE = *buffer_size * 1024
	IO_LIMIT = *io_limit * 1024 * 1024
	DIRECTIO = *directio
	SHOW_PROGRESS = *show_progress
	SHOW_HELP = *show_help
}

func printHelp() {
	println("Usage: gocp [options] <source file path> <destination file path>")
	println("options: ")
	println("  -buffersize <integer>  : set buffer size in KB, default 128KB.")
	println("  -iolimit <integer>  : set IO bandwidth limit in MB.")
	println("  -directio  : use direct IO mode")
	println("  -progress  : show progress")
}

func printProgressBar(ch chan bool, fileSize int64) {

	var bar progress.Bar
	var maxCount int64
	var reminder int64
	var totalStep int64
	var currentStep int64

	maxCount, reminder = fileSize/int64(BUFFER_SIZE), fileSize%int64(BUFFER_SIZE)
	if reminder != 0 {
		maxCount++
	}

	totalStep = maxCount / 100
	bar.NewOption(0, 100)
	for {

		currentStep = WRITE_COUNT / totalStep

		bar.Play(int64(currentStep))
		time.Sleep(1 * time.Second)
		//time.Sleep(500 * time.Millisecond)

		if WRITE_COUNT == maxCount {
			bar.Play(int64(100))
			break
		}
	}
	bar.Finish()
	ch <- true
}

func copyFile(src, dest *os.File, buf []byte) {
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if n == 0 {
			break
		}

		if _, err := dest.Write(buf[:n]); err != nil {
			log.Fatal(err)
		}
	}
}

func copyFileWithWriteCount(src, dest *os.File, buf []byte) {
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if n == 0 {
			break
		}

		if _, err := dest.Write(buf[:n]); err != nil {
			log.Fatal(err)
		}
		WRITE_COUNT++
	}
}

func copyFileWithIoLimit(src, dest *os.File, buf []byte) {
	var l limiter.Limiter
	l.SetLimiter(float64(IO_LIMIT))
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if n == 0 {
			break
		}

		if err := l.WaitN(n); err != nil {
			log.Fatal(err)
		}

		if _, err := dest.Write(buf[:n]); err != nil {
			log.Fatal(err)
		}
		WRITE_COUNT++
	}
}

func main() {

	parseOptions()

	if SHOW_HELP == true {
		printHelp()
		os.Exit(0)
	}
	flag.Usage = printHelp

	optLog := fmt.Sprintf("bufsize=%dKB, iolimit=%d, directio=%t, progress=%t, help=%t", 
		BUFFER_SIZE/1024, IO_LIMIT, DIRECTIO, SHOW_PROGRESS, SHOW_HELP)
	printLog(optLog)

	var srcFile string
	var destFile string
	var srcSize int64
	argsFiles := flag.NArg()
	if argsFiles == 2 {
		srcFile = flag.Arg(0)
		destFile = flag.Arg(1)
	} else {
		log.Fatal("Please input one source path and one destionation path.")
	}

	// check source path
	if srcStat, err := os.Stat(srcFile); os.IsNotExist(err) {
		log.Fatal("Source path ", srcFile, " not exist.")
	} else if srcStat.IsDir() {
		log.Fatal(srcFile, " is not a regular file")
	} else {
		srcSize = srcStat.Size()
	}

	// check destination path
	if destStat, err := os.Stat(destFile); os.IsNotExist(err) {
		if err := os.MkdirAll(destFile, os.ModePerm); err != nil {
			log.Fatal("Can not create destination directory. ", err)
		}
	} else {
		// if dest is a dir, update to filepath with same source filename
		if destStat.IsDir() {
			srcFilename := path.Base(srcFile)
			destFile = path.Join(destFile, srcFilename)
		} 
	}

	var destFP *os.File
	destFP, err := os.Create(destFile) // create file if not exist
	if err != nil {
		log.Fatal(err)
	}
	destFP.Close()
	if DIRECTIO {
		destFP, err = os.OpenFile(destFile, syscall.O_DIRECT|os.O_WRONLY, 0664)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		destFP, err = os.Create(destFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer destFP.Close()

	srcSizeStr := strconv.FormatInt(int64(srcSize), 10)
	srcSizeStr = "(size=" + srcSizeStr + ")"
	printLog("Copying", srcFile, "to", destFile, srcSizeStr)

	var srcFP *os.File
	if DIRECTIO {
		srcFP, err = os.OpenFile(srcFile, syscall.O_DIRECT|os.O_RDONLY, 0664)
	} else {
		srcFP, err = os.Open(srcFile)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer srcFP.Close()

	var buf []byte
	if DIRECTIO {
		buf = buffer.GetAlignedBlock(BUFFER_SIZE)
	} else {
		buf = make([]byte, BUFFER_SIZE)
	}

	if SHOW_PROGRESS || IO_LIMIT != 0 {
		ch := make(chan bool)
		go func() {
			printProgressBar(ch, srcSize)
		}()
		if IO_LIMIT != 0 {
			copyFileWithIoLimit(srcFP, destFP, buf)
		} else {
			copyFileWithWriteCount(srcFP, destFP, buf)
		}
		<-ch
	} else {
		copyFile(srcFP, destFP, buf)
	}
	printLog("Copy completed")
}
