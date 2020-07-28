MySQL Dump with Obfuscate

1. https://golang.org/doc/install#install
2. Fill your obfuscate rules in [obfuscate.go](https://github.com/co-in/mysql-dump-obfuscator/blob/master/obfuscate.go) 
3. Run program with params
```
go run ./main.go ./obfuscator.go ./lib.go -u root -d my_database -h 127.0.0.1 -p 3306
```