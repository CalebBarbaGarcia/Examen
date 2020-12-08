package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"bufio"
	"os"
)

var parar int64

var Mensajito string

var Puerto string

var Puerto2 string

func cliente(hola string)  {
	
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(hola)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func clienteEscuchar()  {
	
	s, err := net.Listen("tcp", ":9997")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleServidor(c)
		s.Close()
		c.Close()
		return
	}
	
}

func hacerPuerto(hola string){
	s, err := net.Listen("tcp", hola)
	Puerto = hola
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleServidor1(c)
	}
}

func hacerPuerto2(hola string){
	s, err := net.Listen("tcp", hola)
	Puerto2 = hola
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleServidor2(c)
	}
}

func handleServidor1(c net.Conn)  {
	var hola string
	err := gob.NewDecoder(c).Decode(&hola)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(hola)
	}
}

func handleServidor2(c net.Conn)  {
	var hola []string
	err := gob.NewDecoder(c).Decode(&hola)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		var i int

		i = 0

		for (i < len(hola)){
			println(hola[i])
			i = i + 1
		}
	}
}

func handleServidor(c net.Conn)  {
	var hola []string
	err := gob.NewDecoder(c).Decode(&hola)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		go hacerPuerto(hola[0])
		go hacerPuerto2(hola[1])
	}
}

func main()  {
	parar = -1
	var auxliarString string
	var hacerMenu bool
	var op int64
	var mensaje string
	fmt.Print("Nick name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	auxliarString = scanner.Text()
	println("Conectando...")
	go cliente(auxliarString)
	go clienteEscuchar()

	hacerMenu = true

	for (hacerMenu){
		println("Menu: ")
		println("1) Escribir mensaje")
		println("2) Mandar archivo")
		println("3) Mostrar chat")
		println("0) Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var mensajeEntero string
			fmt.Print("Escriba el mensaje: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			mensaje = scanner.Text()

			c, err := net.Dial("tcp", ":9998")
			if err != nil {
				fmt.Println(err)
				return
			}

			mensajeEntero = auxliarString + " envio: " + mensaje
			
			err = gob.NewEncoder(c).Encode(mensajeEntero)
			if err != nil {
				fmt.Println(err)
			}

		case 2:

		case 3:
			c, err := net.Dial("tcp", ":9996")
			if err != nil {
				fmt.Println(err)
				return
			}
			
			err = gob.NewEncoder(c).Encode(Puerto2)
			if err != nil {
				fmt.Println(err)
			}

		case 0:
			var aux []string
			c, err := net.Dial("tcp", ":9995")
			if err != nil {
				fmt.Println(err)
				return
			}

			aux = append(aux,auxliarString)
			aux = append(aux,Puerto)

			err = gob.NewEncoder(c).Encode(aux)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Se desconecto")

			hacerMenu = false

		}
	}
	
	parar = 1
} 