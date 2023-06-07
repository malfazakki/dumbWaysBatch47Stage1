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
	Title string
	Content string
	StartDate string
	EndDate string
	Duration string
	NodeJs string
	React string
	Bootstrap string
	Laravel string
}

var dataProject = []Project {
	{
		Title: "Dumbways Web Apps 2023",
		Content: "Content",
		StartDate: "2023/07/01",
		EndDate: "2023/07/06",  
		Duration: "3 months",
		NodeJs: "<i class='fa-brands fa-node-js fa-xl'></i>",
	},
	{
		Title: "Dumbways Web Apps 2023",
		Content: "Content 2",
		StartDate: "2023/07/01",
		EndDate: "2023/07/06",
		Duration: "3 months",
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
	//post
	e.POST("/add-project", AddProject)
	e.POST("/project-delete/:id", deleteProject)
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

	if nodeJs != "" {
		nodeJs = "<i class='fa-brands fa-node-js fa-xl'></i>"
	}
	if react != "" {
		react = "<i class='fa-brands fa-react fa-xl''>"
	}
	if bootstrap != "" {
		bootstrap = "<i class='fa-brands fa-bootstrap fa-xl'></i>"
	}
	if nodeJs != "" {
		nodeJs = "<i class='fa-brands fa-node-js fa-xl'></i>"
	}

	fmt.Println("Title :", title)
	fmt.Println("Duration :", duration)
	fmt.Println("Content :", content)

	var newProject = Project{
		Title: title,
		Content: content,
		StartDate: startDate,
		EndDate: endDate,
		Duration: duration,
		NodeJs: nodeJs,
	}

	dataProject = append(dataProject, newProject)

	fmt.Println(dataProject)
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index :", id)

	dataProject = append(dataProject[:id], dataProject[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}