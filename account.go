package wordnik

// AuthenticationToken as defined by the Wordnik API. Needed for user-specific
// requests.
type AuthenticationToken struct {
	Token         string `json:"token"`
	UserID        int64  `json:"userId"`
	UserSignature string `json:"userSignature"`
}

// APITokenStatus as defined by the Wordnik API.
type APITokenStatus struct {
	Valid           bool   `json:"valid"`
	Token           string `json:"token"`
	ResetsInMillis  int64  `json:"resetsInMillis"`
	RemainingCalls  int64  `json:"remainingCalls"`
	ExpiresInMillis int64  `json:"expiresInMillis"`
	TotalRequests   int64  `json:"totalRequests"`
}

// User as defined by the Wordnik API.
type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Status      int64  `json:"status"`
	FaceBookID  string `json:"faceBookId"`
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
	Password    string `json:"password"`
}
