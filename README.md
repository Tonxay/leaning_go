


how to craete initial project 
# step  start app
1.  go mod init Mymodule                             // it use create  go mod
2.  go get github.com/gofiber/fiber/v2               // this is scrip in install package fiber 


how to access to Db postgres 
# step
1. docker ps                                     // have to seen id imageid active 
2. docker exec -it  imageid  bash
3. psql -U postgres
