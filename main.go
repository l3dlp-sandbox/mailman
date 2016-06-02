package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/briandowns/spinner"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/xuqingfeng/mailman/account"
	"github.com/xuqingfeng/mailman/contacts"
	"github.com/xuqingfeng/mailman/lang"
	"github.com/xuqingfeng/mailman/mail"
	"github.com/xuqingfeng/mailman/smtp"
	"github.com/xuqingfeng/mailman/util"
)

const (
	spinnerCharIndex   = 14
	readLogFileGap     = 5 // second
	HI_THERE           = "HI THERE !"
	minTCPPort         = 0
	maxTCPPort         = 65535
	maxReservedTCPPort = 1024
	// 15M
	maxMemory = 1024 * 1024 * 15
)

var (
	msg            util.Msg
	previewContent = ""
	upgrader       = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	DataIsNotJsonErr = errors.New("data is not json format")
)

type Key struct {
	Key string `json:"key"`
}

func main() {

	cyan := color.New(color.FgCyan).SprintFunc()

	colorName := cyan("NAME:")
	colorUsage := cyan("USAGE:")
	colorVersion := cyan("VERSION:")
	colorAuthor := cyan("AUTHOR")
	colorCommands := cyan("COMMANDS")
	colorGlobalOptions := cyan("GLOBAL OPTIONS")

	cli.AppHelpTemplate = colorName + `
    {{.Name}} - {{.Usage}}
` + colorUsage + `
{{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .Flags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
{{if .Version}}
` + colorVersion + `
{{.Version}}
{{end}}{{if len .Authors}}
` + colorAuthor + `
{{range .Authors}}{{ . }}{{end}}
{{end}}{{if .Commands}}
` + colorCommands + `
{{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
{{end}}{{end}}{{if .Flags}}
` + colorGlobalOptions + `
{{range .Flags}}{{.}}
{{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
{{.Copyright}}
{{end}}
`

	app := cli.NewApp()
	app.Name = "mailman"
	app.Usage = "local email client with customizable SMTP server"
	app.Version = "0.2"
	app.Author = "xuqingfeng"
	app.Action = func(c *cli.Context) {

		portInUse := -1
		portStart := 8000
		portEnd := 8100
		for portStart <= portEnd {
			if isTCPPortAvailable(portStart) {
				portInUse = portStart
				break
			}
			portStart++
		}
		if -1 == portInUse {
			log.Fatal("can't find availiable port")
			os.Exit(1)
		}

		log.Info("open 127.0.0.1:" + strconv.Itoa(portInUse) + " in browser")

		s := spinner.New(spinner.CharSets[spinnerCharIndex], 100*time.Millisecond)
		s.Color("cyan")
		s.Start()

		// util init
		util.CreateConfigDir()

		// listen
		router := mux.NewRouter()

		subRouter := router.PathPrefix("/api").Subrouter()
		subRouter.HandleFunc("/", APIHandler)
		subRouter.HandleFunc("/lang", LangHandler)
		subRouter.HandleFunc("/mail", MailHandler)
		subRouter.HandleFunc("/file", FileHandler)
		subRouter.HandleFunc("/account", AccountHandler)
		subRouter.HandleFunc("/contacts", ContactsHandler)
		subRouter.HandleFunc("/smtpServer", SMTPServerHandler)
		subRouter.HandleFunc("/preview", PreviewHandler)
		subRouter.HandleFunc("/wslog", WSLogHandler)

		// done
		router.PathPrefix("/").Handler(http.FileServer(http.Dir("ui")))

		http.ListenAndServe(":"+strconv.Itoa(portInUse), router)
	}
	app.Commands = []cli.Command{
		{
			Name:        "clean",
			Usage:       "clean up tmp dir",
			Description: "mailman clean",
			Action: func(c *cli.Context) {
				log.Info("*** clean START ***")
				homeDir := util.GetHomeDir()
				tmpPath := filepath.Join(homeDir, util.ConfigPath["tmpPath"])
				err := os.RemoveAll(tmpPath)
				if err != nil {
					log.Error(err)
				}
				util.CreateConfigDir()
				log.Info("*** clean STOP ***")
			},
		},
	}

	app.Run(os.Args)
}

func APIHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "default api route")
}

