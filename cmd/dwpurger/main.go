package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/2785/download-purger/internal/durationutil"
	"github.com/2785/download-purger/internal/fsutil"
)

func main() {
	ignoreWarning := flag.Bool("y", false, "pass in -y to run purger without warning")
	duration := flag.String("duration", "2week", "string representing duration of downloads to clear, e.g.: ")
	folder := flag.String("dir", fmt.Sprintf("~/Downloads"), "path to the folder you wish to purge")
	output := flag.String("o", "~/logs/dwpurger.log", "path to log file")
	flag.Parse()
	dur, err := durationutil.ParseTime(*duration)
	if err != nil {
		fmt.Printf("cannot parse duration: %s\n", err.Error())
		os.Exit(1)
	}

	targetDir, err := fsutil.ParsePath(*folder)
	if err != nil {
		fmt.Printf("cannot parse target dir: %s\n", err.Error())
	}

	outputPath, err := fsutil.ParsePath(*output)
	if err != nil {
		fmt.Printf("cannot parse output dir: %s\n", err.Error())
	}

	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		fmt.Printf("cannot read directory: %s\n", err.Error())
	}

	deletionCandidates := func() []os.FileInfo {
		s := []os.FileInfo{}
		for _, f := range files {
			if time.Since(f.ModTime()) > dur {
				s = append(s, f)
			}
		}
		return s
	}()

	if len(deletionCandidates) == 0 {
		if !*ignoreWarning {
			fmt.Println("nothing to delete, all good")
		}
		os.Exit(0)
	}
	if !*ignoreWarning {
		fmt.Printf("List of files older than %s to delete in %s:\n", dur, *folder)
		for i, v := range deletionCandidates {
			fmt.Printf("#%v: %s\n", i+1, v.Name())
		}
		fmt.Printf("Continue? (Y/n)")

		reader := bufio.NewReader(os.Stdin)
		in, _ := reader.ReadString('\n')
		trimmedIn := strings.TrimSpace(in)

		if trimmedIn == "" || trimmedIn == "Y" {
			fmt.Printf("you said yes, deleting %v items\n", len(deletionCandidates))
		} else {
			fmt.Println("You said no, okay. Exiting :)")
			os.Exit(0)
		}
	}

	deletionOutcome := func() map[string]struct {
		ok  bool
		err error
	} {
		m := make(map[string]struct {
			ok  bool
			err error
		})
		for _, v := range deletionCandidates {
			fileDir := filepath.Join(targetDir, v.Name())
			err := os.RemoveAll(fileDir)
			m[v.Name()] = struct {
				ok  bool
				err error
			}{ok: err != nil, err: err}
		}
		return m
	}()

	logDir := filepath.Dir(outputPath)

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, os.ModePerm)
	}

	logFile, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	if err != nil {
		fmt.Printf("cannot access log file: %s\n", err.Error())
		os.Exit(1)
	}
	_ = logFile

	for k, v := range deletionOutcome {
		logFile.WriteString(fmt.Sprintf("File %s: %s [%s]\n", k, func() string {
			if v.ok {
				return "deleted"
			} else {
				return fmt.Sprintf("failed to delete: %s\n", v.err.Error())
			}
		}(), time.Now()))
	}

}
