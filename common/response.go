package common

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginationResponse struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	Total    int64       `json:"total,omitempty"`
	Page     int         `json:"page,omitempty"`
	PageSize int         `json:"pageSize,omitempty"`
	HasMore  bool        `json:"hasMore,omitempty"`
}
