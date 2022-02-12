package main

import (
	"flag"
	"fmt"
	"gotubemp3/dwl"
)

func main() {
	// get arguments
	help := flag.Bool("help", false, "show help") // -help
	file := flag.String("f", "", "input file")    // -f <file>
	url := flag.String("u", "", "input url")      // -u <url>
	// change argument order in the usage
	flag.Usage = func() {
		fmt.Println("Usage: gotubemp3 [options]")
		fmt.Println("Options:")
		fmt.Println("  -help            Show this message")
		fmt.Println("  -f <file>        Input file")
		fmt.Println("  -u <url>         Input url")
	}
	flag.Parse()
	// check arguments part
	if *help {
		flag.Usage()
		return
	}
	// if file is given, get urls from file
	if *file != "" {
		urls := dwl.GetUrlsFromFile(*file)
		// download mp3 from urls
		for _, url := range urls {
			dwl.DownloadMp3(url)
		}
	} else if *url != "" {
		dwl.DownloadMp3(*url)
	} else {
		// show help
		flag.Usage()
	}
}
