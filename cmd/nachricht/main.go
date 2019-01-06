/*
Copyright 2019 Daniël Franke

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleListParams struct {
	Published bool   `json:"published"`
	Order     string `json:"order"`
}

func main() {
	router := gin.Default()

	apiv1 := router.Group("/v1")
	{
		articles := apiv1.Group("/:site/articles")
		{
			articles.GET("", ArticleList)
			articles.POST("", ArticlePost)
			articles.GET(":articleSlug", ArticleGet)
			articles.PUT(":articleSlug", ArticlePut)
		}
	}

	router.Run(":3000")
}

func ArticleList(c *gin.Context) {
	var params ArticleListParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, params)
}

func ArticlePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ArticlePost not handled yet.",
	})
}

func ArticleGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ArticleGet not handled yet.",
	})
}

func ArticlePut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ArticlePut not handled yet.",
	})
}
