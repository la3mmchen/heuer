package commands

import (
	"fmt"
	"log"

	"github.com/adlio/trello"
	"github.com/la3mmchen/treta/internal/types"
	"github.com/urfave/cli"
)

const (
	indent = " \\_ "
)

func list(cfg types.Configuration) cli.Command {
	cmd := cli.Command{
		Name:  "read",
		Usage: "Read all cards in configured list.",
	}

	cmd.Action = func(c *cli.Context) error {
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

				for _, x := range lists {
					if x.Name == cfg.TrelloList {
						fmt.Printf("cards in %v: \n", cfg.TrelloList)
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
								out += "\n" + indent + "Due to: " + y.Due.Format("_2 Jan 15:04 ")
							}

							if len(y.Desc) > 0 {
								out += "\n" + indent + "Description: " + y.Desc
							}
							fmt.Printf("%v \n", out)
							fmt.Printf("\n")
							out = " "

						}

					}

				}

			}
		}

		fmt.Printf("\n")
		fmt.Printf("\n")

		fmt.Printf("\n")

		fmt.Printf("\n")

		return nil
	}
	return cmd
}

// id of board 58c28b28567dbf42d7819246
