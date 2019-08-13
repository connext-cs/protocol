package api

type (
	ListReq struct {
		Page     int `json:"page"`
		PageSize int `json:"page_size"`
	}
	ListResNode struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Url         string `json:"url"`
		Internalurl string `json:"internalurl"`
		Desc        string `json:"desc"`
	}
	ListRes struct {
		List  []ListResNode `json:"list"`
		Total int           `json:"total"`
	}
)

type (
	AddReq struct {
		Name string `json:"name"`
		Url  string `json:"url"`
		Desc string `json:"desc"`
	}
	AddRes struct {
	}
)

type (
	DeleteReq struct {
		Id int `json:"id"`
	}
	DeleteRes struct {
	}
)

type (
	UpdateReq struct {
		Id   int    `json:"id"`
		Url  string `json:"url"`
		Desc string `json:"desc"`
	}
	UpdateRes struct {
	}
)

type (
	UrlcheckReq struct {
		Url        string `json:"url"`
		ChangeFlag int    `json:"change_flag"`
	}
	UrlcheckRes struct {
	}
)

type (
	NamelistReq struct {
	}
	NamelistResNode struct {
		Name string `json:"name"`
	}
	NamelistRes struct {
		List  []NamelistResNode `json:"list"`
		Total int               `json:"total"`
	}
)
