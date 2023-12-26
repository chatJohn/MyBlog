package dao

import (
	"MyBlog/models"
	"errors"
	"log"
)

func GetCategoryNameById(cid int) (string, error) {
	row := DB.QueryRow("select name from blog_category where cid = ?", cid)
	if row.Err() != nil {
		return "", errors.New("Query Category Name Error")
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName, nil
}
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
