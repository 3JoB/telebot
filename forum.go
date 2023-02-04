package telebot

import (
	"github.com/3JoB/ulib/json"
)

type Forum struct {
	ctx *nativeContext
	ID        int64  `json:"chat_id"`
	IconColor int64  `json:"icon_color"`
	ThreadID  int    `json:"message_thread_id"`
	Name      string `json:"name"`
	EmojiID   string `json:"icon_custom_emoji_id"`
}

func (c *nativeContext) Topic() *Forum {
	r := new(Forum)
	r.ctx = c
	return r
}

func (c *Forum) New(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.ctx.Chat().ID
	}
	_, err := c.ctx.b.Raw("createForumTopic", json.Marshal(r).String())
	return err
}

func (c *Forum) Edit(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.ctx.Chat().ID
	}
	if r.ThreadID == 0 {
		r.ThreadID = c.ctx.Message().ThreadID
	}
	_, err := c.ctx.b.Raw("editForumTopic", json.Marshal(r).String())
	return err
}

func (c *Forum) Delete(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.ctx.Chat().ID
	}
	_, err := c.ctx.b.Raw("deleteForumTopic", json.Marshal(r).String())
	return err
}

func (c *Forum) ReOpen(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.ctx.Chat().ID
	}
	_, err := c.ctx.b.Raw("reopenForumTopic", json.Marshal(r).String())
	return err
}

func (c *Forum) Close(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.ctx.Chat().ID
	}
	if r.ThreadID == 0 {
		r.ThreadID = c.ctx.Message().ThreadID
	}
	_, err := c.ctx.b.Raw("closeForumTopic", json.Marshal(r).String())
	return err
}

func (c *Forum) UnpinAllMessages(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.ctx.Chat().ID
	}
	if r.ThreadID == 0 {
		r.ThreadID = c.ctx.Message().ThreadID
	}
	_, err := c.ctx.b.Raw("unpinAllForumTopicMessages", json.Marshal(r).String())
	return err
}

func (c *Forum) GeneralNameEdit(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.ctx.Chat().ID
	}
	if r.Name == "" {
		r.Name = "GTopic"
	}
	_, err := c.ctx.b.Raw("editGeneralForumTopic", json.Marshal(r).String())
	return err
}

func (c *Forum) GeneralClose(r *Forum) error{
	if r.ID == 0 {
		r.ID = c.ctx.Chat().ID
	}
	_, err := c.ctx.b.Raw("closeGeneralForumTopic", json.Marshal(r).String())
	return err
}


func (c *Forum) GeneralReOpen(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.ctx.Chat().ID
	}
	_, err := c.ctx.b.Raw("reopenGeneralForumTopic", json.Marshal(r).String())
	return err
}

func (c *Forum) GeneralHide(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.ctx.Chat().ID
	}
	_, err := c.ctx.b.Raw("hideGeneralForumTopic", json.Marshal(r).String())
	return err
}

func (c *Forum) GeneralUnHide(r *Forum) error {
	if r.ID == 0 {
		r.ID = c.ctx.Chat().ID
	}
	_, err := c.ctx.b.Raw("unhideGeneralForumTopic", json.Marshal(r).String())
	return err
}