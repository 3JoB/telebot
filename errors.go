package telebot

import (
	"fmt"
	"strings"
)

type (
	Error struct {
		Code        int
		Description string
		Message     string
	}

	FloodError struct {
		err        *Error
		RetryAfter int
	}

	GroupError struct {
		err        *Error
		MigratedTo int64
	}
)

// String returns description of error.
// A tiny shortcut to make code clearer.
func (err *Error) String() string {
	return err.Description
}

// Error implements error interface.
func (err *Error) Error() string {
	msg := err.Message
	if msg == "" {
		split := strings.Split(err.Description, ": ")
		if len(split) == 2 {
			msg = split[1]
		} else {
			msg = err.Description
		}
	}
	return fmt.Sprintf("telegram: %s (%d)", msg, err.Code)
}

// Error implements error interface.
func (err FloodError) Error() string {
	return err.err.Error()
}

// Error implements error interface.
func (err GroupError) Error() string {
	return err.err.Error()
}

// NewError returns new Error instance with given description.
// First element of msgs is Description. The second is optional Message.
func NewError(code int, msgs ...string) *Error {
	err := &Error{Code: code}
	if len(msgs) >= 1 {
		err.Description = msgs[0]
	}
	if len(msgs) >= 2 {
		err.Message = msgs[1]
	}
	return err
}

// General errors
var (
	ErrTooLarge     = NewError(400, "Request Entity Too Large")
	ErrUnauthorized = NewError(401, "Unauthorized")
	ErrNotFound     = NewError(404, "Not Found")
	ErrInternal     = NewError(500, "Internal Server Error")
)

// Bad request errors
var (
	ErrBadButtonData          = NewError(400, "Bad Request: BUTTON_DATA_INVALID")
	ErrBadPollOptions         = NewError(400, "Bad Request: expected an Array of String as options")
	ErrBadURLContent          = NewError(400, "Bad Request: failed to get HTTP URL content")
	ErrCantEditMessage        = NewError(400, "Bad Request: message can't be edited")
	ErrCantRemoveOwner        = NewError(400, "Bad Request: can't remove chat owner")
	ErrCantUploadFile         = NewError(400, "Bad Request: can't upload file by URL")
	ErrCantUseMediaInAlbum    = NewError(400, "Bad Request: can't use the media of the specified type in the album")
	ErrChatAboutNotModified   = NewError(400, "Bad Request: chat description is not modified")
	ErrChatNotFound           = NewError(400, "Bad Request: chat not found")
	ErrEmptyChatID            = NewError(400, "Bad Request: chat_id is empty")
	ErrEmptyMessage           = NewError(400, "Bad Request: message must be non-empty")
	ErrEmptyText              = NewError(400, "Bad Request: text is empty")
	ErrFailedImageProcess     = NewError(400, "Bad Request: IMAGE_PROCESS_FAILED", "Image process failed")
	ErrGroupMigrated          = NewError(400, "Bad Request: group chat was upgraded to a supergroup chat")
	ErrMessageNotModified     = NewError(400, "Bad Request: message is not modified")
	ErrNoRightsToDelete       = NewError(400, "Bad Request: message can't be deleted")
	ErrNoRightsToRestrict     = NewError(400, "Bad Request: not enough rights to restrict/unrestrict chat member")
	ErrNoRightsToSend         = NewError(400, "Bad Request: have no rights to send a message")
	ErrNoRightsToSendGifs     = NewError(400, "Bad Request: CHAT_SEND_GIFS_FORBIDDEN", "sending GIFS is not allowed in this chat")
	ErrNoRightsToSendPhoto    = NewError(400, "Bad Request: not enough rights to send photos to the chat")
	ErrNoRightsToSendStickers = NewError(400, "Bad Request: not enough rights to send stickers to the chat")
	ErrNotFoundToDelete       = NewError(400, "Bad Request: message to delete not found")
	ErrNotFoundToForward      = NewError(400, "Bad Request: message to forward not found")
	ErrNotFoundToReply        = NewError(400, "Bad Request: reply message not found")
	ErrQueryTooOld            = NewError(400, "Bad Request: query is too old and response timeout expired or query ID is invalid")
	ErrSameMessageContent     = NewError(400, "Bad Request: message is not modified: specified new message content and reply markup are exactly the same as a current content and reply markup of the message")
	ErrStickerEmojisInvalid   = NewError(400, "Bad Request: invalid sticker emojis")
	ErrStickerSetInvalid      = NewError(400, "Bad Request: STICKERSET_INVALID", "Stickerset is invalid")
	ErrStickerSetInvalidName  = NewError(400, "Bad Request: invalid sticker set name is specified")
	ErrStickerSetNameOccupied = NewError(400, "Bad Request: sticker set name is already occupied")
	ErrTooLongMarkup          = NewError(400, "Bad Request: reply markup is too long")
	ErrTooLongMessage         = NewError(400, "Bad Request: message is too long")
	ErrUserIsAdmin            = NewError(400, "Bad Request: user is an administrator of the chat")
	ErrWrongFileID            = NewError(400, "Bad Request: wrong file identifier/HTTP URL specified")
	ErrWrongFileIDCharacter   = NewError(400, "Bad Request: wrong remote file id specified: Wrong character in the string")
	ErrWrongFileIDLength      = NewError(400, "Bad Request: wrong remote file id specified: Wrong string length")
	ErrWrongFileIDPadding     = NewError(400, "Bad Request: wrong remote file id specified: Wrong padding in the string")
	ErrWrongFileIDSymbol      = NewError(400, "Bad Request: wrong remote file id specified: can't unserialize it. Wrong last symbol")
	ErrWrongTypeOfContent     = NewError(400, "Bad Request: wrong type of the web page content")
	ErrWrongURL               = NewError(400, "Bad Request: wrong HTTP URL specified")
	ErrForwardMessage         = NewError(400, "Bad Request: administrators of the chat restricted message forwarding")
)

