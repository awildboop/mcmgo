package mcmgo

import (
	"encoding/json"
	"time"
)

/*
* NOTE: All structs from the API are strictly translated from the Objects documentation.
* Save for dates, we *never* alter the specified types. We do not alter data in any way.
*
* Dates are a custom struct (Time) that is just time.Time. It allows for converting from
* unix timestamps (given by the API) to time.Time. This may be changed if it causes any
* problems in the future. This is unlikely, but would be a breaking change.
 */

// Raw response from the API
type Response struct {
	Result string          `json:"result"`
	Data   json.RawMessage `json:"data,omitempty"`
	Error  json.RawMessage `json:"error,omitempty"`
}

// Used to convert unix from the api to time.Time
type Time struct {
	time.Time
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// https://www.mc-market.org/wiki/v1-objects/#alerts-8203
type Alert struct {
	CausedMemberID uint   `json:"caused_member_id"`
	ContentType    string `json:"content_type"`
	ContentID      uint   `json:"content_id"`
	AlertType      string `json:"alert_type"`
	AlertDate      Time   `json:"alert_date"`
}

// https://www.mc-market.org/wiki/v1-objects/#conversations-8203
type Conversation struct {
	ConversationID  uint   `json:"conversation_id"`
	Title           string `json:"title"`
	CreationDate    Time   `json:"creation_date"`
	CreatorID       uint   `json:"creator_id"`
	LastMessageDate Time   `json:"last_message_date"`
	LastReadDate    Time   `json:"last_read_date"`
	Open            bool   `json:"open"`
	ReplyCount      uint   `json:"reply_count"`
	RecipientIds    []uint `json:"recipient_ids"`
}

// https://www.mc-market.org/wiki/v1-objects/#reply-8203
type ConversationReply struct {
	MessageID   uint   `json:"message_id"`
	MessageDate Time   `json:"message_date"`
	AuthorID    uint   `json:"author_id"`
	Message     string `json:"message"`
}

// https://www.mc-market.org/wiki/v1-objects/#member-8203
type Member struct {
	MemberID         uint   `json:"member_id"`
	Username         string `json:"username"`
	JoinDate         Time   `json:"join_date"`
	LastActivityDate Time   `json:"last_activity_date,omitempty"` // optional
	Banned           bool   `json:"banned"`
	Suspended        bool   `json:"suspended"`
	Restricted       bool   `json:"restricted"`
	Disabled         bool   `json:"disabled"`
	PostCount        uint   `json:"post_count"`
	ResourceCount    uint   `json:"resource_count"`
	PurchaseCount    uint   `json:"purchase_count"`
	FeedbackPositive uint   `json:"feedback_positive"`
	FeedbackNeutral  uint   `json:"feedback_neutral"`
	FeedbackNegative uint   `json:"feedback_negative"`
}

// https://www.mc-market.org/wiki/v1-objects/#ban-8203
type Ban struct {
	MemberID   uint   `json:"member_id"`
	BannedByID uint   `json:"banned_by_id"`
	BanDate    Time   `json:"ban_date"`
	Reason     string `json:"reason"`
}

// https://www.mc-market.org/wiki/v1-objects/#profilepost-8203
type ProfilePost struct {
	ProfilePostID uint   `json:"profile_post_id"`
	AuthorID      uint   `json:"author_id"`
	PostDate      Time   `json:"post_date"`
	Message       string `json:"message"`
	CommentCount  uint   `json:"comment_count"`
}

// https://www.mc-market.org/wiki/v1-objects/#basicresource-8203
type BasicResource struct {
	ResourceID uint    `json:"resource_id"`
	AuthorID   uint    `json:"author_id"`
	Title      string  `json:"title"`
	TagLine    string  `json:"tag_line"`
	Price      float64 `json:"price"` // TODO: determine whether this should actually be float64 or whether it should be float32
	Currency   string  `json:"currency"`
}

// https://www.mc-market.org/wiki/v1-objects/#resource-8203
type Resource struct {
	ResourceID       uint    `json:"resource_id"`
	AuthorID         uint    `json:"author_id"`
	Title            string  `json:"title"`
	TagLine          string  `json:"tag_line"`
	Description      string  `json:"description"`
	ReleaseDate      Time    `json:"release_date"`
	LastUpdateDate   Time    `json:"last_update_date"`
	CategoryTitle    string  `json:"category_title"`
	CurrentVersionID uint    `json:"current_version_id"`
	Price            float64 `json:"price"`
	Currency         string  `json:"currency"`
	PurchaseCount    uint    `json:"purchase_count"`
	DownloadCount    uint    `json:"download_count"`
	ReviewCount      uint    `json:"review_count"`
	ReviewAverage    float64 `json:"review_average"`
}

// https://www.mc-market.org/wiki/v1-objects/#version-8203
type Version struct {
	VersionID     uint   `json:"version_id"`
	UpdateID      uint   `json:"update_id"`
	Name          string `json:"name"`
	ReleaseDate   Time   `json:"release_date"`
	DownloadCount uint   `json:"download_count"`
}

// https://www.mc-market.org/wiki/v1-objects/#update-8203
type Update struct {
	UpdateID   uint   `json:"update_id"`
	Title      string `json:"title"`
	Message    string `json:"message"`
	UpdateDate Time   `json:"update_date"`
}

// https://www.mc-market.org/wiki/v1-objects/#review-8203
type Review struct {
	ReviewID   uint   `json:"review_id"`
	ReviewerID uint   `json:"reviewer_id"`
	ReviewDate Time   `json:"review_date"`
	Rating     uint   `json:"rating"`
	Message    string `json:"message"`
	Response   string `json:"response,omitempty"` // optional
}

// https://www.mc-market.org/wiki/v1-objects/#purchase-8203
type Purchase struct {
	PurchaseID     uint    `json:"purchase_id"`
	PurchaserID    uint    `json:"purchaser_id"`
	LicenseID      uint    `json:"license_id"`
	Renewal        bool    `json:"renewal"`
	Status         string  `json:"status"`
	Price          float64 `json:"price"`
	Currency       string  `json:"currency"`
	PurchaseDate   Time    `json:"purchase_date"`
	ValidationDate Time    `json:"validation_date"`
}

// https://www.mc-market.org/wiki/v1-objects/#license-8203
type License struct {
	LicenseID       uint `json:"license_id"`
	PurchaserID     uint `json:"purchaser_id"`
	Validated       bool `json:"validated"`
	Active          bool `json:"active"`
	StartDate       Time `json:"start_date"`
	EndDate         Time `json:"end_date"`
	PreviousEndDate Time `json:"previous_end_date"`
}

// https://www.mc-market.org/wiki/v1-objects/#download-8203
type Download struct {
	DownloadID   uint
	VersionID    uint
	DownloaderID uint
	DownloadDate Time
}

// https://www.mc-market.org/wiki/v1-objects/#basicthread-8203
type Basicthread struct {
	ThreadID        uint   `json:"thread_id"`
	Title           string `json:"title"`
	ReplyCount      uint   `json:"reply_count"`
	ViewCount       uint   `json:"view_count"`
	CreationDate    Time   `json:"creation_date"`
	LastMessageDate Time   `json:"last_message_date"`
}

// https://www.mc-market.org/wiki/v1-objects/#thread-8203
type Thread struct {
	ThreadID     uint   `json:"thread_id"`
	ForumName    string `json:"forum_name"`
	Title        string `json:"title"`
	ReplyCount   uint   `json:"reply_count"`
	ViewCount    uint   `json:"view_count"`
	PostDate     Time   `json:"post_date"`
	ThreadType   string `json:"thread_type"`
	ThreadOpen   bool   `json:"thread_open"`
	LastPostDate Time   `json:"last_post_date"`
}

// no link
type ThreadReply struct {
	ReplyID  uint   `json:"reply_id"`
	AuthorID uint   `json:"author_id"`
	PostDate Time   `json:"post_date"`
	Message  string `json:"message"`
}
