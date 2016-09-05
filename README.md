## http(s) file integrity monitoring using the md5 hash provided by eTag HTTP header

This is avery crude solution to satisfy PCI DSS requirements for monitoring files; used to monitor javascript sources on Amazon S3. It could do more and may one day but for now it doesn't

> PCI Guidance for Requirement no 10.5.5 “File integrity monitoring or change detection systems check for changes to critical files, and notify when such changes are noted. For file integrity monitoring purposes, an entity usually monitors files that don’t regularly change, but when changed indicate a possible compromise“.

> Requirement no 11.5 states that “Deploy a change detection mechanism (for example, file integrity monitoring tools) to alert personnel to unauthorized modification of critical system files, configuration files, or content files; and configure the software to perform critical file comparisons at least weekly“.

> PCI Guidance for Requirement no 11.5 “Change detection solutions such as file integrity monitoring (FIM) tools check for changes to critical files, and notify when such changes are detected. If not implemented properly and the output of the change detection solution monitored, a malicious individual could alter configuration file contents, operating system programs, or application executables. Unauthorized changes, if undetected, could render existing security controls ineffective and/or result in cardholder data being stolen with no perceptible impact to normal processing“.

### Build

`go get github.com/SparkPost/gosparkpost ; go build`

### Required Environment Variables

* Liunx / OS X stanard

```
FIM_SPARKPOST_API_KEY
FIM_NOTIFICATION_TO_EMAIL
FIM_NOTIFICATION_FROM_EMAIL
FIM_URL
```

* Heroku

```
heroku config:set FIM_SPARKPOST_API_KEY="<snip>"
heroku config:set FIM_NOTIFICATION_TO_EMAIL="foo@bob.com"
heroku config:set FIM_NOTIFICATION_FROM_EMAIL="foo@bob.com"
heroku config:set FIM_URL="https://www.google.com/"
```

* Docker

```
docker build -t etag_monitor
docker run -ti -e FIM_SPARKPOST_API_KEY=<snip> -e FIM_NOTIFICATION_TO_EMAIL=<email> -e FIM_NOTIFICATION_FROM_EMAIL=<email> -e FIM_URL=<url>
```

