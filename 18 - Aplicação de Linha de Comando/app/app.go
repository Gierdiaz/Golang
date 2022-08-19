package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidor na internet"
	
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "",	
		},
	}

	app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Busca IPs de endereço na internet",
			Flags: flags,
			Action: BuscarIps,
		},
		{
			Name: "servidores",
			Usage: "Busca o nome dos servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}
	
	return app
}

func BuscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

//Execução de comando
//go run main.go servidores --host google.com
//go run main.go ip --host google.com

//go build
//./linha-de-comando ip --host amazon.com
//./linha-de-comando servidores --host amazon.com