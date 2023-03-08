<h1 align="center">Shell Mate 🤖</h1>
<p align="center">
    Your very own terminal AI assistant<br>
    <a href="https://github.com/AvicennaJr/ShellMate"><img alt="Shell Mate" src="https://img.shields.io/badge/Shell Mate-green"></a>
    <a href="https://github.com/AvicennaJr/ShellMate"><img alt="Shell Mate" src="https://img.shields.io/badge/platform-Linux | Windows | Android-green.svg"></a>
    <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg">
    <img src="https://img.shields.io/github/go-mod/go-version/gomods/athens.svg">
</p>

## Demo
<img align="center" src="./assets/demo.gif">

## Getting Started
- First create an account on [OpenAI](https://platform.openai.com/overview) and create a new [token key](https://platform.openai.com/account/api-keys).

- Place the token key to your environmental variables or you could optionally put it manually when you run the application.

## How To Install
You can download the binary for your specific platform on the release page and run the program on your terminal. Specific instructions below:

### Linux
- Download the zipfile:
```bash
curl -O -L https://github.com/AvicennaJr/ShellMate/releases/download/v1.0.0/shellmate_linux_amd64_v1.0.0.tar.gz
```
- Extract contents of the zip file and place it on the `/usr/local/bin` directory
```
sudo tar -C /usr/local/bin -xzvf shellmate_linux_amd64_v1.0.0.tar.gz
```
- Start the program with: `shellmate`

### Windows
- Download the zip file from the [releases page](https://github.com/AvicennaJr/ShellMate/releases/download/v1.0.0/shellmate_windows_amd64_v1.0.0.tar.gz)
- Extract the binary file and run the program with:
```powershell
shellmate.exe
```
- You could optionally add the executable to path. See detailed instructions [here](https://medium.com/@kevinmarkvi/how-to-add-executables-to-your-path-in-windows-5ffa4ce61a53)

### Mac OS 
Due to Apple's strict policy, you'll have to build the binary file from sources. See instructions on the next section.

### Building From Sources 
- Make sure you have [go installed](https://go.dev/doc/install)
- Clone the repository with:
```bash
git clone github.com/AvicennaJr/ShellMate
```
- Enter the directory:
```bash
cd ShellMate
```
- Build the program with:
```
go build -o shellmate
```
- Run the created binary with: `shellmate`

## Issues and Contributions
For any problems or suggestions open an issue. If you want to contribute, create a pull request.

## Author
Another tool brought to you by [Avicenna](https://github.com/AvicennaJr/ShellMate)
