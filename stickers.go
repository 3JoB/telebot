package telebot

import (
	"strconv"

	"github.com/3JoB/unsafeConvert"
	"github.com/goccy/go-json"
)

type StickerSetType = string

const (
	StickerRegular     = "regular"
	StickerMask        = "mask"
	StickerCustomEmoji = "custom_emoji"
)

// StickerSet represents a sticker set.
type StickerSet struct {
	Type          StickerSetType `json:"sticker_type"`
	Name          string         `json:"name"`
	Title         string         `json:"title"`
	Animated      bool           `json:"is_animated"`
	Video         bool           `json:"is_video"`
	Stickers      []Sticker      `json:"stickers"`
	Thumbnail     *Photo         `json:"thumbnail"`
	PNG           *File          `json:"png_sticker"`
	TGS           *File          `json:"tgs_sticker"`
	WebM          *File          `json:"webm_sticker"`
	Emojis        string         `json:"emojis"`
	ContainsMasks bool           `json:"contains_masks"` // FIXME: can be removed
	MaskPosition  *MaskPosition  `json:"mask_position"`
}

// MaskPosition describes the position on faces where
// a mask should be placed by default.
type MaskPosition struct {
	Feature MaskFeature `json:"point"`
	XShift  float32     `json:"x_shift"`
	YShift  float32     `json:"y_shift"`
	Scale   float32     `json:"scale"`
}

// MaskFeature defines sticker mask position.
type MaskFeature string

const (
	FeatureForehead MaskFeature = "forehead"
	FeatureEyes     MaskFeature = "eyes"
	FeatureMouth    MaskFeature = "mouth"
	FeatureChin     MaskFeature = "chin"
)

// UploadSticker uploads a PNG file with a sticker for later use.
func (b *Bot) UploadSticker(to Recipient, png *File) (*File, error) {
	files := map[string]File{
		"png_sticker": *png,
	}
	params := map[string]any{
		"user_id": to.Recipient(),
	}

	data, err := b.sendFiles("uploadStickerFile", files, params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result File
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return &resp.Result, nil
}

// StickerSet returns a sticker set on success.
func (b *Bot) StickerSet(name string) (*StickerSet, error) {
	data, err := b.Raw("getStickerSet", map[string]string{"name": name})
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result *StickerSet
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}

// CreateStickerSet creates a new sticker set.
func (b *Bot) CreateStickerSet(to Recipient, s StickerSet) error {
	files := make(map[string]File)
	if s.PNG != nil {
		files["png_sticker"] = *s.PNG
	}
	if s.TGS != nil {
		files["tgs_sticker"] = *s.TGS
	}
	if s.WebM != nil {
		files["webm_sticker"] = *s.WebM
	}

	params := map[string]any{
		"user_id":        to.Recipient(),
		"sticker_type":   s.Type,
		"name":           s.Name,
		"title":          s.Title,
		"emojis":         s.Emojis,
		"contains_masks": strconv.FormatBool(s.ContainsMasks),
	}

	if s.MaskPosition != nil {
		data, _ := json.Marshal(&s.MaskPosition)
		params["mask_position"] = unsafeConvert.StringReflect(data)
	}

	_, err := b.sendFiles("createNewStickerSet", files, params)
	return err
}

// Use this method to delete a sticker set that was created by the bot. Returns True on success.
func (b *Bot) DeleteStickerSet(name string) error {
	_, err := b.Raw("deleteStickerSet", map[string]string{"name": name})
	return err
}

// AddSticker adds a new sticker to the existing sticker set.
func (b *Bot) AddSticker(to Recipient, s StickerSet) error {
	files := make(map[string]File)
	if s.PNG != nil {
		files["png_sticker"] = *s.PNG
	} else if s.TGS != nil {
		files["tgs_sticker"] = *s.TGS
	} else if s.WebM != nil {
		files["webm_sticker"] = *s.WebM
	}

	params := map[string]any{
		"user_id": to.Recipient(),
		"name":    s.Name,
		"emojis":  s.Emojis,
	}

	if s.MaskPosition != nil {
		data, _ := json.Marshal(&s.MaskPosition)
		params["mask_position"] = unsafeConvert.StringReflect(data)
	}

	_, err := b.sendFiles("addStickerToSet", files, params)
	return err
}

// SetStickerPosition moves a sticker in set to a specific position.
func (b *Bot) SetStickerPosition(sticker string, position int) error {
	params := map[string]string{
		"sticker":  sticker,
		"position": unsafeConvert.IntToString(position),
	}

	_, err := b.Raw("setStickerPositionInSet", params)
	return err
}

// DeleteSticker deletes a sticker from a set created by the bot.
func (b *Bot) DeleteSticker(sticker string) error {
	_, err := b.Raw("deleteStickerFromSet", map[string]string{"sticker": sticker})
	return err
}

// SetStickerSetThumbnail sets a thumbnail of the sticker set.
// Animated thumbnails can be set for animated sticker sets only.
//
// Thumbnail must be a PNG image, up to 128 kilobytes in size
// and have width and height exactly 100px, or a TGS animation
// up to 32 kilobytes in size.
//
// Animated sticker set thumbnail can't be uploaded via HTTP URL.
func (b *Bot) SetStickerSetThumbnail(to Recipient, s StickerSet) error {
	files := make(map[string]File)
	if s.PNG != nil {
		files["thumbnail"] = *s.PNG
	} else if s.TGS != nil {
		files["thumbnail"] = *s.TGS
	}

	params := map[string]any{
		"name":    s.Name,
		"user_id": to.Recipient(),
	}

	_, err := b.sendFiles("setStickerSetThumbnail", files, params)
	return err
}

// CustomEmojiStickers returns the information about custom emoji stickers by their ids.
func (b *Bot) CustomEmojiStickers(ids []string) ([]Sticker, error) {
	data, _ := json.Marshal(ids)

	params := map[string]string{
		"custom_emoji_ids": unsafeConvert.StringReflect(data),
	}

	data, err := b.Raw("getCustomEmojiStickers", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result []Sticker
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}
