package bosch

import (
	"github.com/ffenix113/zigbee_home/cli/sensor/base"
	"github.com/ffenix113/zigbee_home/cli/types/appconfig"
	dt "github.com/ffenix113/zigbee_home/cli/types/devicetree"
	"github.com/ffenix113/zigbee_home/cli/zcl/cluster"
)

type BME280 struct {
	*base.Base `yaml:",inline"`
}

func (BME280) String() string {
	return "Bosch BME280"
}

func (BME280) TemplatePrefix() string {
	return "bosch_bme280"
}

func (BME280) Clusters() cluster.Clusters {
	return []cluster.Cluster{
		cluster.Temperature{
			MinMeasuredValue: -40,
			MaxMeasuredValue: 85,
			Tolerance:        1,
		},
		cluster.Pressure{
			MinMeasuredValue: 30,
			MaxMeasuredValue: 110,
			Tolerance:        0,
		},
		cluster.RelativeHumidity{
			MinMeasuredValue: 10,
			MaxMeasuredValue: 90,
		},
	}
}

func (BME280) AppConfig() []appconfig.ConfigValue {
	return []appconfig.ConfigValue{
		appconfig.CONFIG_I2C.Required(appconfig.Yes),
		appconfig.CONFIG_SENSOR.Required(appconfig.Yes),
		appconfig.CONFIG_BME280.Required(appconfig.Yes),
	}
}

func (BME280) ApplyOverlay(tree *dt.DeviceTree) error {
	i2c1Node := tree.FindSpecificNode(dt.SearchByLabel(dt.NodeLabelI2c1))
	if i2c1Node == nil {
		return dt.ErrNodeNotFound(dt.NodeLabelI2c1)
	}

	i2c1Node.AddNodes(&dt.Node{
		Name:        "bme280",
		UnitAddress: "76",
		Properties: []dt.Property{
			dt.NewProperty(dt.PropertyNameCompatible, dt.FromValue("bosch,bme280")),
			dt.NewProperty("reg", dt.Angled("0x76")),
			dt.PropertyStatusEnable,
		},
	})

	return nil
}