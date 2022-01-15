package storage

import (
	"context"
	"time"
	"vpub/model"
)

const postQuery = "select p.id, p.subject, p.content, p.created_at, p.topic_id, u.id, u.name, u.picture from posts p left join users u on p.user_id = u.id "

func parseCreatedAt(createdAt string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", createdAt)
}

func (s *Storage) PostsByTopicId(id int64) ([]model.Post, bool, error) {
	rows, err := s.db.Query("select * from postUsers where topic_id=$1 or post_id=$1", id)
	if err != nil {
		return nil, false, err
	}
	var posts []model.Post
	for rows.Next() {
		var post model.Post
		var createdAtStr string
		var topicId *int64
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &createdAtStr, &topicId, &post.BoardId, &post.User.Id, &post.User.Name, &post.User.Picture)
		if err != nil {
			return posts, false, err
		}
		post.CreatedAt, err = parseCreatedAt(createdAtStr)
		if err != nil {
			return posts, false, err
		}
		posts = append(posts, post)
	}
	return posts, false, nil
}

func (s *Storage) CreatePost(post model.Post) (int64, error) {
	var id int64
	err := s.db.QueryRow(`INSERT INTO posts (subject, content, user_id, topic_id, board_id) VALUES ($1, $2, $3, $4, $5) returning id`, post.Title, post.Content, post.User.Id, post.TopicId, post.BoardId).Scan(&id)
	return id, err
}

func (s *Storage) PostById(id int64) (model.Post, error) {
	var post model.Post
	var createdAtStr string
	var topicId *int64
	err := s.db.QueryRow("select * from postUsers where post_id=$1", id).Scan(&post.Id, &post.Title, &post.Content, &createdAtStr, &topicId, &post.BoardId, &post.User.Id, &post.User.Name, &post.User.Picture)
	post.CreatedAt, err = parseCreatedAt(createdAtStr)
	if topicId != nil {
		post.TopicId = *topicId
	} else {
		post.TopicId = post.Id
	}
	return post, err
}

func (s *Storage) DeletePostById(id int64) error {
	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	var isFirstPost bool
	if err := tx.QueryRowContext(ctx, `select true from topics where first_post_id=$1`, id).Scan(&isFirstPost); err != nil {
		tx.Rollback()
		return err
	}
	if isFirstPost {
		if _, err := tx.ExecContext(ctx, `DELETE from posts WHERE id = $1`, id); err != nil {
			tx.Rollback()
			return err
		}
		return tx.Commit()
	}

	var topicId int64
	if err := tx.QueryRowContext(ctx, `DELETE from posts WHERE id = $1 returning topic_id`, id).Scan(&topicId); err != nil {
		tx.Rollback()
		return err
	}
	var boardId int64
	if err := tx.QueryRowContext(ctx, `update topics set replies=replies - 1 where id=$1 returning board_id`, topicId).Scan(&boardId); err != nil {
		tx.Rollback()
		return err
	}
	if _, err := tx.ExecContext(ctx, `update boards set posts=posts - 1 where id=$1`, boardId); err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (s *Storage) UpdatePost(post model.Post) error {
	stmt, err := s.db.Prepare(`UPDATE posts SET subject = $1, content = $2 WHERE id = $3;`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(post.Title, post.Content, post.Id)
	return err
}
