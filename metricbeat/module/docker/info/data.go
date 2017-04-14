package info

import (
	"github.com/elastic/beats/libbeat/common"

	units "github.com/docker/go-units"
	dc "github.com/fsouza/go-dockerclient"
)

func eventMapping(info *dc.DockerInfo) common.MapStr {
	event := common.MapStr{
		"id": info.ID,
		"containers": common.MapStr{
			"total":   info.Containers,
			"running": info.ContainersRunning,
			"paused":  info.ContainersPaused,
			"stopped": info.ContainersStopped,
		},
		"driver": info.Driver,
		"images": info.Images,
	}

	switch driver := info.Driver; driver {
	case "devicemapper":
		driverStatus := common.MapStr{}
		for _, stat := range info.DriverStatus {
			switch statName := stat[0]; statName {
			case "Backing Filesystem":
				driverStatus["backingFilesystem"] = stat[1]
			case "Base Device Size":
				driverStatus["baseDeviceSize"], _ = units.FromHumanSize(stat[1])
			case "Data file":
				driverStatus["dataFile"] = stat[1]
			case "Data Space Available":
				driverStatus["dataSpaceAvailable"], _ = units.FromHumanSize(stat[1])
			case "Data Space Total":
				driverStatus["dataSpaceTotal"], _ = units.FromHumanSize(stat[1])
			case "Data Space Used":
				driverStatus["dataSpaceUsed"], _ = units.FromHumanSize(stat[1])
			case "Metadata file":
				driverStatus["metadataFile"] = stat[1]
			case "Metadata Space Available":
				driverStatus["metadataSpaceAvailable"], _ = units.FromHumanSize(stat[1])
			case "Metadata Space Total":
				driverStatus["metadataSpaceTotal"], _ = units.FromHumanSize(stat[1])
			case "Metadata Space Used":
				driverStatus["metadataSpaceUsed"], _ = units.FromHumanSize(stat[1])
			case "Pool Blocksize":
				driverStatus["poolBlocksize"], _ = units.FromHumanSize(stat[1])
			case "Pool Name":
				driverStatus["poolName"] = stat[1]
			case "Thin Pool Minimum Free Space":
				driverStatus["thinPoolMinimumFreeSpace"], _ = units.FromHumanSize(stat[1])
			}
		}
		event["driverStatus"] = driverStatus
	}

	return event
}
