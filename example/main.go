package main

import (
	"fmt"
	"log"

	"github.com/awildboop/mcmgo"
)

func main() {
	ses, err := mcmgo.NewSession("MlG/PQ4MqQliDZOq2716HY/910UDXg8o", false)
	if err != nil {
		log.Fatalln(err)
	}

	req := ses.HttpClient.R()
	fmt.Println(req.Get("conversations"))
}
