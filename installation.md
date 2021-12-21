# GOLANG INSTALLATION IN UBUNTU
### Download the main file
<code>
curl -OL https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
</code>

### Extract the contents
<code>
sudo tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz
</code>

### GO to .bashrc file using following command
<code>
nano .bashrc
</code>

### Update.bashrc file using following contants and save it using ctrl+x
<code>
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
</code>

### Run following command to apply changes
<code>
source ~/.bashrc
</code>

#### Go has been installed. To check, run following command:
<code>
  go version
</code>


