## http(s) file integrity monitoring using the md5 hash provided by eTag HTTP header

This is avery crude solution to satisfy PCI DSS requirement 10.5.5; used to monitor javascript sources on Amazon S3. It could do more and may one day but for now it doesn't

| PCI Guidance for Requirement no 10.5.5 “File integrity monitoring or change detection systems check for changes to critical files, and notify when such changes are noted. For file integrity monitoring purposes, an entity usually monitors files that don’t regularly change, but when changed indicate a possible compromise“.


### Build

`go get github.com/SparkPost/gosparkpost ; go build`

### Required Environment Variables

FIM_SPARKPOST_API_KEY
FIM_NOTIFICATION_TO_EMAIL
FIM_NOTIFICATION_FROM_EMAIL
FIM_URL
