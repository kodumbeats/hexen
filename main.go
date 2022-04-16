package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

			charMap := make(map[byte]byte)

			for _, kv := range charTblSlice {
				stanzas := strings.Split(kv, "=")
				if len(stanzas) == 2 {
					charFrom := stanzas[0]
					charTo := stanzas[1]
					k, err := strconv.ParseUint(charFrom, 16, 8)
					if err != nil {
						log.Println(err)
					}
					v := []byte(charTo)
					charMap[byte(k)] = v[0]
				}
			}

			rom := filepath.Clean(c.Args().Get(0))
			romPath := filepath.Join(cwd, rom)
			data, err := os.ReadFile(romPath)
			if err != nil {
				log.Println(err)
			}

			toCharFn := func(b byte) byte {
				if char, ok := charMap[b]; ok {
					return char
				} else {
					return '.'
				}
			}

			fmt.Print(hex.Dump(data, toCharFn))

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
