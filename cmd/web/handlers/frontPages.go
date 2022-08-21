package handlers

import (
	"fmt"
	"net/http"
)

func ShowBlogFront(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Blog 前台")
	// http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}

// TODO:
// 优雅关机，处理消息
