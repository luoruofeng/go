package boardcaster_mongo

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

func ShowTime() time.Time {
	return time.Now()
}

func users_page(w http.ResponseWriter, r *http.Request) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<h3>online users</h3>
		
		<hr/>
		{{range .OnlineUsers}}
			<div>
				<a href="message?id={{.Id.Hex}}&name={{.Name}}">{{ .Name }}</a>
			</div>
		{{else}}
			<div><strong>no users</strong></div>
		{{end}}
		
		<hr/>

		<h3>all users</h3>
		{{range .AllUsers}}
			<div>
				<a href="message?id={{.Id.Hex}}&name={{.Name}}">{{ .Name }}</a>
			</div>
		{{else}}
			<div><strong>no users</strong></div>
		{{end}}
		<p>
			NOW:{{st.Format "2006-01-02 15:04:05"}}
		</p>
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t := template.New("userspage")
	t = t.Funcs(template.FuncMap{"st": ShowTime})
	t, err := t.Parse(tpl) //Parse must be after Funcs
	check(err)

	onlineAllUser := make([]*User, 0)
	allUser := GetAllUser()
	for user := range users {
		onlineAllUser = append(onlineAllUser, user)
	}

	data := struct {
		Title       string
		OnlineUsers []*User
		AllUsers    []User
	}{
		Title:       "users",
		OnlineUsers: onlineAllUser,
		AllUsers:    allUser,
	}

	err = t.Execute(w, data)
	check(err)
}

func web() {
	http.HandleFunc("/users", users_page)
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("name")
		userid := r.FormValue("id")

		htmlByte, _ := ioutil.ReadFile("./boardcaster_mongo/message.html")
		t, err := template.New("webpage").Parse(string(htmlByte))
		if err != nil {
			log.Fatal(err)
		}
		data := struct {
			UserName string
			Messages []Message
		}{
			UserName: username,
			Messages: GetMessage(userid),
		}

		err = t.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
