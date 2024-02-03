package response

import "github.com/vsayfb/e-commerce-scrapper/product"

type Response struct {
	Keyword string            `json:"keyword"`
	Page    uint8             `json:"page"`
	Data    []product.Product `json:"data"`
}

func New(keyword string, page uint8, data []product.Product) *Response {

	return &Response{
		Keyword: keyword,
		Data:    data,
		Page:    page,
	}

}

func (r *Response) ConvertToInterfaceSlice() []interface{} {

	var interfaceSlice []interface{}

	for _, v := range r.Data {
		interfaceSlice = append(interfaceSlice, v)
	}

	return interfaceSlice
}
