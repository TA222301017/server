package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"server/api/middlewares"
	"server/api/services"
	"server/api/utils"
	"server/models"
	"server/setup"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterSSERoutes(app *gin.RouterGroup) {
	router := app.Group("/sse")

	router.GET("/token", middlewares.Auth(), func(c *gin.Context) {
		token := rand.Uint32()

		setup.Channel.SSETokens[token] = false
		go func() {
			time.Sleep(time.Minute * 5)
			delete(setup.Channel.SSETokens, token)
		}()

		utils.MakeResponseSuccess(c, token, nil)
	})

	router.GET("/rssi", middlewares.SSE(), func(c *gin.Context) {
		keyword := c.Query("keyword")
		token, err := strconv.ParseUint(c.Query("token"), 10, 32)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		go func() {
			for i := 0; i < 50; i++ {
				setup.Channel.RSSIMessage <- &models.RSSILog{
					PersonelID: 7,
					RSSI:       -77,
					LockID:     1,
					Timestamp:  time.Now(),
					KeyID:      15,
					Personel: models.Personel{
						Name: "Jon",
					},
				}

				setup.Channel.RSSIMessage <- &models.RSSILog{
					PersonelID: 7,
					RSSI:       -77,
					LockID:     18,
					Timestamp:  time.Now(),
					KeyID:      15,
					Personel: models.Personel{
						Name: "Jon",
					},
				}

				setup.Channel.RSSIMessage <- &models.RSSILog{
					PersonelID: 7,
					RSSI:       -77,
					LockID:     26,
					Timestamp:  time.Now(),
					KeyID:      15,
					Personel: models.Personel{
						Name: "Jon",
					},
				}

				time.Sleep(2 * time.Second)
			}
		}()

		if _, ok := setup.Channel.SSETokens[uint32(token)]; !ok {
			utils.ResponseBadRequest(c, errors.New("invalid token"))
			return
		} else {
			delete(setup.Channel.SSETokens, uint32(token))
		}

		clientChan := make(chan *models.RSSILog)

		setup.Channel.NewRSSIClients <- clientChan
		defer func() {
			setup.Channel.ClosedRSSIClients <- clientChan
		}()

		c.Stream(func(w io.Writer) bool {
			if msg, ok := <-clientChan; ok {
				data, err := services.MatchRSSILogEvent(msg, keyword)
				if err != nil {
					return false
				}

				jsonData, err := json.Marshal(data)
				if err != nil {
					return false
				}

				c.SSEvent("rssi", string(jsonData))
				return true
			}
			return false
		})
	})

	router.GET("/access", middlewares.SSE(), func(c *gin.Context) {
		keyword := c.Query("keyword")
		token, err := strconv.ParseUint(c.Query("token"), 10, 32)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		if _, ok := setup.Channel.SSETokens[uint32(token)]; !ok {
			utils.ResponseBadRequest(c, errors.New("invalid token"))
			return
		} else {
			delete(setup.Channel.SSETokens, uint32(token))
		}

		clientChan := make(chan *models.AccessLog)

		setup.Channel.NewAccessClients <- clientChan
		defer func() {
			setup.Channel.ClosedAccessClients <- clientChan
		}()

		c.Stream(func(w io.Writer) bool {
			if msg, ok := <-clientChan; ok {
				data, err := services.MatchAccessLogEvent(msg, keyword)
				if err != nil {
					return false
				}

				jsonData, err := json.Marshal(data)
				if err != nil {
					return false
				}

				c.SSEvent("access", string(jsonData))
				return true
			}
			return false
		})
	})
}
