package utility

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		SMTP   SMTPConf     `yaml:"smtp"`
		Server ServerConfig `yaml:"server"`
	}

	ServerConfig struct {
		JwtKey string `yaml:"jwt_key"`
	}

	SMTPConf struct {
		Host   string `yaml:"host"`
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
		Port   int    `yaml:"port"`
		Sender string `yaml:"sender"`
	}
)

func InitConfig() Config {
	var cfg Config
	env := os.Getenv("APP_ENVIRONMENT")

	rootFile := fmt.Sprintf("./conf/config-%s.yml", env)
	f, err := os.Open(rootFile)
	if err != nil {
		log.Printf("Error: %s", err)
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Printf("Error: %s", err)
	}

	return cfg
}

func ParseHTML(path string, data interface{}) (body string, err error) {
	t := template.New(filepath.Base(path)).Funcs(template.FuncMap{})
	t, err = t.ParseFiles(path)

	if err != nil {
		fmt.Println("Error loading template", err.Error())
		return "", err
	} else {
		var tpl bytes.Buffer

		if err = t.Execute(&tpl, data); err == nil {
			body = tpl.String()
		}
	}

	return
}

func ParseFrontend(r *chi.Mux) {
	frontendsrc := os.Getenv("FRONTENDSRC")

	src := "./templates"
	file := fmt.Sprintf("%s/index.html", src)
	_, err := os.Stat(file)

	if !os.IsNotExist(err) {
		bodyHtml, _ := ParseHTML(file, nil)
		bodyHtml = strings.ReplaceAll(bodyHtml, "<%FRONTENDSRC%>", frontendsrc)

		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(bodyHtml))
		})
	}
}

func FileServer(r chi.Router, src string) {
	fs := http.FileServer(http.Dir(src))
	r.Handle("/", http.StripPrefix("/", fs))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(src + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
			return
		}

		fs.ServeHTTP(w, r)
	})
}

func Mailer(subject, body, email string) {
	cfg := InitConfig()

	host := cfg.SMTP.Host
	port := cfg.SMTP.Port
	username := cfg.SMTP.User
	password := cfg.SMTP.Pass
	sender := cfg.SMTP.Sender

	m := gomail.NewMessage(gomail.SetCharset("UTF-8"))
	m.SetHeader("From", sender)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	m.AddAlternative("text/html", body)
	// m.Attach("")

	d := gomail.NewDialer(host, port, username, password)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error: %s", err)
	}
}
