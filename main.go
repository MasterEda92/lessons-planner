package main

import (
	"embed"

	classsubjects "github.com/MasterEda92/lessons-planner/class_subjects"
	"github.com/MasterEda92/lessons-planner/timetables"

	"github.com/MasterEda92/lessons-planner/classes"
	"github.com/MasterEda92/lessons-planner/db"
	"github.com/MasterEda92/lessons-planner/subjects"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	// setup db client
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	// _, dberr := client.Subject.CreateOne(
	// 	db.Subject.ID.Set(uuid.NewString()),
	// 	db.Subject.Name.Set("Deutsch"),
	// ).Exec(context.Background())
	// if dberr != nil {
	// 	panic(dberr)
	// }

	// _, dberr = client.Subject.CreateOne(
	// 	db.Subject.ID.Set(uuid.NewString()),
	// 	db.Subject.Name.Set("Ev. Religion"),
	// 	db.Subject.Notes.Set("Nur evangelischer Teil der Klasse. Raum 4711A"),
	// ).Exec(context.Background())
	// if dberr != nil {
	// 	panic(dberr)
	// }

	// _, dberr := client.Class.CreateOne(
	// 	db.Class.ID.Set(uuid.NewString()),
	// 	db.Class.Name.Set("8f"),
	// 	db.Class.Grade.Set(8),
	// 	db.Class.PupilCount.Set(23),
	// 	db.Class.PupilCountTimeIncrease.Set(2),
	// ).Exec(context.Background())
	// if dberr != nil {
	// 	panic(dberr)
	// }

	// _, dberr := client.Class.CreateOne(
	// 	db.Class.ID.Set(uuid.NewString()),
	// 	db.Class.Name.Set("5c"),
	// 	db.Class.Grade.Set(5),
	// 	db.Class.PupilCount.Set(21),
	// 	db.Class.PupilCountTimeIncrease.Set(3),
	// ).Exec(context.Background())
	// if dberr != nil {
	// 	panic(dberr)
	// }

	// _, dberr = client.Class.CreateOne(
	// 	db.Class.ID.Set(uuid.NewString()),
	// 	db.Class.Name.Set("8Reli"),
	// 	db.Class.Grade.Set(8),
	// 	db.Class.PupilCount.Set(10),
	// 	db.Class.PupilCountTimeIncrease.Set(1),
	// ).Exec(context.Background())
	// if dberr != nil {
	// 	panic(dberr)
	// }

	// setup repos
	class_repo := classes.NewClassesRepo(client)
	subject_repo := subjects.NewSubjectsRepo(client)
	clsub_repo := classsubjects.NewClassSubjectsRepo(client)
	timetable_repo := timetables.NewTimetablesRepo(client)

	// Create an instance of the app structure
	app := NewApp(class_repo, subject_repo, clsub_repo, timetable_repo)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "myproject",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
