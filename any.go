package crare

type Response[T any] struct {
	Result T `json:"result"`
}

type ProfileStr struct {
	Count  int     `json:"total_count"`
	Photos []Photo `json:"photos"`
}
