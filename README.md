# Linux touch for Windows

touch changes the access and/or modification timestamps of the specified files, similar to the unix touch command

* A FILE argument that does not exist is created empty, unless ```-c``` or ```-h``` is supplied.
* Current time is used if none of ```-u``` or ```-t``` supplied
* If you supply both ```-u``` and ```-t``` only ```-u``` will be used

```
Usage: touch [OPTION]... FILE...

  -u STAMP        Unix timestamp, seconds since Jan 01 1970
  -t TIME         Use time format YYYYMMDDHHMMSS
  -c              Do no create any files
  -a              Change only the access time
  -m              Change only the modification time
```

## Download

You can download the touch.exe executable for Windows (amd64) from the [Release](adambg/win-touch/release/latest/) page


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.