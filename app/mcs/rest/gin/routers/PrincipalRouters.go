package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"../../../../../dao/factory"
	"../../../../../utilities"
	"../../../../../models"
	"../../../../../models/structfaces"

	"log"


	//"database/sql"
	"io/ioutil"
	"bytes"
	"net/http"
)

func GetInstructions(c *gin.Context) {
	c.JSON(200, gin.H{"part1": "welcome to the jungle"})
}

func FinAllRadicacion(c *gin.Context) {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}
	radicacionDao := factory.FactoryDaoRadicacion(config.Engine)
	radicaResult ,_ := radicacionDao.GetAllRadicacion()
	c.JSON(200, radicaResult)
}
func FindByIdRadicacion(c *gin.Context) {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}
	radicacionDao := factory.FactoryDaoRadicacion(config.Engine)
//fmt.Println(c.Request.Body)
	radicaResult ,_ := radicacionDao.FindRadicacion(c.Query("idRadicacion"))
	c.JSON(200, radicaResult)
}

func SetNewRadicacion(c *gin.Context) {
	response :=_struct.Response{}
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}

	var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	radicacion := new(models.Radicacion)
	error := c.Bind(radicacion)
//error := binding.JSON.Bind(c.Request,radicacion)



	/*radicacion := models.Radicacion{}
//	radicacion.NumeroRadicacion  = c.Query("NumeroRadicacion")
	radicacion.NumeroRadicacion = c.PostForm("NumeroRadicacion")
	radicacion.FechaRadicacion,_ = time.Parse(time.RFC3339,c.PostForm("FechaRadicacion"))
	radicacion.FechaDocumento,_ = time.Parse(time.RFC3339,c.PostForm("FechaDocumento"))
	radicacion.Asunto = c.PostForm("Asunto")
	radicacion.IDRemitente = c.PostForm("IDRemitente")
	radicacion.PdfIdPdf = utilities.ToNullInt64(c.PostForm("PdfIdPdf"))
	radicacion.TiempoRespuesta,_ = time.Parse(time.RFC3339,c.PostForm("TiempoRespuesta"))*/

	radicacionDao := factory.FactoryDaoRadicacion(config.Engine)
	affect ,error := radicacionDao.CreateRadicacion(*radicacion);
	if error != nil || affect == 0{
		response.Status = 500
		response.Message ="insertion error in radicacion #"+radicacion.NumeroRadicacion+" &[SetNewRadicacion]"
		response.Error = error
		fmt.Println(error)
		//c.JSON(200, gin.H{"creationError": error})
	}else{
		response.Status = 200
	response.Message = "radicacion #"+radicacion.NumeroRadicacion+" created"
	response.Error = nil
	}

	c.JSON(response.Status, response)

}

func GetDelivery(c *gin.Context) {
	c.XML(200, gin.H{"part2": "If you want it you're gonna bleed but it's the price to pay"})
}


func PostConsoleParams(c *gin.Context) {
	id := c.Query("id")
	name := c.DefaultQuery("name", "0")
	valor1 := c.PostForm("valor1")
	message := c.PostForm("message")
	fmt.Printf("id: %s; valor1: %s; name: %s; message: %s", id, valor1, name, message)
}

func CaptureFile(c *gin.Context)  {
	var filePath = "C:/Users/Software1/Desktop/"
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	if err := c.SaveUploadedFile(file, "C:/Users/Software1/Desktop/"+file.Filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("File save in %s uploaded successfully ", filePath+file.Filename))
}

//---
