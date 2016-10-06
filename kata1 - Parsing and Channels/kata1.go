package main

import "encoding/xml"
import "encoding/json"
import "fmt"
import "time"
import "math/rand"

type Product struct{
	Sku	string `xml:"sku" json:"sku"` 
	Quantity int `xml:"quantity" json:"quantity"` 
}

type Stock struct{
	ProductList []Product `xml:"Product" json:"products"` 
}

func sleep(){
	time.Sleep(time.Duration(rand.Int31n(1000))* time.Millisecond)
}

type Converter struct{
	close chan string
}
/**
 * Converts the given structure to XML and then into JSON
 */
func (c Converter) convert(data []byte,num int) (string,error) {

	fmt.Printf("Starting conversion! #%d\n",num)
	sleep()
	v:= Stock{}
	fmt.Printf("Unmarshal to XML! #%d\n",num)
	err:=xml.Unmarshal(data,&v)
	if nil!=err{
		return "",nil
	}
	sleep()
	fmt.Printf("Marshal to JSON! #%d\n",num)
	output,err := json.Marshal(v)
	if nil!=err{
		return "",nil
	}
	sleep()
	formattedOutput := string(output)
	c.close <-formattedOutput
	return formattedOutput,nil

}

func main(){

rand.Seed(time.Now().Unix())

xmlData := []byte(`<?xml version="1.0" encoding="UTF-8" ?>
<ProductList>
    <Product>
        <sku>ABC123</sku>
        <quantity>2</quantity>
    </Product>
    <Product>
        <sku>ABC124</sku>
        <quantity>20</quantity>
    </Product>
</ProductList>`)

var converter Converter
done:=make(chan string,10)
for i:=1;i<10;i++{
	converter = Converter{done}
	go converter.convert(xmlData,i)
}

//We just get the first one and finish!

	
fmt.Println(<-done)
}