package main

import (
	"bufio"
	"log"
	"os"
	"reflect"

	"github.com/BertoldVdb/go-ais"
	"github.com/BertoldVdb/go-ais/aisnmea"
	"github.com/schollz/progressbar/v3"
)

func main() {
	file, err := os.Open("data/ais-2023-11-17-07.ais")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bar := progressbar.Default(16120558)

	nm := aisnmea.NMEACodecNew(ais.CodecNew(false, false))

	for scanner.Scan() {
		bar.Add(1)
		sentence := scanner.Text()

		decoded, err := nm.ParseSentence(sentence)

		if err != nil || decoded == nil || decoded.Packet == nil {
			continue
		}

		switch decoded.Packet.GetHeader().MessageID {
		case 1, 2, 3:
			immutable := reflect.ValueOf(decoded.Packet)
			immutable.FieldByName("Sog")
		case 5:
			immutable := reflect.ValueOf(decoded.Packet)
			immutable.FieldByName("Destination")
		}
	}
}
