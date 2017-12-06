package vapp

import (
	"os"
	"path/filepath"
	"runtime"
)

//FromAppDir - gives a absolute path from a path relative to app directory
func (app *App) FromAppDir(relPath string) (abs string) {
	home := os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		home = os.Getenv("APPDATA")
	}
	return filepath.Join(home, "."+app.Name, relPath)
}

//AddModule - registers a module with the app
func (app *App) AddModule(module *Module) {
	app.Modules = append(app.Modules, module)
}

//Init - initialize the application, it can initialized only once. Calling run
//auto-initializes app with defaults
func (app *App) Init() {

}

//Run - runs the applications
func (app *App) Run(args []string) (err error) {
	return err
}

// //Run - runs the application
// func (orek *OrekApp) Run(args []string) (err error) {
// 	if runtime.GOOS != "windows" {

// 	}
// 	app := cli.NewApp()
// 	app.ErrWriter = ioutil.Discard
// 	app.Name = "Orek"
// 	app.Version = "0.0.1"
// 	app.Authors = []cli.Author{
// 		cli.Author{
// 			Name:  "Varun Amachi",
// 			Email: "varunamachi@github.com",
// 		},
// 	}
// 	app.Flags = []cli.Flag{
// 		cli.StringFlag{
// 			Name:  "ds",
// 			Value: "sqlite",
// 			Usage: "Datasource name, sqlite|postgres",
// 		},
// 		cli.StringFlag{
// 			Name:  "db-path",
// 			Value: fromOrekDir("orek.db"),
// 			Usage: "Path to SQLite database [Only applicable for SQLite]",
// 		},
// 		cli.StringFlag{
// 			Name:  "ds-host",
// 			Value: "localhost",
// 			Usage: "DataBase host name [Not applicable for SqliteDataSource]",
// 		},
// 		cli.IntFlag{
// 			Name:  "ds-port",
// 			Value: 5432,
// 			Usage: "DataBase port [Not applicable for SqliteDataSource]",
// 		},
// 		cli.StringFlag{
// 			Name:  "db-name",
// 			Value: "orek",
// 			Usage: "DataBase name [Not applicable for SqliteDataSource]",
// 		},
// 		cli.StringFlag{
// 			Name:  "db-user",
// 			Value: "",
// 			Usage: "DataBase username [Not applicable for SqliteDataSource]",
// 		},
// 		cli.StringFlag{
// 			Name:  "db-password",
// 			Value: "",
// 			Usage: "Option db password for testing " +
// 				"[Not applicable for SqliteDataSource]",
// 		},
// 	}
// 	app.Before = func(ctx *cli.Context) (err error) {
// 		argetr := ArgGetter{Ctx: ctx}
// 		ds := argetr.GetRequiredString("ds")
// 		var store data.OrekDataStore
// 		if ds == "sqlite" {
// 			path := argetr.GetRequiredString("db-path")
// 			dirPath := filepath.Dir(path)
// 			if _, err := os.Stat(dirPath); os.IsNotExist(err) {
// 				err = os.Mkdir(dirPath, 0755)
// 				olog.PrintError("Orek", err)
// 			}
// 			store, err = sqlite.Init(&sqlite.Options{
// 				Path: path,
// 			})
// 			if err == nil {
// 				data.SetStore(store)
// 				// err = data.GetStore().Init()
// 				if err != nil {
// 					olog.Fatal("Orek",
// 						"Data Store initialization failed: %v", err)
// 				} else {
// 					olog.Info("Orek", "%s Data Store initialized", store.Name())
// 				}
// 			}
// 		} else if ds == "postgres" {
// 			host := argetr.GetRequiredString("db-host")
// 			port := argetr.GetRequiredInt("db-port")
// 			dbName := argetr.GetRequiredString("db-name")
// 			user := argetr.GetRequiredString("db-user")
// 			pswd := argetr.GetString("db-password")
// 			if len(pswd) == 0 {
// 				fmt.Printf("Password for %s: ", user)
// 				var pbyte []byte
// 				pbyte, err = terminal.ReadPassword(int(syscall.Stdin))
// 				if err != nil {
// 					olog.Fatal("Orek", "Could not retrieve DB password: %v", err)
// 				} else {
// 					pswd = string(pbyte)
// 				}
// 			}
// 			olog.Print("Orek", `Postgres isnt supported yet. Here are the args
// 				Host: %s,
// 				Port: %d,
// 				DbName: %s,
// 				User: %s`, host, port, dbName, user)
// 		} else {
// 			olog.Fatal("Orek", "Unknown datasource %s requested", ds)
// 		}
// 		return err
// 	}
// 	app.Commands = make([]cli.Command, 0, 30)
// 	for _, cmdp := range orek.CommandProviders {
// 		app.Commands = append(app.Commands, cmdp.GetCommand())
// 	}
// 	err = app.Run(args)
// 	return err
// }

//OrekApp - contains command providers and runs the app
// type OrekApp struct {
// 	CommandProviders []CliCommandProvider
// }
