package twitter

import (
	"fmt"
	"net/http"

	"bamboo-api/app/pkg/entity/dto"
	"bamboo-api/app/service"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoginByTwitter(c *gin.Context) {
	//walletId := c.GetHeader("walletId")
	failedResp := &dto.Response{
		Code: http.StatusBadRequest,
		Msg:  "params invalidÏ",
	}
	walletId, exists := c.GetQuery("walletId")
	if !exists {
		c.JSON(http.StatusBadRequest, failedResp)
		return
	}
	userCallbackUrl, exists := c.GetQuery("callback_url")
	if !exists {
		c.JSON(http.StatusBadRequest, failedResp)
		return
	}
	oc := NewTWClient()
	rt, err := oc.RequestTemporaryCredentials(nil, callbackURL, nil)
	if err != nil {
		log.Errorf("[LoginByTwitter] RequestTemporaryCredentials failed=%+v", err)
		c.JSON(http.StatusBadRequest, failedResp)
		return
	}

	session := sessions.Default(c)
	session.Set("request_token", rt.Token)
	session.Set("request_token_secret", rt.Secret)
	session.Set("wallet_id", walletId)
	session.Set("callback_url", userCallbackUrl)
	session.Save()

	url := oc.AuthorizationURL(rt, nil)

	c.Redirect(http.StatusMovedPermanently, url)
	return
}

func TwitterCallback(c *gin.Context) {
	failedResp := &dto.Response{
		Code: http.StatusBadRequest,
		Msg:  "params invalidÏ",
	}
	tok := c.DefaultQuery("oauth_token", "")
	if tok == "" {
		c.JSON(http.StatusBadRequest, failedResp)
		return
	}

	ov := c.DefaultQuery("oauth_verifier", "")
	if ov == "" {
		c.JSON(http.StatusBadRequest, failedResp)
		return
	}

	session := sessions.Default(c)
	v := session.Get("request_token")
	if v == nil {
		c.JSON(http.StatusBadRequest, failedResp)
		return
	}
	rt := v.(string)
	if tok != rt {
		c.JSON(http.StatusBadRequest, failedResp)
		return
	}
	walletId, success := session.Get("wallet_id").(string)

	if !success {
		c.JSON(http.StatusBadRequest, failedResp)
		return
	}

	userCallbackUrl, success := session.Get("callback_url").(string)

	if !success {
		c.JSON(http.StatusBadRequest, failedResp)
		return
	}
	v = session.Get("request_token_secret")
	if v == nil {
		c.JSON(http.StatusBadRequest, failedResp)
		return
	}
	rts := v.(string)
	if rts == "" {
		c.JSON(http.StatusBadRequest, failedResp)
		return
	}

	code, at, err := GetAccessToken(&oauth.Credentials{Token: rt, Secret: rts}, ov)
	if err != nil {
		c.JSON(http.StatusBadRequest, failedResp)
		return
	}

	account := struct {
		ID         string `json:"id_str"`
		ScreenName string `json:"screen_name"`
	}{}
	code, err = GetMe(at, &account)
	if err != nil {
		c.JSON(code, failedResp)
		return
	}

	// TODO use id to make user login.
	fmt.Println(account)
	err = service.UserService.BindTwitter(walletId, account.ID, fmt.Sprintf("https://twitter.com/%v", account.ScreenName), at.Token, at.Secret)

	if nil != err {
		log.Errorf("[TwitterCallback] call user.bindtwitter service failed. err=%+v", err)
	}
	c.Redirect(http.StatusMovedPermanently, userCallbackUrl)
	return
}
