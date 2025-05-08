package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v3"
)

var threadUnsafeCounter = 0

var kbotCmd = &cobra.Command{
	Use:   "kbot",
	Short: "Lorem Ipsum",
	Long: `Lorem ipsum dolor sit amet, 
		consectetur adipiscing elit`,
	Aliases: []string{"start"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")

		pref := tele.Settings{
			Token:  os.Getenv("TELE_TOKEN"),
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		}

		b, err := tele.NewBot(pref)
		if err != nil {
			fmt.Printf("Error creating bot: %v\n", err)
			return
		}

		b.Handle(tele.OnText, func(m tele.Context) error {
			log.Printf("Received message: %s", m.Text())
			payload := m.Message().Payload

			switch payload {
			case "hello":
				msg := m.Send(fmt.Sprintf("hello world telegram bots. counter: %d", threadUnsafeCounter))
				threadUnsafeCounter++
				return msg
			default:
				return m.Send("Usage: /s hello")
			}
		})

		b.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	if _, ok := os.LookupEnv("TELE_TOKEN"); !ok {
		panic("no token found")
	}
}
