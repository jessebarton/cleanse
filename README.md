# Cleanse

Clean up directories on your system of duplicate files, with option to either remove them or move to a different directory.

## Usage

```
  -delete
    	Bool - Delete files: default false
  -directory string
    	String - Directory to Walk: default empty
  -move
    	Bool - Move files to duplicate directory: default false
```

Examples:
```
cleanse -directory=files/

cleanse -directory=files/ -remove=true

cleasne -directory=files/ -move=true

```

## Task List

- [x] Add command line flag for removeing files.
- [x] Add command line flag for moving files to duplicate folder for review.
- [ ] Add ability to move files from inter folders instead of single directories.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/jessebarton/cleanse/tags).

## Authors

* **Jesse Barton** - *Initial work* - [jessebarton](https://github.com/jessebarton)

See also the list of [contributors](https://github.com/jessebarton/cleanse/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
