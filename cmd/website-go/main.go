package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
    "path/filepath"
    
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    // Получаем корневую директорию проекта
    rootPath, err := getRootPath()
    if err != nil {
        log.Fatal("Ошибка получения корневой директории:", err)
    }
    
    r := chi.NewRouter()
    
    // Middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    
    // Статические файлы
    staticPath := filepath.Join(rootPath, "static")
    r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))
    
    // Маршруты
    r.Get("/", homeHandler(rootPath))
    r.Get("/about", aboutHandler(rootPath))
    r.Get("/contact", contactHandler(rootPath))
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    log.Printf("Сервер запущен на http://localhost:%s", port)
    log.Printf("Корневая папка: %s", rootPath)
    log.Fatal(http.ListenAndServe(":"+port, r))
}

func getRootPath() (string, error) {
    // Если запускаем из cmd/website-go, поднимаемся на два уровня вверх
    return filepath.Abs("../..")
}

func homeHandler(rootPath string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, rootPath, "index.html", nil)
    }
}

func aboutHandler(rootPath string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data := struct {
            Title string
        }{
            Title: "О нас",
        }
        renderTemplate(w, rootPath, "index.html", data)
    }
}

func contactHandler(rootPath string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data := struct {
            Title string
        }{
            Title: "Контакты",
        }
        renderTemplate(w, rootPath, "index.html", data)
    }
}

func renderTemplate(w http.ResponseWriter, rootPath, tmpl string, data any) {
    templatePath := filepath.Join(rootPath, "templates", tmpl)
    t, err := template.ParseFiles(templatePath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    err = t.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}