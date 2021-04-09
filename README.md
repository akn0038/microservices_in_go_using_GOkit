# microservices_in_go_using_GOkit

```bash
go run main.go
```
now run the command for uppercase services
```bash
arvind@arvind-HP-Pavilion-Notebook:~$ curl -XPOST -d'{"s":" aba "}' localhost:8080/uppercase
{"v":" ABA "}
arvind@arvind-HP-Pavilion-Notebook:~$ curl -XPOST -d'{"s":" aba "}' localhost:8080/count
{"v":5}
arvind@arvind-HP-Pavilion-Notebook:~$ curl -XPOST -d'{"s":" aba "}' localhost:8080/reverse
{"v":" aba "}
arvind@arvind-HP-Pavilion-Notebook:~$ curl -XPOST -d'{"s":" aba "}' localhost:8080/ispelindrome
{"Result":"True"}
arvind@arvind-HP-Pavilion-Notebook:~$ 
```
