# FileServerComAutenticacaoCriptografada

Servidor de arquivos com acesso limitado a usuários previamente cadastrados e com senha criptografada em MD5.


1) O usuário que desejar modificar para acesso, se faz nessesário alterar onde está identificado como "USERNAME"
2) Para cria a senha criptografada, basta acessar o site https://unix4lyfe.org/crypt/ escrever uma senha nos campos em vermelho e copiar o código o codigo gerado no campo MD5 que está em verde na imagem abaixo, agora você pode colar esse código no campo return onde está identificado como pastSecret.

![image](https://user-images.githubusercontent.com/129122229/228089794-53f98c49-551a-44e7-ad1b-aad6c7dbfcd4.png)


3) Crie uma pasta/diretório no disco C: (ou em outro local de sua preferência em caso de linux), como exemplo pode ser uma pasta com o nome "compartilhados"
e escolha uma porta de sua preferencia que não impacte em outras aplicações que possam estar sendo executadas, no nosso exemplo coloquei a porta 1200, então
para executar o program, abra seu Visual Studio Code e na CLI digite:

 ~> go run main.go   c:\compartilhados  1200
 
 
 4) Vá no seu browser e digite localhost:1200, um pop-up deverá abrir solicitando a senha que você criou no site  https://unix4lyfe.org/crypt/
        
