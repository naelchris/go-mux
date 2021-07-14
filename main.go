package main

import "os"

func main(){
	a := App{} //defining a byte array with the App object (similar to Array list of App)
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		)

	a.Run(":8010")
}