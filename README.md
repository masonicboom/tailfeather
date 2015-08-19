Tailfeather is a command line tool for intelligently pretty-printing log files.

It colors repeated values within each field consistently, recycling colors as needed when novel values are encountered. This aids the eye in seeing patterns and anomalies, when tailing live logs.

## Dependencies

  https://github.com/daviddengcn/go-colortext

## Building

Install dependencies, then run `make`.

## Usage

    Usage of ./tlft:
      -input-delimiter=" ": Input field delimiter.
      -output-delimiter="\t": Output field delimiter.

Pass input in via STDIN.

For example, if you have a comma-delimited logfile named LOGFILE.log, that you want to continuously follow, printing with fields tab-delimited, do: `tail -f LOGFILE.log | ./tlft --input-delimiter "," --output-delimiter "\t"`.

## Credits

Example log data from http://ita.ee.lbl.gov/html/contrib/NASA-HTTP.html.
