package wordnik

type WordObject struct {
	ID            int64    `json:"id"`
	Word          string   `json:"word"`
	OriginalWord  string   `json:"originalWord"`
	Suggestions   []string `json:"suggestions"`
	CanonicalForm string   `json:"canonicalForm"`
	Vulgar        string   `json:"vulgar"`
}
