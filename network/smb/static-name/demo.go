package main

import (
	"io"
	"net"
	"os"

	"github.com/hirochachacha/go-smb2"
)

type SMBClient struct {
	conn    net.Conn
	dialer  *smb2.Dialer
	session *smb2.Session
	share   *smb2.Share
}

func NewSMBClient(server, username, password, sharename string) (*SMBClient, error) {
	conn, err := net.Dial("tcp", server+":445")
	if err != nil {
		return nil, err
	}

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     username,
			Password: password,
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		return nil, err
	}

	share, err := s.Mount(sharename)
	if err != nil {
		return nil, err
	}

	return &SMBClient{
		conn:    conn,
		dialer:  d,
		session: s,
		share:   share,
	}, nil
}

func (c *SMBClient) Close() {
	c.share.Umount()
	c.session.Logoff()
	c.conn.Close()
}

func (c *SMBClient) Upload(localPath, remotePath string) error {
	srcFile, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := c.share.Create(remotePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func (c *SMBClient) Download(remotePath, localPath string) error {
	srcFile, err := c.share.Open(remotePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func (c *SMBClient) Delete(remotePath string) error {
	return c.share.Remove(remotePath)
}

func (c *SMBClient) Rename(oldPath, newPath string) error {
	return c.share.Rename(oldPath, newPath)
}

func (c *SMBClient) Search(pattern string) ([]string, error) {
	f, err := c.share.Open(".")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	entries, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}

	var results []string
	for _, entry := range entries {
		if entry.Name() == pattern {
			results = append(results, entry.Name())
		}
	}

	return results, nil
}

func main() {
	client, err := NewSMBClient("192.168.50.69", "user", "123456", "LANdrive")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	err = client.Upload("/Users/evan/Downloads/SMB Usecase.drawio.png", "test/SMB Usecase.drawio.png")
	if err != nil {
		panic(err)
	}

	err = client.Download("test/SMB Usecase.drawio.png", "/Users/evan/Downloads/SMB Usecase.drawio.png.download")
	if err != nil {
		panic(err)
	}

	err = client.Rename("test/SMB Usecase.drawio.png", "test/SMB Usecase.drawio.png.d")
	if err != nil {
		panic(err)
	}

	/*	err = client.Delete("test/SMB Usecase.drawio.png")
		if err != nil {
			panic(err)
		}

		err = client.Rename("test/SMB Usecase.drawio.png", "test/SMB Usecase.drawio.png.d")
		if err != nil {
			panic(err)
		}

		results, err := client.Search("search_pattern")
		if err != nil {
			panic(err)
		}

		for _, result := range results {
			fmt.Println("File found:", result)
		}*/
}
