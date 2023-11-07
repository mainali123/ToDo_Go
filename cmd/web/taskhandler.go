package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (app *application) taskHandlerGet(c *gin.Context) {
	type tasks struct {
		TaskID          int       `json:"taskID"`
		TaskTitle       string    `json:"taskTitle"`
		TaskDescription string    `json:"taskDescription"`
		DueDateTime     time.Time `json:"dueDateTime"`
		Priority        string    `json:"priority"`
		TaskStatus      string    `json:"taskStatus"`
		ProjectID       int       `json:"projectID"`
		UserID          int       `json:"userID"`
		ParentTaskID    int       `json:"parentTaskID"`
	}
	getTasks := "SELECT * FROM Tasks WHERE UserID=?"
	rows, err := app.database.DB.Query(getTasks, USERID)
	if err == sql.ErrNoRows {
		app.infoLog.Println("Returned empty json value since there is no data.")
		c.JSON(http.StatusOK, tasks{})
	} else if err != nil {
		app.errorLog.Fatal(err.Error())
	}
	defer rows.Close()

	var allTasks []tasks
	for rows.Next() {
		var task tasks
		err := rows.Scan(&task.TaskID, &task.TaskTitle, &task.TaskDescription, &task.DueDateTime, &task.Priority, &task.TaskStatus, &task.ProjectID, &task.UserID, &task.ParentTaskID)
		if err != nil {
			app.errorLog.Fatal(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		allTasks = append(allTasks, task)
	}
	c.JSON(http.StatusOK, allTasks)
	/*getTasks := "SELECT * FROM Tasks WHERE UserID=?"
	row := app.database.DB.QueryRow(getTasks, USERID)

	var task tasks
	err := row.Scan(&task.TaskID, &task.TaskTitle, &task.TaskDescription, &task.DueDateTime, &task.Priority, &task.TaskStatus, &task.ProjectID, &task.UserID, &task.ParentTaskID)
	if err != nil {
		app.errorLog.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, task)*/
}
