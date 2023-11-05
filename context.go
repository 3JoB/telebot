package crare

import (
	"errors"
	"strings"
	"time"

	"github.com/cornelk/hashmap"
)

// Context wraps an update and represents the context of current event.
type Context struct {
	b     *Bot
	u     Update
	next  bool
	store *hashmap.Map[string, any]
}

// Bot returns the bot instance.
func (c *Context) Bot() *Bot {
	return c.b
}

// Update returns the original update.
func (c *Context) Update() Update {
	return c.u
}

// Message returns stored message if such presented.
func (c *Context) Message() *Message {
	switch {
	case c.u.Message != nil:
		return c.u.Message
	case c.u.Callback != nil:
		return c.u.Callback.Message
	case c.u.EditedMessage != nil:
		return c.u.EditedMessage
	case c.u.ChannelPost != nil:
		if c.u.ChannelPost.PinnedMessage != nil {
			return c.u.ChannelPost.PinnedMessage
		}
		return c.u.ChannelPost
	case c.u.EditedChannelPost != nil:
		return c.u.EditedChannelPost
	default:
		return nil
	}
}

// Callback returns stored callback if such presented.
func (c *Context) Callback() *Callback {
	return c.u.Callback
}

// Query returns stored query if such presented.
func (c *Context) Query() *Query {
	return c.u.Query
}

// InlineResult returns stored inline result if such presented.
func (c *Context) InlineResult() *InlineResult {
	return c.u.InlineResult
}

// ShippingQuery returns stored shipping query if such presented.
func (c *Context) ShippingQuery() *ShippingQuery {
	return c.u.ShippingQuery
}

// PreCheckoutQuery returns stored pre checkout query if such presented.
func (c *Context) PreCheckoutQuery() *PreCheckoutQuery {
	return c.u.PreCheckoutQuery
}

// ChatMember returns chat member changes.
func (c *Context) ChatMember() *ChatMemberUpdate {
	switch {
	case c.u.ChatMember != nil:
		return c.u.ChatMember
	case c.u.MyChatMember != nil:
		return c.u.MyChatMember
	default:
		return nil
	}
}

// ChatJoinRequest returns chat member join request.
func (c *Context) ChatJoinRequest() *ChatJoinRequest {
	return c.u.ChatJoinRequest
}

// Poll returns stored poll if such presented.
func (c *Context) Poll() *Poll {
	return c.u.Poll
}

// PollAnswer returns stored poll answer if such presented.
func (c *Context) PollAnswer() *PollAnswer {
	return c.u.PollAnswer
}

// Migration returns both migration from and to chat IDs.
func (c *Context) Migration() (int64, int64) {
	return c.u.Message.MigrateFrom, c.u.Message.MigrateTo
}

// Sender returns the current recipient, depending on the context type.
// Returns nil if user is not presented.
func (c *Context) Sender() *User {
	switch {
	case c.u.Callback != nil:
		return c.u.Callback.Sender
	case c.Message() != nil:
		return c.Message().Sender
	case c.u.Query != nil:
		return c.u.Query.Sender
	case c.u.InlineResult != nil:
		return c.u.InlineResult.Sender
	case c.u.ShippingQuery != nil:
		return c.u.ShippingQuery.Sender
	case c.u.PreCheckoutQuery != nil:
		return c.u.PreCheckoutQuery.Sender
	case c.u.PollAnswer != nil:
		return c.u.PollAnswer.Sender
	case c.u.MyChatMember != nil:
		return c.u.MyChatMember.Sender
	case c.u.ChatMember != nil:
		return c.u.ChatMember.Sender
	case c.u.ChatJoinRequest != nil:
		return c.u.ChatJoinRequest.Sender
	default:
		return nil
	}
}

// Chat returns the current chat, depending on the context type.
// Returns nil if chat is not presented.
func (c *Context) Chat() *Chat {
	switch {
	case c.Message() != nil:
		return c.Message().Chat
	case c.u.MyChatMember != nil:
		return c.u.MyChatMember.Chat
	case c.u.ChatMember != nil:
		return c.u.ChatMember.Chat
	case c.u.ChatJoinRequest != nil:
		return c.u.ChatJoinRequest.Chat
	default:
		return nil
	}
}

