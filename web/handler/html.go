// Code generated by go generate; DO NOT EDIT.

package handler

var TplMap = map[string]string{
	"confirm_remove_post": `{{ define "content" }}
    Are you sure you you want to delete the following post?
    <p>{{ .post.Content }}</p>
    <form action="/posts/{{ .post.Id }}/remove" method="post">
        {{ .csrfField }}
        <input type="submit" value="Submit">
    </form>
{{ end }}`,
	"confirm_remove_reply": `{{ define "content" }}
    Are you sure you you want to delete the following reply?
    <p>{{ .reply.Content }}</p>
    <form action="/replies/{{ .reply.Id }}/remove" method="post">
        {{ .csrfField }}
        <input type="submit" value="Submit">
    </form>
{{ end }}`,
	"create_post": `{{ define "title" }}New Post{{ end }}

{{ define "content" }}
    <h2>New Post</h2>
    <br>
    <form action="/posts/save" method="post">
        {{ .csrfField }}
        {{ template "post_form" .form }}
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
	"index": `{{ define "content" }}
<div class="auth">
{{ if .logged }}
<p>{{ .logged }} (<a href="/logout">Logout</a>)</p>
{{ else }}
<p><a href="/login">Login</a> <a href="/register">Register</a></p>
{{ end }}
</div>
{{ .motd }}
<p><a href="/posts/new">Write post</a> <a href="">Subscribe via Atom</a></p>

<section>
{{ range .posts -}}
    <article>
        <div><a href="/posts/{{ .Id }}">{{ .Title }}</a></div>
        <div><a href="/~{{ .User }}">{{ .User }}</a></div>
        <div>{{ .Date }}</div>
    </article>
{{ end }}
</section>
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
	"post": `{{ define "title "}}Post{{ end }}

{{ define "content"}}
    <h1>{{ .post.Title }}</h1>
    <p>
        <a href="/~{{ .post.User }}">~{{ .post.User }}</a>
        {{- if eq .logged .post.User }}
            <a href="/posts/{{ .post.Id }}/edit">Edit</a>
            <a href="/posts/{{ .post.Id }}/remove">Remove</a>
        {{- end }}
    </p>
    {{ .content }}
    <form action="/posts/{{ .post.Id }}/reply" method="post">
        {{ .csrfField }}
        <div class="field">
            <textarea name="reply"></textarea>
        </div>
        <input type="submit" value="Reply">
    </form>
    {{ template "reply" .replies }}
{{ end }}`,
	"register": `{{ define "title" }}Register{{ end }}

{{ define "content" }}
    <h2>Register</h2>
    <form action="/register" method="post" class="auth-form">
        {{ .csrfField }}
        <div class="field">
            <label for="name">Username</label>
            <input type="text" id="name" name="name" autocomplete="off"/>
        </div>
        <div class="field">
            <label for="password">Password</label>
            <input type="password" id="password" name="password"/>
        </div>
        <div class="field">
            <label for="key">Key</label>
            <input type="text" id="key" name="key"/>
        </div>
        <input type="submit" value="Submit">
    </form>
{{ end }}
`,
	"reply": `{{ define "content" }}
    <article>
        From: <a href="/~{{ .reply.Author }}">{{ .reply.Author }}</a><br>
        In: <a href="/posts/{{ .post.Id }}">{{ .post.Title }}</a><br>
        {{ if .reply.ParentId }}
            <a href="/replies/{{ .reply.ParentId }}">Parent</a><br>
        {{ end }}
        {{ .reply.Content }}
    </article>
    <section>
        <form action="/replies/{{ .reply.Id }}/save" method="post">
            {{ .csrfField }}
            <div class="field">
                <label for="content">Reply</label>
                <textarea name="reply" autofocus></textarea>
            </div>
            <input type="submit" value="Submit">
        </form>
        {{ template "reply" .reply.Thread }}
    </section>
{{ end }}
`,
}
