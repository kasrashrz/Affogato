package jwt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/utils/jwt"
	"net/http"
	"net/url"
)

func HandleGoogleLogin(ctx *gin.Context) {
	fmt.Println("HI")
	ctx.Redirect(http.StatusPermanentRedirect, "https://igm.games/api/mc/login")
}

type LoginRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	TwoFactor string `json:"2fa_code"`
}

type ExternalAPIResponse struct {
	Success   string `json:"success"`
	Token     string `json:"token"`
	SecretKey string `json:"secretKey" json:"-"`
	Email     string `json:"email"`
	Balance   int    `json:"balance"`
	Level     int    `json:"level"`
	Width     int    `json:"width"`
	XP        int    `json:"XP"`
	District  int    `json:"district"`
	Error     string `json:"error"`
}

func FinalLogin(c *gin.Context) {
	// Parse the request body to extract the user's credentials
	var loginReq LoginRequest
	var userDAO domain.User
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Make a request to the external API
	apiURL := "https://igm.games/api/mc/login"
	apiReq := LoginRequest{
		Email:     loginReq.Email,
		Password:  loginReq.Password,
		TwoFactor: loginReq.TwoFactor,
	}
	apiReqBytes, err := json.Marshal(apiReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal API request"})
		return
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(apiReqBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call external API"})
		return
	}
	defer resp.Body.Close()

	// Read the response from the external API
	var apiResp ExternalAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse API response"})
		return
	}

	var token string
	if apiResp.Error == "" {
		token, _ = jwt.GenerateJWT(apiResp.Email, "affogato")
	}

	// Save token in Database
	if token != "" {
		_ = userDAO.ChangeToken(apiResp.Email, apiResp.Token, apiResp.SecretKey)
	}

	// Send the API response back to the user
	c.JSON(http.StatusOK, gin.H{"token": token, "error": apiResp.Error})
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HII")
	URL, err := url.Parse("https://igm.games/api/mc/login")
	if err != nil {
		fmt.Println("Parse: " + err.Error())
	}
	fmt.Println(URL.String())
	parameters := url.Values{}
	//parameters.Add("client_id", oauthConf.ClientID)
	//parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	//parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	//parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		//secret := ctx.GetHeader("Secret")

		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}
		err := jwt.ValidateToken(tokenString, "affogato")
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
