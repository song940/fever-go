package fever

type Group struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type GroupsResponse struct {
	Groups      []Group       `json:"groups"`
	FeedsGroups []FeedsGroups `json:"feeds_groups"`
}

type FeedsResponse struct {
	Feeds       []Feed        `json:"feeds"`
	FeedsGroups []FeedsGroups `json:"feeds_groups"`
}

type FaviconsResponse struct {
	Favicons []Favicon `json:"favicons"`
}

type ItemRequest struct {
	WithIDs string `json:"with_ids"`
	SinceId string `json:"since_id"`
}

type ItemsResponse struct {
	Items []Item `json:"items"`
	Total int    `json:"total_items"`
}

type UnreadResponse struct {
	ItemIDs string `json:"unread_item_ids"`
}

type SavedResponse struct {
	ItemIDs string `json:"saved_item_ids"`
}

type FeedsGroups struct {
	GroupID int64  `json:"group_id"`
	FeedIDs string `json:"feed_ids"`
}

type Feed struct {
	ID          int64  `json:"id"`
	FaviconID   int64  `json:"favicon_id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	SiteURL     string `json:"site_url"`
	IsSpark     int    `json:"is_spark"`
	LastUpdated int64  `json:"last_updated_on_time"`
}

type Item struct {
	ID        int64  `json:"id"`
	FeedID    int64  `json:"feed_id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	HTML      string `json:"html"`
	URL       string `json:"url"`
	IsSaved   int    `json:"is_saved"`
	IsRead    int    `json:"is_read"`
	CreatedAt int64  `json:"created_on_time"`
}

type Favicon struct {
	ID   int64  `json:"id"`
	Data string `json:"data"`
}

type MarkRequest struct {
	Type string
	Id   string
	As   string
}

type MarkResponse struct {
}
