/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/progress"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		success := make(chan bool, 1)
		finished := Progress(success, "Exporting data", time.Second*5)
		t := table.NewWriter()
		t.AppendHeader(table.Row{"Blockchain", "Network"})
		for _, blockchain := range config.Blockchains {
			t.AppendRow(table.Row{blockchain.Name, strings.Join(blockchain.Networks, "\n")})
			t.AppendSeparator()
		}
		os.WriteFile("./output.html", []byte(t.RenderHTML()), os.ModePerm)
		os.WriteFile("./output.csv", []byte(t.RenderCSV()), os.ModePerm)
		success <- true
		<-finished
	},
}

func init() {
	blockchainCmd.AddCommand(exportCmd)
}

func Progress(success chan bool, message string, minSleep time.Duration) (finished chan bool) {
	pw := progress.NewWriter()
	total := 100
	incrementPerCycle := 1
	var units *progress.Units = &progress.Units{
		Notation: "",
		Formatter: func(value int64) string {
			return ""
		},
	}
	pw.SetAutoStop(false)
	pw.ShowTime(true)
	pw.ShowTracker(false)
	pw.ShowValue(false)
	pw.SetMessageWidth(40)
	pw.SetNumTrackersExpected(1)
	pw.SetStyle(progress.StyleDefault)
	pw.SetTrackerPosition(progress.PositionRight)
	pw.SetUpdateFrequency(time.Millisecond * 100)
	pw.Style().Colors = progress.StyleColorsExample
	pw.Style().Options.PercentFormat = "%4.1f%%"
	go pw.Render()

	tracker := progress.Tracker{Message: message, Total: int64(total), Units: *units}

	pw.AppendTracker(&tracker)
	finished = make(chan bool, 1)
	go func() {
		for !tracker.IsDone() {
			if tracker.PercentDone() < 99 {
				tracker.Increment(int64(incrementPerCycle))
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		isSuccess := <-success
		if isSuccess {
			tracker.MarkAsDone()
			for pw.LengthDone() != 1 {
				time.Sleep(time.Millisecond * 100)
			}
			pw.Stop()
		} else {
			pw.Stop()
		}
		for pw.IsRenderInProgress() {
			time.Sleep(time.Millisecond * 100)
		}
		finished <- true
	}()
	time.Sleep(minSleep)
	return
}
