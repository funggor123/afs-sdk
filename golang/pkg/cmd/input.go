package cmd

import (
	"os"
	"github.com/urfave/cli"
)

type (
    // Config information.
    Config struct {
        Address   string
		FilePath  string
		FileName  string
		Field 	  string
		Method    int
		ExpDays   int
		BlockSize int
		Action    string
    }
)
 
var Configs Config
 
func Main() {
    app := cli.NewApp()
    app.Name = "AFS SDK"
	app.Usage = "Interface to AFS"
	app.Version = "1.0.0"
    app.Action = run
    app.Flags = []cli.Flag{
		cli.StringFlag {
			Name: "action, a",
			Value: "upload",
			Usage: "action to perform (upload / download)",
		},
		cli.StringFlag {
			Name: "address, addr",
			Value: "http://39.108.80.53:8074",
			Usage: "upload node address",
		},
		cli.StringFlag {
			Name: "filepath, fp",
			Value: "./",
			Usage: "upload file path",
		},
		cli.StringFlag {
			Name: "filename, fn",
			Value: "48_FL1_Prog.pdf",
			Usage: "upload file name",
		},
		cli.IntFlag {
			Name: "method, m",
			Value: 0,
			Usage: "0 seed upload, 1 raw upload ",
		},
		cli.StringFlag {
			Name: "field, f",
			Value: "",
			Usage: "arfs / afs / empty for upload both ",
		},
		cli.IntFlag {
			Name: "expdays, exp",
			Value: 7,
			Usage: "expire days ",
		},
		cli.IntFlag {
			Name: "size, s",
			Value: 1,
			Usage: "block size (In MB) ",
		},
    }
 
    app.Run(os.Args)
}
 
func run(c *cli.Context) error {
    Configs = Config{
		Action:   c.String("action"),
        Address:  c.String("address"),
		FilePath: c.String("filepath"),
		FileName: c.String("filename"),
		Field:    c.String("field"),
		Method:   c.Int("method"),
		ExpDays:  c.Int("expdays"), 
		BlockSize: c.Int("size"), 
    }
    return exec()
}
 
func exec() error {
    return nil
}