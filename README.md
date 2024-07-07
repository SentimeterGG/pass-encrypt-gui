# Pass Encrypt GUI

Pass Encrypt GUI is a graphical user interface (GUI) for the pass-encrypt tool, designed to facilitate the creation of complex passwords, ensuring enhanced security against potential hacking attempts.
## Features

- User-friendly GUI for easy password generation.
- Creation of complex, secure passwords.
- No command-line knowledge required.

## Getting Started
### Prerequisites

No prerequisites are required to run Pass Encrypt GUI.
## Installation

- Download the latest release package from the Releases section.
- Run the executable file to start the application.

### FOR OTHER DISTRIBUTION

1. **Install Golang:** <br />
   Follow the official Golang installation guide: Golang Installation
   
2. **Install Docker:** <br />
Follow the official Docker installation guide: Docker Installation

3. **Download the Source Code:** <br />
Download the latest source code of Pass-Encrypt GUI from the repository.

4. **Extract the Source Code:** <br />
Extract the downloaded source code archive to a directory of your choice.

### Setting Up the Environment

Navigate to the directory where you extracted the source code. Run the following commands to set up the environment:
```
go get fyne.io/fyne/v2@latest
go install fyne.io/fyne/v2/cmd/fyne@latest
go install github.com/fyne-io/fyne-cross@latest
go get github.com/fyne-io/fyne-cross
```
### Building the Application

After setting up the environment, you can build the application using fyne-cross. Run the following command, replacing <your os>, <your architecture>, and <whatever you like> with your desired values:

```
fyne-cross <your os> -arch=<your architecture> --name <whatever you like> --app-id com.app.encrypt
```
### Locating the Built Application

After the build process is complete, you can find the built application in the `fyne-cross/bin` directory.

### Example

Here’s an example of the build command for a Windows 64-bit architecture:
```
fyne-cross windows -arch=amd64 --name PassEncrypt --app-id com.app.encrypt
```
In this example, the built application will be located in `fyne-cross/bin/windows-amd64`.


## License

This project is licensed under the GNU General Public License v3.0. For more details, see the LICENSE file.
