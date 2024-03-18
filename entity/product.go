package entity

type Product struct {
	ID   int    `json:"id" gorm:"column:id;primaryKey"`
	Name string `json:"name" gorm:"column:name"`
}
