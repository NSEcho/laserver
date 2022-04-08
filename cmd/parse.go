package cmd

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/lateralusd/laserver/db"
	"github.com/spf13/cobra"
)

type Target struct {
	Name  string
	Email string
	URL   string
}

type Attack struct {
	URL     string
	Targets []Target
}

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse lateralus json file",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("in.json")
		if err != nil {
			return err
		}
		defer f.Close()

		var a Attack
		if err := json.NewDecoder(f).Decode(&a); err != nil {
			return err
		}

		db := db.NewDB("data.db")

		return renderTable(db, &a)
	},
}

func renderTable(db *db.DB, atk *Attack) error {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredYellowWhiteOnBlack)

	t.AppendHeader(table.Row{"Name", "Email", "UUID", "Opened"})

	for _, tgt := range atk.Targets {
		uuid := extractUUID(atk.URL, tgt.URL)
		found, err := db.Exists(uuid)
		if err != nil {
			return err
		}
		opened := "NO"
		if found {
			opened = "YES"
		}
		t.AppendRow(table.Row{tgt.Name, tgt.Email, uuid, opened})
	}
	t.SetTitle("Stats")
	t.Style().Title.Align = text.AlignCenter
	t.Render()
	return nil
}

func extractUUID(mainUrl, targetUrl string) string {
	idx := strings.Index(mainUrl, "<CHANGE>")
	if idx < len(targetUrl) {
		return targetUrl[idx:]
	}
	return ""
}

func init() {
	RootCmd.AddCommand(parseCmd)
}
