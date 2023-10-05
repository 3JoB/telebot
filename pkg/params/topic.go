package params

type (
	Topic struct {
		ChatID string `json:"chat_id"`
		ID     int    `json:"message_thread_id"`
	}
)
