package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meilisearch/meilisearch-go"
	"kma_score_api/utils"
	"net/url"
)

func Search(c *fiber.Ctx) error {
	client := utils.MeilisearchClient
	query, err := url.QueryUnescape(c.Query("query", ""))

	if query == "" {
		return c.Status(200).JSON(utils.ApiResponse(200, "OK", fiber.Map{}))
	}
	search, err := client.Index("students").Search(query, &meilisearch.SearchRequest{})

	if err != nil {
		return c.Status(500).JSON(utils.ApiResponse(500, "Internal Server Error", err))
	}

	return c.Status(200).JSON(utils.ApiResponse(200, "OK", search.Hits))
}
