package config

type ConfigChannel struct {
	ChannelId   int     `json:"channel_id"`
	ChannelName string  `json:"channel_name"`
	ExpRate     float32 `json:"exp_rate"`
	DropRate    float32 `json:"drop_rate"`
	MoneyRate   float32 `json:"money_rate"`
	Port        int     `json:"port"`
	MaxPlayer   int     `json:"max_player"`
}
