package main

import (
	"salu2/backend/src/server"
)

type Data struct {
	Message string
	Nunce   string
}

func main() {
	// Manejador para servir archivos estáticos desde el directorio actual.
	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./dist"))))
	//http.HandleFunc("/index.html", MainHandler)
	//
	//// Manejador para la ruta específica "/hello" que devuelve el archivo HTML.
	//http.HandleFunc("/hello", helloHandler)
	//
	//// Puerto en el que escuchará el servidor.
	//port := 8989
	//
	//fmt.Printf("Servidor escuchando en el puerto %d...\n", port)
	//
	//// Iniciar el servidor y manejar posibles errores.
	//err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	//if err != nil {
	//	fmt.Println("Error al iniciar el servidor:", err)
	//	os.Exit(1)
	//}

	serverConfig := server.NewConfig("8080")
	server.New(serverConfig)
}

//
//func MainHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodGet {
//		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
//		return
//	}
//
//	data := Data{
//		Nunce: "Hola Mundo!",
//	}
//	// Abrir el archivo HTML.
//	tmpl, err := template.ParseFiles("./dist/index.html")
//	if err != nil {
//		http.Error(w, "Error al abrir el archivo HTML", http.StatusInternalServerError)
//		return
//	}
//
//	// Configurar la cabecera de la respuesta para indicar que es un archivo HTML.
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	cookie := http.Cookie{
//		Name:       "Juanlu",
//		Value:      "asdf",
//		Path:       "/hello",
//		Domain:     "localhost",
//		Expires:    time.Time{},
//		RawExpires: "",
//		MaxAge:     0,
//		Secure:     true,
//		HttpOnly:   false,
//		SameSite:   0,
//		Raw:        "",
//		Unparsed:   nil,
//	}
//
//	http.SetCookie(w, &cookie)
//	// Copiar el contenido del archivo al cuerpo de la respuesta.
//	err = tmpl.Execute(w, data)
//	if err != nil {
//		http.Error(w, "Error al enviar el contenido del archivo HTML", http.StatusInternalServerError)
//		return
//	}
//}
//
//func helloHandler(w http.ResponseWriter, r *http.Request) {
//	// Comprobar si la solicitud es de tipo GET.
//	if r.Method != http.MethodGet {
//		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
//		return
//	}
//
//	data := Data{
//		Message: "Hola desde el servidor",
//	}
//
//	// Abrir el archivo HTML.
//	tmpl, err := template.ParseFiles("./dist/index.html")
//	if err != nil {
//		http.Error(w, "Error al abrir el archivo HTML", http.StatusInternalServerError)
//		return
//	}
//
//	fmt.Println(r.Cookie("Juanlu"))
//
//	// Configurar la cabecera de la respuesta para indicar que es un archivo HTML.
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//
//	// Copiar el contenido del archivo al cuerpo de la respuesta.
//	err = tmpl.Execute(w, data)
//	if err != nil {
//		http.Error(w, "Error al enviar el contenido del archivo HTML", http.StatusInternalServerError)
//		return
//	}
//}
