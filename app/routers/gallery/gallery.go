package gallery

import (
	"bamboo-api/app/pkg/entity/dto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"

	"bamboo-api/app/clients"
)

func GetGalleryListProxy(c *gin.Context) {
	// use cache
	if galleryList, found := clients.Localcache.Get("gallery"); found {
		c.JSON(http.StatusOK, &dto.Response{
			Code: http.StatusOK,
			Msg:  "success",
			Data: galleryList.(string),
		})
		return
	}

	url := "https://api.ghostmarket.io/api/v1/assets?order_by=bid_price&order_direction=desc&offset=0&limit=10&with_total=0&fiat_currency=USD&auction_state=auction_ongoing&auction_started=started&chain=&grouping=1&only_verified=0&status=active&nsfw_mode=only_safe&price_similar=0&price_similar_delta=0&light_mode=0&blacklisted_mode=not_blacklisted&burned_mode=not_burned&auction_type=classic,reserve,dutch"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, &dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "get gallery err",
		})
		return
	}
	res, err := client.Do(req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, &dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "get gallery err",
		})
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, &dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "get gallery err",
		})
		return
	}

	// set cache
	clients.Localcache.Set("gallery", string(body), time.Minute*10)

	c.JSON(http.StatusOK, &dto.Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: string(body),
	})
}
