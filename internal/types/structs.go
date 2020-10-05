package types

// Configuration <tbd>
type Configuration struct {
	TrelloToken    string
	TrelloAppKey   string
	TrelloUserName string
	TrelloBoard    string
	TrelloList     []string
	Debug          string
}

// WantList is a mapping from string to empty struct to describe the set of wanted strings.
// We use an empty struct instead of a boolean since this saves us 1 byte of memory per value :-D
// https://play.golang.org/p/fc14T5vr_xS
type WantList map[string]struct{}

func NewWantList(values []string) WantList {
	w := make(WantList, len(values))
	for _, s := range values {
		w[s] = struct{}{}
	}
	return w
}

// IsWanted returns true if the given string is in the WantList
// We use the multi return value form of getting a value from a dict
// The second return value is a boolean indicating if the key is in the map.
func (w WantList) IsWanted(s string) bool {
	_, found := w[s]
	return found
}
