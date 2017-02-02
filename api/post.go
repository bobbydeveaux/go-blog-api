package api

import (
	pb "github.com/bobbydeveaux/go-blog-proto/post"
)

func (a *API) GetPosts() []*pb.PostReply_Post {
	var posts []*pb.PostReply_Post
	post1 := &pb.PostReply_Post{
		PostID:    1,
		Content:   "My first post",
		Published: "2017-02-02 13:46:00"}

	post2 := &pb.PostReply_Post{
		PostID:    2,
		Content:   "My second post",
		Published: "2017-02-02 14:16:00"}

	posts = append(posts, post1, post2)

	return posts
}
