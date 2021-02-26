package lib

type NewsModel struct {
	NewsID      uint32 `gorm:"column:id" json:"id"`
	NewsTitle   string `gorm:"column:newstitle" json:"title"`
	NewsDesc    string `gorm:"column:newsdesc" json:"desc"`
	NewsContent string `gorm:"column:newscontent" json:"content"`
	NewsViews   uint32 `gorm:"column:views" json:"views"`
	NewsTime    string `gorm:"column:addtime" json:"time"`
}

func NewNewsModel() *NewsModel {
	return &NewsModel{}
}
