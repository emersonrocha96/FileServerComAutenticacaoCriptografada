# FileServerComAutenticacaoCriptografada

Servidor de arquivos com acesso limitado a usuários previamente cadastrados e com senha criptografada em MD5.


1) O usuário que desejar modificar para acesso, se faz nessesário alterar onde está identificado como "USERNAME"
2) Para cria a senha criptografada, basta acessar o site https://unix4lyfe.org/crypt/ escrever uma senha nos campos em vermelho e copiar o código o codigo gerado no campo MD5 que está em verde na imagem abaixo, agora você pode colar esse código no campo return onde está identificado como pastSecret.

![image](https://user-images.githubusercontent.com/129122229/228089794-53f98c49-551a-44e7-ad1b-aad6c7dbfcd4.png)


3) Crie uma pasta/diretório no disco C: (ou em outro local de sua preferência em caso de linux), como exemplo pode ser uma pasta com o nome "compartilhados" enão crie aquivos aleatórios como teste.txt, apenas para serem visualizados, então escolha uma porta de sua preferencia que não impacte em outras aplicações que possam estar sendo executadas, no nosso exemplo coloquei a porta 1200, então
para executar o program, abra seu Visual Studio Code e na CLI digite:

 ~> go run main.go   c:\compartilhados  1200
 
 
4) Vá no seu browser e digite localhost:1200, um pop-up deverá abrir solicitando o usuário que você colocou no passo 1) e a senha que você criou no site  https://unix4lyfe.org/crypt/ e colou o código gerado no campo pastSecret, passo 2).

![image](https://user-images.githubusercontent.com/129122229/228090408-6755b1b3-56c9-4e81-92e0-ed1594574a0e.png)

        
5) Assim fica a página após a liberação de acesso  aos arquivos:

![image](https://user-images.githubusercontent.com/129122229/228090999-e3a76e6e-a7e2-45ba-a0a2-a8841ca67910.png)
