package post

type PostBody struct {
	Content  string `json:"content" form:"content"`
	PostType string `json:"post_type" form:"post_type" binding:"omitempty,oneof=profile group page"`
	Privacy  string `json:"privacy" form:"privacy" binding:"omitempty,oneof=onlyme friend follower public"`
}

type ReactBody struct {
	Type string `json:"type"`
}

type CommentBody struct {
	PostID    uint           `json:"post_id" form:"post_id"`
	ParentID  *uint          `json:"parent_id" form:"parent_id"`
	Content   string         `json:"content" form:"content"`
}