package cmd

import (
	"fmt"
	"github.com/imwithye/gitignore/template"
	_ "github.com/imwithye/gitignore/template/all"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

var addCmdQuiet, addCmdVerbose bool
var addCmd = &cobra.Command{
	Use:   "add [templates]",
	Short: "Add multi gitignore templates to .gitignore",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
		ignores := []string{}
		for _, arg := range args {
			t, err := template.Get(arg)
			if err != nil {
				if addCmdQuiet {
					continue
				}
				fmt.Println("*", err)
				continue
			}
			if !addCmdQuiet && addCmdVerbose {
				fmt.Println("-", t.Path)
			}
			ignores = append(ignores, t.Ignore)
		}
		if len(ignores) > 0 {
			ioutil.WriteFile(".gitignore", []byte(strings.Join(ignores, "\n")), 0644)
		} else if !addCmdQuiet {
			fmt.Println("Nothing will be added.")
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVarP(&addCmdQuiet, "quiet", "q", false, "quiet mode")
	addCmd.Flags().BoolVarP(&addCmdVerbose, "verbose", "v", false, "verbose mode")
}
