# fish-history-ui

Fish shell history visualize Web UI.

![screenshot calendar](https://user-images.githubusercontent.com/29811106/130643544-cb9ccdc7-fcff-427b-8211-90c84b373449.png)

## Install

Check [releases](https://github.com/luma-dev/fish-history-ui/releases) and download and install corresponding binary.

### Linux Install Example

```bash
VERSION=1.0.0
ARCH=$(arch)
wget -O fish-history-ui.tar.gz "https://github.com/luma-dev/fish-history-ui/releases/download/v${VERSION}/fish-history-ui_${VERSION}_Linux_${ARCH}.tar.gz"
tar -xvf ./fish-history-ui.tar.gz
sudo install fish-history-ui /usr/local/bin/
rm ./fish-history-ui.tar.gz ./fish-history-ui

./fish-history-ui --version
```
