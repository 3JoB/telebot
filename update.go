package crare

import (
	"strings"
)

// Update object represents an incoming update.
type Update struct {
	ID int `json:"update_id"`

	Message           *Message          `json:"message,omitempty"`
	EditedMessage     *Message          `json:"edited_message,omitempty"`
	ChannelPost       *Message          `json:"channel_post,omitempty"`
	EditedChannelPost *Message          `json:"edited_channel_post,omitempty"`
	Callback          *Callback         `json:"callback_query,omitempty"`
	Query             *Query            `json:"inline_query,omitempty"`
	InlineResult      *InlineResult     `json:"chosen_inline_result,omitempty"`
	ShippingQuery     *ShippingQuery    `json:"shipping_query,omitempty"`
	PreCheckoutQuery  *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
	Poll              *Poll             `json:"poll,omitempty"`
	PollAnswer        *PollAnswer       `json:"poll_answer,omitempty"`
	MyChatMember      *ChatMemberUpdate `json:"my_chat_member,omitempty"`
	ChatMember        *ChatMemberUpdate `json:"chat_member,omitempty"`
	ChatJoinRequest   *ChatJoinRequest  `json:"chat_join_request,omitempty"`
}

// ProcessUpdate processes a single incoming update.
// A started bot calls this function automatically.
func (b *Bot) ProcessUpdate(u Update) bool {
	c := b.NewContext(u)

	if u.Message != nil {
		m := u.Message

		if m.PinnedMessage != nil {
			return b.handle(OnPinned, c)
		}

		// Commands
		if m.Text != "" {
			// Filtering malicious messages
			if m.Text[0] == '\a' {
				return true
			}

			// match := cmdRx.FindAllStringSubmatch(m.Text, -1)
			command, botName, payload := process(m.Text)
			if command != "" {
				if botName != "" && !strings.EqualFold(b.Me.Username, botName) {
					return false
				}
				m.Payload = payload
				if b.handle(command, c) {
					return true
				}
			}

			// 1:1 satisfaction
			if b.handle(m.Text, c) {
				return true
			}

			return b.handle(OnText, c)
		}

		if b.handleMedia(c) {
			return true
		}

		if m.Contact != nil {
			return b.handle(OnContact, c)
		}
		if m.Location != nil {
			return b.handle(OnLocation, c)
		}
		if m.Venue != nil {
			return b.handle(OnVenue, c)
		}
		if m.Game != nil {
			return b.handle(OnGame, c)
		}
		if m.Dice != nil {
			return b.handle(OnDice, c)
		}
		if m.Invoice != nil {
			return b.handle(OnInvoice, c)
		}
		if m.Payment != nil {
			return b.handle(OnPayment, c)
		}
		if m.TopicClosed != nil {
			return b.handle(OnTopicCreated, c)
		}
		if m.TopicReopened != nil {
			return b.handle(OnTopicReopened, c)
		}
		if m.TopicClosed != nil {
			return b.handle(OnTopicClosed, c)
		}
		if m.TopicEdited != nil {
			return b.handle(OnTopicEdited, c)
		}
		if m.GeneralTopicHidden != nil {
			return b.handle(OnGeneralTopicHidden, c)
		}
		if m.GeneralTopicUnhidden != nil {
			return b.handle(OnGeneralTopicUnhidden, c)
		}
		if m.WriteAccessAllowed != nil {
			return b.handle(OnWriteAccessAllowed, c)
		}

		wasAdded := (m.UserJoined != nil && m.UserJoined.ID == b.Me.ID) ||
			(m.UsersJoined != nil && isUserInList(b.Me, m.UsersJoined))
		if m.GroupCreated || m.SuperGroupCreated || wasAdded {
			return b.handle(OnAddedToGroup, c)
		}

		if m.UserJoined != nil {
			return b.handle(OnUserJoined, c)
		}

		if m.UsersJoined != nil {
			for _, user := range m.UsersJoined {
				m.UserJoined = &user
				b.handle(OnUserJoined, c)
			}
			return true
		}

		if m.UserLeft != nil {
			return b.handle(OnUserLeft, c)
		}

		if m.NewGroupTitle != "" {
			return b.handle(OnNewGroupTitle, c)
		}

		if m.NewGroupPhoto != nil {
			return b.handle(OnNewGroupPhoto, c)
		}

		if m.GroupPhotoDeleted {
			return b.handle(OnGroupPhotoDeleted, c)
		}

		if m.GroupCreated {
			return b.handle(OnGroupCreated, c)
		}

		if m.SuperGroupCreated {
			return b.handle(OnSuperGroupCreated, c)
		}

		if m.ChannelCreated {
			return b.handle(OnChannelCreated, c)
		}

		if m.MigrateTo != 0 {
			m.MigrateFrom = m.Chat.ID
			return b.handle(OnMigration, c)
		}

		if m.VideoChatStarted != nil {
			return b.handle(OnVideoChatStarted, c)
		}

		if m.VideoChatEnded != nil {
			return b.handle(OnVideoChatEnded, c)
		}

		if m.VideoChatParticipants != nil {
			return b.handle(OnVideoChatParticipants, c)
		}

		if m.VideoChatScheduled != nil {
			return b.handle(OnVideoChatScheduled, c)
		}

		if m.WebAppData != nil {
			b.handle(OnWebApp, c)
		}

		if m.ProximityAlert != nil {
			return b.handle(OnProximityAlert, c)
		}

		if m.AutoDeleteTimer != nil {
			return b.handle(OnAutoDeleteTimer, c)
		}
	}

	if u.EditedMessage != nil {
		return b.handle(OnEdited, c)
	}

	if u.ChannelPost != nil {
		m := u.ChannelPost

		if m.PinnedMessage != nil {
			return b.handle(OnPinned, c)
		}

		return b.handle(OnChannelPost, c)
	}

	if u.EditedChannelPost != nil {
		return b.handle(OnEditedChannelPost, c)
	}

	if u.Callback != nil {
		if data := u.Callback.Data; data != "" && data[0] == '\f' {
			match := cbackRx.FindAllStringSubmatch(data, -1)
			if match != nil {
				unique, payload := match[0][1], match[0][3]
				if handler, ok := b.handlers["\f"+unique]; ok {
					u.Callback.Unique = unique
					u.Callback.Data = payload
					b.runHandler(handler, c)
					return true
				}
			}
		}

		return b.handle(OnCallback, c)
	}

	if u.Query != nil {
		return b.handle(OnQuery, c)
	}

	if u.InlineResult != nil {
		return b.handle(OnInlineResult, c)
	}

	if u.ShippingQuery != nil {
		return b.handle(OnShipping, c)
	}

	if u.PreCheckoutQuery != nil {
		return b.handle(OnCheckout, c)
	}

	if u.Poll != nil {
		return b.handle(OnPoll, c)
	}

	if u.PollAnswer != nil {
		return b.handle(OnPollAnswer, c)
	}

	if u.MyChatMember != nil {
		return b.handle(OnMyChatMember, c)
	}

	if u.ChatMember != nil {
		return b.handle(OnChatMember, c)
	}

	if u.ChatJoinRequest != nil {
		return b.handle(OnChatJoinRequest, c)
	}

	return false
}

