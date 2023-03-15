package form

import (
	"net/http"
	"strconv"
	"strings"
)

type PostForm struct {
	UserIsAdmin bool
	Subject string
	Content string
	TopicId int64
	TopicIdNew int64
	Topics map[int64]string
}

func NewPostForm(r *http.Request) PostForm {
	TopicId, _ := strconv.ParseInt(r.FormValue("topicId"), 10, 64)
	TopicIdNew, _ := strconv.ParseInt(r.FormValue("topicIdNew"), 10, 64)
	return PostForm{
		Subject: strings.TrimSpace(r.FormValue("subject")),
		Content: r.FormValue("content"),
		TopicId: TopicId,
		TopicIdNew: TopicIdNew,
	}
}
