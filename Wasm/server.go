package main

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "io/ioutil"
)
var directorio string

func main() {
    
    if len(os.Args) < 2 {
        fmt.Println("Uso: go run server.go [directorio]")
        os.Exit(1)
    }
    // Servir archivos estáticos desde el directorio actual
    directorio := os.Args[1]
    fs := http.FileServer(http.Dir(directorio))
    http.Handle("/", fs)
    http.HandleFunc("/leer-html", leerContenidoHTMLHandler)

    fmt.Println("Servidor iniciado en http://localhost:8080") 
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error al iniciar el servidor: %s\n", err)
        os.Exit(1)
    }
}

func leerContenidoHTMLHandler(w http.ResponseWriter, r *http.Request) {
    // Obtener el nombre del archivo desde la query string
    nombreArchivo := r.URL.Query().Get("nombreArchivo")
    if nombreArchivo == "" {
        http.Error(w, "Falta el parámetro nombreArchivo", http.StatusBadRequest)
        return
    }
    directorio := os.Args[1]
    rutaArchivo := filepath.Join(directorio, nombreArchivo)

    // Verificar si el archivo existe
    if _, err := os.Stat(rutaArchivo); os.IsNotExist(err) {
        http.Error(w, "Archivo no encontrado", http.StatusNotFound)
        return
    }

    // Leer el contenido del archivo
    contenido, err := ioutil.ReadFile(rutaArchivo)
    if err != nil {
        http.Error(w, "Error al leer el archivo", http.StatusInternalServerError)
        return
    }

    // Escribir el contenido al response
    w.Write(contenido)
}