// Recipient combines both Sender and Chat functions. If there is no user
// the chat will be returned. The native context cannot be without sender,
// but it is useful in the case when the context created intentionally
// by the NewContext constructor and have only Chat field inside.
func (c *Context) Recipient() Recipient {
	chat := c.Chat()
	if chat != nil {
		return chat
	}
	return c.Sender()
}

// Text returns the message text, depending on the context type.
// In the case when no related data presented, returns an empty string.
func (c *Context) Text() string {
	m := c.Message()
	if m == nil {
		return ""
	}
	if m.Caption != "" {
		return m.Caption
	}
	return m.Text
}

// Entities returns the message entities, whether it's media caption's or the text's.
// In the case when no entities presented, returns a nil.
func (c *Context) Entities() Entities {
	m := c.Message()
	if m == nil {
		return nil
	}
	if len(m.CaptionEntities) > 0 {
		return m.CaptionEntities
	}
	return m.Entities
}

// Data returns the current data, depending on the context type.
// If the context contains command, returns its arguments string.
// If the context contains payment, returns its payload.
// In the case when no related data presented, returns an empty string.
func (c *Context) Data() string {
	switch {
	case c.u.Message != nil:
		return c.u.Message.Payload
	case c.u.Callback != nil:
		return c.u.Callback.Data
	case c.u.Query != nil:
		return c.u.Query.Text
	case c.u.InlineResult != nil:
		return c.u.InlineResult.Query
	case c.u.ShippingQuery != nil:
		return c.u.ShippingQuery.Payload
	case c.u.PreCheckoutQuery != nil:
		return c.u.PreCheckoutQuery.Payload
	default:
		return ""
	}
}

// Args returns a raw slice of command or callback arguments as strings.
// The message arguments split by space, while the callback's ones by a "|" symbol.
func (c *Context) Args() []string {
	switch {
	case c.u.Message != nil:
		payload := strings.Trim(c.u.Message.Payload, " ")
		if payload != "" {
			return strings.Split(payload, " ")
		}
	case c.u.Callback != nil:
		return strings.Split(c.u.Callback.Data, "|")
	case c.u.Query != nil:
		return strings.Split(c.u.Query.Text, " ")
	case c.u.InlineResult != nil:
		return strings.Split(c.u.InlineResult.Query, " ")
	}
	return nil
}

// Send sends a message to the current recipient.
// See Send from bot.go.
func (c *Context) Send(what any, opts ...any) (*Message, error) {
	e, err := c.b.Send(c.Recipient(), what, opts...)
	return e, err
}

// SendAlbum sends an album to the current recipient.
// See SendAlbum from bot.go.
func (c *Context) SendAlbum(a Album, opts ...any) error {
	_, err := c.b.SendAlbum(c.Recipient(), a, opts...)
	return err
}

// Reply replies to the current message.
// See Reply from bot.go.
func (c *Context) Reply(what any, opts ...any) (*Message, error) {
	msg := c.Message()
	if msg == nil {
		return nil, ErrBadContext
	}
	return c.b.Reply(msg, what, opts...)
}

// Forward forwards the given message to the current recipient.
// See Forward from bot.go.
func (c *Context) Forward(msg Editable, opts ...any) error {
	_, err := c.b.Forward(c.Recipient(), msg, opts...)
	return err
}

// ForwardTo forwards the current message to the given recipient.
// See Forward from bot.go
func (c *Context) ForwardTo(to Recipient, opts ...any) error {
	msg := c.Message()
	if msg == nil {
		return ErrBadContext
	}
	_, err := c.b.Forward(to, msg, opts...)
	return err
}

// Edit edits the current message.
// See Edit from bot.go.
func (c *Context) Edit(what any, opts ...any) error {
	if c.u.InlineResult != nil {
		_, err := c.b.Edit(c.u.InlineResult, what, opts...)
		return err
	}
	if c.u.Callback != nil {
		_, err := c.b.Edit(c.u.Callback, what, opts...)
		return err
	}
	return ErrBadContext
}

