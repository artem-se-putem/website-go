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
    r.Post("/contact", contactHandler(rootPath))
    
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
        log.Printf("=== ПОЛУЧЕН ЗАПРОС ===")
        log.Printf("Метод: %s", r.Method)
        
        // Парсим форму
        err := r.ParseForm()
        if err != nil {
            log.Printf("❌ Ошибка парсинга: %v", err)
            http.Error(w, "Ошибка формы", http.StatusBadRequest)
            return
        }
        
        // Логируем ВСЕ что пришло
        log.Printf("Все данные формы: %v", r.Form)
        log.Printf("Заголовки: %v", r.Header)
        
        // Получаем данные
        name := r.FormValue("name")
        email := r.FormValue("email")
        phone := r.FormValue("phone")
        agreement := r.FormValue("agreement")
        
        log.Printf("name: '%s'", name)
        log.Printf("email: '%s'", email)
        log.Printf("phone: '%s'", phone)
        log.Printf("agreement: '%s'", agreement)
        
        // ОБЯЗАТЕЛЬНО отправляем ответ!
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`
            <html>
                <body>
                    <h1>Форма успешно отправлена!</h1>
                    <p>Мы свяжемся с вами в ближайшее время.</p>
                    <a href="/">Вернуться на сайт</a>
                </body>
            </html>
        `))
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