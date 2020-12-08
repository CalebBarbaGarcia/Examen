package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"strconv"
)

var parar int64

var puerto int

var holas []string

var Mensajes []string

func MandarMensaje(m string, p string){
	c, err := net.Dial("tcp", m)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	err = gob.NewEncoder(c).Encode(p)
	if err != nil {
		fmt.Println(err)
	}
}

func servidor3()  {
	s, err := net.Listen("tcp", ":9995")
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
		handleServidorTercero(c)

	}
}

func servidor()  {
	s, err := net.Listen("tcp", ":9999")
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
		handleServidorPrimero(c)

		c, err = net.Dial("tcp", ":9997")
		if err != nil {
			fmt.Println(err)
			return
		}
		var holaTamano2 []string
		var hola string
		var hola2 string
		hola = ":" + strconv.Itoa(puerto)
		hola2 = ":" + strconv.Itoa(puerto-1)
		holaTamano2 = append(holaTamano2,hola)
		holaTamano2 = append(holaTamano2,hola2)
		puerto = puerto - 2
		holas = append(holas,hola)
		err = gob.NewEncoder(c).Encode(holaTamano2)
		if err != nil {
			fmt.Println(err)
		}
		
		c.Close()
	}
}

func handleServidorPrimero(c net.Conn)  {
	var hola string
	err := gob.NewDecoder(c).Decode(&hola)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		println("Se conecto: ",hola)
	}
}

func servidor1()  {
	s, err := net.Listen("tcp", ":9998")
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

		handleClient(c)
	}
}

func servidor2() {
	s, err := net.Listen("tcp", ":9996")
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
		handleServidorSegunda(c)
		
	}
}

func handleServidorSegunda(c net.Conn)  {
	var hola string
	err := gob.NewDecoder(c).Decode(&hola)

	c, err = net.Dial("tcp", hola)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(Mensajes)
	if err != nil {
		fmt.Println(err)
	}

}

func handleServidorTercero(c net.Conn)  {
	var hola []string
	var AuxArreglo[] string
	gob.NewDecoder(c).Decode(&hola)

	println("Se desconecto ",hola[0])

	var i int

	i = 0

	for (i < len(holas)){
		if (hola[1] == holas[i]){

		} else {
			AuxArreglo = append(AuxArreglo,holas[i])
		}
		i = i +1
	}

	holas = AuxArreglo
}

func handleClient(c net.Conn)  {
	var proceso string
	err := gob.NewDecoder(c).Decode(&proceso)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		Mensajes = append(Mensajes,proceso)

		var f int
		f = 0
		fmt.Println(proceso)
		for (f < len(holas)){
			MandarMensaje(holas[f],proceso)
			f = f + 1
		}
	}
}


func main()  {	
	puerto = 9993

	go servidor()
	go servidor1()
	go servidor2()
	go servidor3()

	fmt.Scanln(&parar)
} 