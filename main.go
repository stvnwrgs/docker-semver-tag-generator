package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/coreos/go-semver/semver"
	"github.com/urfave/cli"
)

func run(versionArg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	actualVersion := semver.New(versionArg)
	fmt.Println("Actual Version", actualVersion)
	var versions []*semver.Version

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		v := semver.New(scanner.Text())
		versions = append(versions, v)
	}

	semver.Sort(versions)
	for _, s := range versions {
		var latest *semver.Version
		cmp := s.Compare(*actualVersion)
		if cmp == -1 {
			fmt.Println("less", s)

		} else if cmp == 1 {
			fmt.Println("latest is", latest)
			if latest != nil {
				fmt.Println("latest is", latest)
				if latest.Compare(*actualVersion) >= 0 {
					fmt.Println("winner ", actualVersion)
				}
			}
			fmt.Println("bigger", s)
		} else {
			fmt.Println("same", s)
		}
		latest = s
		fmt.Println("latest set to", latest)
	}

	if scanner.Err() != nil {
		// handle error.
	}
	return "asd"
}

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"

	app.Commands = []*cli.Command{
		{
			Name:  "major",
			Usage: "gets the major if the major is bigger",
			Action: func(c *cli.Context) error {
				fmt.Println("major")
				fmt.Println(run(c.Args().First()))
				return nil
			},
		},
		{
			Name:  "minor",
			Usage: "gets the minor if the minor is bigger",
			Action: func(c *cli.Context) error {
				fmt.Println("minor")
				fmt.Println(run(c.Args().First()))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
