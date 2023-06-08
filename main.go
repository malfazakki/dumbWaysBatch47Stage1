package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id int
	Title string
	Content string
	StartDate string
	EndDate string
	Duration string
	NodeJs bool
	React bool
	Bootstrap bool
	Laravel bool
}

var dataProject = []Project {
	{
		Title: "Dumbways Web Apps 2023",
		Content: "Content",
		StartDate: "2023-05-08",
		EndDate: "2023-06-08",
		Duration: "1 month",
		NodeJs: true,
		React: true,
		Bootstrap: true,
		Laravel: true,
	},
	{
		Title: "Dumbways Web Apps 2023",
		Content: "Content 2",
		StartDate: "2023-05-08",
		EndDate: "2023-06-08",
		Duration: "1 month",
		NodeJs: true,
		React: true,
		Bootstrap: true,
		Laravel: true,
	},
	{
		Title: "Dumbways Web Apps 2023",
		Content: "Content 3",
		StartDate: "2023-05-08",
		EndDate: "2023-06-08",
		Duration: "1 month",
		NodeJs: true,
		React: true,
		Bootstrap: false,
		Laravel: false,
	},
}
 
func main() {
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
	//post
	e.POST("/add-project", AddProject)
	e.POST("/project-delete/:id", deleteProject)
	e.POST("/update-project/:id", updateProject)
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func homePage(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{} {
		"Projects": dataProject,
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

	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				Title: data.Title,
				Content: data.Content,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Duration: data.Duration,
				NodeJs: data.NodeJs,
				React: data.React,
				Bootstrap: data.Bootstrap,
				Laravel: data.Laravel,
			}
		}
	}

	data := map[string]interface{} {
		"Project": ProjectDetail,
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func updateProjectPage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				Id: id,
				Title: data.Title,
				Content: data.Content,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Duration: data.Duration,
				NodeJs: data.NodeJs,
				React: data.React,
				Bootstrap: data.Bootstrap,
				Laravel: data.Laravel,
			}
		}
	}

	data := map[string]interface{} {
		"Project": ProjectDetail,
	}

	var tmpl, err = template.ParseFiles("views/update-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
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

func AddProject(c echo.Context) error {
	title := c.FormValue("input-project-name")
	startDate := c.FormValue("input-date-start")
	endDate := c.FormValue("input-date-end")
	content := c.FormValue("input-deskripsi")
	duration := calculateDuration(startDate, endDate)
	nodeJs := c.FormValue("node-js")
	react := c.FormValue("react")
	bootstrap := c.FormValue("bootstrap")
	laravel := c.FormValue("laravel")


	fmt.Println("Title :", title)
	fmt.Println("Duration :", duration)
	fmt.Println("Content :", content)
	fmt.Println(nodeJs)
	fmt.Println(react)
	fmt.Println(bootstrap)
	fmt.Println(laravel)

	var newProject = Project{
		Title: title,
		Content: content,
		StartDate: startDate,
		EndDate: endDate,
		Duration: duration,
		NodeJs: (nodeJs == "nodeJs"),
		React: (react == "react"),
		Bootstrap: (bootstrap == "bootstrap"),
		Laravel: (laravel == "laravel"),
	}

	dataProject = append(dataProject, newProject)

	fmt.Println(dataProject)
	
	return c.Redirect(http.StatusMovedPermanently, "/#my-project")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index :", id)

	dataProject = append(dataProject[:id], dataProject[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/#my-project")
}

func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index :", id)

	title := c.FormValue("input-project-name")
	startDate := c.FormValue("input-date-start")
	endDate := c.FormValue("input-date-end")
	content := c.FormValue("input-deskripsi")
	duration := calculateDuration(startDate, endDate)
	nodeJs := c.FormValue("node-js")
	react := c.FormValue("react")
	bootstrap := c.FormValue("bootstrap")
	laravel := c.FormValue("laravel")


	var updateProject = Project{
		Title: title,
		Content: content,
		StartDate: startDate,
		EndDate: endDate,
		Duration: duration,
		NodeJs: (nodeJs == "nodeJs"),
		React: (react == "react"),
		Bootstrap: (bootstrap == "bootstrap"),
		Laravel: (laravel == "laravel"),
	}

	dataProject[id] = updateProject

	return c.Redirect(http.StatusMovedPermanently, "/#my-project")
}