package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/coreos/go-semver/semver"
	"github.com/urfave/cli"
)

func run(versionArg string) (versions []*semver.Version, versionIdx int) {
	scanner := bufio.NewScanner(os.Stdin)
	actualVersion := semver.New(versionArg)
	
	versions = append(versions, actualVersion)

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		v := semver.New(scanner.Text())
		versions = append(versions, v)
	}

	if scanner.Err() != nil {
		// handle error.
		panic("scan fail")
	}

	semver.Sort(versions)

	versionIdx = indexOf(versions, actualVersion)

	if versionIdx< 0 {
		panic("oh no! how could that even happen?")
	}
	
	return
}

func indexOf(versions []*semver.Version, v *semver.Version) int {
	for idx, s := range versions {
		if s == v {
			return idx
		}
	}
	return -1
}

func major(versions []*semver.Version, versionIdx int) string {
	v := versions[versionIdx]
	if versionIdx == len(versions)-1 {
		return fmt.Sprintf("%v",v.Major)
	}
	nextV := versions[versionIdx+1]
	if nextV.Major>v.Major {
		return fmt.Sprintf("%v",v.Minor)
	}
	// in other cases we shouldnt release major
	// no major version applies
	return ""
}

func minor(versions []*semver.Version, versionIdx int) string {
	v := versions[versionIdx]
	if versionIdx == len(versions)-1 {
		return fmt.Sprintf("%v.%v",v.Major,v.Minor)
	}
	nextV := versions[versionIdx+1]
	if nextV.Major>v.Major || nextV.Minor > v.Minor {
		return fmt.Sprintf("%v.%v",v.Major,v.Minor)
	}

	// no minor version applies

	return ""
} 

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"

	app.Commands = []cli.Command{
		{
			Name:  "major",
			Usage: "gets the major if the major is bigger",
			Action: func(c *cli.Context) error {
				fmt.Println("major", c.Args().First())
				fmt.Println(major(run(c.Args().First())))
				return nil
			},
		},
		{
			Name:  "minor",
			Usage: "gets the minor if the minor is bigger",
			Action: func(c *cli.Context) error {
				fmt.Println("minor", c.Args().First())
				fmt.Println(minor(run(c.Args().First())))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
