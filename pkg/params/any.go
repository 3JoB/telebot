package params

type (
	Any struct {
		Name string `json:"name"`
	}

	OnlyID struct {
		ChatID string `json:"chat_id"`
	}
)
