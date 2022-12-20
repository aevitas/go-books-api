package api

func (s *Server) RegisterRoutes() {
	s.Gin.GET("/books/:isbn", s.HandleGetByISBN)
}
