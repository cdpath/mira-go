package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yourusername/mira-go/internal/mira"
)

var debug bool

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mira",
		Short: "Mira is a tool for controlling Boox Mira e-ink monitors",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Skip device check for version and help commands
			if cmd.Name() == "help" || cmd.Name() == "version" {
				return
			}
			if debug {
				log.SetFlags(log.Ltime | log.Lmicroseconds)
				log.Println("Debug mode enabled")
			}
		},
	}

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")

	var refreshCmd = &cobra.Command{
		Use:   "refresh",
		Short: "Refresh the screen",
		RunE: func(cmd *cobra.Command, args []string) error {
			device, err := mira.NewDevice()
			if err != nil {
				return fmt.Errorf("no Mira device found: %w", err)
			}
			defer device.Close()
			return device.Refresh()
		},
	}

	var antishakeCmd = &cobra.Command{
		Use:   "antishake",
		Short: "Anti-shake automatically",
		RunE: func(cmd *cobra.Command, args []string) error {
			device, err := mira.NewDevice()
			if err != nil {
				return fmt.Errorf("no Mira device found: %w", err)
			}
			defer device.Close()
			return device.SetAutoDitherMode(mira.AutoDitherModeHigh)
		},
	}

	var (
		speed       int
		contrast    int
		refreshMode string
		ditherMode  int
		blackFilter int
		whiteFilter int
		coldLight   int
		warmLight   int
	)

	var settingsCmd = &cobra.Command{
		Use:   "settings",
		Short: "Apply settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			device, err := mira.NewDevice()
			if err != nil {
				return fmt.Errorf("no Mira device found: %w", err)
			}
			defer device.Close()

			if cmd.Flags().Changed("refresh-mode") {
				var mode mira.RefreshMode
				switch refreshMode {
				case "a2":
					mode = mira.RefreshModeA2
				case "direct":
					mode = mira.RefreshModeDirectUpdate
				case "gray":
					mode = mira.RefreshModeGrayUpdate
				default:
					return fmt.Errorf("invalid refresh mode")
				}
				if err := device.SetRefreshMode(mode); err != nil {
					return err
				}
			}

			if cmd.Flags().Changed("speed") {
				if err := device.SetSpeed(speed); err != nil {
					return err
				}
			}

			if cmd.Flags().Changed("contrast") {
				if err := device.SetContrast(contrast); err != nil {
					return err
				}
			}

			if cmd.Flags().Changed("dither-mode") {
				if err := device.SetDitherMode(ditherMode); err != nil {
					return err
				}
			}

			if cmd.Flags().Changed("black-filter") || cmd.Flags().Changed("white-filter") {
				if err := device.SetColorFilter(whiteFilter, blackFilter); err != nil {
					return err
				}
			}

			if cmd.Flags().Changed("cold-light") {
				if err := device.SetColdLight(coldLight); err != nil {
					return err
				}
			}

			if cmd.Flags().Changed("warm-light") {
				if err := device.SetWarmLight(warmLight); err != nil {
					return err
				}
			}

			return nil
		},
	}

	settingsCmd.Flags().IntVar(&speed, "speed", 0, "The refresh speed (1-7)")
	settingsCmd.Flags().IntVar(&contrast, "contrast", 0, "The contrast (0-15)")
	settingsCmd.Flags().StringVar(&refreshMode, "refresh-mode", "", "The refresh mode (a2, direct, gray)")
	settingsCmd.Flags().IntVar(&ditherMode, "dither-mode", 0, "The dither mode (0-3)")
	settingsCmd.Flags().IntVar(&blackFilter, "black-filter", 0, "The black filter level (0-254)")
	settingsCmd.Flags().IntVar(&whiteFilter, "white-filter", 0, "The white filter level (0-254)")
	settingsCmd.Flags().IntVar(&coldLight, "cold-light", 0, "The cold backlight level (0-254)")
	settingsCmd.Flags().IntVar(&warmLight, "warm-light", 0, "The warm backlight level (0-254)")

	rootCmd.AddCommand(refreshCmd, antishakeCmd, settingsCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
