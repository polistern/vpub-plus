// Code generated by go generate; DO NOT EDIT.

package handler

var TplMap = map[string]string{
	"account": `{{ define "title" }}Account{{ end }}

{{ define "content" }}
<h1>Account</h1>

<form action="/save-account" method="post">
    {{ .csrfField }}
    <div class="field">
        <label for="picture">Picture</label>
        <input type="url" name="picture" id="picture" value="{{ .user.Picture }}">
    </div>
    <div class="field">
        <label for="about">About</label>
        <textarea name="about" id="about" autofocus>{{ .user.About }}</textarea>
    </div>
    <input type="submit" value="Submit">
</form>
</section>
{{ end }}
`,
	"admin": `{{ define "breadcrumb" }} > Admin{{ end }}
{{ define "content"}}
<h1>Admin</h1>
<nav>
  <ul>
    <li><a href="/admin/settings/edit">Edit settings</a></li>
    <li><a href="/admin/keys">Manage keys</a></li>
    <li><a href="/admin/boards">Manage boards</a></li>
    <li><a href="/admin/forums">Manage forums</a></li>
    <li><a href="/admin/users">Manage users</a></li>
  </ul>
</nav>
{{ end }}`,
	"admin_board": `{{ define "content"}}
<h1><a href="/admin">Admin</a> > Boards</h1>
<p>
    {{ if .hasForums }}
    <a href="/admin/boards/new">New board</a>
    {{ else }}
    <a href="/admin/forums">Create a forum</a> to create boards
    {{ end }}
</p>
<table>
    <thead>
    <tr>
        <th class="grow">Board</th>
        <th>Edit</th>
    </tr>
    </thead>
    <tbody>
    {{ if .forums }}
    {{ range .forums }}
    <tr class="forum">
        <td colspan="4">{{ .Name }}</td>
    </tr>
    {{ range .Boards }}
    <tr>
        <td colspan="grow">
            <a href="/boards/{{ .Id }}">{{ .Name }}</a><br>{{ .Description }}
        </td>
        <td class="center"><a href="/admin/boards/{{ .Id }}/edit">Edit</a></td>
    </tr>
    {{ end }}
    {{ end }}
    {{ else }}
    <tr>
        <td colspan="2">No boards yet.</td>
    </tr>
    {{ end }}
    </tbody>
</table>
{{ end }}`,
	"admin_board_create": `{{ define "breadcrumb" }} > <a href="/admin">Admin</a> > <a href="/admin/boards">Boards</a>{{ end }}
{{ define "title" }}New board{{ end }}
{{ define "content" }}
<h1><a href="/admin">Admin</a> > <a href="/admin/boards">Boards</a> > Create board</h1>
<form action="/admin/boards/save" method="post">
  {{ .csrfField }}
  {{ template "board_form" .form }}
  <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"admin_board_edit": `{{ define "breadcrumb" }} > <a href="/admin">Admin</a> > <a href="/admin/boards">Boards</a>{{ end }}
{{ define "title" }}Edit board{{ end }}
{{ define "content" }}
<h1><a href="/admin">Admin</a> > <a href="/admin/boards">Boards</a> > Edit board</h1>
<form action="/admin/boards/{{ .board.Id }}/update" method="post">
    {{ .csrfField }}
    {{ template "board_form" .form }}
    <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"admin_forum": `{{ define "content"}}
<h1><a href="/admin">Admin</a> > Forums</h1>
<p><a href="/admin/forums/new">New forum</a></p>
<table>
    <thead>
    <tr>
        <th class="grow">Forum</th>
        <th>Edit</th>
    </tr>
    </thead>
    <tbody>
    {{ if .forums }}
    {{ range .forums }}
    <tr>
        <td colspan="grow">
            {{ .Name }}
        </td>
        <td class="center"><a href="/admin/forums/{{ .Id }}/edit">Edit</a></td>
    </tr>
    {{ end }}
    {{ else }}
    <tr>
        <td colspan="2">No boards yet.</td>
    </tr>
    {{ end }}
    </tbody>
</table>
{{ end }}`,
	"admin_forum_create": `{{ define "breadcrumb" }} > <a href="/admin">Admin</a> > <a href="/admin/boards">Boards</a>{{ end }}
{{ define "title" }}New forum{{ end }}
{{ define "content" }}
<h1><a href="/admin">Admin</a> > <a href="/admin/forums">Forums</a> > Create forum</h1>
<form action="/admin/forums/save" method="post">
    {{ .csrfField }}
    {{ template "forum_form" .form }}
    <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"admin_forum_edit": `{{ define "title" }}Edit forum{{ end }}
{{ define "content" }}
<h1><a href="/admin">Admin</a> > <a href="/admin/forums">Forums</a> > Edit forum</h1>
<form action="/admin/forums/{{ .forum.Id }}/update" method="post">
    {{ .csrfField }}
    {{ template "forum_form" .form }}
    <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"admin_keys": `{{ define "breadcrumb" }} > <a href="/admin">Admin</a> > Keys{{ end }}
{{ define "title" }}Keys{{ end }}
{{ define "content" }}
<h1><a href="/admin">Admin</a> > Keys</h1>
<form action="/admin/keys/save" method="post">
    {{ .csrfField }}
    <input type="submit" value="Create key">
</form>

<table>
    <thead>
    <tr>
        <th class="grow">Key</th>
        <th>Created</th>
    </tr>
    </thead>
    <tbody>
    {{ range .keys }}
    <tr>
        <td colspan="grow">{{ .Key }}</td>
        <td class="center">{{ iso8601 .CreatedAt }}</td>
    </tr>
    {{ end }}
    </tbody>
</table>
{{ end }}
`,
	"admin_settings_edit": `{{ define "breadcrumb" }} > <a href="/admin">Admin</a> > Settings{{ end }}
{{ define "content"}}
<h1><a href="/admin">Admin</a> > Edit Settings</h1>
<form action="/admin/settings/update" method="post">
    {{ .csrfField }}
    <div class="field">
        <label for="name">Name</label>
        <input type="text" name="name" id="name" value="{{ .form.Name }}" autocomplete="off" maxlength="120" required autofocus/>
    </div>
    <div class="field">
        <label for="css">CSS</label>
        <textarea class="editor" name="css" id="css">{{ .form.Css }}</textarea>
    </div>
    <input type="submit" value="Submit">
</form>
</table>
{{ end }}`,
	"admin_user": `{{ define "breadcrumb" }} > <a href="/admin">Admin</a> > Users{{ end }}
{{ define "content"}}
<h1><a href="/admin">Admin</a> > Users</h1>
<table>
    <thead>
    <tr>
        <th class="grow">User</th>
        <th>Edit</th>
        <th>Password</th>
    </tr>
    </thead>
    <tbody>
    {{ range .users }}
    <tr>
        <td colspan="grow">
            <a href="/~{{ .Name }}">{{ .Name }}</a>
        </td>
        <td class="center"><a href="/admin/users/{{ .Name }}/edit">Edit</a></td>
        <td class="center"><a href="/reset-password?hash={{ .Hash }}">Reset</a></td>
    </tr>
    {{ end }}
    </tbody>
</table>
{{ end }}`,
	"admin_user_edit": `{{ define "breadcrumb" }} > <a href="/admin">Admin</a> > <a href="/admin/users">Users</a>{{ end }}
{{ define "title" }}Edit user{{ end }}
{{ define "content" }}
<h2>Edit user</h2>
<form action="/admin/users/{{ .user.Name }}/update" method="post">
    {{ .csrfField }}
    <div class="field">
        <label for="name">Name</label>
        <input type="text" name="name" id="name" value="{{ .form.Username }}" autocomplete="off" maxlength="120" required autofocus/>
    </div>
    <div class="field">
        <label for="about">About</label>
        <textarea class="editor" name="about" id="about">{{ .form.About }}</textarea>
    </div>
    <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"board": `{{ define "breadcrumb" }}<a href="/">boards</a> > {{ .board.Name }}{{ end }}
{{ define "content" }}
<h1><a href="/">boards</a> > {{ .board.Name }}</h1>

<!--<p>{{ .board.Description }}</p>-->

<nav class="actions">
    <p>
        {{ if .logged }}
        <a href="/boards/{{ .board.Id }}/new-topic">new topic</a>
        {{ end }}
<!--        <a href="TODO">follow</a>-->
    </p>
</nav>

<section>
    <table>
        <thead>
        <tr>
            <th class="grow">Subject</th>
            <th>Author</th>
            <th>Replies</th>
            <th>Updated</th>
        </tr>
        </thead>
        <tbody>
        {{ if .topics }}
        {{ range .topics }}
        <tr{{ if .IsSticky }} class="sticky"{{ end }}>
            <td colspan="grow"><a href="/topics/{{ .Id }}">{{ .Post.Subject }}</a></td>
            <td class="center"><a href="/~{{ .Post.User.Id }}">{{ .Post.User.Name }}</a></td>
            <td class="center">{{ .Replies }}</td>
            <td class="center">{{ iso8601 .UpdatedAt }}</td>
        </tr>
        {{ end }}
        {{ else }}
        <tr>
            <td colspan="4">No topics yet.</td>
        </tr>
        {{ end }}
        </tbody>
    </table>
</section>
{{ end }}`,
	"confirm_remove_post": `{{ define "content" }}
    Are you sure you you want to delete the following post?
    <p>{{ gmi2html .post.Content }}</p>
    <form action="/posts/{{ .post.Id }}/remove" method="post">
        {{ .csrfField }}
        <input type="submit" value="Submit">
    </form>
{{ end }}`,
	"confirm_remove_reply": `{{ define "content" }}
    Are you sure you you want to delete the following reply?
    <p>{{ gmi2html .reply.Content }}</p>
    <form action="/replies/{{ .reply.Id }}/remove" method="post">
        {{ .csrfField }}
        <input type="submit" value="Submit">
    </form>
{{ end }}`,
	"create_topic": `{{ define "title" }}New Thread{{ end }}
{{ define "breadcrumb" }} > <a href="/boards/{{ .board.Id }}">{{ .board.Name }}</a>{{ end }}
{{ define "content" }}
<h2>New topic</h2>
<form action="/boards/{{ .board.Id }}/save-topic" method="post">
    {{ .csrfField }}
    {{ template "post_form" .form.PostForm }}
    <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"edit_post": `{{ define "title" }}Edit Post{{ end }}

{{ define "content" }}
    <h2>Edit Post</h2>
    <form action="/posts/{{ .post.Id }}/update" method="post">
        {{ .csrfField }}
        {{ template "post_form" .form }}
        <input type="submit" value="Update">
    </form>
{{ end }}
`,
	"edit_reply": `{{ define "content" }}
    <form action="/replies/{{ .reply.Id }}/update" method="post" class="form">
        {{ .csrfField }}
        <label for="reply">reply</label><textarea name="reply" id="reply">{{ .form.Content }}</textarea>
        <input type="submit" value="Submit">
    </form>
{{ end }}
`,
	"index": `{{ define "content"}}
<h1>Boards</h1>
<table>
    <thead>
        <tr>
            <th class="grow">Board</th>
            <th>Topics</th>
            <th>Posts</th>
            <th>Updated</th>
        </tr>
    </thead>
    <tbody>
        {{ if .forums }}
        {{ range .forums }}
        <tr class="forum">
            <td colspan="4">{{ .Name }}</td>
        </tr>
        {{ range .Boards }}
        <tr>
            <td colspan="grow">
                <a href="/boards/{{ .Id }}">{{ .Name }}</a><br>{{ .Description }}
            </td>
            <td class="center">{{ .Topics }}</td>
            <td class="center">{{ .Posts }}</td>
            <td class="center">{{ iso8601 .UpdatedAt }}</td>
        </tr>
        {{ end }}
        {{ end }}
        {{ else }}
        <tr>
            <td colspan="4">No boards yet.</td>
        </tr>
        {{ end }}
    </tbody>
</table>
{{ end }}`,
	"login": `{{ define "title" }}Login{{ end }}

{{ define "content" }}
<h2>Login</h2>
<form action="/login" method="post" class="auth-form">
    {{ .csrfField }}
    <div class="field">
        <label for="name">Username</label>
        <input type="text" name="name" id="name" autocomplete="off" required/>
    </div>
    <div class="field">
        <label for="password">Password</label>
        <input type="password" name="password" id="password" required/>
    </div>
    <input type="submit" value="Login">
</form>
{{ end }}`,
	"notifications": `{{ define "content" }}
<h1>New replies</h1>
<p><a href="/notifications/mark-all-read">mark all as read</a></p>
{{ range .notifications }}
<div>
    <div class="meta">
        <ul class="key-value">
            <li><span class="key">From: </span><span class="value"><a href="/~{{ .Reply.User }}">{{ .Reply.User }}</a></span></li>
            <li><span class="key">On: </span><span class="value">{{ .Reply.Date }}</span></li>
            <li><span class="key">Post: </span><span class="value"><a href="/posts/{{ .Reply.PostId }}">{{ .Reply.PostTitle }}</a></span></li>
            <li><span class="key">Parent: </span><span class="value">
                        {{ if .Reply.ParentId }}<a href="/replies/{{ .Reply.ParentId }}">view</a>{{ else }}view{{end}}
            </span></li>
        </ul>
    </div>
    <div class="content">{{ gmi2html .Reply.Content }}</div>
    <p>
        <a href="/replies/{{ .Reply.Id }}">reply</a>
        <a href="/notifications/{{ .Id }}/mark-read">mark as read</a>
    </p>
</div>
{{ end }}
{{ end }}`,
	"paginate": `{{ define "content" }}
<p>Page {{ .page }}{{ if .topic }} of <a href="/topics/{{ .topic }}">{{ .topic }}</a>{{ end }}</p>
<section>
    {{ template "posts" .posts }}

    {{ if .hasMore }}
    {{ if .topic }}
    <a href="/page/{{ .nextPage }}?topic={{ .topic }}">More</a>
    {{ else }}
    <a href="/page/{{ .nextPage }}">More</a>
    {{ end }}
    {{ end }}
</section>
{{ end }}`,
	"post": `{{ define "breadcrumb" }} > <a href="/topics/{{ .post.Topic }}">{{ .post.Topic }}</a>{{ end }}
{{ define "content"}}
<!--<table class="thread">-->
<!--    <tr>-->
<!--        <td>{{ .post.User }}</td>-->
<!--        <td>{{ gmi2html .content }}</td>-->
<!--    </tr>-->
<!--    {{ range .replies }}-->
<!--    <tr>-->
<!--        <td>{{ .User }}</td>-->
<!--        <td>{{ gmi2html .Content }}</td>-->
<!--    </tr>-->
<!--    {{ end }}-->
<!--</table>-->
<!--<form action="/posts/{{ .post.Id }}/reply" method="post">-->
<!--    {{ .csrfField }}-->
<!--    <div class="field">-->
<!--        <textarea name="reply"></textarea>-->
<!--    </div>-->
<!--    <input type="submit" value="Reply">-->
<!--</form>-->
<h1>{{ .post.Subject }}</h1>
<!--<ol class="thread">-->
<!--    <li class="post">-->
<!--        <table>-->
<!--            <tr>-->
<!--                <td class="post-aside">{{ .post.User }}</td>-->
<!--                <td class="post-content">{{ gmi2html .content }}</td>-->
<!--            </tr>-->
<!--        </table>-->
<!--    </li>-->
<!--    {{ range .replies }}-->
<!--    <li class="post">-->
<!--        <table>-->
<!--            <tr>-->
<!--                <td class="post-aside">{{ .User }}</td>-->
<!--                <td class="post-content">{{ gmi2html .Content }}</td>-->
<!--            </tr>-->
<!--        </table>-->
<!--    </li>-->
<!--    {{ end }}-->
<!--</ol>-->
<table class="thread">
    <tr class="post">
        <td class="post-aside">
            <p>{{ .post.User }}</p>
            <p>{{ timeAgo .post.CreatedAt }}</p>
        </td>
        <td class="post-content">
            {{ gmi2html .content }}
        </td>
    </tr>
    {{ range .replies }}
    <tr class="post">
        <td class="post-aside">
            <p>{{ .User }}</p>
            <p>{{ timeAgo .CreatedAt }}</p>
        </td>
        <td class="post-content">
            {{ gmi2html .Content }}
        </td>
    </tr>
    {{ end }}
</table>
<form action="/posts/{{ .post.Id }}/reply" method="post">
    {{ .csrfField }}
    <div class="field">
        <textarea name="reply"></textarea>
    </div>
    <input type="submit" value="Reply">
</form>
<!--<h1>{{ .post.Subject }}</h1>-->
<!--<div class="meta">-->
<!--    {{ with .post }}-->
<!--    <ul class="key-value">-->
<!--        <li><span class="key">From: </span><span class="value"><a href="/~{{ .User }}">{{ .User }}</a></span></li>-->
<!--        <li><span class="key">On: </span><span class="value">{{ timeAgo .CreatedAt }} ({{ .Date }})</span></li>-->
<!--        {{ if .Topic }}<li><span class="key">Topic: </span><span class="value"><a href="/topics/{{ .Topic }}">{{ .Topic }}</a></span></li>{{ end }}-->
<!--    </ul>-->
<!--    {{ end }}-->
<!--</div>-->
<!--<div class="content">{{ gmi2html .content }}</div>-->
<!--{{- if eq .logged .post.User }}-->
<!--<p>-->
<!--    <a href="/posts/{{ .post.Id }}/edit">Edit</a>-->
<!--    <a href="/posts/{{ .post.Id }}/remove">Remove</a>-->
<!--</p>-->
<!--{{- end }}-->
<!--{{ if .logged }}-->
<!--<form action="/posts/{{ .post.Id }}/reply" method="post">-->
<!--    {{ .csrfField }}-->
<!--    <div class="field">-->
<!--        <textarea name="reply"></textarea>-->
<!--    </div>-->
<!--    <input type="submit" value="Reply">-->
<!--</form>-->
<!--{{ end }}-->
<!--{{ template "reply" .replies }}-->
{{ end }}`,
	"register": `{{ define "title" }}Register{{ end }}

{{ define "content" }}
    <h2>Register</h2>
    {{ if .error }}
    <p class="error">{{ .error }}</p>
    {{ end }}
    <form action="/register" method="post" class="auth-form">
        {{ .csrfField }}
        <div class="field">
            <label for="name">Username</label>
            <input type="text" id="name" name="name" autocomplete="off" value="{{ .form.Username }}" maxlength="15"/>
        </div>
        <div class="field">
            <label for="password">Password</label>
            <input type="password" id="password" name="password"/>
        </div>
        <div class="field">
            <label for="confirm">Confirm password</label>
            <input type="password" id="confirm" name="confirm" required/>
        </div>
        <div class="field">
            <label for="key">Key</label>
            <input type="text" id="key" name="key"/>
        </div>
        <input type="submit" value="Submit">
    </form>
{{ end }}
`,
	"reset_password": `{{ define "title" }}Reset password{{ end }}

{{ define "content" }}
<h2>Reset password</h2>
{{ if .error }}
<p class="error">{{ .error }}</p>
{{ end }}
<form action="/reset-password" method="post" class="auth-form">
    {{ .csrfField }}
    <input name="hash" type="hidden" value="{{ .hash }}">
    <div class="field">
        <label for="password">New password</label>
        <input type="password" id="password" name="password"/>
    </div>
    <div class="field">
        <label for="confirm">Confirm password</label>
        <input type="password" id="confirm" name="confirm" required/>
    </div>
    <input type="submit" value="Submit">
</form>
{{ end }}`,
	"topic": `{{ define "breadcrumb" }}<a href="/">boards</a> > <a href="/boards/{{ .board.Id }}">{{ .board.Name }}</a>{{ end }}
{{ define "content"}}
<h1><a href="/">boards</a> > <a href="/boards/{{ .board.Id }}">{{ .board.Name }}</a></h1>
<table class="post">
    {{ range .posts }}
    <tr class="header" id="{{ .Id }}">
        <td>{{ .User.Name }} {{ timeAgo .CreatedAt }}</td>
    </tr>
    <tr>
        <td>
            {{ if eq $.topic.Id .Id }}<h1>{{ .Subject }}</h1>{{ end }}
            {{ gmi2html .Content }}
            {{ if or (hasPermission .User.Name) $.logged.IsAdmin }}
            <p><a href="/posts/{{ .Id }}/edit">edit</a> <a href="/posts/{{ .Id }}/remove">remove</a></p>
            {{ end }}
        </td>
    </tr>
    {{ end }}
</table>
{{ if not .topic.IsLocked }}
<section style="margin-top: 1em;">
    <form action="/posts/save" method="post">
        {{ .csrfField }}
        <input type="hidden" name="topicId" value="{{ .topic.Id }}">
        <input type="hidden" name="boardId" value="{{ .board.Id }}">
        <input type="hidden" name="subject" value="Re: {{ .topic.Post.Subject }}">
        <div class="field">
            <label for="content">Reply to this topic</label>
            <textarea name="content" id="content" style="height: 150px;"></textarea>
        </div>
        <input type="submit" value="Reply">
    </form>
</section>
{{ end }}
{{ end }}`,
	"user_posts": `{{ define "content" }}
<h1>{{ .user.Name }}</h1>
<div class="content">{{ gmi2html .user.About }}</div>
<section class="posts">
{{ template "posts" .posts }}

{{if .showMore }}
<a href="/~{{ .user.Name }}?page={{ .nextPage }}">More</a>
{{ end }}
</section>
{{ end }}
`,
}
