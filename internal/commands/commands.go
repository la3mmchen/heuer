package commands

import (
	"fmt"
	"log"

	"github.com/adlio/trello"
	"github.com/la3mmchen/treta/internal/types"
	"github.com/urfave/cli"
	"github.com/xeonx/timeago"
)

const (
	indent = " \\_ "
)

func read(cfg types.Configuration) cli.Command {
	cmd := cli.Command{
		Name:  "read",
		Usage: "Read all cards in configured list.",
	}

	cmd.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "list, l",
			Usage: "Trello lists to obtain cards from. Can be repeated for multiple lists.",
			//Value: &cli.StringSlice{"tomorrow", "today"},
		},
	}

	cmd.Action = func(c *cli.Context) error {
		// do not do anything without any lists
		if len(c.StringSlice("list")) == 0 {
			fmt.Printf("Please provide name of lists you want to read cards of. \n")
			return nil
		}

		// read from user provided flags
		cfg.TrelloList = c.StringSlice("list")

		// client object to interact with trello
		client := trello.NewClient(cfg.TrelloAppKey, cfg.TrelloToken)

		member, err := client.GetMember(cfg.TrelloUserName, trello.Defaults())
		if err != nil {
			log.Fatal(err)
		}
		boards, err := member.GetBoards(trello.Defaults())
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range boards {
			if v.Name == cfg.TrelloBoard {

				board, err := client.GetBoard(v.ID, trello.Defaults())
				if err != nil {
					log.Fatal(err)
				}

				lists, err := board.GetLists(trello.Defaults())
				if err != nil {
					log.Fatal(err)
				}
				var found bool = false
				for _, x := range lists {
					for _, wantedList := range cfg.TrelloList {
						if x.Name == wantedList {
							found = true
							fmt.Printf("cards in [%v]: \n", wantedList)
							list, err := client.GetList(x.ID, trello.Defaults())
							if err != nil {
								log.Fatal(err)
							}

							cards, err := list.GetCards(trello.Defaults())
							if err != nil {
								log.Fatal(err)
							}

							for _, y := range cards {
								var out string
								if y.DueComplete {
									out += "[done] "
								} else {
									out += "[    ] "
								}
								out += y.Name

								if y.Due != nil {
									out += "\n" + indent + "Due: " + timeago.English.Format(*y.Due)
								}

								if len(y.Desc) > 0 {
									out += "\n" + indent + "Description: " + y.Desc
								}
								fmt.Printf("%v \n", out)
								fmt.Printf("\n")
								out = " "

							}
							fmt.Printf("___ \n")
						}
					}
				}
				if !found {
					fmt.Printf("Sorry, the wanted lists are not found in your trello account.. \n")
				}
			}
		}
		return nil
	}
	return cmd
}
