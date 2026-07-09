// package main

// import (
// 	"fmt"
// 	"log"
// 	"torrent-client/client"
// 	"torrent-client/torrentfile"
// )

// func main() {
// 	tf, err := torrentfile.Open("/home/yuvraj/projects/torrent-client/torrentfile/testdata/ubuntu-26.04-desktop-amd64.iso.torrent")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%+v\n", tf)

// 	peerID := [20]byte{}

// 	peers, err := tf.RequestPeers(peerID, 6881)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(peers)

// 	//testing connection
// 	c, err := client.New(peers[0], peerID, tf.InfoHash)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Handshake successful!")
// 	fmt.Println("Peer:", peers[0])
// 	fmt.Println("Choked:", c.Choked)
// 	fmt.Println("Bitfield length:", len(c.Bitfield))
// }

package main

import (
	"log"
	"os"

	"torrent-client/torrentfile"
)

func main() {
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
