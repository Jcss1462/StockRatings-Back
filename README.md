# **Stock Ratings API 📈**  
API para la gestión de acciones y recomendaciones de inversión, desarrollada con **Go 1.24.0** y **Gin**. Incluye documentación con **Swagger** y utiliza **CockroachDB** como base de datos.  

## 🚀 **Requisitos Previos**  

Antes de ejecutar el proyecto, asegúrate de tener instalados:  
- **Go 1.24.0** → [Descargar Go](https://go.dev/dl/)  
- **Docker** → [Instalar Docker](https://docs.docker.com/get-docker/)  

## 🛠 **Instalación y Configuración**  

### **1️⃣ Clonar el repositorio**  
```sh
git clone https://github.com/Jcss1462/StockRatings-Back.git
cd StockRatings-Back
```

### **2️⃣ Configurar variables de entorno**  
El proyecto ya incluye un archivo **`.env`** con la configuración de la base de datos lista para funcionar, por lo que no necesitas modificar nada para conectarte a **CockroachDB**.  

Si deseas revisarlo o cambiar valores, edita el archivo `.env`:  

```sh
nano .env  # Para editarlo en Linux/macOS
notepad .env  # Para editarlo en Windows
```

📌 **Nota:** No es necesario crear manualmente el archivo `.env`, ya que está incluido en el repositorio.  

### **3️⃣ Levantar la base de datos en Docker**  

Ejecuta el siguiente comando para iniciar **CockroachDB** en un contenedor de Docker:  

```sh
docker run -d --name=cockroachdb --hostname=cockroachdb -p 26257:26257 -p 8080:8080 cockroachdb/cockroach:v24.3.6 start-single-node --insecure
```

### **📌 Verificar si la base de datos está corriendo**  
Para comprobar el estado del contenedor:  
```sh
docker ps  # Muestra los contenedores en ejecución
```
Si `cockroachdb` aparece en la lista, significa que está corriendo correctamente.  

### **📌 Detener la base de datos si es necesario**  
```sh
docker stop cockroachdb  # Detiene el contenedor
```

### **📌 Reiniciar la base de datos**  
Si el contenedor ya está creado pero detenido, puedes iniciarlo con:  
```sh
docker start cockroachdb  # Vuelve a iniciar CockroachDB
```

### **4️⃣ Generar la base de datos y las tablas automáticamente**  

Una vez que **CockroachDB** esté corriendo, el proyecto generará automáticamente la base de datos y las tablas al iniciarse.  

Si necesitas crear la base de datos y tablas manualmente, puedes acceder a la consola interactiva de CockroachDB con:  

```sh
docker exec -it cockroachdb cockroach sql --insecure
```

Luego, dentro de la consola SQL, ejecutar:  

```sql
CREATE DATABASE stockratings;
USE stockratings;
```

📌 **Nota:** No es necesario ejecutar estos comandos si simplemente corres la API, ya que la configuración del proyecto se encarga de crear la base de datos y tablas automáticamente.

## 🚀 **Ejecutar la API**  

Para iniciar el servidor, simplemente ejecuta:  

```sh
go run .
```

Esto iniciará la API en `http://localhost:8081/` y, si estás en **modo desarrollo**, abrirá automáticamente **Swagger** en `http://localhost:8081/swagger/index.html`.  

## 📖 **Documentación de la API**  

La API cuenta con documentación generada con **Swagger**. Para verla, inicia la aplicación y accede a:  

🔗 **Swagger UI** → [http://localhost:8081/swagger/index.html](http://localhost:8081/swagger/index.html)  

## 🔧 **Comandos útiles**  

### **Compilar la aplicación**  
```sh
go build
```

### **Ejecutar la aplicación compilada**  
```sh
.\StockRatings-Back.exe
```
---
