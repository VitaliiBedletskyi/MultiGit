# MultiGit: Effortless Git Repositories Management Tool

**MultiGit** is a simple CLI tool designed to streamline and automate batch processing of commands across multiple Git repositories. Ideal for developers and teams managing multiple repositories, MultiGit provides an intuitive way to handle various Git actions, saving you time and reducing repetitive tasks.

## Key Features

- **Batch Operations**: Run Git commands across all repositories defined in a single configuration file (`.mgitrc`), allowing you to efficiently manage multiple projects.
- **Configuration-Driven**: Use `mgit init` to set up repositories and initialize a `.mgitrc` file to keep your setup organized and consistent.

## Commands

- **`add`**: Add a new repository and clone it locally.
- **`checkout`**: Switch all repositories to the specified branch. If no branch is specified, MultiGit will default to the branch defined in the `.mgitrc` configuration.
- **`clone`**: Clone all repositories listed in your `.mgitrc` file.
- **`completion`**: Generate autocompletion scripts for popular shells, improving command-line efficiency.
- **`help`**: View detailed help information for any command.
- **`init`**: Initialize a new `.mgitrc` configuration file in your specified directory.
- **`pull`**: Pull the latest changes for all configured repositories.
- **`run`**: Execute any custom command across all repositories, supporting batch execution of scripts or Git commands.
