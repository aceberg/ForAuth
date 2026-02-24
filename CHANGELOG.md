# Change Log
All notable changes to this project will be documented in this file.

## [0.2.1] - 2026-02-24
### Added
- Multiuser for targets
- Sessions stored in file to restore after app reloads
- Optional client IP info (uses https://ipinfo.io)
- Test notifications
- Second notification option

### Fixed
- Redirect to Referer() instead of "/"

## [0.1.5] - 2025-09-12
### Changed
- Upgraded to `go-1.25.1`
- Moved to maintained `Shoutrrr`: [github.com/nicholas-fedor/shoutrrr](https://github.com/nicholas-fedor/shoutrrr)

### Fixed
- Bug when Shoutrrr notification failed

## [0.1.4] - 2025-03-11
### Added
- Show current sessions on Config page

## [0.1.3] - 2025-03-10
### Added
- Log INFO when user session expires

### Fixed
- Session bug: concurrent map writes

## [0.1.2] - 2025-02-01
### Added
- Multiple targets
- Logs and notifications text updated

### Fixed
- Logout bug

## [0.1.1] - 2024-11-02
### Added
- Login page for Config
- Version file
- Notify on login