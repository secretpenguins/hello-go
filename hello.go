package main

import (
	"log"
	"github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/martini-contrib/sessions"
	"net/http"
	"login"
	"github.com/coopernurse/gorp"
	"data"
	"config"
)

type Test struct {
	Count int;
	MemcachePath string;
}

func checkSession() martini.Handler {
	return func(res http.ResponseWriter, r *http.Request, c martini.Context, rend render.Render, session sessions.Session) {
		v := session.Get("user_id")
		if v == nil {
			rend.Redirect("/")
		}
	}
}

var dbMap *gorp.DbMap

func init() {
	//data.Setup()
}
func main() {
	//config.Setup()
	dbMap = data.GetDbMap()
	defer dbMap.Db.Close()

	log.Println(dbMap)

	m := martini.Classic()

    store := sessions.NewCookieStore([]byte("secret123"))
    m.Use(sessions.Sessions("my_session", store))

    m.Use(render.Renderer(render.Options{
    	Directory: "templates",
    	Layout: "layout",
    }))

    mywhiteList := []string{"/login", "/test"}
    m.Use(login.Setup(mywhiteList))

    m.Get("/login", func(r render.Render, session sessions.Session) {
    	var login = &data.Login{}
    	options := &render.HTMLOptions {
	    	Layout: "simple",
	    }
    	r.HTML(200, "login", login, *options)
	})

	m.Post("/login", binding.Bind(data.Login{}), func(l data.Login, r render.Render, s sessions.Session) {
		var dbLogin data.Login = data.Login{ LoginId: -1 }

	    dbMap.SelectOne(&dbLogin, "SELECT * FROM Logins WHERE UserName = ?", l.UserName)
	    
	    if (dbLogin.LoginId != -1) {
	    	var passwordMatch = login.ComparePassword(l.Password, dbLogin.Password)
	    	if (passwordMatch) {
	    		s.Set("user_id", string(dbLogin.LoginId))
	    		r.Redirect("/")
	    	}
	    }
	    
   		r.HTML(200, "login", l)
	})

	m.Get("/test", func(r render.Render) {
		posts := data.GetPosts()
		postCount := len(posts)
		log.Println("Count: ", postCount)
		test := &Test{ 
			Count: postCount,
			MemcachePath: config.GetConfig().MemcachePath,
		}
		r.HTML(200, "test", test)
	})

	m.Get("/new", func(r render.Render, session sessions.Session) {
		var post = &data.Post{}
		r.HTML(200, "new", post)
	})

	m.Get("/:id", func(args martini.Params, r render.Render) {
		log.Println("Getting a post")
		post := data.GetPost(args["id"])

		r.HTML(200, "view", post)

	})

	m.Get("/:id/edit", func(args martini.Params, r render.Render) {
		var post = data.GetPost(args["id"])
		r.HTML(200, "edit", post)
	})

	m.Get("/:id/delete", func(args martini.Params, r render.Render) {
		_, err := dbMap.Exec("delete from posts where postId = ? ", args["id"])
		checkErr(err, "Delete failed")
		r.Redirect("/")

	});

	m.Get("/", func (r render.Render) {
		var posts = data.GetPosts()
		r.HTML(200, "index", posts)
	})

	m.Post("/", binding.Bind(data.Post{}), func(post data.Post, r render.Render) {
		log.Println(post)
		/*np := &Post{
			Title: post.Title, 
			Content: post.Content,
		}*/


		var err = dbMap.Insert(&post)
		checkErr(err, "Insert failed")

		r.Redirect("/")
	})

	m.Post("/save", binding.Bind(data.Post{}), func(post data.Post, r render.Render) {
		post.Save()
		r.Redirect("/")
	})

	m.Run()
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Println(msg, err)
    }
}
