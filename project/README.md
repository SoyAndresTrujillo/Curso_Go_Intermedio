# Definiendo Workers, Jobs y Dispatchers

Golang permite utilizar conceptos como los de worker pools y, en combinación con buffered channels, la creación de job
queues que utilizando concurrencia nos permitirán tener un alto rendimiento a la hora de la creación de muchas tareas.

# Creando web server para procesar jobs
Con la librería estándar de Go y utilizando el paquete net, somos capaces de crear un servidor que será el que atienda
las peticiones y asignará los nuevos workers para que lleven a cabo los trabajos que se está buscando conseguir.

# Test API
Use postman to test the API: POST http://localhost:8081/fib
body {
    "name" : "name",
    "value": value,
    "delay": 2s
}