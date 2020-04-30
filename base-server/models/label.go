package models

//-----------------------------------------------------------------------------

type Label struct {
	OriginId int    `gorm:"primary_key" json:"origin_id"`
	Text     string `json:"text"`
	Audio    string `json:"audio"`
}

func AddLabel(data map[string]interface{}) error {
	label := Label{
		OriginId: data["id"].(int),
		Text:     data["text"].(string),
		Audio:    data["audio"].(string),
	}
	err := db.Create(&label).Error
	return err
}
