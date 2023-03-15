package handler

import (
	"net/http"
	"vpub/web/handler/form"
	"vpub/web/handler/request"
	"strings"
)

func (h *Handler) showEditPostView(w http.ResponseWriter, r *http.Request) {
	user := request.GetUserContextKey(r)

	post, err := h.storage.PostById(RouteInt64Param(r, "postId"))
	if err != nil {
		serverError(w, err)
		return
	}

	topic_map := make(map[int64]string)

	if (user.IsAdmin) {
		topics, err := h.storage.Topics()
		if err != nil {
			serverError(w, err)
			return
		}

		for _, topic_v := range topics {
			post_t, err := h.storage.PostById(topic_v.Post.Id)
			board_t, err := h.storage.BoardById(topic_v.BoardId)

			if err != nil {
				notFound(w)
				return
			}

			var sb strings.Builder

			sb.WriteString(" / ")
			sb.WriteString(board_t.Forum.Name)
			sb.WriteString(" / ")
			sb.WriteString(board_t.Name)
			sb.WriteString(" / ")
			sb.WriteString(post_t.Subject)

			topic_map[topic_v.Id] = sb.String()
		}
	}

	topic, err := h.storage.TopicById(post.TopicId)
	if err != nil {
		notFound(w)
		return
	}

	board, err := h.storage.BoardById(topic.BoardId)

	postForm := form.PostForm{
		UserIsAdmin: 	user.IsAdmin,
		Subject: 			post.Subject,
		Content: 			post.Content,
		TopicId: 			post.TopicId,
		Topics: 			topic_map,
	}

	v := NewView(w, r, "edit_post")
	v.Set("navigation", navigation{
		Forum: board.Forum,
		Board: board,
		Topic: topic.Post.Subject,
	})
	v.Set("form", postForm)
	v.Set("post", post)
	v.Set("topic", topic)
	v.Set("board", board)
	v.Render()
}
