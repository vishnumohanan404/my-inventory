package main

func main() {
	app := App{}
	app.Initialize(DbUser, DbName)
	app.Run("localhost:10000")
}
