package types

type Topic struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
	ThreadID          int    `json:"message_thread_id"`
}

type (
	TopicCreated struct{ Topic }

	TopicClosed struct{}

	TopicDeleted struct{ Topic }

	TopicReopened struct{ Topic }

	TopicEdited struct{ Topic }

	GeneralTopicHidden struct{}

	GeneralTopicUnhidden struct{}
)