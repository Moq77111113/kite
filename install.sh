#!/bin/sh
set -e

# Kite installer script
# Usage: curl -sfL https://raw.githubusercontent.com/Moq77111113/kite/main/install.sh | sh


REPO="Moq77111113/kite"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"


OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "$OS" in
  linux*)   OS="linux" ;;
  darwin*)  OS="macOS" ;;
  *)        echo "Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
  x86_64)   ARCH="amd64" ;;
  aarch64)  ARCH="arm64" ;;
  arm64)    ARCH="arm64" ;;
  *)        echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

echo "Installing Kite for $OS/$ARCH..."

echo "Fetching latest release..."
LATEST_VERSION=$(curl -sf "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_VERSION" ]; then
  echo "Error: Could not fetch latest version"
  exit 1
fi

echo "Latest version: $LATEST_VERSION"

FILENAME="kite_${LATEST_VERSION#v}_${OS}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_VERSION/$FILENAME"

echo "Downloading from: $DOWNLOAD_URL"

TMP_DIR=$(mktemp -d)
trap "rm -rf $TMP_DIR" EXIT

cd "$TMP_DIR"
curl -sfL "$DOWNLOAD_URL" -o kite.tar.gz

echo "Extracting..."
tar -xzf kite.tar.gz

echo "Installing to $INSTALL_DIR..."
if [ -w "$INSTALL_DIR" ]; then
  mv kite "$INSTALL_DIR/kite"
  chmod +x "$INSTALL_DIR/kite"
else
  echo "Requesting sudo access to install to $INSTALL_DIR..."
  sudo mv kite "$INSTALL_DIR/kite"
  sudo chmod +x "$INSTALL_DIR/kite"
fi

echo "✓ Kite installed successfully!"
echo ""
echo "Run 'kite --help' to get started."
