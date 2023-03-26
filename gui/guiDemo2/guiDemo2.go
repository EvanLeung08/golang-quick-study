package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andlabs/ui"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	err := ui.Main(func() {
		inputUsername := ui.NewEntry()
		inputPassword := ui.NewPasswordEntry()
		inputKeyFile := ui.NewEntry()
		inputKeyPassword := ui.NewPasswordEntry()
		inputHostname := ui.NewEntry()
		inputPort := ui.NewEntry()
		inputPath := ui.NewEntry()
		button := ui.NewButton("Connect")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Username:"), false)
		box.Append(inputUsername, false)
		box.Append(ui.NewLabel("Password:"), false)
		box.Append(inputPassword, false)
		box.Append(ui.NewLabel("Key File:"), false)
		box.Append(inputKeyFile, false)
		box.Append(ui.NewLabel("Key Password:"), false)
		box.Append(inputKeyPassword, false)
		box.Append(ui.NewLabel("Hostname:"), false)
		box.Append(inputHostname, false)
		box.Append(ui.NewLabel("Port:"), false)
		box.Append(inputPort, false)
		box.Append(ui.NewLabel("Path:"), false)
		box.Append(inputPath, false)
		box.Append(button, false)
		window := ui.NewWindow("SFTP Client", 400, 300, false)
		window.SetMargined(true)
		window.SetChild(box)

		button.OnClicked(func(*ui.Button) {
			username := inputUsername.Text()
			password := inputPassword.Text()
			keyFile := inputKeyFile.Text()
			keyPassword := inputKeyPassword.Text()
			hostname := inputHostname.Text()
			port := inputPort.Text()
			path := inputPath.Text()

			if username == "" || (password == "" && keyFile == "") || hostname == "" || port == "" || path == "" {
				ui.MsgBoxError(window, "Error", "Please fill in all fields.")
				return
			}

			// create SSH client config
			config := &ssh.ClientConfig{
				User: username,
				Auth: []ssh.AuthMethod{},
			}
			if password != "" {
				config.Auth = append(config.Auth, ssh.Password(password))
			}
			if keyFile != "" {
				key, err := ssh.ParsePrivateKey([]byte(keyFile))
				if err != nil {
					ui.MsgBoxError(window, "Error", "Failed to parse key file.")
					return
				}
				if keyPassword != "" {
					config.Auth = append(config.Auth, ssh.PublicKeysCallback(func() ([]ssh.Signer, error) {
						return []ssh.Signer{key}, nil
					}), ssh.Password(keyPassword))
				} else {
					config.Auth = append(config.Auth, ssh.PublicKeys(key))
				}
			}

			// dial SSH server
			conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
			if err != nil {
				ui.MsgBoxError(window, "Error", "Failed to dial SSH server.")
				return
			}

			// open SFTP session
			sftpClient, err := sftp.NewClient(conn)
			if err != nil {
				ui.MsgBoxError(window, "Error", "Failed to open SFTP session.")
				return
			}
			defer sftpClient.Close()

			// download
			localPath := filepath.Join(os.TempDir(), filepath.Base(path))
			localFile, err := os.Create(localPath)
			if err != nil {
				ui.MsgBoxError(window, "Error", "Failed to create local file.")
				return
			}
			defer localFile.Close()

			remoteFile, err := sftpClient.Open(path)
			if err != nil {
				ui.MsgBoxError(window, "Error", "Failed to open remote file.")
				return
			}
			defer remoteFile.Close()

			if _, err := remoteFile.WriteTo(localFile); err != nil {
				ui.MsgBoxError(window, "Error", "Failed to download file.")
				return
			}

			ui.MsgBox(window, "Success", fmt.Sprintf("File downloaded to %s.", localPath))
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
