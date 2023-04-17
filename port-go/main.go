package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

const (
	bytesInLine = 0x10 // Number of bytes to expect in each line
)

var (
	cAddr    uint64
	hexToCh  = make(map[byte]byte)
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
)

func main() {
	for scanner.Scan() {
		line := scanner.Text()
		line = line[:len(line)-1] // Strip the linefeed (we can't strip all white space here, think of a line of 0x20s)
		data, asciiData := parseLine(line)
		strAddr, strData := data[0], data[1]
		addr, err := hex.DecodeString(strAddr[:len(strAddr)-1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse address in line: '%s'\n", line)
			os.Exit(1)
		}
		addrInt := byteOrderUint64(addr)
		if cAddr != addrInt-bytesInLine {
			if cAddr != 0 {
				fmt.Fprintf(os.Stderr, "Unexpected cAddr in line: '%s'\n", line)
				os.Exit(1)
			}
		}
		cAddr = addrInt
		dataBytes, err := hex.DecodeString(strData)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse data bytes in line: '%s'\n", line)
			os.Exit(1)
		}
		if len(dataBytes) != bytesInLine {
			fmt.Fprintf(os.Stderr, "Unexpected number of bytes in line: '%s'\n", line)
			os.Exit(1)
		}
		// Verify that the mapping from hex data to ASCII is consistent (sanity check for transmission errors)
		for i, b := range dataBytes {
			c := asciiData[i]
			if hexToCh[b] != c {
				fmt.Fprintf(os.Stderr, "Inconsistency between hex data and ASCII data in line (or the lines before): '%s'\n", line)
				os.Exit(1)
			}
		}
		writer.Write(dataBytes)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input: %v\n", err)
		os.Exit(1)
	}
	if err := writer.Flush(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write output: %v\n", err)
		os.Exit(1)
	}
}

func parseLine(line string) ([]string, []byte) {
	data := make([]string, 2)
	var asciiData []byte
	_, err := fmt.Sscanf(line, "%s    %s", &data[0], &asciiData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse line: '%s'\n", line)
		os.Exit(1)
	}
	return data, asciiData
}

func byteOrderUint64(b []byte) uint64 {
	var result uint64
	for i := 0; i < len(b); i++ {
		result <<= 8
		result |= uint64(b[i])
	}
	return result
}

