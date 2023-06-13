package main

import (
	"context"
	"fmt"
	"main/connection"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id int
	Title string
	Description string
	StartDate string
	StartTime time.Time
	EndTime time.Time
	EndDate string
	Duration string
	NodeJs bool
	React bool
	Bootstrap bool
	Laravel bool
	Image string
}

 
func main() {
	connection.DatabaseConnect()

	e := echo.New()

	e.Static("/public", "public")

	// routing
	//get
	e.GET("/", homePage)
	e.GET("/contact", contactPage)
	e.GET("/add-project", addProjectPage)
	e.GET("/testimonial", testimonialPage)
	e.GET("/project-detail/:id", projectDetailPage)
	e.GET("/update-project/:id", updateProjectPage)
	e.GET("/login", loginPage)
	e.GET("/register", registerPage)
	//post
	e.POST("/add-project", AddProject)
	e.POST("/project-delete/:id", deleteProject)
	e.POST("/update-project/:id", updateProject)
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func homePage(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, description, start_date, end_date, duration, node_js, react, bootstrap, laravel, image FROM tb_project ORDER BY id ASC")

	var result []Project
	for data.Next() {
		var each = Project{}
		
		err := data.Scan(&each.Id, &each.Title, &each.Description, &each.StartDate, &each.EndDate, &each.Duration, &each.NodeJs, &each.React, &each.Bootstrap, &each.Laravel, &each.Image)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		result = append(result, each)
	}

	projects := map[string]interface{} {
		"Projects": result,
	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
}

func contactPage(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addProjectPage(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimonialPage(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)

}

func projectDetailPage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, description, start_date, end_date, duration, node_js, react, bootstrap, laravel, image FROM tb_project WHERE id=$1", id).Scan(&ProjectDetail.Id, &ProjectDetail.Title, &ProjectDetail.Description, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Duration, &ProjectDetail.NodeJs, &ProjectDetail.React, &ProjectDetail.Bootstrap, &ProjectDetail.Laravel, &ProjectDetail.Image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	StartTime, _ := time.Parse("2006-01-02", ProjectDetail.StartDate)
	EndTime, _ := time.Parse("2006-01-02", ProjectDetail.EndDate)
	ProjectDetail.StartDate = StartTime.Format("2 January 2006")
	ProjectDetail.EndDate = EndTime.Format("2 January 2006")
	fmt.Println(ProjectDetail.StartDate, ProjectDetail.EndDate)

	data := map[string]interface{} {
		"Project": ProjectDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/project-detail.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func updateProjectPage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, description, start_date, end_date, duration, node_js, react, bootstrap, laravel, image FROM tb_project WHERE id=$1", id).Scan(&ProjectDetail.Id, &ProjectDetail.Title, &ProjectDetail.Description, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Duration, &ProjectDetail.NodeJs, &ProjectDetail.React, &ProjectDetail.Bootstrap, &ProjectDetail.Laravel, &ProjectDetail.Image)

	data := map[string]interface{} {
		"Project": ProjectDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/update-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}



func calculateDuration(startDate, endDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " week"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " days"
				} else {
					duration = strconv.Itoa(durationDays) + " day"
				}
			}
		}
	}

	return duration
}

func loginPage(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)

}
func registerPage(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func AddProject(c echo.Context) error {
	title := c.FormValue("input-project-name")
	description := c.FormValue("input-deskripsi")
	startDate := c.FormValue("input-date-start")
	endDate := c.FormValue("input-date-end")
	duration := calculateDuration(startDate, endDate)
	nodeJs := (c.FormValue("node-js") == "nodeJs")
	react := (c.FormValue("react") == "react")
	bootstrap := (c.FormValue("bootstrap") == "bootstrap")
	laravel := (c.FormValue("laravel") == "laravel")
	image := c.FormValue("input-image")

	fmt.Println(title, duration, description, nodeJs, react, bootstrap, laravel, image)

	_, err:= connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (title, description, start_date, end_date, duration, node_js, react, bootstrap, laravel, image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", title, description, startDate, endDate, duration, nodeJs, react, bootstrap, laravel, image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.Redirect(http.StatusMovedPermanently, "/#my-project")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Id :", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/#my-project")
}

func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Id :", id)

	title := c.FormValue("input-project-name")
	startDate := c.FormValue("input-date-start")
	endDate := c.FormValue("input-date-end")
	description := c.FormValue("input-deskripsi")
	duration := calculateDuration(startDate, endDate)
	nodeJs := (c.FormValue("node-js") == "nodeJs")
	react := (c.FormValue("react") == "react")
	bootstrap := (c.FormValue("bootstrap") == "bootstrap")
	laravel := (c.FormValue("laravel") == "laravel")

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET title=$1, description=$2, start_date=$3, end_date=$4, duration=$5, node_js=$6, react=$7, bootstrap=$8, laravel=$9 WHERE id=$10", title, description, startDate, endDate, duration, nodeJs, react, bootstrap, laravel, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/#my-project")
}