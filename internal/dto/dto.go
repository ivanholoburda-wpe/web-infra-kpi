package dto

import "time"

type CreateSiteRequest struct {
	Name string `json:"name"`
	URL  string `json:"url" binding:"required,url"`
}

type SiteResponse struct {
	ID           uint      `json:"id"`
	URL          string    `json:"url"`
	HttpStatus   int       `json:"http_status"`
	ResponseTime int       `json:"response_time"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
