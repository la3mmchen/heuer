# heuer

a command line cli to display one list of trello. I use it to access a list called `today` so i can see what I've got to do today.

## config & usage

First you need to setup a personal token to access your account via the trello API. Read at the trello documentation <https://trello.com/app-key>.

You should end with an AppKey and a Token. Save them to an config file in your home folder (`~/.heuer/config.json`) or at the directory you execute the binary.

The config should look like this: 

```test
{
    "TrelloUserName": "<your username>", // username is needed to retrieve your boards */
    "TrelloBoard": "<board name>", // the board in which to search for the list; your user needs permission on this board
    "TrelloAppKey": "...",
    "TrelloToken": "..."
}
```

Build a go binary:

```bash
go build -o bin/heuer
```

And run it:

```bash
./bin/heuer read -l today -l next
```
