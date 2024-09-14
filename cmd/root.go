package cmd

import (
	"fmt"
	"os"
	debg "runtime/debug"

	"github.com/apex/log"
	"github.com/spf13/cobra"
)

var (
	debug   bool
	quiet   bool
	version bool

	shaLen = 7

	// Version stores the build version of VHS at the time of package through -ldflags.
	//
	// go build -ldflags "-s -w -X=main.Version=$(VERSION)"
	Version string

	// CommitSHA stores the git SHA at the time of package through -ldflags.
	CommitSHA string
)

var rootCmd = &cobra.Command{
	Use:   "go-jwt",
	Short: "Decode and validate your JWT tokens with ease.",
	Long: `Decode and validate your JWT tokens with ease.

WARNING: This is a work in progress. The CLI is not ready yet.`,
	PersistentPreRunE: func(_ *cobra.Command, _ []string) (err error) {
		if debug && !quiet {
			log.SetLevel(log.DebugLevel)
			log.WithField("flag", "debug").Debug("debug flag on")
		}
		if quiet {
			log.SetLevel(log.ErrorLevel)
		}
		return err
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) == 0 {
			fmt.Println("Error: no command specified")
			err = cmd.Help()
		}
		return err
	},
}

// Execute executes the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "verbose", "V", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Enable silent output")
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "Show version information")
	if len(CommitSHA) > shaLen {
		vt := rootCmd.VersionTemplate()
		rootCmd.SetVersionTemplate(vt[:len(vt)-1] + " (" + CommitSHA[:shaLen] + ")\n")
	}
	if Version == "" {
		if info, ok := debg.ReadBuildInfo(); ok && info.Main.Sum != "" {
			Version = info.Main.Version
		} else {
			Version = "unknown (built from source)"
		}
	}
	rootCmd.Version = Version
}
