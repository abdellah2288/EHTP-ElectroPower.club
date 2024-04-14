package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type (
	Server struct{}
	Member struct {
		FullName         string
		LinkedInProfile  string
		ProfilePhotoPath string
		Role             string
	}
)
type News struct {
	Index        int
	ID           string
	Caption      string
	Link         string
	PreviewImage string
}
type TeamData struct {
	teamName       string
	teamSchoolName string
	teamCount      string
	teamLeadName   string
	teamLeadPhone  string
	teamLeadMail   string
}
type Sponsor struct {
	Link string
	Logo string
}
type ServerData struct {
	BoardMembers []Member
	LatestNews   []News
	Sponsors     []Sponsor
}

var mainData ServerData

func (server Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryDB()
	if r.Method == "GET" {
		switch r.URL.Path {
		case "/":
			t, _ := template.ParseFiles("static/templates/index.html")
			t.Execute(w, mainData)
		case "/register":
			t, _ := template.ParseFiles("static/templates/signup.html")
			t.Execute(w, nil)
		case "/projects":
			fmt.Fprint(w, "Page under construction")
		case "/circuitjam":
			t, _ := template.ParseFiles("static/templates/countdown.html")
			t.Execute(w, nil)
		case "/osaker":
			http.Redirect(w, r, "https://www.youtube.com/watch?v=OaMPGbWFaX8", http.StatusSeeOther)
		case "/palestine":
			t, _ := template.ParseFiles("static/templates/palestine.html")
			t.Execute(w, nil)
		default:
			fmt.Fprint(w, "Page not found")
		}
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		var team TeamData
		team.teamName = r.FormValue("teamName")
		team.teamSchoolName = r.FormValue("teamSchoolName")
		team.teamCount = r.FormValue("teamCount")
		team.teamLeadName = r.FormValue("teamLeadName")
		team.teamLeadPhone = r.FormValue("teamLeadPhone")
		team.teamLeadMail = r.FormValue("teamLeadMail")
		registerTeamInDB(team)
		template, err := template.ParseFiles("static/templates/signup.html")
		checkError(err)
		template.Execute(w, "Signup Successful")
	}
}

func queryDB() {
	/*Database  Queries*/
	count := 0
	database, err := sql.Open("sqlite3", "./mainDB")
	checkError(err)
	countQuery, err := database.Query("SELECT COUNT(*) FROM boardMembers")
	checkError(err)
	rows, err := database.Query("SELECT fullName,linkedInProfile,imagePath,role FROM boardMembers")
	checkError(err)

	if countQuery.Next() {
		countQuery.Scan(&count)
	}
	mainData.BoardMembers = make([]Member, count, count)
	index := 0
	for rows.Next() {
		currentMember := &(mainData.BoardMembers[index])
		rows.Scan(&(currentMember.FullName), &(currentMember.LinkedInProfile), &(currentMember.ProfilePhotoPath), &(currentMember.Role))
		index++
	}
	index = 0

	countQuery, err = database.Query("SELECT COUNT(*) FROM news")
	checkError(err)
	rows, err = database.Query("SELECT ID,Link,PreviewImage,Caption FROM news")
	checkError(err)

	if countQuery.Next() {
		countQuery.Scan(&count)
	}
	mainData.LatestNews = make([]News, count, count)
	for rows.Next() {
		currentNews := &(mainData.LatestNews[index])
		rows.Scan(&(currentNews.ID), &(currentNews.Link), &(currentNews.PreviewImage), &(currentNews.Caption))
		currentNews.Index = index
		index++
	}

	index = 0

	countQuery, err = database.Query("SELECT COUNT(*) FROM sponsors")
	checkError(err)
	rows, err = database.Query("SELECT link,logo FROM sponsors")
	checkError(err)

	if countQuery.Next() {
		countQuery.Scan(&count)
	}
	mainData.Sponsors = make([]Sponsor, count, count)
	for rows.Next() {
		currentSponsor := &(mainData.Sponsors[index])
		rows.Scan(&(currentSponsor.Link), &(currentSponsor.Logo))
		index++
	}

	database.Close()
}

func registerTeamInDB(teamData TeamData) {
	database, err := sql.Open("sqlite3", "./mainDB")
	defer database.Close()
	checkError(err)
	statement, err := database.Prepare(" INSERT INTO circuitJamTeams(teamName,teamCount,teamSchool,teamLeadName,teamLeadPhone,teamLeadMail) VALUES (?,?,?,?,?,?)")
	checkError(err)
	_, err = statement.Exec(teamData.teamName, teamData.teamCount, teamData.teamSchoolName, teamData.teamLeadName, teamData.teamLeadPhone, teamData.teamLeadMail)
	checkError(err)
}

func main() {
	var server Server
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/static/", func(wr http.ResponseWriter, req *http.Request) {
		// Determine mime type based on the URL
		if strings.HasSuffix(req.URL.Path, ".css") {
			wr.Header().Set("Content-Type", "text/css")
		}
		http.StripPrefix("/static/", fs).ServeHTTP(wr, req)
	})
	http.Handle("/", server)
	http.Handle("/circuitjam", server)
	http.Handle("/register", server)
	http.HandleFunc("/registerTeam", handleSignup)
	http.ListenAndServe(":8080", nil)
}
