package views

type ProductSearch struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Article string `json:"article"`
}

type Product struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Article     string     `json:"article"`
	Brand       Brand      `json:"brand"`
	Category    Category   `json:"category"`
	Country     Country    `json:"country"`
	Width       int        `json:"width"`
	Height      int        `json:"height"`
	Depth       int        `json:"depth"`
	Materials   []Material `json:"materials"`
	Colors      []Color    `json:"colors"`
	Photos      []string   `json:"photos"`
	Seems       []Product  `json:"seems"`
	Price       int        `json:"price"`
	Description string     `json:"description"`
}

type ProductId struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Article     string   `json:"article"`
	Brand       string   `json:"brand"`
	Category    string   `json:"category"`
	Country     string   `json:"country"`
	Width       int      `json:"width"`
	Height      int      `json:"height"`
	Depth       int      `json:"depth"`
	Materials   []string `json:"materials"`
	Colors      []string `json:"colors"`
	Photos      []string `json:"photos"`
	Seems       []string `json:"seems"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
}

type Brand struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Uri   string `json:"uri"`
	Img   string `json:"img"`
}

type Country struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Friendly string `json:"friendly"`
}

type Material struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type Color struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

type ProductFilter struct {
	Brand     []string `json:"brand"`
	Country   []string `json:"country"`
	Category  []string `json:"category"`
	MinWidth  int      `json:"min_width"`
	MaxWidth  int      `json:"max_width"`
	MinHeight int      `json:"min_height"`
	MaxHeight int      `json:"max_height"`
	MinDepth  int      `json:"min_depth"`
	MaxDepth  int      `json:"max_depth"`
	Materials []string `json:"materials"`
	Colors    []string `json:"colors"`
	MinPrice  int      `json:"min_price"`
	MaxPrice  int      `json:"max_price"`
	SortBy    string   `json:"sort_by"`
	SortOrder string   `json:"sort_order"`
	Offset    int      `json:"offset"`
	Limit     int      `json:"limit"`
}

type Dictionaries struct {
	Brands     []Brand    `json:"brands"`
	Categories []Category `json:"categories"`
	Countries  []Country  `json:"countries"`
	Materials  []Material `json:"materials"`
	Colors     []Color    `json:"colors"`

	MinPrice  int `json:"min_price"`
	MaxPrice  int `json:"max_price"`
	MinWidth  int `json:"min_width"`
	MaxWidth  int `json:"max_width"`
	MinHeight int `json:"min_height"`
	MaxHeight int `json:"max_height"`
	MinDepth  int `json:"min_depth"`
	MaxDepth  int `json:"max_depth"`
}
type ProductColorPhotos struct {
	ProductId string   `json:"product_id"`
	ColorId   string   `json:"color_id"`
	Photos    []string `json:"photos"`
}
