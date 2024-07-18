# Install

## GO

[Guide](https://go.dev/doc/install)

1. Click the download button [here](https://go.dev/doc/install)

    Copy and paste your system's download link and `wget` it

```shell
wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
```

2. Remove any previous Go installation

```shell
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
```

3. Add /usr/local/go/bin to the PATH environment variable

    You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

```shell
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$HOME/go/bin
```

4. Verify that you've installed Go

```shell
go version
```

5. Restart terminal
