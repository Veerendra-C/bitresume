package pointshandlers

import (
	activitygraph "bitresume/api/dashboard/activity_graph"
	"bitresume/config"
	"bitresume/models"
	"math"

	"github.com/gin-gonic/gin"
)

func HandlePointlogs(c *gin.Context) {
	var data models.Points_Logs
	if err := c.ShouldBindJSON(&data); err != nil{
		c.JSON(500, err.Error())
		return
	}
	rollno := data.RollNo
	source := data.Source
	points := data.Points
	desc := data.Description
	sem := data.Sem
	currdate := data.Currdate
	var newpoints float64
	rank , rankerr := activitygraph.FetchDataRank(rollno)

	if rankerr != nil {
		c.JSON(500, rankerr.Error())
		return
	}

	if source == "PS" {
		if rank.Current_rank == "TITANIUM"{
			if points > 0{
				newpoints = float64(points) *  0.5/300.0
			}else if points == 0{
				newpoints = 0
			}else {
				newpoints = -1
			}
		}else if rank.Current_rank == "GOLD" {
			if points > 0{
				newpoints = float64(points) *  1/300.0
			}else if points == 0{
				newpoints = 0
			}else {
				newpoints = -0.5
			}
		}else {
			if points > 0{
				newpoints = float64(points) *   2/300.0
			}else if points == 0{
				newpoints = 0
			}else {
				newpoints = -0.5
			}
		}
	}

	newpoints = math.Round(newpoints*100)/100

	stmp ,reqerr := config.DB.Prepare("INSERT INTO points_logs(rollno,source,points,description,sem,currdate) values (?,?,?,?,?,?)")

	if reqerr != nil {
		c.JSON(500, reqerr.Error())
		return
	}

	_,execErr := stmp.Exec(rollno,source,newpoints,desc,sem,currdate)

	if execErr != nil {
		c.JSON(500,execErr.Error())
	}

	activitygraph.HandleActivityGraphPoints(rollno,sem,currdate)
}