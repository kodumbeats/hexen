package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kodumbeats/hex/v2"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			cwd, err := os.Getwd()
			if err != nil {
				log.Println(err)
			}

			tbl := filepath.Clean(c.Args().Get(1))
			tblPath := filepath.Join(cwd, tbl)
			buf, err := os.ReadFile(tblPath)
			if err != nil {
				log.Println(err)
			}

			charTblSlice := strings.Split(string(buf), "\n")

			charMap := make(map[string][]byte)

			for _, kv := range charTblSlice {
				stanzas := strings.Split(kv, "=")
				if len(stanzas) == 2 {
					// k := []byte(stanzas[0])
					if err != nil {
						log.Println(err)
					}
					v := []byte(stanzas[1])
					charMap[stanzas[0]] = v
				}
			}

			fmt.Print(charMap)

			rom := filepath.Clean(c.Args().Get(0))
			romPath := filepath.Join(cwd, rom)
			data, err := os.ReadFile(romPath)
			if err != nil {
				log.Println(err)
			}

			for i := 0; i < len(data); i = i + 16 {
				line := data[i : i+16]

				for j := 0; j < 16; j++ {
					fmt.Print(hex.EncodeToString(line[j:j+1]) + " ")
					if j == 7 {
						fmt.Print(" ")
					}
				}
				fmt.Print("\n")
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
