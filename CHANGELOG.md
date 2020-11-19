# Changelog

## v1.1.2
**11/19/2020**
- Fixed bug where commands.Parse was not passing in the right flags

## v1.1.1
**11/18/2020**
- Using Go 1.15
- Added error base-types for all errors returned
- Updated errors returned to be wrapped errors for better usage
- Added usage function to flags
- Updated usage function for commands

## v1.1.0
**11/18/2020**
- Added help and version flags/commands
- Removed `Print` function from `Commander` interface
- Added `Usage` function
- Began some documentation

## v1.0.0
- Initial release
- Didn't keep this changelog then, oops
- Added flags for:
    - Int
    - String
    - Bool
    - Uint -- (jcmdln)
- Added commands
