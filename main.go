package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	 auth "github.com/abbot/go-http-auth"   // Importando o pacote de autenticação do http, o mesmo foi baixado
	                                       // no CLI assim: go get github.com/abbot/go-http-auth
										  // este pacote tem suporte para criptorgrafia de senhas 
)

func Secret (user, realm string) string {
	    
	    // Nome do usuário
		if user == "userName" { 
			
		// Senha criptografada para o programa
		return "pastSecret"

		/* A senha foi criada no site https://unix4lyfe.org/crypt/
		usando a opção de MD5 */

		}

		return ""
}


func main() {

	// Aqui será considerado que 3 argumentos devem ser passados e não pode ser diferente disso
	if len(os.Args) != 3 { 
		                  
		// Caso seja diferente de 3, sendo menor ou maior, então o bloco abaixo será executado 
		fmt.Println("Como executar: go run main.go  <caminho-do-diretorio>  <porta>")
	

		// Será informado erro e será encerrado o programa
		os.Exit(1) 

	}

	// Argumento da posição 1 será o caminho do diretório dos arquivos
	CaminhoDir := os.Args[1] 
	
	// Argumento da posição 2 será a porta onde o serviço será executado.
	portServer := os.Args[2] 
	
	/* A variavel authenticator vai receber os dados inseridos na função Secret, como 
	nome do usuário e senha, porém, a senha já estará criptografada no argumento da função.
	O campo "myserver" é o realm, ou seja, apersistencia dos dados, é onde as credenciais estarão*/
	authenticator := auth.NewBasicAuthenticator("myserver.com", Secret) 


	http.HandleFunc("/", //Essa função associal um url, ou seja, ele autentica aquela url/caminho, neste exemplo a raiz do sistema.
	authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest){ // Chamada da autenticação com Request e Response
	http.FileServer(http.Dir(CaminhoDir)).ServeHTTP(w, &r.Request) }))

	fmt.Printf( "Subindo servidor na porta %s ...", portServer)
	log.Fatal(http.ListenAndServe( // Funciona como o start do serviço, pode receber também logs de erros em caso de problemas
	":"+portServer, nil))  // Aqui é recebido o valor da porta que será execuatda no servidor, essa porta será passado como parametro na variavel portServer


	/*

	A forma padrão de executar essa aplicação no cli é assim:

	         go run main.go   c:\compartilhados    1200

                       ^              ^              ^
				  Argumento 1	Argumento 2	   Argumento 3	
				  
				Esses são os 3 argumentos considerados no bloco IF para saber se o programa está sendo
			  executado da maneira correta. 

	*/
}

