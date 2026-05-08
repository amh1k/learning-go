package blogposts

import (
	"errors"

	"io/fs"
)


type StubFailingFs struct {

}
func (s StubFailingFs) Open(name string)(fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}
func NewPostsFromFs(fileSystem fs.FS) ([]Post, error){
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts [] Post
	for  _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func getPost(fileSystem fs.FS, fileName string)(Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}



