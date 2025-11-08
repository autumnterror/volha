package views

type (
	SWGSuccessResponse struct {
		Answer string `json:"answer" example:"operation successful"`
	}
	SWGIdResponse struct {
		Id string `json:"id" example:"dfhsa12341324"`
	}

	SWGErrorResponse struct {
		Error string `json:"error" example:"something went wrong"`
	}

	SWGFileUploadResponse struct {
		Name string `json:"name"`
		Mime string `json:"mime"`
	}

	SWGBrandListResponse struct {
		Brands []Brand `json:"brands"`
	}
	SWGCategoryListResponse struct {
		Categories []Category `json:"categories"`
	}
	SWGColorListResponse struct {
		Colors []Color `json:"colors"`
	}
	SWGCountryListResponse struct {
		Countries []Country `json:"countries"`
	}
	SWGMaterialListResponse struct {
		Materials []Material `json:"materials"`
	}
	SWGProductListResponse struct {
		Products []Product `json:"products"`
	}
	SWGProductColorPhotosListResponse struct {
		Items []ProductColorPhotos `json:"items"`
	}
	ProductColorPhotosId struct {
		ProductId string `json:"product_id"`
		ColorId   string `json:"color_id"`
	}
)
