# bm Directory Bookmarking Tool (bm)

`bm` is a cross-platform CLI tool written in Go for efficient directory bookmarking and navigation. It allows you to save directories with aliases, quickly navigate to them, and manage your bookmarks seamlessly.

## Installation

### Prerequisites
- Go 1.20 or higher
- Git

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/farhancdr/go-bookmark.git
   cd go-bookmark
   ```
2. Build the tool:
   ```bash
   make build
   ```
   This creates the `bm` binary in the `bin/` directory.
3. Install the tool globally (optional):
   ```bash
   make install
   ```
   This places `bm` in your `$GOPATH/bin` or `$HOME/go/bin`. Ensure this directory is in your `$PATH`.
4. Set up the shell function for `goto` (see below).

### Shell Setup
To enable directory navigation with `bm goto`, you need a shell function to capture the output path and execute `cd`. Add the following to your shell configuration file:

#### Bash/Zsh (`~/.bashrc` or `~/.zshrc`)
```bash
function bm_goto() {
    local target_dir=$(bm goto "$1")
    if [ -n "$target_dir" ] && [ -d "$target_dir" ]; then
        cd "$target_dir"
    else
        echo "Error: $target_dir"
    fi
}
alias bm_g='bm_goto'
```

#### Fish (`~/.config/fish/config.fish`)
```fish
function bm_goto
    set target_dir (bm goto $argv[1])
    if test -n "$target_dir" -a -d "$target_dir"
        cd $target_dir
    else
        echo "Error: $target_dir"
    end
end
```

After adding, reload your shell configuration:
```bash
source ~/.bashrc  # or ~/.zshrc or ~/.config/fish/config.fish
```

## Usage

### Commands
- **Save a directory**:
  ```bash
  bm save my-project
  ```
  Saves the current directory with the alias `my-project`. Prompts for confirmation if the alias exists.

- **Navigate to a bookmark**:
  ```bash
  bm_g my-project
  ```
  Navigates to the directory associated with `my-project`. Requires the shell function above.

- **List all bookmarks**:
  ```bash
  bm list
  ```
  Displays aliases, paths, last updated times, and existence status in a table.

- **Delete a bookmark**:
  ```bash
  bm delete my-project
  ```
  Removes the `my-project` bookmark after confirmation.

- **Clear all bookmarks**:
  ```bash
  bm clear
  ```
  Deletes all bookmarks after confirmation.

- **Update a bookmark**:
  ```bash
  bm update my-project /new/path
  ```
  Updates `my-project` to point to `/new/path`. If no path is provided, uses the current directory.

- **Rename a bookmark**:
  ```bash
  bm rename my-project new-project
  ```
  Renames the alias `my-project` to `new-project`.

- **View bookmark info**:
  ```bash
  bm info
 my-project
  ```
  Shows detailed information about the `my-project` bookmark.

### Configuration File
Bookmarks are stored in a JSON file at:
- `~/.config/bm/bookmarks.json` (Linux/macOS)
- `%APPDATA%\`bm\bookmarks.json` (Windows)

Example JSON structure:
```json
{
  "bookmarks": [
    {
      "alias": "my-project",
      "path": "/home/user/projects/my-project",
      "created_at": "2023-10-01T12:00:00Z",
      "updated_at": "2023-10-01T12:00:00Z"
    }
  ]
}
```

## Troubleshooting
- **Command not found**: Ensure `bm` is in your `$PATH` after running `make install`.
- **Navigation not working**: Verify the shell function (`bm_goto`) is correctly added and sourced.
- **Bookmark file errors**: Check permissions for `~/.config/bm/bookmarks.json`. Delete the file to reset bookmarks.
- **Invalid alias**: Aliases must contain only alphanumeric characters, underscores, or hyphens.

## Development
- **Build**: `make build`
- **Install**: `make install`
- **Clean**: `make clean`

Dependencies:
- [spf13/cobra](https://github.com/spf13/cobra) for CLI parsing

## License
MIT License
```

This implementation provides a fully functional `bm` CLI tool with all requested features, including:
- Cobra-based CLI structure
- JSON-based bookmark storage with mutex-protected file operations
- Platform-appropriate configuration directory
- Comprehensive error handling
- Clear user feedback
- Shell integration instructions
- Build and installation scripts

The code follows Go best practices, is modular, and includes a detailed README for user adoption. Let me know if you need further enhancements or clarifications!
