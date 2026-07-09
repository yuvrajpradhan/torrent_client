package main

import (
	"fmt"
	"log"
	"os"

	"torrent-client/torrentfile"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: torrent-client <path_to_torrent> <output_path>")
		os.Exit(1)
	}

	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
