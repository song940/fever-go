package fever

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Handler interface {
	FeverAuthenticate(apiKey string) bool
	FeverGroups() GroupsResponse
	FeverFeeds() FeedsResponse
	FeverItems(req *ItemRequest) ItemsResponse
	FeverSavedItemIds() SavedResponse
	FeverUnreadItemIds() UnreadResponse
	FeverMark(req *MarkRequest) MarkResponse
}

type Fever struct {
	Handler
}

func New(handler Handler) *Fever {
	return &Fever{
		Handler: handler,
	}
}

func (f *Fever) Handle(input url.Values) (response map[string]interface{}) {
	response = make(map[string]interface{})
	response["last_refreshed_on_time"] = time.Now().Unix()
	response["api_version"] = 3
	if !input.Has("api_key") {
		response["auth"] = 0
		response["error"] = "missing api_key"
		return
	}
	apiKey := input.Get("api_key")
	if f.Handler.FeverAuthenticate(apiKey) {
		response["auth"] = 1
	} else {
		response["auth"] = 0
		return
	}
	var payload any
	payload = make(map[string]interface{})
	switch {
	case input.Has("groups"):
		payload = f.Handler.FeverGroups()
	case input.Has("feeds"):
		payload = f.Handler.FeverFeeds()
	case input.Has("unread_item_ids"):
		payload = f.Handler.FeverUnreadItemIds()
	case input.Has("items"):
		req := &ItemRequest{}
		req.SinceId = input.Get("since_id")
		req.WithIDs = input.Get("with_ids")
		payload = f.Handler.FeverItems(req)
	case input.Has("saved_item_ids"):
		payload = f.Handler.FeverSavedItemIds()
	case input.Has("mark"):
		req := &MarkRequest{}
		req.Type = input.Get("mark")
		req.As = input.Get("as")
		req.Id, _ = strconv.Atoi(input.Get("id"))
		payload = f.Handler.FeverMark(req)
	default:
		log.Println("fever unknown request", input)
	}
	ja, _ := json.Marshal(payload)
	json.Unmarshal(ja, &response)
	return
}

func (f *Fever) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// log.Println("fever request", r.Method, r.URL, r.Form)
	response := f.Handle(r.Form)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
