# GitHub Activity Fetcher

This Go program provides a simple HTTP server that fetches and displays a user's public GitHub activity. It leverages the GitHub API to retrieve recent events and presents them in a user-friendly format within a web browser.

## Functionality

The server does the following:

*   **Fetches GitHub Activity:** Retrieves the public activity feed for a specified GitHub username using the GitHub API.
*   **Displays Activity:** Presents a summarized list of the user's recent GitHub events (e.g., pushed commits, created repositories, etc.) in the browser.
*   **Handles Errors:** Gracefully handles cases where the GitHub username is not provided or if there are issues communicating with the GitHub API.
*   **Listens on Port 8080:** By default, the server listens for incoming HTTP requests on port 8080.

**Prerequisites:**

*   **Go:** Ensure that Go (version 1.18 or newer) is installed on your system. You can download it from the official Go website: [https://go.dev/dl/](https://go.dev/dl/).
*   **Environment Setup (If needed):** If you are not working with Go modules, you may need to set your `GOPATH` environment variable correctly. Note that this is not required when using Go modules.

## Build and Run

1.  **Clone the repository:**
    

