package cmd

import (
	"fmt"
	debg "runtime/debug"

	"github.com/spf13/cobra"
)

var (
	debug  bool
	silent bool

	// Version stores the build version of VHS at the time of package through -ldflags.
	//
	// go build -ldflags "-s -w -X=main.Version=$(VERSION)"
	Version string

	// CommitSHA stores the git SHA at the time of package through -ldflags.
	CommitSHA string
)

// Entry point for the root command.
var RootCmd = &cobra.Command{
	Use:   "go-jwt",
	Short: "Decode and validate your JWT tokens with ease.",
	Long: `Decode and validate your JWT tokens with ease.

This is a work in progress. The CLI is not ready yet.

To check the version of go-jwt, run:

	go-jwt version
`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) == 0 {
			fmt.Println("Error: No command specified. Use 'go-jwt --help' for usage.")
			err = cmd.Help()
		} else {
			if args[0] == "version" {
				if len(CommitSHA) > 7 {
					Version = fmt.Sprintf("commit=%s", CommitSHA[:7])
				}
				if Version != "" {
					if info, ok := debg.ReadBuildInfo(); ok && info.Main.Version != "" {
						Version = info.Main.Version
					} else {
						Version = "unknown (built from source)"
					}
				} else {
					Version = "unknown"
				}
				fmt.Println(Version)
				return nil
			}
			err = fmt.Errorf("Unknown command %q", args[0])
		}
		return err
	},
}

// Execute executes the root command.
func init() {
	RootCmd.PersistentFlags().BoolVarP(&debug, "verbose", "V", false, "Enable verbose output")
	RootCmd.PersistentFlags().BoolVarP(&silent, "quiet", "q", false, "Enable silent output")
}
