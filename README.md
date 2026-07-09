# Torrent Client (Go)

A lightweight BitTorrent client written in Go from scratch. This project implements the BitTorrent protocol to download files from peers. It handles decoding `.torrent` files, communicating with trackers, connecting to peers, managing piece downloads concurrently, and verifying data integrity using SHA-1 hashes.

## Features

- **Bencode Parsing**: Decodes the bencoded format of `.torrent` files.
- **Tracker Communication**: Sends HTTP requests to the tracker to get a list of peers.
- **Peer-to-Peer Protocol**: Implements the BitTorrent wire protocol, including handshakes, piece requests, and choking/unchoking logic.
- **Concurrent Downloading**: Uses Go channels and goroutines to download multiple pieces from different peers concurrently.
- **Data Integrity**: Verifies downloaded pieces against the SHA-1 hashes provided in the torrent file.
- **Memory Efficiency**: Streams pieces directly to disk, allowing you to download massive files without exhausting your RAM.
- **Robustness**: Graceful error handling and early exit detection preventing deadlocks if all peers disconnect.

## Usage

Build the client:
```bash
go build -o torrent-client main.go
```

Run the client:
```bash
./torrent-client <path_to_torrent> <output_path>
```

### Examples

Download the Ubuntu ISO using the provided test torrent file:
```bash
./torrent-client ubuntu-26.04-desktop-amd64.iso.torrent ubuntu.iso
```

## Differences from Original Project

This project was inspired by and built upon the excellent guide by [veggiedefender/torrent-client](https://github.com/veggiedefender/torrent-client/).

However, this implementation introduces key improvements over the original for better robustness and real-world usage:
- **Disk Streaming vs RAM Buffering**: The original implementation allocated a single byte array in memory, storing the entire file in RAM before writing it to disk. This version opens the output file immediately and streams pieces directly to disk as they arrive, allowing you to download massive files without exhausting system memory.
- **Deadlock Prevention**: The original implementation could deadlock and hang indefinitely if all peer connections dropped before the download finished. This version introduces a signaling channel to detect when all worker goroutines have exited, ensuring the client cleanly terminates with an error instead of hanging.