func LangHandler(rw http.ResponseWriter, r *http.Request) {

	if "GET" == r.Method {
		lg, _ := lang.GetLang()

		switch lg {
		case "en", "zh":
			sendSuccess(rw, lg, "get lang success")
		default:
			sendSuccess(rw, "en", "get lang success")
		}

	} else if "POST" == r.Method {

		data := r.PostFormValue("data")
		var lg lang.Lang
		err := json.Unmarshal([]byte(data), &lg)
		if err != nil {

			sendError(rw, DataIsNotJsonErr.Error())
		} else if err = lang.SaveLang(lg); err != nil {
			sendError(rw, "save lang fail "+err.Error())
		} else {
			l, err := lang.GetLang()
			if err != nil {

				sendError(rw, "get lang fail "+err.Error())
			} else {

				sendSuccess(rw, l, "save lang success")
			}
		}
	}
}

func MailHandler(rw http.ResponseWriter, r *http.Request) {

	if "GET" == r.Method {

		sendSuccess(rw, struct{}{}, HI_THERE)
	} else if "POST" == r.Method {

		data := r.PostFormValue("data")
		var m mail.Mail

		err := json.Unmarshal([]byte(data), &m)

		if err != nil {

			sendError(rw, DataIsNotJsonErr.Error())
		} else if err = mail.SendMail(m); err != nil {

			sendError(rw, "send mail fail "+err.Error())
		} else {
			// empty struct
			sendSuccess(rw, struct{}{}, "send mail success")
		}
	}
}

func FileHandler(rw http.ResponseWriter, r *http.Request) {

	if "GET" == r.Method {

		sendSuccess(rw, struct{}{}, HI_THERE)
	} else if "POST" == r.Method {

		if err := r.ParseMultipartForm(maxMemory); err != nil {
			sendError(rw, err.Error())
		}

		token := ""
		for k, vs := range r.MultipartForm.Value {
			for _, v := range vs {
				if "token" == k {
					token += v
				}
			}
		}

		for _, fileHeaders := range r.MultipartForm.File {
			for _, fileHeader := range fileHeaders {
				f, _ := fileHeader.Open()
				fileContent, _ := ioutil.ReadAll(f)
				err := mail.SaveAttachment(fileContent, token, fileHeader.Filename)
				if err != nil {
					sendError(rw, "save attachment fail")
					// todo multi
					break
				}
			}
		}
	}
}

func AccountHandler(rw http.ResponseWriter, r *http.Request) {

	if "GET" == r.Method {

		emails, err := account.GetAccountEmail()
		if err != nil {
			sendError(rw, "get account email fail "+err.Error())
		} else {
			// empty []string
			sendSuccess(rw, emails, "get account email success")
		}
	} else if "POST" == r.Method {

		data := r.PostFormValue("data")
		var at account.Account

		err := json.Unmarshal([]byte(data), &at)
		if err != nil {

			sendError(rw, DataIsNotJsonErr.Error())
		} else if err = account.SaveAccount(at); err != nil {

			sendError(rw, "save account fail "+err.Error())
		} else {
			emails, err := account.GetAccountEmail()
			if err != nil {

				sendError(rw, "get account email fail "+err.Error())
			} else {

				sendSuccess(rw, emails, "save account success")
			}
		}
	} else if "DELETE" == r.Method {

		var k Key
		err := json.NewDecoder(r.Body).Decode(&k)
		if err != nil {

			sendError(rw, DataIsNotJsonErr.Error()+" "+err.Error())
		} else if err = account.DeleteAccount(k.Key); err != nil {

			sendError(rw, "delete account fail "+err.Error())
		} else {
			emails, err := account.GetAccountEmail()
			if err != nil {

				sendError(rw, "get account email fail "+err.Error())
			} else {

				sendSuccess(rw, emails, "delete account success")
			}
		}
	}
}

