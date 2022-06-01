package main

import (
	"flag"
	"fmt"
	buffer "gocp/bufferAlign"
	progress "gocp/progressBar"
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
var BW_LIMIT int
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
	buffer_size := flag.Int("buffersize", 4096, "set buffer size")
	bw_limit := flag.Int("bwlimit", 0, "limit IO bandwidth in MB, no limit by value 0")
	directio := flag.Bool("directio", false, "direct IO mode")
	show_progress := flag.Bool("progress", false, "show progress")
	show_help := flag.Bool("help", false, "show command help info")

	flag.Parse()

	BUFFER_SIZE = *buffer_size
	BW_LIMIT = *bw_limit
	DIRECTIO = *directio
	SHOW_PROGRESS = *show_progress
	SHOW_HELP = *show_help
}

func printHelp() {
	println("usage: gocp [options] <source file path> <destination file path>")
	println("options: ")
	println("  -buffersize <integer>  : set buffer size")
	println("  -bwlimit <integer>  : set IO bandwidth limit in MB.")
	println("  -directio  : use direct IO mode")
	println("  -progress  : show progress")
	println("  -help  : show help")
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

func main() {

	parseOptions()

	optLog := fmt.Sprintf("bufsize=%d, bwlimit=%d, directio=%t, progress=%t, help=%t", BUFFER_SIZE, BW_LIMIT, DIRECTIO, SHOW_PROGRESS, SHOW_HELP)
	printLog(optLog)

	if SHOW_HELP == true {
		printHelp()
		os.Exit(0)
	}

	var srcFile string
	var destFile string

	argsFiles := flag.NArg()
	if argsFiles == 2 {
		srcFile = flag.Arg(0)
		destFile = flag.Arg(1)
	} else {
		log.Fatal("Invalid file arguments")
	}

	srcStat, err := os.Stat(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	if !srcStat.Mode().IsRegular() {
		log.Fatal("%s is not a regular file", srcFile)
	}

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

	// if dest is a dir, update to filepath with same source filename
	destStat, err := os.Stat(destFile)
	if destStat.IsDir() {
		srcFilename := path.Base(srcFile)
		destFile = path.Join(destFile, srcFilename)
	}

	var destFP *os.File
	if DIRECTIO {
		destFP, err = os.Create(destFile) // create file if not exist
		if err != nil {
			log.Fatal(err)
		}
		destFP.Close()
		destFP, err = os.OpenFile(destFile, syscall.O_DIRECT|os.O_WRONLY, 0664)
	} else {
		destFP, err = os.Create(destFile)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer destFP.Close()

	srcSize := srcStat.Size()
	srcSizeStr := strconv.FormatInt(int64(srcSize), 10)
	srcSizeStr = "(size=" + srcSizeStr + ")"
	printLog("Copying", srcFile, "to", destFile, srcSizeStr)

	var buf []byte
	if DIRECTIO {
		buf = buffer.GetAlignedBlock(BUFFER_SIZE)
	} else {
		buf = make([]byte, BUFFER_SIZE)
	}

	if SHOW_PROGRESS {
		ch := make(chan bool)
		go func() {
			printProgressBar(ch, srcSize)
		}()
		copyFileWithWriteCount(srcFP, destFP, buf)
		<-ch
	} else {
		copyFile(srcFP, destFP, buf)
	}
	printLog("Copy completed")
}
