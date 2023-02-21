package spider

type JavbusSpider struct {
	//主页
	BaseUrl string `json:"base_url"`
	//搜索url
	SearchFilmUrl string `json:"search_film_url"`
}

var (
	JavbusSpiderClient = JavbusSpider{
		BaseUrl:       "https://javbus.com",
		SearchFilmUrl: "https://www.javbus.com/search/{}&type=&parent=ce",
	}
)

// QueryFilmItemByCode
//
//	@Description: 根据番号查询作品栏
//	@receiver javbusSpider
//	@param filmCode
//	@return FilmItem
func (javbusSpider *JavbusSpider) QueryFilmItemByCode(filmCode string) FilmItem {
	return FilmItem{
		ItemImageUrl:  "",
		FilmTitle:     "",
		ResolutionTag: "",
		SubtitleTag:   "",
		FilmCode:      "",
		PublishData:   "",
	}
}
