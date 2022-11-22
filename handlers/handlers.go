package handlers

import(
	"encoding/json"
	"io/ioutil"
	"github.com/balajisa09/url-shortener/pkg/apis/shortener"
	"github.com/gin-gonic/gin"
	"github.com/balajisa09/url-shortener/models"
	"regexp"
)

var ServerURL string

func HealthCheck(ctx *gin.Context){
	ctx.IndentedJSON(200,"App is Healthy")
}

func ShortenURL(ctx *gin.Context){
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.IndentedJSON(400,"Invalid Request")
		return
	}
	var req models.Request
	json.Unmarshal(data,&req)
	link := req.URL
	//validate link
	re := regexp.MustCompile(`http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
	matched := re.MatchString(link)
	if(!matched){
		ctx.IndentedJSON(400,"Invalid URL")
		return
	}
	//check if an empty request
	if link == ""{
		ctx.IndentedJSON(400,"Invalid Request")
		return
	}
	//shorten the url from client input
	resultPtr,err := shortener.Short(link)
	result := *resultPtr
	if(err!= nil){
		ctx.IndentedJSON(500,"Something went wrong")
		return
	}
	//take first six characters from the generated hash
	result = result[:6]
	ctx.IndentedJSON(200,ServerURL+result)
}

func GetURLHashMap(ctx *gin.Context){
	
	var respList []models.Response	
	readData, err := ioutil.ReadFile("db/values.json")
	if(err != nil){
		ctx.IndentedJSON(500,"Something went wrong")
		return
	}
	var readResp map[string]string
	json.Unmarshal(readData,&readResp)
	for key, value := range readResp{
		resp := models.Response{
			 URL: key,
			 Hash: value,
		}
		respList = append(respList, resp)
	}
	ctx.IndentedJSON(200,respList)
}

func init(){
	ServerURL = "http://localhost:8080/"
}