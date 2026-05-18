package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"projetosaude/app/handlers"
	"projetosaude/app/utils"
)

func main() {
	// Conecta ao banco de dados utilizando a função definida em utils
	utils.ConnectToDB()

	// Cria um file server para servir arquivos estáticos da pasta "./static"
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)

	// Define rotas da API
	http.HandleFunc("/api/pacientes", handlers.ListPacientesHandler) // GET
	http.HandleFunc("/api/paciente/create", handlers.CreatePacienteHandler) // POST
	http.HandleFunc("/api/paciente/update", handlers.UpdatePacienteHandler) // PUT
	http.HandleFunc("/api/paciente/delete", handlers.DeletePacienteHandler) // DELETE

	// Obtém os endereços de rede disponíveis na máquina para exibir onde o servidor está rodando
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	var localIP string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			localIP = ipNet.IP.String()
			break
		}
	}

	port := "3000"
	if localIP == "" {
		localIP = "127.0.0.1"
	}

	fmt.Printf("Servidor de Saúde rodando em: http://%s:%s/\n", localIP, port)

	// Inicia o servidor HTTP
	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		log.Fatal(err)
	}
}