func (b *Bot) handle(end string, c *Context) bool {
	if handler, ok := b.handlers[end]; ok {
		b.runHandler(handler, c)
		return true
	}
	return false
}

func (b *Bot) handleMedia(c *Context) bool {
	var (
		m     = c.Message()
		fired bool
	)

	switch {
	case m.Photo != nil:
		fired = b.handle(OnPhoto, c)
	case m.Voice != nil:
		fired = b.handle(OnVoice, c)
	case m.Audio != nil:
		fired = b.handle(OnAudio, c)
	case m.Animation != nil:
		fired = b.handle(OnAnimation, c)
	case m.Document != nil:
		fired = b.handle(OnDocument, c)
	case m.Sticker != nil:
		fired = b.handle(OnSticker, c)
	case m.Video != nil:
		fired = b.handle(OnVideo, c)
	case m.VideoNote != nil:
		fired = b.handle(OnVideoNote, c)
	default:
		return fired
	}

	if !fired {
		return b.handle(OnMedia, c)
	}

	return fired
}

func (b *Bot) runHandler(h *Handle, c *Context) {
	f := func() {
		if err := h.doMiddleware(c); err != nil {
			b.OnError(err, c)
		}
		if err := h.do(c); err != nil {
			b.OnError(err, c)
		}
		c.releaseContext()
	}
	if b.synchronous {
		f()
	} else {
		go f()
	}
}

func isUserInList(user *User, list []User) bool {
	for _, user2 := range list {
		if user.ID == user2.ID {
			return true
		}
	}
	return false
}
