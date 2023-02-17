package users

// User represents a user from a Jira instance to further work with in the CLI app. A user is
// retrieved from a Jira instance through a service function which calls the Jira API.
type User struct {
	Id        int
	Firstname string
	Lastname  string
	Username  string
	Email     string
}

var dummyUserList []User

func init() {
	// Todo: Temporary ... Remove once I can read users from a Jira instance

	dummyUserList = []User{
		{
			Id:        1,
			Firstname: "Admin",
			Lastname:  "Admin",
			Username:  "admin.admin",
			Email:     "admin@localhost",
		},
		{
			Id:        2,
			Firstname: "Jim",
			Lastname:  "Panse",
			Username:  "jim.panse",
			Email:     "jim.panse@localhost",
		},
		{
			Id:        3,
			Firstname: "Homer",
			Lastname:  "Simpson",
			Username:  "homer.simpson",
			Email:     "homer.simpson@localhost",
		},
		{
			Id:        4,
			Firstname: "Barney",
			Lastname:  "Gumble",
			Username:  "barney.gumble",
			Email:     "barney.gumble@localhost",
		},
		{
			Id:        5,
			Firstname: "Lenny",
			Lastname:  "Lennard",
			Username:  "lenny.lennard",
			Email:     "lenny.lennard@localhost",
		},
		{
			Id:        6,
			Firstname: "Carl",
			Lastname:  "Carlson",
			Username:  "carl.carlson",
			Email:     "carl.carlson@localhost",
		},
		{
			Id:        7,
			Firstname: "Moe",
			Lastname:  "Szyslak",
			Username:  "moe.szyslak",
			Email:     "moe.szyslak@localhost",
		},
	}
}
