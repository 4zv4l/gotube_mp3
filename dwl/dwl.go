package dwl

import (
	"bufio"
	"fmt"
	"os"
)

// return the youtube url from file
func GetUrlsFromFile(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	line := bufio.NewScanner(f)
	var urls []string
	for line.Scan() {
		// check url format
		if !checkURL(line.Text()) {
			fmt.Println("invalid url format:", line.Text())
			continue
		}
		urls = append(urls, line.Text())
	}
	return urls
}

// check the url format
func checkURL(url string) bool {
	// check url format
	if len(url) < 11 {
		return false
	}
	if url[:11] != "https://www.y" {
		return false
	}
	return true
}

// get the title of the video from youtube url
func getTitle(url string) string {
	// TODO: get the title of the video
	return ""
}

// Download mp3 from youtube url
func DownloadMp3(url string) {
	// output is the titile of the video
	output := getTitle(url)
	// download mp3
	fmt.Println("Downloading mp3 from", url, "to", output)
	// TODO: download mp3
}
