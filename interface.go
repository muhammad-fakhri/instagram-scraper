package main

type JsonHTTPResponse struct {
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

type ProfileInfo struct {
	CreatedTime int64 `json:"created_time"`
	Info        struct {
		Biography            string `json:"biography"`
		BusinessCategoryName string `json:"business_category_name"`
		BusinessEmail        string `json:"business_email"`
		BusinessPhoneNumber  string `json:"business_phone_number"`
		ExternalUrl          string `json:"external_url"`
		FollowersCount       uint   `json:"followers_count"`
		FollowingCount       uint   `json:"following_count"`
		FullName             string `json:"full_name"`
		ID                   string `json:"id"`
		IsBusinessAccount    bool   `json:"is_business_account"`
		IsJoinedRecently     bool   `json:"is_joined_recently"`
		IsPrivate            bool   `json:"is_private"`
		PostsCount           int    `json:"posts_count"`
		ProfilePicURL        string `json:"profile_pic_url"`
	}
	Username string `json:"username"`
}

type ThumbnailResource struct {
	ConfigHeight uint   `json:"config_height"`
	ConfigWidth  uint   `json:"config_width"`
	Source       string `json:"src"`
}

type GraphImage struct {
	TypeName         string `json:"__typename"`
	CommentsDisabled bool   `json:"comments_disabled"`
	Dimensions       struct {
		Height uint `json:"height"`
		Width  uint `json:"width"`
	} `json:"dimensions"`
	DisplayURL           string `json:"display_url"`
	EdgeMediaPreviewLike struct {
		Count uint `json:"count"`
	} `json:"edge_media_preview_like"`
	EdgeMediaToCaption struct {
		Edges []*struct {
			Node struct {
				Text string `json:"text"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"edge_media_to_caption"`
	EdgeMediaToComment struct {
		Count uint `json:"count"`
	} `json:"edge_media_to_comment"`
	GatingInfo   string `json:"gating_info"`
	ID           string `json:"id"`
	IsVideo      bool   `json:"is_video"`
	MediaPreview string `json:"media_preview"`
	Owner        struct {
		ID string `json:"id"`
	} `json:"owner"`
	ShortCode          string               `json:"shortcode"`
	Tags               []string             `json:"tags"`
	TakenAtTimestamp   int64                `json:"taken_at_timestamp"`
	ThumbnailResources []*ThumbnailResource `json:"thumbnail_resources"`
	ThumbnailSource    string               `json:"thumbnail_src"`
	URLs               []string             `json:"urls"`
	Username           string               `json:"username"`
}

type JsonData struct {
	GraphImages      []*GraphImage `json:"GraphImages"`
	GraphProfileInfo *ProfileInfo  `json:"GraphProfileInfo"`
}

type ProfileResponse struct {
	UserID      string `json:"id"`
	FullName    string `json:"full_name"`
	Username    string `json:"username"`
	Biography   string `json:"biography"`
	PostsCount  int    `json:"posts_count"`
	CreatedTime int64  `json:"created_time"`
}
