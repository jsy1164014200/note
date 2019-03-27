```go
package main 
 
import ( 
    "bytes" 
    "encoding/json" 
    "fmt" 
    "github.com/gin-gonic/gin" 
    "io/ioutil" 
    "net/http" 
    "strings" 
) 
 
type AccessTokenResponse struct { 
    AccessToken string  `json:"access_token"` 
    ExpiresIn   float64 `json:"expires_in"` 
} 
 
type UserAccessToken struct { 
    UsAccessToken string `json:"access_token"` 
    ExpiresIn float64   `json:"expires_in"` 
    RefreshToken string `json:"refresh_token"` 
    OpenId string   `json:"openid"` 
    Scope string    `json:"scope"` 
 
} 
type UserInfo struct { 
    OpenId string   `json:"openid"` 
    NickName string `json:"nickname"` 
    Sex int         `json:"sex"` 
    Province string `json:"province"` 
    City string     `json:"city"` 
    Country string  `json:"country"` 
    HeadImgUrl string `json:"headimgurl"` 
 
} 
 
var AccessToken string = "" 
 
//var AppID string = "wx77a0cc6147eabc1e" 
//var AppSecret string = "332bc9d9cbfe2f49f041b52831ff5817" 
var AppID string = "wxb0a4d6588f1cd28b" 
var AppSecret string = "d0ac1b22de43ff9ff62f65dbf2aa6280" 
var MenuStr string = `{ 
    "button":[ 
        { 
            "name": "login", 
            "type": "view", 
            "url" : "http://jiangshiyi.top/login" 
        }, 
        { 
            "name": "managment", 
            "sub_button": [ 
                { 
                    "name":"user center", 
                    "type":"click", 
                    "key":"molan_user_center" 
                }, 
                { 
                    "name":"publish", 
                    "type":"click", 
                    "key":"molan_institution" 
                }] 
        } 
    ] 
}` 
 
func main() { 
    r := gin.Default() 
    r.GET("/wx", func(c *gin.Context) { 
        GetAndUpdateAccessToken() 
        fmt.Println("AccessToken:", AccessToken) 
        fmt.Println(AccessToken) 
        PushWxMenuCreate(AccessToken, []byte(MenuStr)) 
 
    }) 
    r.GET("/login",func(c *gin.Context) { 
        url := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" +AppID + "&redirect_uri=" +"ht
tp%3a%2f%2fjiangshiyi.top%2fhandleLogin" + "&response_type=code&scope=snsapi_userinfo#wechat_redirect" 
        c.Redirect(302,url) 
 
    }) 
    r.GET("/handleLogin",func(c *gin.Context) { 
        code := c.Query("code") 
 
 
        fmt.Println(code) 
        requestLine := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + AppID +"&secret=" + Ap
pSecret + "&code=" + code + "&grant_type=authorization_code" 
 
        resp, err := http.Get(requestLine) 
        if err != nil || resp.StatusCode != http.StatusOK{ 
            fmt.Println("�~O~Q�~@~A请�~B�~T~Y误",err) 
            c.String(http.StatusOK, "�~N��~O~V信�~A�失败") 
         
        } 
        defer resp.Body.Close() 
        body , _ := ioutil.ReadAll(resp.Body) 
 
        fmt.Println(string(body)) 
        fmt.Println("=========================================") 
 
        //fmt.Println(body) 
        userAccessToken := UserAccessToken{} 
        err = json.Unmarshal(body,&userAccessToken) 
        if err != nil { 
            fmt.Println("UserInfo json error") 
            c.String(http.StatusOK,"�~N��~O~V信�~A�fail") 
 
        }  
         
        //fmt.Print(userAccessToken) 
        info := GetReqBody("https://api.weixin.qq.com/sns/userinfo?access_token=" + userAccessToken.UsAcce
ssToken + "&openid=" + userAccessToken.OpenId + "&lang=zh_CN")  
        fmt.Println("=============================") 
        fmt.Println(string(info)) 
        userinfo := UserInfo{} 
        err = json.Unmarshal(info,&userinfo) 
        if err != nil{ 
            fmt.Println(" Userinfo json error") 
            c.String(http.StatusOK,"failed") 
        } 
         
        userStr := "�| �~Z~D�~T��~H��~P~M" + userinfo.NickName + "\n" + "�| �~Z~D头�~C~O" + userinfo.Hee
adImgUrl  
        c.String(http.StatusOK,userStr) 
 
    }) 
    r.Run(":80") // listen on 80 
} 
 
func GetReqBody(url string) []byte { 
        requestLine := url 
 
        resp, err := http.Get(requestLine) 
        if err != nil || resp.StatusCode != http.StatusOK{ 
            fmt.Println("�~O~Q�~@~A请�~B�~T~Y误",err) 
            return nil  
        } 
        defer resp.Body.Close() 
        body , _ := ioutil.ReadAll(resp.Body) 
        return body 
} 
 
 
func FetchAccessToken(appID, appSecret, accessTokenFetchUrl string) (string, error) { 
    requestLine := strings.Join([]string{accessTokenFetchUrl, "?grant_type=client_credential&appid=", 
        appID, 
        "&secret=", 
        appSecret}, "") 
 
    resp, err := http.Get(requestLine) 
    if err != nil || resp.StatusCode != http.StatusOK { 
        fmt.Println("�~O~Q�~@~A请�~B�~T~Y误", err) 
        return "", err 
    } 
    defer resp.Body.Close() 
    body, err := ioutil.ReadAll(resp.Body) 
    if err != nil { 
        fmt.Println("�~O~Q�~@~Aget请�~B�~N��~O~V atoken 读�~O~V�~T�~[~^body�~T~Y误", err) 
        return "", err 
    } 
 
    //fmt.Println(body) 
    //fmt.Println(string(body)) 
    //fmt.Println([]byte("access_token")) 
    if bytes.Contains(body, []byte("access_token")) { 
        atr := AccessTokenResponse{} 
        //  fmt.Println(body) 
        err = json.Unmarshal(body, &atr) 
        if err != nil { 
            fmt.Println("json�~U��~M�解�~^~P�~T~Y误") 
            return "", err 
        } 
        return atr.AccessToken, nil 
    } else { 
        fmt.Println("�~O~Q�~@~Aget�~L微信�~T�~[~^error") 
        return "", nil 
    } 
} 
 
func GetAndUpdateAccessToken() error { 
    accessToken, err := FetchAccessToken(AppID, AppSecret, "https://api.weixin.qq.com/cgi-bin/token") 
    if err != nil { 
        return err 
    } 
    AccessToken = accessToken 
    return nil 
} 
 
func PushWxMenuCreate(accessToken string, menuJsonBytes []byte) error { 
    postReq, err := http.NewRequest("POST", 
        strings.Join([]string{"https://api.weixin.qq.com/cgi-bin/menu/create", "?access_token=", accessTok
en}, ""), 
        bytes.NewReader(menuJsonBytes)) 
    if err != nil { 
        fmt.Println("request weixin fail") 
        return err 
    } 
 
    postReq.Header.Set("Content-Type", "application/json;encoding=utf-8") 
    client := &http.Client{} 
    resp, err := client.Do(postReq) 
    if err != nil { 
        fmt.Println("client request fail") 
        return err 
    } 
 
    body, err := ioutil.ReadAll(resp.Body) 
    if err != nil { 
        fmt.Println("client request fail@2 ") 
        return err 
    } 
    fmt.Println(string(body)) 
    fmt.Println("menu build success") 
    defer resp.Body.Close() 
 
    return nil 
}
```

