# GitHub Activity CLI

A simple command-line tool written in Go that displays the recent activity of a GitHub user.

## Features

- Fetches and displays a user's recent GitHub activity
- Shows different types of events (pushes, pull requests, issues, etc.)
- Handles errors gracefully
- No external dependencies beyond the Go standard library

## Installation

1. Make sure you have Go installed on your system.
2. Clone this repository or download the source code.
3. Build the executable:

```bash
go build -o github-activity
```

4. Move the executable to a location in your PATH (optional):

```bash
# On Linux/macOS
mv github-activity /usr/local/bin/

# On Windows, move to a directory in your PATH
```

## Usage

```bash
github-activity <username>
```

Replace `<username>` with the GitHub username you want to check.

### Example

```bash
$ github-activity kamranahmedse
- Pushed 3 commits to kamranahmedse/developer-roadmap
- Opened issue in kamranahmedse/developer-roadmap: Add Docker roadmap
- Starred microsoft/vscode
```

## Supported Event Types

The tool supports displaying the following GitHub event types:

- Push events
- Pull request events
- Issue events
- Repository creation/deletion
- Stars (Watch events)
- Forks
- Issue comments
- Commit comments
- Releases
- Repository publicity changes
- Collaborator additions

## Error Handling

The tool handles various error scenarios:

- Missing username argument
- Invalid/non-existent GitHub username
- GitHub API errors
- Network issues

## Notes

- The GitHub API has rate limits for unauthenticated requests. If you hit the rate limit, you may need to wait before making more requests.
- This tool uses the public GitHub API and doesn't require authentication.