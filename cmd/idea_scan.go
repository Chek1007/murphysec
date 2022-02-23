package cmd

import (
	"github.com/spf13/cobra"
	"murphysec-cli-simple/api"
	"murphysec-cli-simple/inspector"
	"murphysec-cli-simple/logger"
	"murphysec-cli-simple/utils/must"
	"path/filepath"
)

func ideaScanCmd() *cobra.Command {
	var dir string
	c := &cobra.Command{
		Hidden: true,
		Use:    "ideascan --dir ProjectDir",
		Run: func(cmd *cobra.Command, args []string) {
			_, e := inspector.Scan(must.String(filepath.Abs(dir)), api.TaskTypeIdea)
			if e != nil {
				SetGlobalExitCode(1)
				logger.Err.Println("idea plugin scan failed.", e.Error())
			}
		},
	}
	c.Flags().StringVar(&dir, "dir", "", "project base dir")
	c.Args = cobra.NoArgs
	must.Must(c.MarkFlagRequired("dir"))
	must.Must(c.MarkFlagDirname("dir"))
	return c
}
