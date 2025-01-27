package storage

import (
	"context"
	"errors"
	"vpub/model"
)

func (s *Storage) PostsByTopicId(id int64) ([]model.Post, bool, error) {
	rows, err := s.db.Query("select topic_id, post_id, subject, content, created_at, updated_at, user_id, name, picture, about from posts_full where topic_id=$1 order by created_at", id)
	if err != nil {
		return nil, false, err
	}
	var posts []model.Post
	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.TopicId, &post.Id, &post.Subject, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.User.Id, &post.User.Name, &post.User.Picture, &post.User.About)
		if err != nil {
			return posts, false, err
		}
		posts = append(posts, post)
	}
	return posts, false, nil
}

func (s *Storage) Posts(page int64) ([]model.Post, bool, error) {
	var posts []model.Post
	settings, err := s.Settings()
	if err != nil {
		return posts, false, err
	}
	rows, err := s.db.Query(`
select
       topic_id,
       post_id,
       subject,
       content,
       created_at,
       updated_at,
       user_id,
       name
from posts_full 
order by created_at desc
offset $1
limit $2`, settings.PerPage*(page-1), settings.PerPage+1)
	if err != nil {
		return nil, false, err
	}
	for rows.Next() {
		var post model.Post
		err := rows.Scan(
			&post.TopicId,
			&post.Id,
			&post.Subject,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.User.Id,
			&post.User.Name,
		)
		if err != nil {
			return posts, false, err
		}
		posts = append(posts, post)
	}
	if len(posts) > int(settings.PerPage) {
		return posts[0:settings.PerPage], true, err
	}
	return posts, false, nil
}

func (s *Storage) PostsByUserId(id, page int64) ([]model.Post, bool, error) {
	var posts []model.Post
	settings, err := s.Settings()
	if err != nil {
		return posts, false, err
	}
	rows, err := s.db.Query(`
select
       topic_id,
       post_id,
       subject,
       content,
       created_at,
       updated_at,
       user_id,
       name
from posts_full 
where user_id=$1
order by created_at desc
offset $2
limit $3`, id, settings.PerPage*(page-1), settings.PerPage+1)
	if err != nil {
		return nil, false, err
	}
	for rows.Next() {
		var post model.Post
		err := rows.Scan(
			&post.TopicId,
			&post.Id,
			&post.Subject,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.User.Id,
			&post.User.Name,
		)
		if err != nil {
			return posts, false, err
		}
		posts = append(posts, post)
	}
	if len(posts) > int(settings.PerPage) {
		return posts[0:settings.PerPage], true, err
	}
	return posts, false, nil
}

func (s *Storage) CreatePost(userId, topicId int64, request model.PostRequest) (int64, error) {
	var id int64

	query := `
		INSERT INTO	posts 
		    (subject, content, user_id, topic_id) 
		VALUES ($1, $2, $3, $4)
		returning id
    `

	if err := s.db.QueryRow(
		query,
		request.Subject,
		request.Content,
		userId,
		topicId,
	).Scan(
		&id,
	); err != nil {
		return id, errors.New("unable to create post: " + err.Error())
	}

	return id, nil
}

func (s *Storage) PostById(id int64) (model.Post, error) {
	var post model.Post
	err := s.db.QueryRow("select * from posts_full where post_id=$1", id).Scan(&post.TopicId, &post.Id, &post.Subject, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.User.Id, &post.User.Name, &post.User.Picture, &post.User.About)
	if err != nil {
		return post, err
	}
	return post, err
}

func (s *Storage) DeletePost(post model.Post) error {
	stmt, err := s.db.Prepare(`delete from posts where id=$1 and (user_id = $2 or (select is_admin from users where id=$2))`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(post.Id, post.User.Id)
	return err
}

func (s *Storage) UpdatePost(id, userId int64, request model.PostRequest) error {
	query := `
        UPDATE posts
        SET 
            subject=$1,
            content=$2,
            updated_at=now() 
        WHERE 
            id=$3 and (user_id=$4 or (select is_admin from users where id=$4))
    `

	if _, err := s.db.Exec(
		query,
		request.Subject,
		request.Content,
		id,
		userId,
	); err != nil {
		return errors.New("unable to update post: " + err.Error())
	}

	return nil
}

func (s *Storage) UpdateAndMovePost(id, userId, newTopicId, oldTopicId int64, request model.PostRequest) error {
	ctx := context.Background()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, `
		UPDATE posts
        SET 
            subject=$1,
            content=$2,
            updated_at=now(),
            topic_id=$3 
        WHERE 
            id=$4 AND (SELECT is_admin FROM users WHERE id=$5)
    `,
		request.Subject,
		request.Content,
		newTopicId,
		id,
		userId,
	); err != nil {
		return errors.New("Unable to move or update post: " + err.Error())
	}

	if _, err := tx.ExecContext(ctx, `
        UPDATE topics SET
            posts_count = (SELECT posts_count FROM topics WHERE id=$1) + 1
        WHERE id=$1 AND (SELECT is_admin FROM users WHERE id=$2)
    `,
		newTopicId,
		userId,
	); err != nil {
		tx.Rollback()
		return errors.New("Unable to increase posts count: " + err.Error())
	}

	if _, err := tx.ExecContext(ctx, `
        UPDATE topics SET
            posts_count = (SELECT posts_count FROM topics WHERE id=$1) - 1
        WHERE id=$1 AND posts_count > 0 AND (SELECT is_admin FROM users WHERE id=$2)
    `,
		oldTopicId,
		userId,
	); err != nil {
		tx.Rollback()
		return errors.New("Unable to decrease posts count: " + err.Error())
	}

	err = tx.Commit()
	return err
}

func (s *Storage) NewestPostFromTopic(topicId int64) (int64, error) {
	var id int64

	query := `
        select id from posts where topic_id=$1 order by created_at desc
    `

	err := s.db.QueryRow(
		query,
		topicId,
	).Scan(
		&id,
	)

	return id, err
}
