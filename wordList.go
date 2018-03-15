package wordnik

// WordLists is a container type for WordList, for the purpose of
// unmarshalling JSON.
type WordLists struct {
	WordLists []WordList
}

// WordList as defined by the Wordnik API.
type WordList struct {
	ID                int64  `json:"id,omitempty"`
	Permalink         string `json:"permalink,omitempty"`
	Name              string `json:"name,omitempty"`
	CreatedAt         string `json:"createdAt,omitempty"`
	UpdatedAt         string `json:"updatedAt,omitempty"`
	LastActivityAt    string `json:"lastActivityAt,omitempty"`
	Username          string `json:"username,omitempty"`
	UserID            int64  `json:"userId,omitempty"`
	Description       string `json:"description,omitempty"`
	NumberWordsInList int64  `json:"numberWordsInList,omitempty"`
	Type              string `json:"type,omitempty"`
}

// stringValue is provided for convenient JSON marshalling in PostWordListWords.
type stringValue struct {
	Word string `json:"word,omitempty"`
}

// WordListWords is a container type for WordListWord, for the purpose of
// unmarshalling JSON.
type WordListWords struct {
	WordListWords []WordListWord
}

// WordListWord as defined by the Wordnik API.
type WordListWord struct {
	ID                   int64  `json:"id"`
	Word                 string `json:"word"`
	Username             string `json:"username"`
	UserID               int64  `json:"userId"`
	CreatedAt            string `json:"createdAt"`
	NumberCommentsOnWord int64  `json:"numberCommentsOnWord"`
	NumberLists          int64  `json:"numberLists"`
}
