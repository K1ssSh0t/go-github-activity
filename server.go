
package main

// Import necessary packages
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Event struct to represent a GitHub event
type Event struct {
	Type      string    `json:"type"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

// Repo struct to represent a GitHub repository
type Repo struct {
	Name string `json:"name"`
}

// Payload struct to represent the payload of an event
type Payload struct {
	Commits []Commit `json:"commits"`
	Action  string   `json:"action"`
	Issue   Issue    `json:"issue"`
}

// Commit struct to represent a commit
type Commit struct {
	Message string `json:"message"`
}

// Issue struct to represent an issue
type Issue struct {
	Title string `json:"title"`
}

// main function: entry point of the program
func main() {
	// Check if a username is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <username>")
		return
	}

	// Get the username from the command-line arguments
	username := os.Args[1]
	// Fetch GitHub activity for the given username
	events, err := fetchGitHubActivity(username)
	// Handle any errors during the fetch process
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Display the fetched activity
	displayActivity(events)
}

// fetchGitHubActivity function: fetches GitHub activity for a given username
func fetchGitHubActivity(username string) ([]Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response into the events slice
	var events []Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, err
	}

	// Return the events
	return events, nil
}

func displayActivity(events []Event) {
	if len(events) == 0 {
		fmt.Println("No activity found for this user.")
		return
	}
	// Iterate over each event and display it
	for _, event := range events {
		date := event.CreatedAt.Format("2006-01-02 15:04")

		switch event.Type {
		case "PushEvent":
			fmt.Printf("- Pushed %d commits to %s (%s)\n", len(event.Payload.Commits), event.Repo.Name, date)
		case "IssuesEvent":
			fmt.Printf("- %s an issue '%s' in %s (%s)\n", event.Payload.Action, event.Payload.Issue.Title, event.Repo.Name, date)
		case "WatchEvent":
			fmt.Printf("- Starred  %s (%s)\n", event.Repo.Name, date)
		case "CreateEvent":
			fmt.Printf("- Created a repository called %s (%s)\n", event.Repo.Name, date)
		case "PullRequestEvent":
			fmt.Printf("- Opened a pull request in %s (%s)\n", event.Repo.Name, date)
		default:
			fmt.Printf("- %s in %s (%s)\n", event.Type, event.Repo.Name, date)
		}
	}
}
