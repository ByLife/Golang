package module

import "html/template"

type Templ struct { //Struct sent to api
	Post     []Post        //displays posts
	User     User          // user info
	JsonUser []jsonUser    // store github api info (about us)
	Comments []Comment     //displays comments under post
	Report   []Report      //stores reports
	Mods     []Mod         // store moderators list
	Inter    []Interaction // display interactions between users (like...)
	Notif    Notification  // display notifications
}

type Post struct {
	Id       int
	Creator  string
	Title    string
	Content  string
	Picture  template.HTML
	Likes    int
	Dislikes int
	Slash    string
	Date     string
}

type Notification struct {
	Post  Post
	Com   []Comment
	Inter []Interaction
}
type Interaction struct {
	Postid string
	Uuid   string
	Karma  int
	Date   string
}

type User struct {
	Username     string
	Creationdate string
	Level        int
	Karma        int
}

type Mod struct {
	Uuid         string
	Email        string
	Username     string
	Creationdate string
	Level        int
	Karma        int
}

type Report struct {
	ReportContent string
	ReportId      string
	User          string
	Postid        string
}

type Comment struct {
	Creator          string
	Postid           string
	Comment          string
	Likes            int
	Dislikes         int
	File             string
	Creationdateuser string
	Username         string
	Level            int
	Date             string
}

type jsonUser struct {
	Name     string `json:"login"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar_url"`
	Location string `json:"location"`
}
