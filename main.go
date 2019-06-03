package main

import (
	"log"
	"os"
	"strings"
	"time"
)

// fake data bas
var UserDataBase = []User{
{
		"Mooney",
		"Rich",
		"GONKLE",
		"mooneyrich@gonkle.com"},
	{
		"Cathryn",
		 "Hampton",
		 "GEEKOLA",
		 "cathrynhampton@geekola.com"},
	{
		"Cardenas",
		 "Suarez",
		 "COLAIRE",
		 "cardenassuarez@colaire.com"},
	{
		"Rich",
		 "Newman",
		 "STREZZO",
		 "richnewman@strezzo.com"},
	{
		"Vincent",
		 "Dotson",
		 "CYTREX",
		 "vincentdotson@cytrex.com"},
	{
		"Ada",
		 "Bauer",
		 "SHOPABOUT",
		 "adabauer@shopabout.com"},
	{
		"Tammy",
		 "Garner",
		 "ELECTONIC",
		 "tammygarner@electonic.com"},
		{
		"Erin",
		 "Fuentes",
		 "ENVIRE",
		 "erinfuentes@envire.com"},
	{
		"Mays",
		 "Green",
		 "XSPORTS",
		 "maysgreen@xsports.com"},
	{
		"Duran",
		 "Caldwell",
		 "DIGIQUE",
		 "durancaldwell@digique.com"},
	{
		"Tucker",
		 "Lawson",
		 "IDETICA",
		 "tuckerlawson@idetica.com"},
	{
		"Rhoda",
		 "Holmes",
		 "BICOL",
		 "rhodaholmes@bicol.com"},
	{
		"Bates",
		"Hoover",
		"SHADEASE",
		"bateshoover@shadease.com" },
	{
		"Tabitha",
		"Park",
		"GROK",
		"tabithapark@grok.com"},
	{
		"Boone",
		"Rush",
		"GEEKOL",
		"boonerush@geekol.com"},
	{
		"Wallace",
		"Browning",
		"GEOFORM",
		"wallacebrowning@geoform.com"},
	{
		"Moore",
		"Trujillo",
		"KINETICUT",
		"mooretrujillo@kineticut.com"},
	{
		"Charlotte",
		"Finley",
		"HINWAY",
		"charlottefinley@hinway.com"},
	{
		"Chrystal",
		"Ochoa",
		"ZENTIME",
		"chrystalochoa@zentime.com"},
	{
		"Lily",
		"Rivers",
		"BLURRYBUS",
		"lilyrivers@blurrybus.com"},
	{
		"Jami",
		"Floyd",
		"MEDALERT",
		"jamifloyd@medalert.com"},
	{
		"Lina",
		"Kelley",
		"ORBOID",
		"linakelley@orboid.com"},
	{
		"Chang",
		"Aguirre",
		"STOCKPOST",
		"changaguirre@stockpost.com"},
	{
		"Eva",
		"Camacho",
		"ACCRUEX",
		"evacamacho@accruex.com"},
	{
		"Ida",
		"Morse",
		"XURBAN",
		"idamorse@xurban.com"},
	{
		"Evelyn",
		"Peters",
		"ZAGGLE",
		"evelynpeters@zaggle.com"},
	{
		"Oconnor",
		"Stanton",
		"LIQUIDOC",
		"oconnorstanton@liquidoc.com"},
	{
		"Janna",
		"Santos",
		"ECLIPSENT",
		"jannasantos@eclipsent.com"},
	{
		"Isabel",
		"Wynn",
		"PORTICA",
		"isabelwynn@portica.com"},
	{
		"Schneider",
		"Summers",
		"QUIZMO",
		"schneidersummers@quizmo.com"},
}

// User type
type User struct {
	firstName string
	lastName string
	company string
	email string
}

// Worker type
type Worker struct {
	users []User
	ch chan *User
}

// Create a new worker
func NewWorker(users []User, ch chan *User) *Worker {
	return &Worker{users, ch}
}

// Find a user in the user collection
func (w *Worker) Find(email string, workerName string) {
	for index := range w.users {
		user := &w.users[index]
		if strings.Contains(user.email, email)  {
			log.Printf(workerName)
			w.ch <- user
		}
	}
}

func main() {

	// get argument from command line
	userInput := os.Args[1]

	// create a channel
	ch := make(chan *User)

	// add workers
	log.Printf("We looking for %s ...", userInput)
	go NewWorker(UserDataBase[:15], ch).Find(userInput, "Worker #1")
	go NewWorker(UserDataBase[15:], ch).Find(userInput, "Worker #2")

	for {

		select {
			case user := <- ch:
				log.Printf("User find is %s", user.firstName)
			case <- time.After(1 * time.Second):
				return
		}

	}




}