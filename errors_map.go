package telebot

var ErrMap map[uint32]*Error = map[uint32]*Error{
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