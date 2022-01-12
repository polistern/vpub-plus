package handler

import (
	"github.com/gorilla/csrf"
	"net/http"
)

func (h *Handler) showTopicView(w http.ResponseWriter, r *http.Request) {
	user, _ := h.session.Get(r)
	topic, err := h.storage.TopicById(RouteInt64Param(r, "topicId"))
	if err != nil {
		notFound(w)
		return
	}
	board, err := h.storage.BoardById(topic.BoardId)
	posts, _, err := h.storage.PostsByTopicId(topic.Id)
	h.renderLayout(w, "topic", map[string]interface{}{
		"board":          board,
		"topic":          topic,
		"posts":          posts,
		csrf.TemplateTag: csrf.TemplateField(r),
	}, user)
}