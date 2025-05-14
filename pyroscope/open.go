package pyroscope

import (
	"github.com/grafana/pyroscope-go"
	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/app"
	"github.com/pterm/pterm"
	"os"
)

// docker run -itd -p 4040:4040 pyroscope/pyroscope:latest server
func load(address string) {
	if address == "" {
		address = "http://127.0.0.1:4040"
	}

	log.Info("pyroscope address:", address)

	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: app.Name,
		Tags:            map[string]string{"hostname": os.Getenv("HOSTNAME")},
		ServerAddress:   address,
		UploadRate:      0,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
		DisableGCRuns: true,
	})
	if err != nil {
		pterm.Error.Printfln("start pyroscope err:%v", err)
	}
}
