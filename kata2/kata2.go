package main

import "fmt"
import "flag"
import "github.com/gin-gonic/gin"
import "strconv"
import "math/rand"
import "time"
import "bytes"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lengthLetters = len(letterRunes)

func init() {
	rand.Seed(time.Now().UnixNano())
}

/**
 * Returns a random string with the given length
 */
func randomString(length int) string {
	var buffer bytes.Buffer

	for i := 0; i < length; i++ {
		num := rand.Intn(lengthLetters)
		buffer.WriteRune(letterRunes[num])
	}

	return buffer.String()
}

func randomProductXML() string {

	return fmt.Sprintf("\n\t<Product>\n\t\t<sku>%s</sku>\n\t\t<quantity>%d</quantity>\n\t</Product>", randomString(40), rand.Intn(100))
}

func fakeLoad() {

	probability := rand.Intn(100)
	var delay int
	if probability < 5 {
		delay = 0
	} else if (probability >= 5) && (probability < 25) {
		delay = rand.Intn(10)
	} else if probability >= 25 && probability <= 75 {
		delay = rand.Intn(50) + 50
	} else {
		delay = rand.Intn(500) + 200
	}

	//fmt.Printf("Probability was %d and delay was %d\n", probability, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))

}

func buildRandomXML(products int) string {
	fmt.Printf("Will generated %d products",products)
	var buffer bytes.Buffer
	buffer.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	buffer.WriteString("<ProductList>")
	for i := 0; i < products; i++ {
		buffer.WriteString(randomProductXML())
	}
	buffer.WriteString("\n</ProductList>")
	return buffer.String()

}

func runServer(port int) {
	r := gin.Default()
	r.GET("/products", func(c *gin.Context) {
		errRand := rand.Intn(100)
		if errRand < 10 {
			c.String(500, "Internal server error")
		} else {
			fakeLoad()
			c.String(200, buildRandomXML(rand.Intn(10)))
		}

	})
	r.Run(":" + strconv.Itoa(port))
}

func main() {
	var port, rc int

	//Parse the port value from the command line
	flag.IntVar(&port, "port", 8081, "The port in which the server will run")
	flag.IntVar(&rc, "rc", 50, "The number of random characters to generate")

	flag.Parse()

	fmt.Printf("Starting server at port %d\n", port)
	fmt.Printf("Asked to generate a string with %d chars\n", rc)

	randomString := randomString(rc)

	fmt.Printf("Generated random string = %s\n", randomString)

	fmt.Printf("Generated product string = %s\n", randomProductXML())
	runServer(port)

}
