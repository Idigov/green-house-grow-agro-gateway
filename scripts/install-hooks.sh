#!/usr/bin/env bash
# Копирует хуки из scripts/git-hooks/ в .git/hooks/
set -e
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
HOOKS_SRC="$ROOT/scripts/git-hooks"
HOOKS_DST="$ROOT/.git/hooks"
if [[ ! -d "$HOOKS_DST" ]]; then
  echo "Not a git repo or .git/hooks missing. Run from repo root."
  exit 1
fi
for f in pre-commit pre-push; do
  cp "$HOOKS_SRC/$f" "$HOOKS_DST/$f"
  chmod +x "$HOOKS_DST/$f"
  echo "Installed $f"
done
echo "Done. Hooks: go fmt, go vet (pre-commit); block push to main (pre-push)."
