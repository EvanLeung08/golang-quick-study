package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"regexp"
	"time"

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

func (c *SMBClient) getMatchingFiles(basePath, pattern string) ([]string, error) {
	var results []string
	regex := regexp.MustCompile(pattern)

	fis, err := c.share.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	for _, fi := range fis {
		if !fi.IsDir() && regex.MatchString(fi.Name()) {
			results = append(results, basePath+"/"+fi.Name())
		}
	}

	return results, nil
}

func (c *SMBClient) Download(basePath, remotePattern, localPath string) error {
	matchingFiles, err := c.getMatchingFiles(basePath, remotePattern)
	if err != nil {
		return err
	}

	for _, remoteFIle := range matchingFiles {
		srcFile, err := c.share.Open(remoteFIle)
		if err != nil {
			return err
		}
		defer srcFile.Close()
		stat, err := srcFile.Stat()
		if err != nil {
			return err
		}
		dstFile, err := os.Create(localPath + stat.Name())
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *SMBClient) Delete(basePath, remotePattern string) error {
	matchingFiles, err := c.getMatchingFiles(basePath, remotePattern)
	if err != nil {
		return err
	}

	for _, remotePath := range matchingFiles {
		err := c.share.Remove(remotePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *SMBClient) Rename(basePath, oldPattern, suffix string) error {
	matchingFiles, err := c.getMatchingFiles(basePath, oldPattern)
	if err != nil {
		return err
	}

	for _, oldPath := range matchingFiles {
		err := c.share.Rename(oldPath, oldPath+suffix)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *SMBClient) Search(basePath, pattern string) ([]string, error) {
	return c.getMatchingFiles(basePath, pattern)
}

func main() {
	client, err := NewSMBClient("192.168.50.69", "user", "123456", "LANdrive")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	err = client.Upload("/Users/evan/Downloads/downloaded_file.txt", "test/downloaded_file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Upload files to {%s} successfully\n", "test/")

	err = client.Download("test/", ".*$", "/Users/evan/Downloads/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Download files to {%s} successfully\n", "/Users/evan/Downloads/")

	err = client.Rename("test/", ".*$", "."+time.Now().Format("20060102150405")+".d")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rename files from {%s} successfully\n", "test/")

	results, err := client.Search("test", ".*")
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println("File found:", result)
	}

	err = client.Delete("test/", ".*$")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Delete files from {%s} successfully\n", "test/")
}
