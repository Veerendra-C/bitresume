package activitygraph

import (
	"bitresume/config"
	"bitresume/models"
	"log"
)

func FetchDataRank(rollno string) (models.ActGph,error){
	rows, err := config.DB.Prepare("select current_rank from activity_graph where rollno = ? order by currdate desc limit 1")
	var r models.ActGph
	if err != nil {
		return r, err
	}

	row := rows.QueryRow(rollno)

	err = row.Scan(&r.Current_rank)
	if err != nil {
		return r, err
	}
	return r, nil
}

func FetchLastPoints(rollno string) (models.ActGph,error){
	rows, err := config.DB.Prepare("select current_point from activity_graph where rollno = ? order by currdate desc limit 1")
	var r models.ActGph
	if err != nil {
		return r, err 
	}

	row := rows.QueryRow(rollno)

	err = row.Scan(&r.Current_point)
	if err != nil {
		return r, err
	}
	return r, nil
}

func HandleActivityGraphPoints(rollno string,sem int, currdate string){
	rows, err := config.DB.Prepare("SELECT SUM(points) FROM points_logs WHERE rollno = ? AND currdate = ?")
	var r models.ActGph
	if err != nil {
		return
	}
	row := rows.QueryRow(rollno,currdate)

	err = row.Scan(&r.Current_point)

	if err != nil {
		log.Fatal("Failed to scan")
		return
	}
	prevPoints, _ := FetchLastPoints(rollno)

	var newpoints float64
	var rank string
	newpoints = float64(prevPoints.Current_point) + float64(r.Current_point)
	
	if newpoints > 100{
		newpoints = 100 
	}
	if newpoints < 70 {
		newpoints = 70
	}
	
	if newpoints >= 90 {
		rank = "TITANIUM"
	}else if newpoints >= 80{
		rank = "GOLD"
	}else{
		rank = "SILVER"
	}

	stmp ,reqerr := config.DB.Prepare("INSERT INTO activity_graph(rollno,current_point,current_rank,sem,currdate) values (?,?,?,?,?)")

	if reqerr != nil {
		log.Fatal("Failed to insert")
		return
	}

	_,execErr := stmp.Exec(rollno,newpoints,rank,sem,currdate)

	if execErr != nil {
		log.Fatal("Failed to Push")
		return
	}
}