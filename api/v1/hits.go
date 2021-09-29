package v1

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"hits/api/prisma/db"
	"hits/api/utils"
	. "hits/api/badge"
	. "hits/api/utils"
	"regexp"
	"strconv"
	"fmt"
	"strings"
)

func GetHits(c *fiber.Ctx) error {
	var url = c.Query("url")
	var svgQuery, _ = strconv.ParseBool(c.Query("svg"))
	var bgColorQuery = c.Query("bg")
	var client = utils.GetPrisma()
	var ctx = context.Background()
	const regex = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	match, _ := regexp.MatchString(regex, url)

	if url != "" && !match {
		return c.Status(400).JSON(Response{
			Success: false,
			Message: "Invalid URL",
		})
	} else if url == "" {
		return c.Status(400).JSON(Response{
			Success: false,
			Message: "URL cannot be empty",
		})
	}

	if bgColorQuery == fmt.Sprint(0) {
		bgColorQuery = "97ca00"
	}

	hit, err := client.Hits.FindUnique(
		db.Hits.URL.Equals(url),
	).Update(
		db.Hits.Hits.Increment(1),
	).Exec(ctx)

	if err != nil && err.Error() == "ErrNotFound" {
		createHit, createHitError := client.Hits.CreateOne(
			db.Hits.URL.Set(url),
		).Exec(ctx)

		if createHitError != nil {
			return c.Status(500).JSON(Response{
				Success: false,
				Message: "An internal server error occurred!",
			})
		}

		hit = createHit

	} else if err != nil {
		return c.Status(500).JSON(Response{
			Success: false,
			Message: "An internal server error occurred!",
		})
	}

	if svgQuery == true {
		svg := GenerateBadge(strconv.Itoa(hit.Hits), "000", fmt.Sprintf("#%s", strings.Trim(bgColorQuery, "\"")))
		c.Set(fiber.HeaderContentType, "image/svg+xml;charset=utf-8")
		c.Set(fiber.HeaderCacheControl, "max-age=0, s-maxage=0, must-revalidate, no-cache, no-store")
		return c.Send(svg)
	}

	return c.JSON(Response{
		Success: true,
		Message: "Successfully requested hit!",
		Data:    hit,
	})
}
