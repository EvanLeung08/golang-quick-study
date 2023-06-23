package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hirochachacha/go-smb2"
	"io"
	"net"
	"os"
	"regexp"
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

	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		localPath := c.PostForm("localPath")
		remotePath := c.PostForm("remotePath")
		fmt.Printf("localPath: %s, remotePath: %s\n", localPath, remotePath)
		err := client.Upload(localPath, remotePath)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Upload successful"})
	})

	router.GET("/download", func(c *gin.Context) {
		basePath := c.Query("basePath")
		remotePattern := c.Query("remotePattern")
		localPath := c.Query("localPath")

		err := client.Download(basePath, remotePattern, localPath)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Download successful"})
	})

	router.PUT("/rename", func(c *gin.Context) {
		basePath := c.Query("basePath")
		oldPattern := c.Query("oldPattern")
		suffix := c.Query("suffix")

		err := client.Rename(basePath, oldPattern, suffix)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Rename successful"})
	})

	router.GET("/search", func(c *gin.Context) {
		basePath := c.Query("basePath")
		pattern := c.Query("pattern")

		results, err := client.Search(basePath, pattern)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"files": results})
	})

	router.DELETE("/delete", func(c *gin.Context) {
		basePath := c.Query("basePath")
		remotePattern := c.Query("remotePattern")

		err := client.Delete(basePath, remotePattern)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Delete successful"})
	})

	router.Run(":8080")
}
