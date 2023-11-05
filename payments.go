package crare

import (
	"math"

	"github.com/3JoB/unsafeConvert"
)

// ShippingQuery contains information about an incoming shipping query.
type ShippingQuery struct {
	Sender  *User           `json:"from"`
	ID      string          `json:"id"`
	Payload string          `json:"invoice_payload"`
	Address ShippingAddress `json:"shipping_address"`
}

// ShippingAddress represents a shipping address.
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// ShippingOption represents one shipping option.
type ShippingOption struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Prices []Price `json:"prices"`
}

// Payment contains basic information about a successful payment.
type Payment struct {
	Currency         string `json:"currency"`
	Payload          string `json:"invoice_payload"`
	OptionID         string `json:"shipping_option_id"`
	TelegramChargeID string `json:"telegram_payment_charge_id"`
	ProviderChargeID string `json:"provider_payment_charge_id"`
	Order            Order  `json:"order_info"`
	Total            int    `json:"total_amount"`
}

// PreCheckoutQuery contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	Sender   *User  `json:"from"`
	ID       string `json:"id"`
	Currency string `json:"currency"`
	Payload  string `json:"invoice_payload"`
	OptionID string `json:"shipping_option_id"`
	Total    int    `json:"total_amount"`
	Order    Order  `json:"order_info"`
}

// Order represents information about an order.
type Order struct {
	Name        string          `json:"name"`
	PhoneNumber string          `json:"phone_number"`
	Email       string          `json:"email"`
	Address     ShippingAddress `json:"shipping_address"`
}

// Invoice contains basic information about an invoice.
type Invoice struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Payload     string `json:"payload"`
	Currency    string `json:"currency"`
	Token       string `json:"provider_token"`
	Data        string `json:"provider_data"`

	// Unique deep-linking parameter that can be used to
	// generate this invoice when used as a start parameter (0).
	Start string `json:"start_parameter"`

	Prices []Price `json:"prices"`

	Photo     *Photo `json:"photo"`
	PhotoSize int    `json:"photo_size"`

	// Shows the total price in the smallest units of the currency.
	// For example, for a price of US$ 1.45 pass amount = 145.
	Total int `json:"total_amount"`

	MaxTipAmount        int   `json:"max_tip_amount"`
	SuggestedTipAmounts []int `json:"suggested_tip_amounts"`

	NeedName            bool `json:"need_name"`
	NeedPhoneNumber     bool `json:"need_phone_number"`
	NeedEmail           bool `json:"need_email"`
	NeedShippingAddress bool `json:"need_shipping_address"`
	SendPhoneNumber     bool `json:"send_phone_number_to_provider"`
	SendEmail           bool `json:"send_email_to_provider"`
	Flexible            bool `json:"is_flexible"`
}

func (i Invoice) params() map[string]any {
	params := map[string]any{
		"title":                         i.Title,
		"description":                   i.Description,
		"start_parameter":               i.Start,
		"payload":                       i.Payload,
		"provider_token":                i.Token,
		"provider_data":                 i.Data,
		"currency":                      i.Currency,
		"max_tip_amount":                i.MaxTipAmount,
		"need_name":                     i.NeedName,
		"need_phone_number":             i.NeedPhoneNumber,
		"need_email":                    i.NeedEmail,
		"need_shipping_address":         i.NeedShippingAddress,
		"send_phone_number_to_provider": i.SendPhoneNumber,
		"send_email_to_provider":        i.SendEmail,
		"is_flexible":                   i.Flexible,
	}
	if i.Photo != nil {
		if i.Photo.FileURL != "" {
			params["photo_url"] = i.Photo.FileURL
		}
		if i.PhotoSize > 0 {
			params["photo_size"] = i.PhotoSize
		}
		if i.Photo.Width > 0 {
			params["photo_width"] = i.Photo.Width
		}
		if i.Photo.Height > 0 {
			params["photo_height"] = i.Photo.Height
		}
	}
	if len(i.Prices) > 0 {
		data, _ := defaultJson.Marshal(i.Prices)
		params["prices"] = unsafeConvert.StringPointer(data)
	}
	if len(i.SuggestedTipAmounts) > 0 {
		data, _ := defaultJson.Marshal(i.SuggestedTipAmounts)
		params["suggested_tip_amounts"] = unsafeConvert.StringPointer(data)
	}
	return params
}

// Price represents a portion of the price for goods or services.
type Price struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

// Currency contains information about supported currency for payments.
type Currency struct {
	Code         string `json:"code"`
	Title        string `json:"title"`
	Symbol       string `json:"symbol"`
	Native       string `json:"native"`
	ThousandsSep string `json:"thousands_sep"`
	DecimalSep   string `json:"decimal_sep"`
	SymbolLeft   bool   `json:"symbol_left"`
	SpaceBetween bool   `json:"space_between"`
	Exp          int    `json:"exp"`
	MinAmount    any    `json:"min_amount"`
	MaxAmount    any    `json:"max_amount"`
}

func (c Currency) FromTotal(total int) float64 {
	return float64(total) / math.Pow(10, float64(c.Exp))
}

func (c Currency) ToTotal(total float64) int {
	return int(total) * int(math.Pow(10, float64(c.Exp)))
}

// Send delivers invoice through bot b to recipient.
func (i *Invoice) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := i.params()
	params["chat_id"] = to.Recipient()
	b.embedSendOptions(params, opt)

	data, err := b.Raw("sendInvoice", params)
	if err != nil {
		return nil, err
	}

	return extractMessage(data)
}
