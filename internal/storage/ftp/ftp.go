package ftp

import (
	"context"
	"fmt"
	"io"

	"github.com/jlaffaye/ftp"
)

type FTPConf struct {
	Host     string
	Port     int
	Username string
	Password string
	Path     string
}

type FTP struct {
	conn *ftp.ServerConn
	conf *FTPConf
}

func NewFTP(conf *FTPConf) (*FTP, error) {
	conn, err := ftp.Dial(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
	if err != nil {
		return nil, fmt.Errorf("ftp.Dial: %v", err)
	}

	if err := conn.Login(conf.Username, conf.Password); err != nil {
		return nil, fmt.Errorf("conn.Login: %v", err)
	}

	if err := conn.ChangeDir(conf.Path); err != nil {
		return nil, fmt.Errorf("conn.ChangeDir: %v", err)
	}

	fmt.Println("FTP Connection Established")

	return &FTP{conn: conn, conf: conf}, nil
}

func (f *FTP) Store(ctx context.Context, key string, data io.Reader) (string, error) {
	if err := f.conn.Stor(key, data); err != nil {
		return "Err", fmt.Errorf("f.conn.Stor: %v", err)
	}

	return f.conf.Path + "/" + key, nil
}
