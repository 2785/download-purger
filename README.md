_I do not know how to do the fancy link to the fancy pipeline status_

# Download Purger

Download purger is a go application that targets at keeping your downloads folder clean and worry free. It is able to clean up files in the specified folder, and keep a log of the files it deleted. 

## Get it

Linux binary can be downloaded with github releases with the latest tags. Binaries on other platforms will have to be manually built from source. Clone the repo, get [go](https://golang.org/) and build it for your system. 

## Usage

```
Usage of /home/shiqi/go/bin/downloadpurger:
  -dir string
        path to the folder you wish to purge (default "~/Downloads")
  -duration string
        string representing duration of downloads to clear, e.g.:  (default "2week")
  -o string
        path to log file (default "~/logs/dwpurger.log")
  -y    pass in -y to run purger without warning
```
or do `/path/to/your/binary -help`, it will show the same thing. 

- dir specifies what directory to look and purge. Defaults to the home/Downloads of your current folder. Might have weird behaviors on non-linux environments as I didn't bother testing. 
  ```sh
  ./downloadpurger -dir="~/path/to/a/different/folder"
  ```
- duration specifies the age of files that the app purges, app understands `hour`, `day`, `week`, `month` and `year`, defaults to two weeks
  ```sh
  ./downloadpurger -duration="10year" # why even bother cleaning downloads at that point
  ```
- the flag o specifies what file to log the deletion records
  ```sh
  ./downloadpurger -o="./your/path/to/a/log/file.log"
  ```
- pass in `-y` if you do not want a confirmation, e.g., when you run the purger on computer start up and _REALLY_ don't care about the things in your downloads folder :)
  ```sh
  ./downloadpurger -y
  ```
## A setup to run purger on startup and forget about it forever

*DON'T*, if for whatever reasons you depend on the Download folder for crucial tasks. 

On linux: 

1. Grab the binary
2. Make a script that calls the binary
   ```sh
   #!/bin/sh
   /path/to/the/binary -y # and whatever args you want to put in
   ```
3. Add it to cronjob with the `@reboot` tag
   ```sh
   crontab -e
   ```
   And append the line to the crontab file
   ```sh
   @reboot /path/to/the/script
   ```
   save, done. 

## Contribution

Appreciate pull requests or issues :)