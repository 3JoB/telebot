package params

type (
	CreateStickerSet struct {
		UserID       string `json:"user_id"`
		Type         string `json:"sticker_type"`
		Name         string `json:"name"`
		Title        string `json:"title"`
		Emojis       string `json:"emojis"`
		MaskPosition string `json:"mask_position"`

		PNG  any `json:"png_sticker,omitempty"`
		TGS  any `json:"tgs_sticker,omitempty"`
		WEBM any `json:"webm_sticker,omitempty"`
	}
)
