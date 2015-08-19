![After](https://raw.githubusercontent.com/masonicboom/tailfeather/master/after.png)

Tailfeather is a command line tool that pretty-prints log files, and other delimited text files.

It colors repeated values within each field consistently (recycling colors as needed). When tailing live logs, this helps you see patterns and anomalies.

## Screenshots

### Before

![Before](https://raw.githubusercontent.com/masonicboom/tailfeather/master/before.png)

### After

![After](https://raw.githubusercontent.com/masonicboom/tailfeather/master/after.png)

## System Requirements

Should build on any system with [Go](http://golang.org).

## Dependencies

  https://github.com/daviddengcn/go-colortext

## Building

Install dependencies (see [go get]( https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) documentation), then run `make`.

## Usage

    Usage of ./tlft:
      -input-delimiter=" ": Input field delimiter.
      -output-delimiter="\t": Output field delimiter.

Pass input in via STDIN.

For example, if you have a comma-delimited logfile named LOGFILE.log, that you want to continuously follow, printing with fields tab-delimited, do: `tail -f LOGFILE.log | ./tlft --input-delimiter "," --output-delimiter "\t"`.

## Credits

Example log data from http://ita.ee.lbl.gov/html/contrib/NASA-HTTP.html.
