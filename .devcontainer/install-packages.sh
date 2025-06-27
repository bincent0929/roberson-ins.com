#!/bin/bash
set -e

# Update package lists and upgrade installed packages
sudo apt-get update

# Install required packages
sudo apt-get install -y --no-install-recommends tmux watchman

# Clean up apt caches and remove package lists
sudo apt-get clean
sudo rm -rf /var/lib/apt/lists/*

# Create tmux sessions
tmux has-session -t caddy || tmux new-session -d -s caddy -- caddy run
tmux has-session -t tailwind || tmux new-session -d -s tailwind -- tailwindcss -o ./styles/tailwind-output.css --watch