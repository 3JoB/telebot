package telebot

import (
	"github.com/3JoB/unsafeConvert"
	"github.com/goccy/go-json"
)

type Forum struct {
	ctx       *nativeContext `json:"-"`
	ID        int64          `json:"chat_id"`
	IconColor int64          `json:"icon_color"`
	ThreadID  int            `json:"message_thread_id"`
	Name      string         `json:"name"`
	EmojiID   string         `json:"icon_custom_emoji_id"`
}

func (c *nativeContext) Topic() *Forum {
	r := new(Forum)
	r.ctx = c
	return r
}

func (c *Forum) New(r *Forum) error {
	if r.ID == 0 {
		if !c.ctx.Chat().IsForum {
			return Err("Not Forum")
		}
		r.ID = c.ctx.Chat().ID
	}
	if !c.ctx.Chat().IsForum {
		return Err("Not Forum")
	}
	data, _ := json.Marshal(r)
	_, err := c.ctx.b.Raw("createForumTopic", unsafeConvert.StringReflect(data))
	return err
}

func (c *Forum) Edit(r *Forum) error {
	if r.ID == 0 {
		if !c.ctx.Chat().IsForum {
			return Err("Not Forum")
		}
		r.ID = c.ctx.Chat().ID
	}
	if r.ThreadID == 0 {
		r.ThreadID = c.ctx.Message().ThreadID
	}
	data, _ := json.Marshal(r)
	_, err := c.ctx.b.Raw("editForumTopic", unsafeConvert.StringReflect(data))
	return err
}

func (c *Forum) Delete(r *Forum) error {
	if r.ID == 0 {
		if !c.ctx.Chat().IsForum {
			return Err("Not Forum")
		}
		r.ID = c.ctx.Chat().ID
	}
	data, _ := json.Marshal(r)
	_, err := c.ctx.b.Raw("deleteForumTopic", unsafeConvert.StringReflect(data))
	return err
}

func (c *Forum) ReOpen(r *Forum) error {
	if r.ID == 0 {
		if !c.ctx.Chat().IsForum {
			return Err("Not Forum")
		}
		r.ID = c.ctx.Chat().ID
	}
	data, _ := json.Marshal(r)
	_, err := c.ctx.b.Raw("reopenForumTopic", unsafeConvert.StringReflect(data))
	return err
}

func (c *Forum) Close(r *Forum) error {
	if r.ID == 0 {
		if !c.ctx.Chat().IsForum {
			return Err("Not Forum")
		}
		r.ID = c.ctx.Chat().ID
	}
	if r.ThreadID == 0 {
		r.ThreadID = c.ctx.Message().ThreadID
	}
	data, _ := json.Marshal(r)
	_, err := c.ctx.b.Raw("closeForumTopic", unsafeConvert.StringReflect(data))
	return err
}

func (c *Forum) UnpinAllMessages(r *Forum) error {
	if r.ID == 0 {
		if !c.ctx.Chat().IsForum {
			return Err("Not Forum")
		}
		r.ID = c.ctx.Chat().ID
	}
	if r.ThreadID == 0 {
		r.ThreadID = c.ctx.Message().ThreadID
	}
	data, _ := json.Marshal(r)
	_, err := c.ctx.b.Raw("unpinAllForumTopicMessages", unsafeConvert.StringReflect(data))
	return err
}

func (c *Forum) GeneralNameEdit(r *Forum) error {
	if r.ID == 0 {
		if !c.ctx.Chat().IsForum {
			return Err("Not Forum")
		}
		r.ID = c.ctx.Chat().ID
	}
	if r.Name == "" {
		r.Name = "GTopic"
	}
	data, _ := json.Marshal(r)
	_, err := c.ctx.b.Raw("editGeneralForumTopic", unsafeConvert.StringReflect(data))
	return err
}

func (c *Forum) GeneralClose(r *Forum) error {
	if r.ID == 0 {
		if !c.ctx.Chat().IsForum {
			return Err("Not Forum")
		}
		r.ID = c.ctx.Chat().ID
	}
	data, _ := json.Marshal(r)
	_, err := c.ctx.b.Raw("closeGeneralForumTopic", unsafeConvert.StringReflect(data))
	return err
}

func (c *Forum) GeneralReOpen(r *Forum) error {
	if r.ID == 0 {
		if !c.ctx.Chat().IsForum {
			return Err("Not Forum")
		}
		r.ID = c.ctx.Chat().ID
	}
	data, _ := json.Marshal(r)
	_, err := c.ctx.b.Raw("reopenGeneralForumTopic", unsafeConvert.StringReflect(data))
	return err
}

func (c *Forum) GeneralHide(r *Forum) error {
	if r.ID == 0 {
		if !c.ctx.Chat().IsForum {
			return Err("Not Forum")
		}
		r.ID = c.ctx.Chat().ID
	}
	data, _ := json.Marshal(r)
	_, err := c.ctx.b.Raw("hideGeneralForumTopic", unsafeConvert.StringReflect(data))
	return err
}

func (c *Forum) GeneralUnHide(r *Forum) error {
	if r.ID == 0 {
		if !c.ctx.Chat().IsForum {
			return Err("Not Forum")
		}
		r.ID = c.ctx.Chat().ID
	}
	data, _ := json.Marshal(r)
	_, err := c.ctx.b.Raw("unhideGeneralForumTopic", unsafeConvert.StringReflect(data))
	return err
}
