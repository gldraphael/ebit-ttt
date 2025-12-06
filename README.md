# ebit-ttt

A simple Tic-Tac-Toe game built with [Ebitengine](https://ebitengine.org/) for fun.

## About

This project is a classic tic-tac-toe (also known as noughts and crosses) game implementation using the Ebitengine game library. It's a learning project to explore game development in Go with Ebitengine.

## Features

- Interactive 3x3 game board
- Two-player gameplay (X and O)
- Mouse-based input for making moves
- Built with Ebitengine v2

## Requirements

- Go 1.24.2 or higher
- Ebitengine v2.8.8

## Installation

```bash
go get github.com/gldraphael/ebit-ttt
```

## Usage

Clone the repository and run:

```bash
go run main.go
```

Click on the grid to place your mark (X or O). Players alternate turns.

## Project Structure

- `main.go` - Entry point of the application
- `game/` - Game loop and rendering logic
  - `RunGame.go` - Game initialization and window setup
  - `screen.go` - Screen rendering and input handling
  - `tile.go` - Tile rendering
- `core/` - Game logic
  - `GameState.go` - Game state management
  - `BoardData.go` - Board data structure
  - `TileValue.go` - Tile value constants

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
