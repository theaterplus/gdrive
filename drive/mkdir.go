package drive

import (
    "google.golang.org/api/drive/v3"
)

const DirectoryMimeType = "application/vnd.google-apps.folder"

type MkdirArgs struct {
    Name string
    Parent string
    Share bool
}

func (self *Drive) Mkdir(args MkdirArgs) {
    dstFile := &drive.File{Name: args.Name, MimeType: DirectoryMimeType}

    // Set parent folder if provided
    if args.Parent != "" {
        dstFile.Parents = []string{args.Parent}
    }

    // Create folder
    f, err := self.service.Files.Create(dstFile).Do()
    errorF(err, "Failed to create folder: %s", err)

    PrintFileInfo(PrintFileInfoArgs{File: f})

    //if args.Share {
    //    self.Share(TODO)
    //}
}