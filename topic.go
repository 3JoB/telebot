package telebot

import (
	"strconv"
)

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
	params := map[string]string{
		"chat_id": chat.Recipient(),
		"name":    forum.Name,
	}

	if forum.IconColor != 0 {
		params["icon_color"] = strconv.Itoa(forum.IconColor)
	}
	if forum.IconCustomEmojiID != "" {
		params["icon_custom_emoji_id"] = forum.IconCustomEmojiID
	}

	_, err := b.Raw("createForumTopic", params)
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

	_, err := b.Raw("editForumTopic", params)
	return err
}

// CloseTopic closes an open topic in a forum supergroup chat.
func (b *Bot) CloseTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id":           chat.Recipient(),
		"message_thread_id": forum.ThreadID,
	}

	_, err := b.Raw("closeForumTopic", params)
	return err
}

// ReopenTopic reopens a closed topic in a forum supergroup chat.
func (b *Bot) ReopenTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id":           chat.Recipient(),
		"message_thread_id": forum.ThreadID,
	}

	_, err := b.Raw("reopenForumTopic", params)
	return err
}

// DeleteTopic deletes a forum topic along with all its messages in a forum supergroup chat.
func (b *Bot) DeleteTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id":           chat.Recipient(),
		"message_thread_id": forum.ThreadID,
	}

	_, err := b.Raw("deleteForumTopic", params)
	return err
}

// UnpinAllTopicMessages clears the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup.
func (b *Bot) UnpinAllTopicMessages(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id":           chat.Recipient(),
		"message_thread_id": forum.ThreadID,
	}

	_, err := b.Raw("unpinAllForumTopicMessages", params)
	return err
}

// TopicIconStickers gets custom emoji stickers, which can be used as a forum topic icon by any user.
func (b *Bot) TopicIconStickers() ([]Sticker, error) {
	params := map[string]string{}

	data, err := b.Raw("getForumTopicIconStickers", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result []Sticker
	}
	if err := b.json.Unmarshal(data, &resp); err != nil {
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

	_, err := b.Raw("editGeneralForumTopic", params)
	return err
}

// CloseGeneralTopic closes an open 'General' topic in a forum supergroup chat.
func (b *Bot) CloseGeneralTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id": chat.Recipient(),
	}

	_, err := b.Raw("closeGeneralForumTopic", params)
	return err
}

// ReopenGeneralTopic reopens a closed 'General' topic in a forum supergroup chat.
func (b *Bot) ReopenGeneralTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id": chat.Recipient(),
	}

	_, err := b.Raw("reopenGeneralForumTopic", params)
	return err
}

// HideGeneralTopic hides the 'General' topic in a forum supergroup chat.
func (b *Bot) HideGeneralTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id": chat.Recipient(),
	}

	_, err := b.Raw("hideGeneralForumTopic", params)
	return err
}

// UnhideGeneralTopic unhides the 'General' topic in a forum supergroup chat.
func (b *Bot) UnhideGeneralTopic(chat *Chat, forum *Topic) error {
	params := map[string]any{
		"chat_id": chat.Recipient(),
	}

	_, err := b.Raw("unhideGeneralForumTopic", params)
	return err
}

// Use this method to clear the list of pinned messages in a General forum topic.
// The bot must be an administrator in the chat for this to work and must have the
// can_pin_messages administrator right in the supergroup.
func (b *Bot) UnpinAllGeneralForumTopicMessages(chat *Chat) error {
	params := map[string]any{
		"chat_id": chat.Recipient(),
	}

	_, err := b.Raw("unpinAllGeneralForumTopicMessages", params)
	return err
}
