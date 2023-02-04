package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", server.AddMiddleware(PostRequest, Logging()))
	server.Handle("POST", "/user", server.AddMiddleware(UserPostRequest, Logging()))
	server.Handle("POST", "/api", server.AddMiddleware(HandleAPI, CheckAuth(), Logging()))
	server.Listen()
}
