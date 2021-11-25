package handlers

import (
	"net/http"
)

func (h *BaseHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// var projects []model.Project
	// result, err := db.Query("SELECT id, title from posts")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer result.Close()
	// for result.Next() {
	// 	var post Post
	// 	err := result.Scan(&post.ID, &post.Title)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	posts = append(posts, post)
	// }
	// json.NewEncoder(w).Encode(posts)
}
