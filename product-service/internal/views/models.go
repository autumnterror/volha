package views

type Product struct {
	Id          string
	Title       string
	Article     string
	Brand       Brand
	Category    Category
	Country     Country
	Width       int
	Height      int
	Depth       int
	Materials   []Material
	Colors      []Color
	Photos      []string
	Seems       []Product
	Price       int
	Description string
}

type ProductId struct {
	Id          string
	Title       string
	Article     string
	Brand       string
	Category    string
	Country     string
	Width       int
	Height      int
	Depth       int
	Materials   []string
	Colors      []string
	Photos      []string
	Seems       []string
	Price       int
	Description string
}
type Brand struct {
	Id   string
	Name string
}
type Category struct {
	Id    string
	Title string
	Uri   string
	Img   string
}

type Country struct {
	Id       string
	Title    string
	Friendly string
}

type Material struct {
	Id    string
	Title string
}

type Color struct {
	Id   string
	Name string
	Hex  string
}

type ProductFilter struct {
	Brand     []string
	Country   []string
	Category  []string
	MinWidth  int
	MaxWidth  int
	MinHeight int
	MaxHeight int
	MinDepth  int
	MaxDepth  int
	Materials []string
	Colors    []string
	MinPrice  int
	MaxPrice  int
	SortBy    string
	SortOrder string
	Offset    int
	Limit     int
}

type Dictionaries struct {
	Brands     []Brand
	Categories []Category
	Countries  []Country
	Materials  []Material
	Colors     []Color

	MinPrice  int
	MaxPrice  int
	MinWidth  int
	MaxWidth  int
	MinHeight int
	MaxHeight int
	MinDepth  int
	MaxDepth  int
}

type ProductSearch struct {
	Id      string
	Title   string
	Article string
}
type ProductColorPhotos struct {
	ProductId string
	ColorId   string
	Photos    []string
}
