package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

var (
	gitCommit  string
	app        = cli.NewApp()
	OutputFlag = cli.StringFlag{
		Name:  "o",
		Usage: "output to file instead of stdout",
	}
	InputFlag = cli.StringFlag{
		Name:  "i",
		Usage: "input file(s) instead of stdin",
	}
	DecodeFlag = cli.BoolFlag{
		Name:  "d",
		Usage: "decode (default is to encode)",
	}
)

func init() {
	app.Usage = "hex stream/file encoder/decoder"
	app.Flags = append(app.Flags, []cli.Flag{InputFlag, OutputFlag, DecodeFlag}...)
	app.Action = streamer
	app.Version = "0.0.1"
	app.Name = "ahex"
	app.HelpName = "ahex help"
	app.ArgsUsage = ""
	app.UsageText = ""
	app.Copyright = "Copyright 2018  aerth <aerth@riseup.net>\n   GPLv3 - https://github.com/aerth/ahex"
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Println("fatal:", err)
		os.Exit(111)
	}
}

func streamer(ctx *cli.Context) (err error) {
	var (
		input  io.Reader
		output io.Writer
	)
	if ctx.NArg() != 0 {
		return app.Run([]string{os.Args[0], "-h"})
	}
	input = os.Stdin
	if ctx.IsSet("i") {
		input, err = os.Open(ctx.String("i"))
		if err != nil {
			return err
		}
	}

	output = os.Stdout
	if ctx.IsSet("o") {
		filename := ctx.String("o")
		fmt.Println("opening", filename)
		output, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
		if err != nil {
			return err
		}
	}
	if ctx.Bool("d") {
		return streamerDecode(ctx, input, output)
	}
	return streamerEncode(ctx, input, output)
}

func streamerEncode(ctx *cli.Context, input io.Reader, output io.Writer) (err error) {
	encoder := hex.NewEncoder(output)
	_, err = io.Copy(encoder, input)
	return err
}

func streamerDecode(ctx *cli.Context, input io.Reader, output io.Writer) (err error) {
	decoder := hex.NewDecoder(input)
	_, err = io.Copy(output, decoder)
	return err
}
