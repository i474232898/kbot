package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var kbotCmd = &cobra.Command{
	Use:   "kbot",
	Short: "Lorem Ipsum",
	Long:  `Lorem ipsum dolor sit amet, 
		consectetur adipiscing elit`,
	Aliases: []string{"start"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world telegram bots")
	},
}

func init(){
	rootCmd.AddCommand(kbotCmd)
}
