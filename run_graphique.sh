#!/usr/bin/env bash
set -e
cd "$(dirname "$0")"

echo "Lancement de DUNGEON CRAFT (graphique)..."
GALLIUM_DRIVER=d3d12 go run ./graphique
