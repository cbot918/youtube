package api

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// class data member
type Server struct {
	Port       string
	StaticPath string
	DB         *sql.DB
	DAO        *Dao
}

// class constructor
func New(cfg *Config, dao *Dao) *Server {
	return &Server{
		Port:       cfg.PORT,
		StaticPath: cfg.StaticPath,
		DAO:        dao,
	}
}

// class function member
func (s *Server) Run() {
	// new fiber instance
	app := fiber.New()

	// allow * cors
	app.Use(SetCorsAllowAll())

	// serve react page
	app.Static("/", s.StaticPath)

	// api router here
	app = s.RegistRoute(app)

	// listen for incoming http request
	app.Listen(s.Port)
}

func (s *Server) RegistRoute(a *fiber.App) *fiber.App {

	a.Post("/login", s.Login)
	a.Post("/register", s.Register)

	return a
}

// Controllers
func (s *Server) Login(c *fiber.Ctx) error {
	// get request data
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		fmt.Println("Error parsing request body:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// process data
	table := "users"
	if !s.DAO.EmailExists(user.Email, table) {
		fmt.Println("login: user is not existed")
		return c.SendString("failed: could't found user , please regist first")
	}
	if !s.DAO.PasswordCorrect(user.Email, user.Password, table) {
		fmt.Println("login: wrong password")
		return c.SendString("failed: wrong password")
	}

	// return to client
	fmt.Println("login: success")
	return c.SendString("Login success")
}

func (s *Server) Register(c *fiber.Ctx) error {
	// get request data
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		fmt.Println("Error parsing request body:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// process data
	table := "users"
	if s.DAO.EmailExists(user.Email, table) {
		// email exists
		fmt.Println("register: email exists")
		return c.SendString("failed: email already existed, try another one")
	} else {
		// add user
		ok := s.DAO.AddUser(user.Email, user.Password, table)
		if !ok {
			// db insert user failed
			fmt.Println("register: user to db failed")
			return c.SendString("Regist failed")
		}
	}

	// return to client
	fmt.Println("register: success")
	return c.SendString("Regist success")
}
