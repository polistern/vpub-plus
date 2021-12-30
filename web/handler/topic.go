package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"pboard/model"
	"strconv"
)

func contains(list []string, val string) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func (h *Handler) showTopicView(w http.ResponseWriter, r *http.Request) {
	user, _ := h.session.Get(r)

	var page int64 = 0
	if val, ok := r.URL.Query()["page"]; ok && len(val[0]) == 1 {
		page, _ = strconv.ParseInt(val[0], 10, 64)
	}

	topic := mux.Vars(r)["topic"]
	if !contains(h.topics, topic) {
		notFound(w)
		return
	}

	posts, hasMore, err := h.storage.PostsTopic(topic, page, h.perPage)
	if err != nil {
		serverError(w, err)
		return
	}

	h.renderLayout(w, "topic", map[string]interface{}{
		"topic":   topic,
		"posts":   posts,
		"topics":  h.topics,
		"hasMore": hasMore,
	}, user)
}

func (h *Handler) showPageNumber(w http.ResponseWriter, r *http.Request) {
	user, _ := h.session.Get(r)

	var page int64 = 0
	if val, ok := mux.Vars(r)["nb"]; ok {
		page, _ = strconv.ParseInt(val, 10, 64)
	}

	var topic string
	if val, ok := r.URL.Query()["topic"]; ok && len(val) == 1 {
		topic = val[0]
	}
	if !contains(h.topics, topic) && topic != "" {
		notFound(w)
		return
	}

	var posts []model.Post
	var hasMore bool
	var err error
	if topic != "" {
		posts, hasMore, err = h.storage.PostsTopic(topic, page, h.perPage)
	} else {
		posts, hasMore, err = h.storage.Posts(page, h.perPage)
	}

	if err != nil {
		serverError(w, err)
		return
	}

	h.renderLayout(w, "paginate", map[string]interface{}{
		"topic":    topic,
		"posts":    posts,
		"page":     page,
		"topics":   h.topics,
		"hasMore":  hasMore,
		"nextPage": page + 1,
	}, user)
}
