package params

type Updates struct {
	Offset         int      `json:"offset"`
	Timeout        int      `json:"timeout"`
	Limit          int      `json:"limit,omitempty"`
	AllowedUpdates []string `json:"allowed_updates"`
}
