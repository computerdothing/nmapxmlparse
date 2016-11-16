package util

import (
  "fmt"
  "golang.org/x/sys/unix"
  "os"
)

func Exists(fn string) bool {
        if _, err := os.Stat(fn); os.IsNotExist(err) {

                return false
        }
        return true
}

func DoFlock(fd int) error {
        // Using unix.Flock here to enable testing of SafeOpenFile
        if err := unix.Flock(fd, unix.LOCK_EX); err != nil {
                return err
        }
        return nil
}


func SafeOpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
        var prestat unix.Stat_t
        err := unix.Lstat(name, &prestat)
        if err == unix.ENOENT {
                flag = flag | os.O_EXCL

                f, err := os.OpenFile(name, flag, perm)
                if err != nil {
                        return nil, err
                }

                return f, nil

        } else if err != nil {
                fmt.Println(err)
                return nil, err
        }

        if (prestat.Mode&unix.S_IFMT)&unix.S_IFLNK == unix.S_IFLNK {
                return nil, fmt.Errorf("File is a symlink and we don't like that.")
        }

        truncd := false
        if flag&os.O_TRUNC != 0 {
                flag = flag &^ os.O_TRUNC
                truncd = true
        }

        f, err := os.OpenFile(name, flag, perm)
        if err != nil {
                return nil, err
        }

        fd := f.Fd()
        if err := DoFlock(int(fd)); err != nil {
                return nil, err
        }

        var openstat unix.Stat_t
        if err := unix.Fstat(int(fd), &openstat); err != nil {
                return nil, err
        }

        if prestat.Ino != openstat.Ino {
                return nil, fmt.Errorf("Inode doesn't match what we were expecting")
        }

        if truncd {
                f.Truncate(0)
        }
        return f, nil
}