// EditCaption edits the caption of the current message.
// See EditCaption from bot.go.
func (c *Context) EditCaption(caption string, opts ...any) error {
	if c.u.InlineResult != nil {
		_, err := c.b.EditCaption(c.u.InlineResult, caption, opts...)
		return err
	}
	if c.u.Callback != nil {
		_, err := c.b.EditCaption(c.u.Callback, caption, opts...)
		return err
	}
	return ErrBadContext
}

// EditOrSend edits the current message if the update is callback,
// otherwise the content is sent to the chat as a separate message.
func (c *Context) EditOrSend(what any, opts ...any) error {
	err := c.Edit(what, opts...)
	if err == ErrBadContext {
		_, err := c.Send(what, opts...)
		return err
	}
	return err
}

// EditOrReply edits the current message if the update is callback,
// otherwise the content is replied as a separate message.
func (c *Context) EditOrReply(what any, opts ...any) (*Message, error) {
	err := c.Edit(what, opts...)
	if err == ErrBadContext {
		return c.Reply(what, opts...)
	}
	return nil, err
}

// Delete removes the current message.
// See Delete from bot.go.
func (c *Context) Delete() error {
	msg := c.Message()
	if msg == nil {
		return ErrBadContext
	}
	return c.b.Delete(msg)
}

// DeleteAfter waits for the duration to elapse and then removes the
// message. It handles an error automatically using b.OnError callback.
// It returns a Timer that can be used to cancel the call using its Stop method.
func (c *Context) DeleteAfter(d time.Duration) *time.Timer {
	return time.AfterFunc(d, func() {
		if err := c.Delete(); err != nil {
			c.b.OnError(err, c)
		}
	})
}

// Next pass control to the next middleware/ctx function.
func (c *Context) Next() error {
	c.next = true

	return nil
}

// Next pass control to the next middleware/ctx function.
func (c *Context) Done() error {
	return nil
}

// Notify updates the chat action for the current recipient.
// See Notify from bot.go.
func (c *Context) Notify(action ChatAction) error {
	return c.b.Notify(c.Recipient(), action)
}

// Ship replies to the current shipping query.
// See Ship from bot.go.
func (c *Context) Ship(what ...any) error {
	if c.u.ShippingQuery == nil {
		return errors.New("telebot: context shipping query is nil")
	}
	return c.b.Ship(c.u.ShippingQuery, what...)
}

// Accept finalizes the current deal.
// See Accept from bot.go.
func (c *Context) Accept(errorMessage ...string) error {
	if c.u.PreCheckoutQuery == nil {
		return errors.New("telebot: context pre checkout query is nil")
	}
	return c.b.Accept(c.u.PreCheckoutQuery, errorMessage...)
}

// Respond sends a response for the current callback query.
// See Respond from bot.go.
func (c *Context) Respond(resp ...*CallbackResponse) error {
	if c.u.Callback == nil {
		return errors.New("telebot: context callback is nil")
	}
	return c.b.Respond(c.u.Callback, resp...)
}

// Answer sends a response to the current inline query.
// See Answer from bot.go.
func (c *Context) Answer(resp *QueryResponse) error {
	if c.u.Query == nil {
		return errors.New("telebot: context inline query is nil")
	}
	return c.b.Answer(c.u.Query, resp)
}

// Set saves data in the context.
func (c *Context) Set(k string, v any) {
	if c.store == nil {
		c.store = hashmap.New[string, any]()
	}
	c.store.Set(k, v)
}

// Get retrieves data from the context.
func (c *Context) Get(k string) any {
	if c.store == nil {
		c.store = hashmap.New[string, any]()
	}
	v, _ := c.store.Get(k)
	return v
}

// Release the Context. After it is released,
// the previous Context should not be continued to be used.
func (n *Context) releaseContext() {
	if n == nil {
		return
	}
	n.next = false
	if n.store != nil {
		n.store.Range(func(k string, v any) bool {
			n.store.Del(k)
			return true
		})
	}
	n.b = nil
	n.u = Update{}
	ctxPool.Put(n)
}
