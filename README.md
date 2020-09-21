# procuptime
Convert Linux /proc/uptime into JSON and other formats.

# Description

This program reads one line from the stdin, then interprets and outputs the uptime and reboot time.

# Usage

```
procuptime [-h|-j] 
```
* -h prints help
* -j output in JSON

# Output Format

The program outputs these fields on stdout:

* booted - Unix epoch time of last reboot
* uptime - Seconds since reboot
* datebooted - RFC3339 formated date time of last reboot
* uptimeduration - Duration since reboot hours, min etc

The default format has these fields in the order above on a line. e.g.
```
1600484057 200330 2020-09-19T12:54:17+10:00 55h38m50.46s
```
This is the JSON format

```json
{
    "booted": 1600484057,
    "datebooted": "2020-09-19T12:54:17+10:00",
    "uptime": 200552,
    "uptimeduration": "55h42m32.93s"
}
```
# Installation

This is a Go program so can be run directly from go 

```shell
    $ go get github.com/birchb1024/procuptime
    # And then build in the procuptime directory
    $ go build
```
Or get binary [on GitHub](https://github.com/birchb1024/procuptime/releases) 

# Examples

## Get local uptime

```shell
$ </proc/uptime ./procuptime 
1600484057 200330 2020-09-19T12:54:17+10:00 55h38m50.46s
```

## Get local uptime in JSON

```shell
$ </proc/uptime ./procuptime -j
{
  "booted": 1600484057,
  "datebooted": "2020-09-19T12:54:17+10:00",
  "uptime": 200552,
  "uptimeduration": "55h42m32.93s"
}
```

## Get Uptimes for a list of remote servers 

```shell
for X in server1 server2 server3
do
    echo  $X $(ssh go@$X cat /proc/uptime | ./procuptime ) 
done

server1 1598208885 2475258 2020-08-24T04:54:45+10:00 687h34m18.04s
server2 1598208642 2475501 2020-08-24T04:50:42+10:00 687h38m21.16s
server3 1598208694 2475577 2020-08-24T04:51:34+10:00 687h39m37.9s
```

## Extract the date booted field with JMESPATH tool 'jp'
```shell
$ </proc/uptime ./procuptime -j | jp -u 'datebooted'
2020-09-19T12:54:17+10:00
```

