package sysinit

import (
	"fmt"
	"io/ioutil"
	v1 "k8s.io/api/networking/v1"
	"log"
	"sigs.k8s.io/yaml"
)

type Server struct {
	Port int
}

type Config struct {
	Server
	Ingress []v1.Ingress
}

var ServerConfig = &Config{}

func Conf() {
	//f, err := os.Open("ingress_test.yaml")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fb, err := io.ReadAll(f)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	f, err := ioutil.ReadFile("ingress_test.yaml")
	if err != nil {
		log.Fatal(err)
		//return  err
	}
	if err := yaml.Unmarshal(f, ServerConfig); err != nil {
		log.Fatal(err)
	}
	fmt.Println(ServerConfig)

	//

	ParseRules()

}
