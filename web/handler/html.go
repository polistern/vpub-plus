// Code generated by go generate; DO NOT EDIT.

package handler

var TplMap = map[string]string{
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
	"index": `{{ define "content" }}
{{ if .logged }}
<p>Logged as {{ .logged }}</p>
<p><a href="/logout">Logout</a></p>
<p><a href="/posts/new">Write post</a></p>
{{ else }}
<p><a href="/register">Register</a></p>
<p><a href="/login">Login</a></p>
{{ end }}
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
    here{{ .content }}
    <form action="/todo" method="post">
        {{ .csrf }}
        <div class="field">
            <textarea name="reply"></textarea>
        </div>
        <input type="submit" value="Reply">
    </form>
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
}
