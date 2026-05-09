package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	blogrenderer "github.com/amh1k/blogrenderer"
	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var(
		aPost = blogrenderer.Post {
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("it converts a single file intp html", func(t *testing.T) {
		buf := bytes.Buffer{}
		
		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())


	})
// 	t.Run("converts single post to html", func(t *testing.T) {
// 		buf := bytes.Buffer{}
// 		err := blogrenderer.Render(&buf, aPost)
// 		if err != nil {
// 			t.Fatal(err)

// 		}
// 		got := buf.String()
// 		want := `<h1>hello world</h1>
// <p>This is a description</p>
// Tags: <ul><li>go</li><li>tdd</li></ul>`
// 		if got != want {
// 			t.Errorf("got '%s' want '%s'", got, want)

// 		}
// 	})


	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		postRenderer.Render(io.Discard, aPost)
	}
}

