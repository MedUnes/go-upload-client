package main

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/Myra-Security-GmbH/go-upload-client/api"

	arg "github.com/alexflint/go-arg"
	config "github.com/Myra-Security-GmbH/go-upload-client/config"
)

var args struct {
	Source    string `arg:"positional,required,help:Source file or folder"`
	Target    string `arg:"positional,required,help:Target folder"`
	Init      bool   `arg:"--init,help:Creates a default configuration file"`
	Domain    string `arg:"-d,required,help:Domain"`
	Bucket    string `arg:"-b,required,help:Bucket"`
	Recursive bool   `arg:"-r,help:Upload folder recursive"`
	Silent    bool   `arg:"-s,help:No progress output"`

	Configfile string `arg:"-c,help:Configfile to use"`
}

//
// main
//
func main() {
	homePath, homeExist := os.LookupEnv("HOME")

	if !homeExist {
		fmt.Println("HOME environment variable is not set.")
	}

	_, e := os.Stat(homePath + "/config.yml")
	_, e2 := os.Stat("./config.yml")

	if homePath != "" && e == nil {
		args.Configfile = homePath + "/config.yml"
	} else if e2 == nil {
		args.Configfile = "./config.yml"
	}

	arg.MustParse(&args)

	if args.Init {
		_, err := os.Stat(args.Configfile)

		if err == nil {
			fmt.Println("File already exists")

			os.Exit(1)
		}

		fmt.Println("Creating config file")
		cfg := &config.Config{
			Endpoint: "http://upload.myracloud.com",
			Proxy:    "",
			Login: config.User{
				APIKey: "apiKey",
				Secret: "secret",
				User:   "username",
			},
		}

		err = config.SaveConfigFile(args.Configfile, cfg)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Config file successfully created")
		fmt.Println(": " + args.Configfile)

		os.Exit(0)
	}

	cfg, err := config.ReadConfigFile(args.Configfile)

	if err != nil {
		fmt.Printf("Error reading config file [%s]", args.Configfile)
		fmt.Println(err)
		os.Exit(1)
	}

	up, err := api.NewAPI(
		cfg.Login.APIKey,
		cfg.Login.Secret,
		cfg.Endpoint,
		cfg.Proxy,
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !args.Recursive {
		fi, err := os.Stat(args.Source)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if !fi.Mode().IsRegular() {
			fmt.Println("Given path is not a regular file")
			os.Exit(1)
		}

		fp, err := os.Open(args.Source)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		mimeType := mime.TypeByExtension(filepath.Ext(args.Source))

		buildPath := args.Target

		if args.Target[len(args.Target)-1:] == "/" {
			buildPath += filepath.Base(args.Source)
		}

		err = up.UploadFile(args.Domain, args.Bucket, buildPath, mimeType, fp)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	err = filepath.Walk(args.Source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if info.IsDir() {
			return nil
		}

		relativeTargetPath := "/" + strings.Trim(args.Target, "/") + "/" + strings.Trim(
			strings.Replace(path, args.Source, "", 1), "/",
		)

		fp, err := os.Open(path)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		mimeType := mime.TypeByExtension(filepath.Ext(path))

		err = up.UploadFile(args.Domain, args.Bucket, relativeTargetPath, mimeType, fp)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if !args.Silent {
			fmt.Printf("%s => %s [mime=%s]\n", path, relativeTargetPath, mimeType)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
