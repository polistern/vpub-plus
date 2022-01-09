// Code generated by go generate; DO NOT EDIT.

package assets

var AssetsMap = map[string]string{
	"style": `body {
	max-width: 800px;
	margin: auto;
	padding: 5px;
}

/* Lists ************************************************************/

ol.posts, ol.replies {
	padding: 0;
	list-style: none;
}

/*ol.posts > li:not(:last-child),*/
ol.replies > li:not(:last-child) {
	margin-bottom: 1em;
}

/* Posts ************************************************************/

table {
	border-collapse: collapse;
	border: 1px solid;
	width: 100%;
}
tr, td, th {
	vertical-align: top;
	border: 1px solid;
	padding: .5em;
}
.post-aside {
	text-align: center;
	width: 150px;
}

.content h1 { font-size: 1.5em; }
.content h2 { font-size: 1.2em; }
.content h3 { font-size: 1em; }
.content { margin: 1em 0; }

/* Navigation *******************************************************/

header > nav {
	float: right;
}

body > footer {
	margin-top: 1em;
	border-top: 1px solid lightgrey;
	color: grey;
	text-align: center;
}

/* Forms ************************************************************/

.auth-form {
	max-width: 200px;
}

.field {
	margin-bottom: 1em;
}

.field label {
	display: block;
}

input[type=text], input[type=password] {
	width: 100%;
	box-sizing: border-box;
}

textarea {
	width: 100%;
	height: 250px;
	display: block;
	box-sizing: border-box;
}

/* Misc *************************************************************/

blockquote {
	margin: 0;
	color: green;
	font-style: italic;
}

.center { text-align: center; }
.grow { width: 100%; }`,
}
