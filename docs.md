# Starfield Simulation

## Overview
This project is a 3D starfield simulation implemented in Go using the Ebiten game engine. It creates an immersive experience where stars move toward the viewer, with speed controlled by mouse position. Stars regenerate when they pass the viewer, creating an endless starfield effect.

## Installation
1. Install [Go](https://golang.org/dl/) (version 1.16+ recommended)
2. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/starfield-simulation.git
   cd starfield-simulation
   ```
3. Install dependencies:
   ```bash
   go get github.com/hajimehoshi/ebiten/v2
   ```
4. Build and run:
   ```bash
   go run main.go
   ```

## Usage
- Move your mouse horizontally across the screen to control star speed:
  - Left side: Slowest movement
  - Right side: Fastest movement
- Close the window to exit

## Code Structure
### Main Components
1. `main.go` - Entry point
   - Initializes window dimensions
   - Creates game instance
   - Starts game loop

2. `starfield/game.go` - Core game logic
   - `Game` struct: Manages starfield state
   - `Update()`: Handles star movement and regeneration
   - `Draw()`: Renders stars and their motion trails

3. `starfield/stars.go` - Star implementation
   - `Star` struct: 3D coordinates (x, y, z)
   - `NewStar()`: Creates stars with random positions

4. `utils/utils.go` - Helper functions
   - `Map()`: Linear interpolation utility
