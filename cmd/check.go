package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gouted/pkg/checker"
	"io/ioutil"
	"log"
)

func CheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "check [path]",
		Args: cobra.MinimumNArgs(1),
		Run:  handleCheckCmd,
	}
	return cmd
}

func handleCheckCmd(cmd *cobra.Command, args []string) {
	content, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	packages, err := checker.Check(content)
	if err != nil {
		log.Fatal(err)
	}

	for _, pkg := range packages {
		fmt.Printf("%s\t%s → %s\n", pkg.Name, pkg.CurrentVersion, pkg.LatestVersion)
	}
}
