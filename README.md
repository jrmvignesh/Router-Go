# cmpe-273-fall15-lab2
Example of REST services using JSON request and response in GO using httprouter multiplexer.


Execution steps:
POST request using curl command
 curl  -H "Content-Type: application/json"  -d '{"Name" : "World"}' http://127.0.0.1:8080/user
 
 Response :
{"greetings":"Hello,World!"}
