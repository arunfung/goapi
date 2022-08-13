package category

import (
    "goapi/pkg/app"
    "goapi/pkg/database"
    "goapi/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idStr string) (category Category) {
    database.DB.Where("id", idStr).First(&category)
    return
}

func GetBy(field, value string) (category Category) {
    database.DB.Where("? = ?", field, value).First(&category)
    return
}

func All() (categories []Category) {
    database.DB.Find(&categories)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Category{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Category{}),
        &categories,
        app.V1URL(database.TableName(&Category{})),
        perPage,
    )
    return
}