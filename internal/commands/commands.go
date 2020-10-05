package commands

import (
	"fmt"
	"log"

	"github.com/xeonx/timeago"

	"github.com/adlio/trello"
	"github.com/la3mmchen/heuer/internal/types"
	"github.com/urfave/cli"
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

		wanted := types.NewWantList(cfg.TrelloList)
		lists, err := getListsFromBoard(client, cfg)
		if err != nil {
			log.Fatal(err)
		}

		found := renderLists(client, lists, wanted)
		if !found {
			fmt.Printf("Sorry, the wanted lists are not found in your trello account.. \n")
		}
		return nil
	}
	return cmd
}

// getListsFromBoard returns the lists of the specified trello board
func getListsFromBoard(client *trello.Client, cfg types.Configuration) ([]*trello.List, error) {
	member, err := client.GetMember(cfg.TrelloUserName, trello.Defaults())
	if err != nil {
		// Use the %w in the format string. This is a new feature introduced in go 1.13 to wrap errors.
		// See https://blog.golang.org/go1.13-errors for details
		return nil, fmt.Errorf("failed to get trello member struct: %w", err)
	}

	boards, err := member.GetBoards(trello.Defaults())
	if err != nil {
		return nil, fmt.Errorf("failed to list all trello boards: %w", err)
	}
	var boardID string
	for _, v := range boards {
		if v.Name == cfg.TrelloBoard {
			boardID = v.ID
			break
		}
	}
	if boardID == "" {
		return nil, fmt.Errorf("did not find the specified trello board in your account")
	}
	board, err := client.GetBoard(boardID, trello.Defaults())
	if err != nil {
		return nil, fmt.Errorf("failed to get board details: %w ", err)
	}

	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return nil, fmt.Errorf("failed to get lists of the specified trello board: %w", err)
	}
	return lists, nil
}

// renderLists renders all wanted trello lists to stdout
func renderLists(client *trello.Client, lists []*trello.List, wanted types.WantList) bool {
	var found bool = false
	for _, x := range lists {
		if wanted.IsWanted(x.Name) {
			found = true
			fmt.Printf("cards in [%v]: \n", x.Name)
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
	return found
}
