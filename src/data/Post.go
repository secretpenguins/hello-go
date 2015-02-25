package data

import (
	"log"
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"strconv"
)

type Post struct {
	PostId int64 `form:"PostId"`
	Title string `form:"Title" binding:"required"`
	Content  string `form:"Content"`
}

func GetPost(id string) *Post {
	var post Post

	var key string = "post-" + id
		
	item, _ := mc.Get(key)
	if item == nil {
		err := dbMap.SelectOne(&post, "select postId, title, content from Posts where postId=?", id)
 		if err != nil {
 			log.Println(err)
 		}

		bytes, _ := json.Marshal(&post)

		item := &memcache.Item {
			Key: key,
			Value: bytes,
		}
		mc.Set(item)

		log.Println("Cache Miss")
	} else {
		json.Unmarshal(item.Value, &post)
		log.Println("Cache Hit")
	}

	return &post
}

func GetPosts() []Post {
	var posts []Post
	_, err := dbMap.Select(&posts, "select * from Posts")
	if (err != nil) {
		log.Println(err)
	}
	return posts
}

func (post *Post) Save() {
	var key string = "post-" + strconv.Itoa((int)(post.PostId))
	log.Println("key ", key)
	log.Println("")
	dbMap.Update(&post)
	mc.Delete(key)
}