package dao

import (
	"MyBlog/models"
	"log"
)

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("Query Category Error: ", err)
	}
	var categorys []models.Category
	for rows.Next() {
		var category = models.Category{}
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("Get Category Error:", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