func ContactsHandler(rw http.ResponseWriter, r *http.Request) {

	if "GET" == r.Method {

		contacts, err := contacts.GetContacts()
		if err != nil {

			sendError(rw, "get contacts fail "+err.Error())
		} else {

			sendSuccess(rw, contacts, "get contacts success")
		}
	} else if "POST" == r.Method {

		data := r.PostFormValue("data")
		var ct contacts.Contacts

		err := json.Unmarshal([]byte(data), &ct)
		if err != nil {

			sendError(rw, DataIsNotJsonErr.Error())
		} else if err = contacts.SaveContacts(ct); err != nil {

			sendError(rw, "save contacts fail "+err.Error())
		} else {
			contacts, err := contacts.GetContacts()
			if err != nil {

				sendError(rw, "get contacts fail "+err.Error())
			} else {

				sendSuccess(rw, contacts, "save contacts success")
			}
		}
	} else if "DELETE" == r.Method {
		var k Key
		err := json.NewDecoder(r.Body).Decode(&k)
		if err != nil {

			sendError(rw, DataIsNotJsonErr.Error()+" "+err.Error())
		} else if err = contacts.DeleteContacts(k.Key); err != nil {

			sendError(rw, "delete contacts fail "+err.Error())
		} else {
			c, err := contacts.GetContacts()
			if err != nil {

				sendError(rw, "get contacts fail "+err.Error())
			} else {

				sendSuccess(rw, c, "delete contacts success")
			}
		}
	}
}

func SMTPServerHandler(rw http.ResponseWriter, r *http.Request) {

	if "GET" == r.Method {

		customSMTPServer, err := smtp.GetCustomSMTPServer()
		if err != nil {

			sendError(rw, "get custom SMTP server fail"+err.Error())
		} else {

			sendSuccess(rw, customSMTPServer, "get custom SMTP Server success")
		}
	} else if "POST" == r.Method {

		data := r.PostFormValue("data")
		var smtpServer smtp.SMTPServer

		err := json.Unmarshal([]byte(data), &smtpServer)
		if err != nil {

			sendError(rw, DataIsNotJsonErr.Error())
		} else if err = smtp.SaveSMTPServer(smtpServer); err != nil {

			sendError(rw, err.Error())
		} else {

			customSMTPServer, err := smtp.GetCustomSMTPServer()
			if err != nil {

				sendError(rw, err.Error())
			} else {

				sendSuccess(rw, customSMTPServer, "save SMTP Server success")
			}
		}
	} else if "DELETE" == r.Method {
		var k Key
		err := json.NewDecoder(r.Body).Decode(&k)
		if err != nil {

			sendError(rw, DataIsNotJsonErr.Error()+" "+err.Error())
		} else if err = smtp.DeleteSMTPServer(k.Key); err != nil {

			sendError(rw, "delete SMTPServer fail "+err.Error())
		} else {
			server, err := smtp.GetCustomSMTPServer()
			if err != nil {

				sendError(rw, "get custom SMTP server fail "+err.Error())
			} else {

				sendSuccess(rw, server, "delete SMTP server success")
			}
		}
	}
}

func PreviewHandler(rw http.ResponseWriter, r *http.Request) {

	if "GET" == r.Method {

		rw.Header().Set("Content-Type", "text/html")
		rw.Write([]byte(previewContent))
	} else if "POST" == r.Method {

		data := r.PostFormValue("data")
		type Body struct {
			Body string `json:"body"`
		}
		var body Body
		err := json.Unmarshal([]byte(data), &body)
		if err != nil {

			sendError(rw, DataIsNotJsonErr.Error())
		} else {

			previewContent = mail.ParseMailContent(body.Body)
			sendSuccess(rw, struct{}{}, previewContent)
		}
	}
}

func WSLogHandler(rw http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Error(err.Error())
	}

	homeDir := util.GetHomeDir()
	logFilePath := filepath.Join(homeDir, util.ConfigPath["logPath"], util.LogName)
	logFile, err := os.Open(logFilePath)
	if err != nil {
		log.Error(err.Error())
	}
	reader := bufio.NewReader(logFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err.Error())
		} else if err == io.EOF {
			// wait
			time.Sleep(readLogFileGap * time.Second)
		} else {
			if err = conn.WriteMessage(1, []byte(line)); err != nil {
				log.Error(err.Error())
			}
		}
	}
}

//***echo JSON START*****
func sendSuccess(rw http.ResponseWriter, data interface{}, message string) {

	msg = util.Msg{
		Success: true,
		Data:    data,
		Message: message,
	}

	msgInByteSlice, _ := json.Marshal(msg)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(msgInByteSlice)
}
func sendError(rw http.ResponseWriter, message string) {

	msg = util.Msg{
		Success: false,
		Message: message,
	}

	msgInByteSlice, _ := json.Marshal(msg)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(msgInByteSlice)
}

//***echo JSON END*****

//***check TCP port is available START***
func isTCPPortAvailable(port int) bool {

	if port < minTCPPort || port > maxTCPPort {
		return false
	}
	conn, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

//***check TCP port is available END***
