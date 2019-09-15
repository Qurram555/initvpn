# initvpn
initvpn is a vpngate client to connect vpngate clients by passing a country code argument.

**THIS PROJECT IS UNMAINTAINED.**  
I wrote this a long time ago and I do not use or maintain this script any more.
Using this will very likely cause permanent nuclear winter everywhere, so just
look for something better, alright? I'll likely delete the project soon so that
I don't distribute potentially insecure tools.

`autovpn` is a tool to automatically connect you to a random VPN in a country
of your choice. It uses [openvpn][openvpn] to connect you to a server obtained
from [VPN Gate](http://www.vpngate.net/en/).

### Compiling

First clone the repo and `cd` into the directory:

```bash
$ git clone https://github.com/Qurram555/initvpn
$ cd ~/initvpn-master/src
```

Then run this to generate the executable:

```bash
$ go build initvpn.go
```

It's Golang. What do you expect?

### Requirements

This requires [openvpn].

If you're on a `apt`-based distro:

```bash
$ sudo apt-get install openvpn
```

### Usage

Simply run:

```bash
$ ./initvpn
```

and you're done. You'll be connected to a server in the US. Welcome to the US!

You can give a country if you want. For example, if you want to connect to a server
in South Korea:

```bash
$ ./initvpn KR
```

You may need superuser privileges when you arent using linux/deb. Don't worry, I'm not running `rm -rf --no-preserve-root /`
underneath. It's for `openvpn`.

### To run executable on Linux/Ubuntu/Deb
```
$ sudo cp ~/initvpn-master/bin /usr/local/bin

$ sudo initvpn KR 
```

### Contributing

All patches welcome!

### Disclaimer

This is completely insecure. Please do not use this for anything important. Get a
real and secure VPN. This is mostly a fun tool to get a VPN for a few minutes. To 
make it secure add your own certificates at the end of file. 

### License

```
    Since it is a fun tool. So no license.
