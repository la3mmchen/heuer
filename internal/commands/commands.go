package commands

import (
	"fmt"
	"log"

	"github.com/adlio/trello"
	"github.com/la3mmchen/treta/internal/types"
	"github.com/urfave/cli"
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
						list, err := client.GetList(x.ID, trello.Defaults())
						if err != nil {
							log.Fatal(err)
						}

						cards, err := list.GetCards(trello.Defaults())
						if err != nil {
							log.Fatal(err)
						}

						for _, y := range cards {
							fmt.Printf("card %v ; %v ; %v \n", y.Name, y.Desc, y.Due)
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
