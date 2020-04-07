package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/2785/download-purger/internal/durationutil"
)

func main() {
	ignoreWarning := flag.Bool("y", false, "pass in -y to run purger without warning")
	_ = ignoreWarning
	duration := flag.String("duration", "2w", "string representing duration of downloads to clear, e.g.: ")

	usr, _ := user.Current()

	folder := flag.String("dir", fmt.Sprintf("%s/Downloads", usr.HomeDir), "path to the folder you wish to purge")
	flag.Parse()
	dur, err := durationutil.ParseTime(*duration)
	_ = dur
	if err != nil {
		fmt.Printf("Encountered an error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("directory: %s\n", *folder)

	// reader := bufio.NewReader(os.Stdin)

	// fmt.Printf("Input duration: %s, continue? (Y,n)\n", dur)
	// in, _ := reader.ReadString('\n')
	// trimmedIn := strings.TrimSpace(in)
	// if trimmedIn == "" || trimmedIn == "Y" {
	// 	fmt.Println("yes")
	// } else {
	// 	fmt.Println("no")
	// }
}
