#!/bin/bash
set -e

# Default installation directory
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
REPO="farhancdr/go-bookmark"
BINARY_NAME="bm"
GITHUB_API_URL="https://api.github.com/repos/${REPO}/releases/latest"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
  darwin) OS="macOS" ;;
  linux) OS="Linux" ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
  x86_64) ARCH="x86_64" ;;
  aarch64 | arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Get the latest release version
VERSION=$(curl -s "$GITHUB_API_URL" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
if [ -z "$VERSION" ]; then
  echo "Error: Could not fetch latest release version"
  exit 1
fi

# Construct download URL
FILE_NAME="${BINARY_NAME}_${VERSION}_${OS}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/farhancdr/go-bookmark/releases/download/v1.1.2/bm_1.1.2_macOS_arm64.tar.gz"
echo "Downloading ${DOWNLOAD_URL}"

# https://github.com/farhancdr/go-bookmark/releases/download/v1.1.2/bm_v1.1.2_macOS_arm64.tar.gz
# https://github.com/farhancdr/go-bookmark/releases/download/v1.1.2/bm_1.1.2_macOS_arm64.tar.gz
# Download and extract the binary
TEMP_DIR=$(mktemp -d)
trap 'rm -rf "$TEMP_DIR"' EXIT

echo "Downloading ${BINARY_NAME} ${VERSION} for ${OS}/${ARCH}..."
curl -fsSL "$DOWNLOAD_URL" -o "$TEMP_DIR/$FILE_NAME"
if [ $? -ne 0 ]; then
  echo "Error: Failed to download ${FILE_NAME}"
  exit 1
fi

tar -xzf "$TEMP_DIR/$FILE_NAME" -C "$TEMP_DIR"
if [ ! -f "$TEMP_DIR/$BINARY_NAME" ]; then
  echo "Error: Binary not found in archive"
  exit 1
fi

# Install the binary
echo "Installing ${BINARY_NAME} to ${INSTALL_DIR}..."
sudo mv "$TEMP_DIR/$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"

echo "${BINARY_NAME} ${VERSION} installed successfully!"
echo "Run 'bm --help' for usage instructions."
echo "Follow the README to set up the 'bm_goto' shell function for navigation."