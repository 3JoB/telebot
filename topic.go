package telebot

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

// CreateTopic creates a topic in a forum supergroup chat.
func (b *Bot) CreateTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id": chat.Recipient(),
		"name":    forum.Name,
	}

	if forum.IconColor != 0 {
		params["icon_color"] = forum.IconColor
	}
	if forum.IconCustomEmojiID != "" {
		params["icon_custom_emoji_id"] = forum.IconCustomEmojiID
	}

	data, err := Raw(b, "createForumTopic", params)
	ReleaseBuffer(data)
	return err
}

// EditTopic edits name and icon of a topic in a forum supergroup chat.
func (b *Bot) EditTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id":           chat.Recipient(),
		"message_thread_id": forum.ThreadID,
	}

	if forum.Name != "" {
		params["name"] = forum.Name
	}
	if forum.IconCustomEmojiID != "" {
		params["icon_custom_emoji_id"] = forum.IconCustomEmojiID
	}

	data, err := Raw(b, "editForumTopic", params)
	ReleaseBuffer(data)
	return err
}

// CloseTopic closes an open topic in a forum supergroup chat.
func (b *Bot) CloseTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id":           chat.Recipient(),
		"message_thread_id": forum.ThreadID,
	}

	data, err := Raw(b, "closeForumTopic", params)
	ReleaseBuffer(data)
	return err
}

// ReopenTopic reopens a closed topic in a forum supergroup chat.
func (b *Bot) ReopenTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id":           chat.Recipient(),
		"message_thread_id": forum.ThreadID,
	}

	data, err := Raw(b, "reopenForumTopic", params)
	ReleaseBuffer(data)
	return err
}

// DeleteTopic deletes a forum topic along with all its messages in a forum supergroup chat.
func (b *Bot) DeleteTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id":           chat.Recipient(),
		"message_thread_id": forum.ThreadID,
	}

	data, err := Raw(b, "deleteForumTopic", params)
	ReleaseBuffer(data)
	return err
}

// UnpinAllTopicMessages clears the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup.
func (b *Bot) UnpinAllTopicMessages(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id":           chat.Recipient(),
		"message_thread_id": forum.ThreadID,
	}

	data, err := Raw(b, "unpinAllForumTopicMessages", params)
	ReleaseBuffer(data)
	return err
}

// TopicIconStickers gets custom emoji stickers, which can be used as a forum topic icon by any user.
func (b *Bot) TopicIconStickers() ([]Sticker, error) {
	params := make(map[string]any)

	data, err := Raw(b, "getForumTopicIconStickers", params)
	defer ReleaseBuffer(data)
	if err != nil {
		return nil, err
	}

	var resp Response[[]Sticker]
	if err := b.json.NewDecoder(data).Decode(&resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}

// EditGeneralTopic edits name of the 'General' topic in a forum supergroup chat.
func (b *Bot) EditGeneralTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id": chat.Recipient(),
		"name":    forum.Name,
	}

	_, err := Raw(b, "editGeneralForumTopic", params)
	return err
}

// CloseGeneralTopic closes an open 'General' topic in a forum supergroup chat.
func (b *Bot) CloseGeneralTopic(chat *Chat, forum *Topic) error {
	params := map[string]string{
		"chat_id": chat.Recipient(),
	}

	_, err := Raw(b, "closeGeneralForumTopic", params)
	return err
}

// ReopenGeneralTopic reopens a closed 'General' topic in a forum supergroup chat.
func (b *Bot) ReopenGeneralTopic(chat *Chat, forum *Topic) error {
	params := map[string]string{
		"chat_id": chat.Recipient(),
	}

	_, err := Raw(b, "reopenGeneralForumTopic", params)
	return err
}

// HideGeneralTopic hides the 'General' topic in a forum supergroup chat.
func (b *Bot) HideGeneralTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id": chat.Recipient(),
	}

	_, err := Raw(b, "hideGeneralForumTopic", params)
	return err
}

// UnhideGeneralTopic unhides the 'General' topic in a forum supergroup chat.
func (b *Bot) UnhideGeneralTopic(chat *Chat, forum *Topic) error {
	params := map[string]string{
		"chat_id": chat.Recipient(),
	}

	_, err := Raw(b, "unhideGeneralForumTopic", params)
	return err
}

// Use this method to clear the list of pinned messages in a General forum topic.
// The bot must be an administrator in the chat for this to work and must have the
// can_pin_messages administrator right in the supergroup.
func (b *Bot) UnpinAllGeneralForumTopicMessages(chat *Chat) error {
	params := map[string]string{
		"chat_id": chat.Recipient(),
	}

	_, err := Raw(b, "unpinAllGeneralForumTopicMessages", params)
	return err
}
