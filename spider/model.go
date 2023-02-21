package spider

// FilmItem
//
//	FilmItem
//	@Description: 电影作品栏简介
type FilmItem struct {
	ItemImageUrl  string `json:"item_image_url"`
	FilmTitle     string `json:"film_title"`
	ResolutionTag string `json:"resolution_tag"`
	SubtitleTag   string `json:"subtitle_tag"`
	FilmCode      string `json:"film_code"`
	PublishData   string `json:"publish_data"`
}

// FilmInfo
//
//	@Description: 影片信息
type FilmInfo struct {
	//封面
	FilmCoverUrl string `json:"film_cover_url"`
	//访问链接
	FilmPageUrl string `json:"film_page_url"`
	//演员作品链接
	StarPageUrl string `json:"star_page_url"`
	//作品标题
	FilmTitle string `json:"film_title"`
	//番号
	FilmCode string `json:"film_code"`
	//发布时间
	PublishDate string `json:"publish_date"`
	//导演
	Director string `json:"director"`
	//制作商
	Producer string `json:"producer"`
	//发行商
	Publisher string `json:"publisher"`
	//标签
	Tags string `json:"tags"`
	//系列
	FilmSeries string `json:"film_series"`
	//演员名
	StarName string `json:"star_name"`
	//磁力链接
	MagnetLinks []MagnetLink `json:"magnet_links"`
	//样品图
	SampleImages []SampleImage `json:"sample_images"`
}

// MagnetLink
//
//	MagnetLink
//	@Description: 磁力链接
type MagnetLink struct {
	MagnetName     string `json:"magnet_name"`
	MagnetFileSize string `json:"magnet_file_size"`
	ShareDate      string `json:"share_date"`
	//清晰度
	MagnetResolutionTag string `json:"magnet_resolution_tag"`
	//字幕
	MagnetSubtitleTag string `json:"magnet_subtitle_tag"`
}

// SampleImage
//
//	SampleImage
//	@Description: 样品图
type SampleImage struct {
	//图片链接
	Url string `json:"url"`
	//图片名
	FileName string `json:"file_name"`
	//图片数据
	ImageData []byte `json:"image_data"`
}