// Forbidden errors
var (
	ErrBlockedByUser        = NewError(403, "Forbidden: bot was blocked by the user")
	ErrKickedFromGroup      = NewError(403, "Forbidden: bot was kicked from the group chat")
	ErrKickedFromSuperGroup = NewError(403, "Forbidden: bot was kicked from the supergroup chat")
	ErrKickedFromChannel    = NewError(403, "Forbidden: bot was kicked from the channel chat")
	ErrNotStartedByUser     = NewError(403, "Forbidden: bot can't initiate conversation with a user")
	ErrUserIsDeactivated    = NewError(403, "Forbidden: user is deactivated")
)

var (
	ErrMap map[uint32]*Error = map[uint32]*Error{
		hash32(ErrTooLarge):     ErrTooLarge,
		hash32(ErrUnauthorized): ErrUnauthorized,
		hash32(ErrNotFound):     ErrNotFound,
		hash32(ErrInternal):     ErrInternal,

		hash32(ErrBadButtonData):          ErrBadButtonData,
		hash32(ErrBadPollOptions):         ErrBadPollOptions,
		hash32(ErrBadURLContent):          ErrBadURLContent,
		hash32(ErrCantEditMessage):        ErrCantEditMessage,
		hash32(ErrCantRemoveOwner):        ErrCantRemoveOwner,
		hash32(ErrCantUploadFile):         ErrCantUploadFile,
		hash32(ErrCantUseMediaInAlbum):    ErrCantUseMediaInAlbum,
		hash32(ErrChatAboutNotModified):   ErrChatAboutNotModified,
		hash32(ErrChatNotFound):           ErrChatNotFound,
		hash32(ErrEmptyChatID):            ErrEmptyChatID,
		hash32(ErrEmptyMessage):           ErrEmptyMessage,
		hash32(ErrEmptyText):              ErrEmptyText,
		hash32(ErrFailedImageProcess):     ErrFailedImageProcess,
		hash32(ErrGroupMigrated):          ErrGroupMigrated,
		hash32(ErrMessageNotModified):     ErrMessageNotModified,
		hash32(ErrNoRightsToDelete):       ErrNoRightsToDelete,
		hash32(ErrNoRightsToRestrict):     ErrNoRightsToRestrict,
		hash32(ErrNoRightsToSend):         ErrNoRightsToSend,
		hash32(ErrNoRightsToSendGifs):     ErrNoRightsToSendGifs,
		hash32(ErrNoRightsToSendPhoto):    ErrNoRightsToSendPhoto,
		hash32(ErrNoRightsToSendStickers): ErrNoRightsToSendStickers,
		hash32(ErrNotFoundToDelete):       ErrNotFoundToDelete,
		hash32(ErrNotFoundToForward):      ErrNotFoundToForward,
		hash32(ErrNotFoundToReply):        ErrNotFoundToReply,
		hash32(ErrQueryTooOld):            ErrQueryTooOld,
		hash32(ErrSameMessageContent):     ErrSameMessageContent,
		hash32(ErrStickerEmojisInvalid):   ErrStickerEmojisInvalid,
		hash32(ErrStickerSetInvalid):      ErrStickerSetInvalid,
		hash32(ErrStickerSetInvalidName):  ErrStickerSetInvalidName,
		hash32(ErrStickerSetNameOccupied): ErrStickerSetNameOccupied,
		hash32(ErrTooLongMarkup):          ErrTooLongMarkup,
		hash32(ErrTooLongMessage):         ErrTooLongMessage,
		hash32(ErrUserIsAdmin):            ErrUserIsAdmin,
		hash32(ErrWrongFileID):            ErrWrongFileID,
		hash32(ErrWrongFileIDCharacter):   ErrWrongFileIDCharacter,
		hash32(ErrWrongFileIDLength):      ErrWrongFileIDLength,
		hash32(ErrWrongFileIDPadding):     ErrWrongFileIDPadding,
		hash32(ErrWrongFileIDSymbol):      ErrWrongFileIDSymbol,
		hash32(ErrWrongTypeOfContent):     ErrWrongTypeOfContent,
		hash32(ErrWrongURL):               ErrWrongURL,
		hash32(ErrForwardMessage):         ErrForwardMessage,

		hash32(ErrBlockedByUser):        ErrBlockedByUser,
		hash32(ErrKickedFromGroup):      ErrKickedFromGroup,
		hash32(ErrKickedFromSuperGroup): ErrKickedFromSuperGroup,
		hash32(ErrKickedFromChannel):    ErrKickedFromChannel,
		hash32(ErrNotStartedByUser):     ErrNotStartedByUser,
		hash32(ErrUserIsDeactivated):    ErrUserIsDeactivated,
	}
)

// Err returns Error instance by given description.
func Err(s string) error {
	if r, ok := ErrMap[hash32p(s)]; ok {
		fmt.Println(ok)
		fmt.Println(r.Description)
		return r
	}
	return nil
}

// wrapError returns new wrapped telebot-related error.
func wrapError(err error) error {
	return fmt.Errorf("telebot: %w", err)
}
