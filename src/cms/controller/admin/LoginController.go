package admin

import (
	"net/http"
	"fmt"
	//"lib/mustache"
	//"cms/initialize"
	"encoding/json"
	"cms/model"
	"lib/jwt-go"
	"time"
	"net/url"
)

//登录验证
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//登录页面显示
func  (own *Login)GET(w http.ResponseWriter, r *http.Request)  {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	//预留参数a  以便选择哪个方法
	if len(queryForm["a"])>0{
		switch  queryForm["a"][0] {
		//主页
		case "main":
			own.Main(w,r)
		}
	}else{
		//登录页
		//testing
		//session, _ := Session.Get(r,"test")
		// Set some session values.
		//session.Values["foo"] = "bar"
		//session.Values[42] = 43
		// Save it before we write to the response/return from the handler.
		//session.Save(r, w)
		//sql := "select * from m_admin_user"
		//results, err := initialize.Orm.Query(sql)
		//fmt.Println(results,err)
		//fmt.Println(session.Options)
		/*
		t,error:=template.ParseFiles("./src/cms/views/admin/login.html")//New("login.html")
		if error!=nil{
			fmt.Fprintf(w,error.Error())
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		error = t.Execute(w, TemplateData)
		if error != nil{
			fmt.Fprintf(w,error.Error())
			return
		}*/
		TemplateData["title"] ="OK"
		Template.Html(w,r,"login",TemplateData)
		//s := mustache.RenderFileInLayout("./src/cms/views/admin/login.html", "./src/cms/views/admin/layout.html.mustache", nil)
		//fmt.Fprintf(w,s)
	}


}

//登录操作 | jwt
func  (own *Login)POST(w http.ResponseWriter, r *http.Request)  {
	var user Login
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request",err)
		return
	}

	//查询数据库
	fmt.Println(r.Body)
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	model_AdminUser := new(model.AdminUser)
	h,m,err := model_AdminUser.GetByUsernameAndPassword(user.Username,user.Password)
	if h==false{
		fmt.Fprintf(w,"Error in username or password")
	}else{
		//jwt token 加密操作
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
		claims["iat"] = time.Now().Unix()
		claims["admin_login_info"] = m
		token.Claims = claims
		tokenString, err := token.SignedString([]byte(JwtSecretKey))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while signing the token")
		}
		response := JwtToken{tokenString}
		JsonResponse(response, w)
	}
}

func (own *Login)Main(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("ok")
	fmt.Fprintf(w,"hello")
}





