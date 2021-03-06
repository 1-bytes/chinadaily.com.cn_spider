package parser

type JsonData struct {
	ID string `json:"id"`
	//ArticleID string    `json:"article_id"`
	//Title       string    `json:"title"`
	//Author      string    `json:"author"`
	//ReleaseDate string    `json:"release_date"`
	//Category    string    `json:"category"`
	SourceDomain string    `json:"source_domain"`
	SourceURL    string    `json:"source_url"`
	Paragraph    paragraph `json:"paragraph"`
}

type DictArticleModel struct {
	ID                  int    `gorm:"type:int; primaryKey; autoIncrement; unsigned; not null" json:"id"`
	Type                int    `gorm:"type:int; not null" json:"type"`
	Title               string `gorm:"type:varchar(255); not null" json:"title"`
	Author              string `gorm:"type:varchar(128); not null" json:"author"`
	ReleaseDate         string `gorm:"type:varchar(128); not null" json:"release_date"`
	MostRecentlyUpdated string `gorm:"type:varchar(128); not null" json:"most_recently_updated"`
	SourceDomain        string `gorm:"type:varchar(255); not null" json:"source_domain"`
	SourceUrl           string `gorm:"type:varchar(255); not null" json:"source_url"`
}

var TypeMap = map[string]int{
	"news":          1, // 新闻
	"science":       2, // 科学、科普
	"humors":        3, // 笑话、幽默
	"novel":         4, // 小说
	"entertainment": 5, // 娱乐
	"poems":         6, // 诗歌、诗集
	"essays":        7, // 散文、文集
	"story":         8, // 故事
	"speech":        9, // 演讲
}
