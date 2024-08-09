package reply_dto


type ReplyCreateDto struct {
	Description string `json:"desc"`	
	BlogID uint64 `json:"blogId"`
}