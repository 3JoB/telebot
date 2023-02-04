package telebot

import (
	"github.com/3JoB/ulib/json"
)

type Forum struct {
	ID        int64  `json:"chat_id"`
	IconColor int64  `json:"icon_color"`
	ThreadID  int    `json:"message_thread_id"`
	Name      string `json:"name"`
	EmojiID   string `json:"icon_custom_emoji_id"`
}

func (c *nativeContext) NewTopic(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.Chat().ID
	}
	_, err := c.b.Raw("createForumTopic", json.Marshal(r).String())
	return err
}

func (c *nativeContext) EditTopic(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.Chat().ID
	}
	if r.ThreadID == 0 {
		r.ThreadID = c.Message().ThreadID
	}
	_, err := c.b.Raw("editForumTopic", json.Marshal(r).String())
	return err
}

func (c *nativeContext) TopicDelete(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.Chat().ID
	}
	_, err := c.b.Raw("deleteForumTopic", json.Marshal(r).String())
	return err
}

func (c *nativeContext) TopicReopen(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.Chat().ID
	}
	_, err := c.b.Raw("reopenForumTopic", json.Marshal(r).String())
	return err
}

func (c *nativeContext) TopicClose(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.Chat().ID
	}
	if r.ThreadID == 0 {
		r.ThreadID = c.Message().ThreadID
	}
	_, err := c.b.Raw("closeForumTopic", json.Marshal(r).String())
	return err
}

func (c *nativeContext) UnpinAllTopicMessages(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.Chat().ID
	}
	if r.ThreadID == 0 {
		r.ThreadID = c.Message().ThreadID
	}
	_, err := c.b.Raw("unpinAllForumTopicMessages", json.Marshal(r).String())
	return err
}

func (c *nativeContext) EditGeneralTopicName(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.Chat().ID
	}
	if r.Name == "" {
		r.Name = "GTopic"
	}
	_, err := c.b.Raw("editGeneralForumTopic", json.Marshal(r).String())
	return err
}

func (c *nativeContext) CloseGeneralTopic(r *Forum) error{
	if r.ID == 0 {
		r.ID = c.Chat().ID
	}
	_, err := c.b.Raw("closeGeneralForumTopic", json.Marshal(r).String())
	return err
}

