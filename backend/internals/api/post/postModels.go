package post

type PostBody struct {
	Content  string `json:"content" binding:"required"`
	PostType string `json:"post_type" binding:"required,oneof=profile group page"`
	Privacy  string `json:"privacy" binding:"required,oneof=onlyme friend follower public"`
}
