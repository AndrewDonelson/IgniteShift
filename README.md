# IgniteShift

**IgniteShift** is a powerful command-line tool designed to assist developers in managing their Cosmos SDK installations. A nod to the Cosmos-SDK command-line tool, Ignite, IgniteShift emphasizes the ability to seamlessly switch between different Cosmos SDK versions.

## Features

1. **Easy Installation**: Install any version of Cosmos SDK with a simple command. The latest version is installed by default.
2. **Self-Updating**: IgniteShift can check for its own updates and the updates for Cosmos SDK. Just run `igniteshift` without parameters.
3. **Version Removal**: Effortlessly remove any version of Cosmos SDK. By default, the latest version is removed.
4. **Seamless Upgrades**: Upgrade to the latest version of Cosmos SDK with ease.
5. **Version Management**: View a list of all available Cosmos SDK versions and see which ones are currently installed.
6. **On-the-fly Version Switching**: Switch between different installed versions of Cosmos SDK as per your development needs.

## Prerequisites

- Golang (latest stable version recommended)

## Installation

To install **IgniteShift**:

```
go install github.com/AndrewDonelson/IgniteShift@latest
```

## Usage

- **Check for Updates**: 
  ```
  igniteshift
  ```
  Run `igniteshift` without any parameters to check for updates to both IgniteShift and Cosmos SDK.

- **Install Cosmos SDK**:
  ```
  igniteshift install [version]
  ```
  If no version is specified, the latest version will be installed.

- **Remove Cosmos SDK**:
  ```
  igniteshift remove [version]
  ```
  If no version is specified, the latest version will be removed.

- **Upgrade Cosmos SDK**:
  ```
  igniteshift upgrade
  ```
  This command upgrades the Cosmos SDK to the latest version.

- **List Available Versions**:
  ```
  igniteshift list
  ```
  This command lists all available Cosmos SDK versions and indicates which ones are currently installed.

- **Switch Between Versions**:
  ```
  igniteshift switch [version]
  ```
  This command allows you to switch between different installed versions of Cosmos SDK.

## Contribution

Feel free to fork the project, open issues, or submit pull requests. All contributions are welcome!

### Repository

[IgniteShift on GitHub](https://github.com/AndrewDonelson/IgniteShift)

## Task List

- [ ] Setup the basic Go project structure.
- [ ] Integrate the `cobra` library for command-line parameter handling.
- [ ] Implement the self-updating feature:
  - [ ] Check for the latest version of IgniteShift.
  - [ ] Check for the latest version of Cosmos SDK.
  - [ ] Prompt the user for upgrades if available.
- [ ] Implement the installation feature:
  - [ ] Fetch the desired Cosmos SDK version.
  - [ ] Install the fetched version at the default location.
- [ ] Implement the removal feature:
  - [ ] Identify the version to be removed.
  - [ ] Remove the selected version.
- [ ] Implement the upgrade feature:
  - [ ] Check for the latest Cosmos SDK version.
  - [ ] Upgrade to the latest version if not already installed.
- [ ] Implement version management:
  - [ ] List all available Cosmos SDK versions.
  - [ ] Indicate which versions are currently installed.
- [ ] Implement on-the-fly version switching:
  - [ ] Allow users to switch between different installed versions.
- [ ] Write comprehensive tests for each feature.
- [ ] Document usage instructions in the README.
- [ ] Publish the first release.

## License

Apache License 2.0