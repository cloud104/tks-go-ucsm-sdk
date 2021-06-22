package mo

import (
	"encoding/xml"
	"net"
)

// Any represents any valid managed object.
type Any interface{}

// TopSystem provides general information about the system, such as the
// name, IP address and current time.
type TopSystem struct {
	XMLName          xml.Name           `xml:"topSystem"`
	Address          net.IP             `xml:"address,attr,omitempty"`
	CurrentTime      string             `xml:"currentTime,attr,omitempty"`
	Description      string             `xml:"descr,attr,omitempty"`
	Dn               string             `xml:"dn,attr,omitempty"`
	Ipv6Addr         string             `xml:"ipv6Addr,attr,omitempty"`
	Mode             string             `xml:"mode,attr,omitempty"`
	Name             string             `xml:"name,attr,omitempty"`
	Owner            string             `xml:"owner,attr,omitempty"`
	Site             string             `xml:"site,attr,omitempty"`
	SystemUptime     string             `xml:"systemUpTime,attr,omitempty"`
	VersionEp        VersionEp          `xml:"versionEp"`
	CommServiceEp    CommServiceEp      `xml:"commSvcEp"`
	EquipmentChassis []EquipmentChassis `xml:"equipmentChassis"`
	ComputeRackUnits []ComputeRackUnit  `xml:"computeRackUnit"`
}

// CommServiceEp contains configuration for various services.
type CommServiceEp struct {
	FiniteStateMachineTask
	XMLName             xml.Name `xml:"commSvcEp"`
	ConfigState         string   `xml:"configState,attr,omitempty"`
	ConfigStatusMessage string   `xml:"configStatusMessage,attr,omitempty"`
	Description         string   `xml:"descr,attr,omitempty"`
	Dn                  string   `xml:"dn,attr,omitempty"`
	IntId               string   `xml:"intId,attr,omitempty"`
	Name                string   `xml:"name,attr,omitempty"`
	PolicyLevel         int      `xml:"policyLevel,attr,omitempty"`
	PolicyOwner         string   `xml:"policyOwner,attr,omitempty"`
	CommDns             CommDns  `xml:"commDns"`
}

// CommDns contains the DNS settings of the UCS system.
type CommDns struct {
	XMLName         xml.Name          `xml:"commDns"`
	AdminState      string            `xml:"adminState,attr,omitempty"`
	Description     string            `xml:"descr,attr,omitempty"`
	Dn              string            `xml:"dn,attr,omitempty"`
	Domain          string            `xml:"domain,attr,omitempty"`
	IntId           string            `xml:"intId,attr,omitempty"`
	Name            string            `xml:"name,attr,omitempty"`
	OperationalPort int               `xml:"operPort,attr,omitempty"`
	PolicyLevel     int               `xml:"policyLevel,attr,omitempty"`
	PolicyOwner     string            `xml:"policyOwner,attr,omitempty"`
	Port            int               `xml:"port,attr,omitempty"`
	Proto           string            `xml:"proto,attr,omitempty"`
	Providers       []CommDnsProvider `xml:"commDnsProvider"`
}

// CommDnsProvider represents a DNS service provider.
type CommDnsProvider struct {
	XMLName     xml.Name `xml:"commDnsProvider"`
	AdminState  string   `xml:"adminState,attr,omitempty"`
	Description string   `xml:"descr,attr,omitempty"`
	Dn          string   `xml:"dn,attr,omitempty"`
	Hostname    string   `xml:"hostname,attr,omitempty"`
	Name        string   `xml:"name,attr,omitempty"`
}

// VersionEp contains version information.
type VersionEp struct {
	XMLName     xml.Name           `xml:"versionEp"`
	ChildAction string             `xml:"childAction,attr,omitempty"`
	Dn          string             `xml:"dn,attr,omitempty"`
	Application VersionApplication `xml:"versionApplication,omitempty"`
}

// VersionApplication contains the application version.
type VersionApplication struct {
	XMLName     xml.Name `xml:"versionApplication"`
	ChildAction string   `xml:"childAction,attr,omitempty"`
	Detail      string   `xml:"detail,attr,omitempty"`
	Dn          string   `xml:"dn,attr,omitempty"`
	Rn          string   `xml:"rn,attr,omitempty"`
	Time        string   `xml:"time,attr,omitempty"`
	Version     string   `xml:"version,attr,omitempty"`
}

// EquipmentChassis represents a physical unit that can accomodate multiple blade servers.
// For example, the Cisco UCS 5108 Blade Server Chassis is six rack units (6RU) high,
// can mount in an industry-standard 19-inch rack and uses front-to-back cooling.
type EquipmentChassis struct {
	FiniteStateMachineTask
	XMLName                    xml.Name             `xml:"equipmentChassis"`
	AckProgressIndicator       string               `xml:"ackProgressIndicator,attr,omitempty"`
	AdminState                 string               `xml:"adminState,attr,omitempty"`
	AssignedToDn               string               `xml:"assignedToDn,attr,omitempty"`
	Association                string               `xml:"association,attr,omitempty"`
	Availability               string               `xml:"availability,attr,omitempty"`
	ConfigState                string               `xml:"configState,attr,omitempty"`
	ConnPath                   string               `xml:"connPath,attr,omitempty"`
	ConnStatus                 string               `xml:"connStatus,attr,omitempty"`
	Discovery                  string               `xml:"discovery,attr,omitempty"`
	DiscoveryStatus            string               `xml:"discoveryStatus,attr,omitempty"`
	Dn                         string               `xml:"dn,attr,omitempty"`
	FabricEpDn                 string               `xml:"fabricEpDn,attr,omitempty"`
	FltAggr                    int                  `xml:"fltAggr,attr,omitempty"`
	Id                         string               `xml:"id,attr,omitempty"`
	LcTimestamp                string               `xml:"lcTs,attr,omitempty"`
	LicGP                      int                  `xml:"licGP,attr,omitempty"`
	LicState                   string               `xml:"licState,attr,omitempty"`
	ManagingInstance           string               `xml:"managingInst,attr,omitempty"`
	ManufacturingTime          string               `xml:"mfgTime,attr,omitempty"`
	Model                      string               `xml:"model,attr,omitempty"`
	OperationalQualifier       string               `xml:"operQualifier,attr,omitempty"`
	OperationalQualifierReason string               `xml:"operQualifierReason,attr,omitempty"`
	OperationalState           string               `xml:"operState,attr,omitempty"`
	Operability                string               `xml:"operability,attr,omitempty"`
	PartNumber                 string               `xml:"partNumber,attr,omitempty"`
	Power                      string               `xml:"power,attr,omitempty"`
	Presence                   string               `xml:"presence,attr,omitempty"`
	Revision                   string               `xml:"revision,attr,omitempty"`
	SeepromOperationalState    string               `xml:"seepromOperState,attr,omitempty"`
	Serial                     string               `xml:"serial,attr,omitempty"`
	ServiceState               string               `xml:"serviceState,attr,omitempty"`
	Thermal                    string               `xml:"thermal,attr,omitempty"`
	ThermanlStateQualifier     string               `xml:"thermalStateQualifier,attr,omitempty"`
	UserLabel                  string               `xml:"usrLbl,attr,omitempty"`
	Vendor                     string               `xml:"vendor,attr,omitempty"`
	VersionHolder              string               `xml:"versionHolder,attr,omitempty"`
	Vid                        string               `xml:"vid,attr,omitempty"`
	ComputeBlades              []ComputeBlade       `xml:"computeBlade"`
	FanModules                 []EquipmentFanModule `xml:"equipmentFanModule"`
}

// ComputePhysical represents a physical specification of an abstract compute item.
// Serves as the base of physical compute nodes (e.g. blade, stand-alone computer or server).
type ComputePhysical struct {
	FiniteStateMachineTask
	AdminPower                          string               `xml:"adminPower,attr,omitempty"`
	AdminState                          string               `xml:"adminState,attr,omitempty"`
	AssignedToDn                        string               `xml:"assignedToDn,attr,omitempty"`
	Association                         string               `xml:"association,attr,omitempty"`
	Availability                        string               `xml:"availability,attr,omitempty"`
	AvailableMemory                     int                  `xml:"availableMemory,attr,omitempty"`
	ChassisId                           string               `xml:"chassisId,attr,omitempty"`
	CheckPoint                          string               `xml:"checkPoint,attr,omitempty"`
	ConnPath                            string               `xml:"connPath,attr,omitempty"`
	ConnStatus                          string               `xml:"connStatus,attr,omitempty"`
	Description                         string               `xml:"descr,attr,omitempty"`
	Discovery                           string               `xml:"discovery,attr,omitempty"`
	DiscoveryStatus                     string               `xml:"discoveryStatus,attr,omitempty"`
	Dn                                  string               `xml:"dn,attr,omitempty"`
	FltAggr                             int                  `xml:"fltAggr,attr,omitempty"`
	IntId                               string               `xml:"intId,attr,omitempty"`
	Lc                                  string               `xml:"lc,attr,omitempty"`
	LcTimestamp                         string               `xml:"lcTs,attr,omitempty"`
	LocalId                             string               `xml:"localId,attr,omitempty"`
	LowVoltageMemory                    string               `xml:"lowVoltageMemory,attr,omitempty"`
	ManagingInstance                    string               `xml:"managingInst,attr,omitempty"`
	MemorySpeed                         string               `xml:"memorySpeed,attr,omitempty"`
	ManufacturingTime                   string               `xml:"mfgTime,attr,omitempty"`
	Model                               string               `xml:"model,attr,omitempty"`
	Name                                string               `xml:"name,attr,omitempty"`
	NumOf40GAdaptorsWithOldFirmware     int                  `xml:"numOf40GAdaptorsWithOldFw,attr,omitempty"`
	NumOf40GAdaptorsWithUnknownFirmware int                  `xml:"numOf40GAdaptorsWithUnknownFw,attr,omitempty"`
	NumOfAdaptors                       int                  `xml:"numOfAdaptors,attr,omitempty"`
	NumOfCores                          int                  `xml:"numOfCores,attr,omitempty"`
	NumOfCoresEnabled                   int                  `xml:"numOfCoresEnabled,attr,omitempty"`
	NumOfCpus                           int                  `xml:"numOfCpus,attr,omitempty"`
	NumOfEthHostInterfaces              int                  `xml:"numOfEthHostIfs,attr,omitempty"`
	NumOfFcHostInterfaces               int                  `xml:"numOfFcHostIfs,attr,omitempty"`
	NumOfThreads                        int                  `xml:"numOfThreads,attr,omitempty"`
	OperationalPower                    string               `xml:"operPower,attr,omitempty"`
	OperationalPowerTransitionSource    string               `xml:"operPwrTransSrc,attr,omitempty"`
	OperationalQualifier                string               `xml:"operQualifier,attr,omitempty"`
	OperationalState                    string               `xml:"operState,attr,omitempty"`
	Operability                         string               `xml:"operability,attr,omitempty"`
	OriginalUuid                        string               `xml:"originalUuid,attr,omitempty"`
	PartNumber                          string               `xml:"partNumber,attr,omitempty"`
	PolicyLevel                         int                  `xml:"policyLevel,attr,omitempty"`
	PolicyOwner                         string               `xml:"policyOwner,attr,omitempty"`
	Presence                            string               `xml:"presence,attr,omitempty"`
	Revision                            string               `xml:"revision,attr,omitempty"`
	ScaledMode                          string               `xml:"scaledMode,attr,omitempty"`
	Serial                              string               `xml:"serial,attr,omitempty"`
	ServerId                            string               `xml:"serverId,attr,omitempty"`
	SlotId                              int                  `xml:"slotId,attr,omitempty"`
	TotalMemory                         int                  `xml:"totalMemory,attr,omitempty"`
	UserLabel                           string               `xml:"usrLbl,attr,omitempty"`
	Uuid                                string               `xml:"uuid,attr,omitempty"`
	Vendor                              string               `xml:"vendor,attr,omitempty"`
	Vid                                 string               `xml:"vid,attr,omitempty"`
	ComputeBoard                        ComputeBoard         `xml:"computeBoard"`
	AdaptorUnits                        []AdaptorUnit        `xml:"adaptorUnit"`
	ManagementController                ManagementController `xml:"mgmtController"`
	FirmwareStatus                      FirmwareStatus       `xml:"firmwareStatus"`
	BiosUnit                            BiosUnit             `xml:"biosUnit"`
}

// ComputeBlade represents a physical compute blade.
// Physical compute item in blade form factor.
type ComputeBlade struct {
	XMLName xml.Name `xml:"computeBlade"`
	ComputePhysical
}

// ComputeRackUnit represents a physical compute RackUnit.
// Physical compute item representing a Rack mountable unit.
type ComputeRackUnit struct {
	XMLName xml.Name `xml:"computeRackUnit"`
	ComputePhysical
}

// ComputeServerUnit represents a server instance on a cartridge.
type ComputeServerUnit struct {
	XMLName xml.Name `xml:"computeServerUnit"`
	ComputePhysical
}

// ComputeItem type represents a container for all compute items,
// which include blades, rack units and stand-alone servers.
type ComputeItem struct {
	XMLName     xml.Name
	Blades      []ComputeBlade      `xml:"computeBlade"`
	RackUnits   []ComputeRackUnit   `xml:"computeRackUnit"`
	ServerUnits []ComputeServerUnit `xml:"computeServerUnit"`
}

// ComputeBoard represents a motherboard contained by physical compute item.
type ComputeBoard struct {
	XMLName                    xml.Name          `xml:"computeBoard"`
	CmosVoltage                string            `xml:"cmosVoltage,attr,omitempty"`
	CpuTypeDescription         string            `xml:"cpuTypeDescription,attr,omitempty"`
	Dn                         string            `xml:"dn,attr,omitempty"`
	FaultQualifier             string            `xml:"faultQualifier,attr,omitempty"`
	Id                         int               `xml:"id,attr,omitempty"`
	LocationDn                 string            `xml:"locationDn,attr,omitempty"`
	Model                      string            `xml:"model,attr,omitempty"`
	OperationalPower           string            `xml:"operPower,attr,omitempty"`
	OperationalQualifierReason string            `xml:"operQualifierReason,attr,omitempty"`
	OperationalState           string            `xml:"operState,attr,omitempty"`
	Operability                string            `xml:"operability,attr,omitempty"`
	Perf                       string            `xml:"perf,attr,omitempty"`
	Power                      string            `xml:"power,attr,omitempty"`
	PowerUsage                 string            `xml:"powerUsage,attr,omitempty"`
	Presence                   string            `xml:"presence,attr,omitempty"`
	Revision                   string            `xml:"revision,attr,omitempty"`
	Serial                     string            `xml:"serial,attr,omitempty"`
	Thermal                    string            `xml:"thermal,attr,omitempty"`
	Vendor                     string            `xml:"vendor,attr,omitempty"`
	Voltage                    string            `xml:"voltage,attr,omitempty"`
	MemoryArray                MemoryArray       `xml:"memoryArray"`
	ProcessorUnits             []ProcessorUnit   `xml:"processorUnit"`
	StorageController          StorageController `xml:"storageController"`
}

// MemoryArray represents an array of memory units.
type MemoryArray struct {
	XMLName                    xml.Name     `xml:"memoryArray"`
	ChildAction                string       `xml:"childAction,attr,omitempty"`
	CpuId                      int          `xml:"cpuId,attr,omitempty"`
	CurrentCapacity            int          `xml:"currCapacity,attr,omitempty"`
	ErrorCorrectionn           string       `xml:"errorCorrection,attr,omitempty"`
	Id                         int          `xml:"id,attr,omitempty"`
	LocationDn                 string       `xml:"locationDn,attr,omitempty"`
	MaxCapacity                int          `xml:"maxCapacity,attr,omitempty"`
	MaxDevices                 int          `xml:"maxDevices,attr,omitempty"`
	Model                      string       `xml:"model,attr,omitempty"`
	OperationalQualifierReason string       `xml:"operQualifierReason,attr,omitempty"`
	OperationalState           string       `xml:"operState,attr,omitempty"`
	Operability                string       `xml:"operability,attr,omitempty"`
	Perf                       string       `xml:"perf,attr,omitempty"`
	Populated                  int          `xml:"populated,attr,omitempty"`
	Power                      string       `xml:"power,attr,omitempty"`
	Presence                   string       `xml:"presence,attr,omitempty"`
	Revision                   string       `xml:"revision,attr,omitempty"`
	Rn                         string       `xml:"rn,attr,omitempty"`
	Serial                     string       `xml:"serial,attr,omitempty"`
	Thermal                    string       `xml:"thermal,attr,omitempty"`
	Vendor                     string       `xml:"vendor,attr,omitempty"`
	Voltage                    string       `xml:"voltage,attr,omitempty"`
	Units                      []MemoryUnit `xml:"memoryUnit"`
}

// MemoryUnit represents a single memory unit in a memory array.
type MemoryUnit struct {
	XMLName                    xml.Name `xml:"memoryUnit"`
	AdminState                 string   `xml:"adminState,attr,omitempty"`
	Array                      int      `xml:"array,attr,omitempty"`
	Bank                       int      `xml:"bank,attr,omitempty"`
	Capacity                   string   `xml:"capacity,attr,omitempty"`
	ChildAction                string   `xml:"childAction,attr,omitempty"`
	Clock                      string   `xml:"clock,attr,omitempty"`
	FormFactor                 string   `xml:"formFactor,attr,omitempty"`
	Id                         int      `xml:"id,attr,omitempty"`
	Latency                    string   `xml:"latency,attr,omitempty"`
	Location                   string   `xml:"location,attr,omitempty"`
	LocationDn                 string   `xml:"locationDn,attr,omitempty"`
	Model                      string   `xml:"model,attr,omitempty"`
	OperationalQualifier       string   `xml:"operQualifier,attr,omitempty"`
	OperationalQualifierReason string   `xml:"operQualifierReason,attr,omitempty"`
	OperationalState           string   `xml:"operState,attr,omitempty"`
	Operability                string   `xml:"operability,attr,omitempty"`
	Perf                       string   `xml:"perf,attr,omitempty"`
	Power                      string   `xml:"power,attr,omitempty"`
	Presence                   string   `xml:"presence,attr,omitempty"`
	Revision                   string   `xml:"revision,attr,omitempty"`
	Rn                         string   `xml:"rn,attr,omitempty"`
	Serial                     string   `xml:"serial,attr,omitempty"`
	Set                        int      `xml:"set,attr,omitempty"`
	Speed                      string   `xml:"speed,attr,omitempty"`
	Thermal                    string   `xml:"thermal,attr,omitempty"`
	Type                       string   `xml:"type,attr,omitempty"`
	Vendor                     string   `xml:"vendor,attr,omitempty"`
	Visibility                 string   `xml:"visibility,attr,omitempty"`
	Voltage                    string   `xml:"voltage,attr,omitempty"`
	Width                      string   `xml:"width,attr,omitempty"`
}

// ProcessorUnit represents a single processor unit.
type ProcessorUnit struct {
	XMLName                    xml.Name `xml:"processorUnit"`
	Arch                       string   `xml:"arch,attr,omitempty"`
	ChildAction                string   `xml:"childAction,attr,omitempty"`
	Cores                      int      `xml:"cores,attr,omitempty"`
	CoresEnabled               int      `xml:"coresEnabled,attr,omitempty"`
	Id                         int      `xml:"id,attr,omitempty"`
	LocationDn                 string   `xml:"locationDn,attr,omitempty"`
	Model                      string   `xml:"model,attr,omitempty"`
	OperationalQualifierReason string   `xml:"operQualifierReason,attr,omitempty"`
	OperationalState           string   `xml:"operState,attr,omitempty"`
	Operability                string   `xml:"operability,attr,omitempty"`
	Perf                       string   `xml:"perf,attr,omitempty"`
	Power                      string   `xml:"power,attr,omitempty"`
	Presence                   string   `xml:"presence,attr,omitempty"`
	Revision                   string   `xml:"revision,attr,omitempty"`
	Rn                         string   `xml:"rn,attr,omitempty"`
	Serial                     string   `xml:"serial,attr,omitempty"`
	SocketDesignation          string   `xml:"socketDesignation,attr,omitempty"`
	Speed                      string   `xml:"speed,attr,omitempty"`
	Stepping                   int      `xml:"stepping,attr,omitempty"`
	Thermal                    string   `xml:"thermal,attr,omitempty"`
	Threads                    int      `xml:"threads,attr,omitempty"`
	Vendor                     string   `xml:"vendor,attr,omitempty"`
	Visibility                 string   `xml:"visibility,attr,omitempty"`
	Voltage                    string   `xml:"voltage,attr,omitempty"`
}

// AdaptorUnit is a managed object representing a network adaptor unit such as a
// card that has NIC and/or HBA, SCSI functionality.
type AdaptorUnit struct {
	XMLName                       xml.Name                       `xml:"adaptorUnit"`
	AdminPowerState               string                         `xml:"adminPowerState,attr,omitempty"`
	BaseMac                       string                         `xml:"baseMac,attr,omitempty"`
	BladeId                       int                            `xml:"bladeId,attr,omitempty"`
	CartridgeId                   int                            `xml:"cartridgeId,attr,omitempty"`
	ChassisId                     string                         `xml:"chassisId,attr,omitempty"`
	ChildAction                   string                         `xml:"childAction,attr,omitempty"`
	ConnPath                      string                         `xml:"connPath,attr,omitempty"`
	ConnStatus                    string                         `xml:"connStatus,attr,omitempty"`
	DiscoveryStatus               string                         `xml:"discoveryStatus,attr,omitempty"`
	FltAggr                       int                            `xml:"fltAggr,attr,omitempty"`
	Id                            int                            `xml:"id,attr,omitempty"`
	Integrated                    string                         `xml:"integrated,attr,omitempty"`
	LocationDn                    string                         `xml:"locationDn,attr,omitempty"`
	ManagingInstance              string                         `xml:"managingInst,attr,omitempty"`
	ManufacturingTime             string                         `xml:"mfgTime,attr,omitempty"`
	Model                         string                         `xml:"model,attr,omitempty"`
	OperationalQualifierReason    string                         `xml:"operQualifierReason,attr,omitempty"`
	OperationalState              string                         `xml:"operState,attr,omitempty"`
	Operability                   string                         `xml:"operability,attr,omitempty"`
	PartNumber                    string                         `xml:"partNumber,attr,omitempty"`
	PciAddress                    string                         `xml:"pciAddr,attr,omitempty"`
	PciSlot                       string                         `xml:"pciSlot,attr,omitempty"`
	Perf                          string                         `xml:"perf,attr,omitempty"`
	Power                         string                         `xml:"power,attr,omitempty"`
	Presence                      string                         `xml:"presence,attr,omitempty"`
	Reachability                  string                         `xml:"reachability,attr,omitempty"`
	Revision                      string                         `xml:"revision,attr,omitempty"`
	Rn                            string                         `xml:"rn,attr,omitempty"`
	Serial                        string                         `xml:"serial,attr,omitempty"`
	Thermal                       string                         `xml:"thermal,attr,omitempty"`
	Vendor                        string                         `xml:"vendor,attr,omitempty"`
	Vid                           string                         `xml:"vid,attr,omitempty"`
	Voltage                       string                         `xml:"voltage,attr,omitempty"`
	AdaptorHostEthernetInterfaces []AdaptorHostEthernetInterface `xml:"adaptorHostEthIf"`
	ManagementController          ManagementController           `xml:"mgmtController"`
}

// AdaptorHostEthernetInterface is a representation of a host-facing Ethernet interface
// on a server adaptor. A server adaptor has network facing interfaces (NIF), which
// provide network connectivity to the network (through the IO Module for UCS blades) and
// server facing interfaces (SIF), which are visible by the Operating System.
type AdaptorHostEthernetInterface struct {
	FiniteStateMachineTask
	XMLName                    xml.Name              `xml:"adaptorHostEthIf"`
	AdminState                 string                `xml:"adminState,attr,omitempty"`
	BootDev                    string                `xml:"bootDev,attr,omitempty"`
	CdnName                    string                `xml:"cdnName,attr,omitempty"`
	ChassisId                  string                `xml:"chassisId,attr,omitempty"`
	ChildAction                string                `xml:"childAction,attr,omitempty"`
	Discovery                  string                `xml:"discovery,attr,omitempty"`
	EpDn                       string                `xml:"epDn,attr,omitempty"`
	FltAggr                    int                   `xml:"fltAggr,attr,omitempty"`
	HostPort                   string                `xml:"hostPort,attr,omitempty"`
	Id                         int                   `xml:"id,attr,omitempty"`
	InterfaceRole              string                `xml:"ifRole,attr,omitempty"`
	InterfaceType              string                `xml:"ifType,attr,omitempty"`
	Lc                         string                `xml:"lc,attr,omitempty"`
	LinkState                  string                `xml:"linkState,attr,omitempty"`
	Locale                     string                `xml:"locale,attr,omitempty"`
	Mac                        string                `xml:"mac,attr,omitempty"`
	Model                      string                `xml:"model,attr,omitempty"`
	Mtu                        int                   `xml:"mtu,attr,omitempty"`
	Name                       string                `xml:"name,attr,omitempty"`
	OperationalQualifierReason string                `xml:"operQualifierReason,attr,omitempty"`
	OperationalState           string                `xml:"operState,attr,omitempty"`
	Operability                string                `xml:"operability,attr,omitempty"`
	Order                      int                   `xml:"order,attr,omitempty"`
	OriginalMac                string                `xml:"originaMac,attr,omitempty"`
	PciAddress                 string                `xml:"pciAddr,attr,omitempty"`
	PciFunc                    int                   `xml:"pciFunc,attr,omitempty"`
	PciSlot                    int                   `xml:"pciSlot,attr,omitempty"`
	PeerChassisId              string                `xml:"peerChassisId,attr,omitempty"`
	PeerDn                     string                `xml:"peerDn,attr,omitempty"`
	PeerPortId                 int                   `xml:"peerPortId,attr,omitempty"`
	PeerSlotId                 int                   `xml:"peerSlotId,attr,omitempty"`
	Perf                       string                `xml:"perf,attr,omitempty"`
	PfDn                       string                `xml:"pfDn,attr,omitempty"`
	PortId                     int                   `xml:"portId,attr,omitempty"`
	Power                      string                `xml:"power,attr,omitempty"`
	Presence                   string                `xml:"presence,attr,omitempty"`
	Purpose                    string                `xml:"purpose,attr,omitempty"`
	Revision                   string                `xml:"revision,attr,omitempty"`
	Rn                         string                `xml:"rn,attr,omitempty"`
	Serial                     string                `xml:"serial,attr,omitempty"`
	Side                       string                `xml:"side,attr,omitempty"`
	SlotId                     int                   `xml:"slotId,attr,omitempty"`
	SwitchId                   string                `xml:"switchId,attr,omitempty"`
	Thermal                    string                `xml:"thermal,attr,omitempty"`
	Transport                  string                `xml:"transport,attr,omitempty"`
	Type                       string                `xml:"type,attr,omitempty"`
	Vendor                     string                `xml:"vendor,attr,omitempty"`
	VirtualizationPreference   string                `xml:"virtualizationPreference,attr,omitempty"`
	VnicDn                     string                `xml:"vnicDn,attr,omitempty"`
	Voltage                    string                `xml:"voltage,attr,omitempty"`
	ManagementInterfaces       []ManagementInterface `xml:"mgmtIf"`
}

// ManagementInterface encapsulates the configuration of a CIMC management interface.
type ManagementInterface struct {
	FiniteStateMachineTask
	XMLName        xml.Name `xml:"mgmtIf"`
	Access         string   `xml:"access,attr,omitempty"`
	AdminState     string   `xml:"adminState,attr,omitempty"`
	AggrPortId     int      `xml:"aggrPortId,attr,omitempty"`
	ChassisId      string   `xml:"chassisId,attr,omitempty"`
	ChildAction    string   `xml:"childAction,attr,omitempty"`
	Discovery      string   `xml:"discovery,attr,omitempty"`
	EpDn           string   `xml:"epDn,attr,omitempty"`
	ExtBroadcast   net.IP   `xml:"extBroadcast,attr,omitempty"`
	ExtGateway     net.IP   `xml:"extGw,attr,omitempty"`
	ExtIp          net.IP   `xml:"extIp,attr,omitempty"`
	ExtNetmask     net.IP   `xml:"extMask,attr,omitempty"`
	Id             int      `xml:"id,attr,omitempty"`
	InterfaceRole  string   `xml:"ifRole,attr,omitempty"`
	InterfaceType  string   `xml:"ifType,attr,omitempty"`
	InstanceId     int      `xml:"instanceId,attr,omitempty"`
	Locale         string   `xml:"locale,attr,omitempty"`
	Ip             net.IP   `xml:"ip,attr,omitempty"`
	Mac            string   `xml:"mac,attr,omitempty"`
	Netmask        net.IP   `xml:"mask,attr,omitempty"`
	Name           string   `xml:"name,attr,omitempty"`
	PeerAggrPortId int      `xml:"peerAggrPortId,attr,omitempty"`
	PeerChassisId  string   `xml:"peerChassisId,attr,omitempty"`
	PeerDn         string   `xml:"peerDn,attr,omitempty"`
	PeerPortId     int      `xml:"peerPortId,attr,omitempty"`
	PeerSlotId     int      `xml:"peerSlotId,attr,omitempty"`
	PortId         int      `xml:"portId,attr,omitempty"`
	Rn             string   `xml:"rn,attr,omitempty"`
	SlotId         int      `xml:"slotId,attr,omitempty"`
	StateQual      string   `xml:"stateQual,attr,omitempty"`
	Subject        string   `xml:"subject,attr,omitempty"`
	SwitchId       string   `xml:"switchId,attr,omitempty"`
	Transport      string   `xml:"transport,attr,omitempty"`
	Type           string   `xml:"type,attr,omitempty"`
	Vnet           int      `xml:"vnet,attr,omitempty"`
}

// EquipmentFanModule represents an inventoried Fan module.
// This object is created implicitly when a Fan module is detected during equipment discovery.
type EquipmentFanModule struct {
	XMLName              xml.Name       `xml:"equipmentFanModule"`
	ChildAction          string         `xml:"childAction,attr,omitempty"`
	Dn                   string         `xml:"dn,attr,omitempty"`
	FltAggr              int            `xml:"fltAggr,attr,omitempty"`
	Id                   int            `xml:"id,attr,omitempty"`
	ManufacturingTime    string         `xml:"mfgTime,attr,omitempty"`
	Model                string         `xml:"model,attr,omitempty"`
	OperationalQualifier string         `xml:"operQualifier,attr,omitempty"`
	OperationalState     string         `xml:"operState,attr,omitempty"`
	Operability          string         `xml:"operability,attr,omitempty"`
	PartNumber           string         `xml:"partNumber,attr,omitempty"`
	Perf                 string         `xml:"perf,attr,omitempty"`
	Power                string         `xml:"power,attr,omitempty"`
	Presence             string         `xml:"presence,attr,omitempty"`
	Revision             string         `xml:"revision,attr,omitempty"`
	Serial               string         `xml:"serial,attr,omitempty"`
	Thermal              string         `xml:"thermal,attr,omitempty"`
	Tray                 int            `xml:"tray,attr,omitempty"`
	Vendor               string         `xml:"vendor,attr,omitempty"`
	Vid                  string         `xml:"vid,attr,omitempty"`
	Voltage              string         `xml:"voltage,attr,omitempty"`
	Fans                 []EquipmentFan `xml:"equipmentFan"`
}

// EquipmentFan represents a fan in a Fan module.
type EquipmentFan struct {
	XMLName                        xml.Name `xml:"equipmentFan"`
	ChildAction                    string   `xml:"childAction,attr,omitempty"`
	FanSpeedPolicyAdminState       string   `xml:"fanSpeedPolicyAdminState,attr,omitempty"`
	FanSpeedPolicyOperationalState string   `xml:"fanSpeedPolicyOperState,attr,omitempty"`
	FltAggr                        int      `xml:"fltAggr,attr,omitempty"`
	Id                             int      `xml:"id,attr,omitempty"`
	InternalType                   string   `xml:"intType,attr,omitempty"`
	Model                          string   `xml:"model,attr,omitempty"`
	Module                         int      `xml:"module,attr,omitempty"`
	OperationalQualifierReason     string   `xml:"operQualifierReason,attr,omitempty"`
	OperationalState               string   `xml:"operState,attr,omitempty"`
	Operability                    string   `xml:"operability,attr,omitempty"`
	Perf                           string   `xml:"perf,attr,omitempty"`
	Power                          string   `xml:"power,attr,omitempty"`
	Presence                       string   `xml:"presence,attr,omitempty"`
	Revision                       string   `xml:"revision,attr,omitempty"`
	Rn                             string   `xml:"rn,attr,omitempty"`
	Serial                         string   `xml:"serial,attr,omitempty"`
	Thermal                        string   `xml:"thermal,attr,omitempty"`
	Tray                           int      `xml:"tray,attr,omitempty"`
	Vendor                         string   `xml:"vendor,attr,omitempty"`
	Voltage                        string   `xml:"voltage,attr,omitempty"`
}

// FiniteStateMachineTask represents the result of an FSM task.
type FiniteStateMachineTask struct {
	FsmDescription             string `xml:"fsmDescr,attr,omitempty"`
	FsmFlags                   string `xml:"fsmFlags,attr,omitempty"`
	FsmPrev                    string `xml:"fsmPrev,attr,omitempty"`
	FsmProgress                int    `xml:"fsmProgr,attr,omitempty"`
	FsmRemoteInvErrCode        string `xml:"fsmRmtInvErrCode,attr,omitempty"`
	FsmRemoteInvErrDescription string `xml:"fsmRmtInvErrDescr,attr,omitempty"`
	FsmRemoteInvResult         string `xml:"fsmRmtInvRslt,attr,omitempty"`
	FsmStageDescription        string `xml:"fsmStageDescr,attr,omitempty"`
	FsmTimestamp               string `xml:"fsmStamp,attr,omitempty"`
	FsmStatus                  string `xml:"fsmStatus,attr,omitempty"`
	FsmTry                     int    `xml:"fsmTry,attr,omitempty"`
}

// FirmwareRunning is a representation of the primary firmware image (currently running).
type FirmwareRunning struct {
	XMLName        xml.Name `xml:"firmwareRunning"`
	Deployment     string   `xml:"deployment,attr,omitempty"`
	Dn             string   `xml:"dn,attr,omitempty"`
	InvTag         string   `xml:"invTag,attr,omitempty"`
	PackageVersion string   `xml:"packageVersion,attr,omitempty"`
	Type           string   `xml:"type,attr,omitempty"`
	Version        string   `xml:"version,attr,omitempty"`
}

// FirmwareUpdatable is a representation of a backup firmware image for the chassis components
// that supports backup image (CMC, BMC, BIOS, Adaptor, etc).
type FirmwareUpdatable struct {
	XMLName                   xml.Name `xml:"firmwareUpdatable"`
	AdminState                string   `xml:"adminState,attr,omitempty"`
	ChildAction               string   `xml:"childAction,attr,omitempty"`
	Deployment                string   `xml:"deployment,attr,omitempty"`
	OperationalState          string   `xml:"operState,attr,omitempty"`
	OperationalStateQualifier string   `xml:"operStateQual,attr,omitempty"`
	PreviousVersion           string   `xml:"prevVersion,attr,omitempty"`
	Rn                        string   `xml:"rn,attr,omitempty"`
	Version                   string   `xml:"version,attr,omitempty"`
}

// FirmwareStatus represents a registered client for monitoring firmware update progress.
type FirmwareStatus struct {
	XMLName          xml.Name `xml:"firmwareStatus"`
	ChildAction      string   `xml:"childAction,attr,omitempty"`
	CimcVersion      string   `xml:"cimcVersion,attr,omitempty"`
	FirmwareState    string   `xml:"firmwareState,attr,omitempty"`
	OperationalState string   `xml:"operState,attr,omitempty"`
	PackageVersion   string   `xml:"packageVersion,attr,omitempty"`
	PldVersion       string   `xml:"pldVersion,attr,omitempty"`
	Rn               string   `xml:"rn,attr,omitempty"`
}

// StorageController represents a storage controller.
type StorageController struct {
	XMLName                    xml.Name             `xml:"storageController"`
	AdminAction                string               `xml:"adminAction,attr,omitempty"`
	AdminActionTrigger         string               `xml:"adminActionTrigger,attr,omitempty"`
	ConfigState                string               `xml:"configState,attr,omitempty"`
	ControllerOperations       string               `xml:"controllerOps,attr,omitempty"`
	ControllerStatus           string               `xml:"controllerStatus,attr,omitempty"`
	DefaultStripSize           string               `xml:"defaultStripSize,attr,omitempty"`
	DeviceRaidSupport          string               `xml:"deviceRaidSupport,attr,omitempty"`
	DiskOperations             string               `xml:"diskOps,attr,omitempty"`
	Dn                         string               `xml:"dn,attr,omitempty"`
	FaultMonitoring            string               `xml:"faultMonitoring,attr,omitempty"`
	HardwareRevision           string               `xml:"hwRevision,attr,omitempty"`
	Id                         int                  `xml:"id,attr,omitempty"`
	IdCount                    string               `xml:"idCount,attr,omitempty"`
	Lc                         string               `xml:"lc,attr,omitempty"`
	LocationDn                 string               `xml:"locationDn,attr,omitempty"`
	Mode                       string               `xml:"mode,attr,omitempty"`
	Model                      string               `xml:"model,attr,omitempty"`
	OnBoardMemoryPresent       string               `xml:"onBoardMemoryPresent,attr,omitempty"`
	OnBoardMemorySize          string               `xml:"onBoardMemorySize,attr,omitempty"`
	OobControllerId            string               `xml:"oobControllerId,attr,omitempty"`
	OobInterfaceSupported      string               `xml:"oobInterfaceSupported,attr,omitempty"`
	OperationalQualifierReason string               `xml:"operQualifierReason,attr,omitempty"`
	OperationalState           string               `xml:"operState,attr,omitempty"`
	Operability                string               `xml:"operability,attr,omitempty"`
	OpromBootStatus            string               `xml:"opromBootStatus,attr,omitempty"`
	PartNumber                 string               `xml:"partNumber,attr,omitempty"`
	PciAddress                 string               `xml:"pciAddr,attr,omitempty"`
	PciSlot                    string               `xml:"pciSlot,attr,omitempty"`
	PciSlotRawName             string               `xml:"pciSlotRawName,attr,omitempty"`
	Perf                       string               `xml:"perf,attr,omitempty"`
	PinnedCacheStatus          string               `xml:"pinnedCacheStatus,attr,omitempty"`
	Power                      string               `xml:"power,attr,omitempty"`
	Presence                   string               `xml:"presence,attr,omitempty"`
	RaidBatteryOperations      string               `xml:"raidBatteryOps,attr,omitempty"`
	RaidSupport                string               `xml:"raidSupport,attr,omitempty"`
	RebuildRate                string               `xml:"rebuildRate,attr,omitempty"`
	Revision                   string               `xml:"revision,attr,omitempty"`
	Serial                     string               `xml:"serial,attr,omitempty"`
	SubOemId                   string               `xml:"subOemId,attr,omitempty"`
	SupportedStripSizes        string               `xml:"supportedStripSizes,attr,omitempty"`
	Thermal                    string               `xml:"thermal,attr,omitempty"`
	Type                       string               `xml:"type,attr,omitempty"`
	VariantType                string               `xml:"variantType,attr,omitempty"`
	Vendor                     string               `xml:"vendor,attr,omitempty"`
	Vid                        string               `xml:"vid,attr,omitempty"`
	VirtualDriveOperations     string               `xml:"virtualDriveops,attr,omitempty"`
	Voltage                    string               `xml:"voltage,attr,omitempty"`
	ManagementController       ManagementController `xml:"mgmtController"`
	FirmwareRunning            []FirmwareRunning    `xml:"firmwareRunning"`
}

// BiosUnit represents a BIOS unit.
type BiosUnit struct {
	XMLName           xml.Name          `xml:"biosUnit"`
	ChildAction       string            `xml:"childAction,attr,omitempty"`
	InitSequence      string            `xml:"initSeq,attr,omitempty"`
	InitTimestamp     string            `xml:"initTs,attr,omitempty"`
	Model             string            `xml:"model,attr,omitempty"`
	Revision          string            `xml:"revision,attr,omitempty"`
	Rn                string            `xml:"rn,attr,omitempty"`
	Serial            string            `xml:"serial,attr,omitempty"`
	Vendor            string            `xml:"vendor,attr,omitempty"`
	FirmwareRunning   FirmwareRunning   `xml:"firmwareRunning"`
	FirmwareUpdatable FirmwareUpdatable `xml:"firmwareUpdatable"`
}

// ManagementController represents an instance of a management controller.
type ManagementController struct {
	FiniteStateMachineTask
	XMLName                          xml.Name              `xml:"mgmtController"`
	DesiredMaintenanceMode           string                `xml:"desiredMaintenanceMode,attr,omitempty"`
	DimmBlackListingOperationalState string                `xml:"dimmBlacklistingOperState,attr,omitempty"`
	DiskZoningState                  string                `xml:"diskZoningState,attr,omitempty"`
	Dn                               string                `xml:"dn,attr,omitempty"`
	Guid                             string                `xml:"guid,attr,omitempty"`
	Id                               string                `xml:"id,attr,omitempty"`
	LastRebootReason                 string                `xml:"lastRebootReason,attr,omitempty"`
	Model                            string                `xml:"model,attr,omitempty"`
	OperationalConnection            string                `xml:"operConn,attr,omitempty"`
	PowerFanSpeedPolicySupported     string                `xml:"powerFanSpeedPolicySupported,attr,omitempty"`
	Revision                         string                `xml:"revision,attr,omitempty"`
	Serial                           string                `xml:"serial,attr,omitempty"`
	StorageOobConfigSupported        string                `xml:"storageOobConfigSupported,attr,omitempty"`
	StorageOobInterfaceSupported     string                `xml:"storageOobInterfaceSupported,attr,omitempty"`
	StorageSubsystemState            string                `xml:"storageSubsystemState,attr,omitempty"`
	Subject                          string                `xml:"subject,attr,omitempty"`
	SupportedCapability              string                `xml:"supportedCapability,attr,omitempty"`
	Vendor                           string                `xml:"vendor,attr,omitempty"`
	FirmwareRunning                  []FirmwareRunning     `xml:"firmwareRunning"`
	FirmwareUpdatable                FirmwareUpdatable     `xml:"firmwareUpdatable"`
	ManagementInterfaces             []ManagementInterface `xml:"mgmtIf"`
}

// StorageItem represents a storage item.
type StorageItem struct {
	XMLName          xml.Name `xml:"storageItem"`
	AlarmType        string   `xml:"alarmType,attr,omitempty"`
	ChildAction      string   `xml:"childAction,attr,omitempty"`
	Name             string   `xml:"name,attr,omitempty"`
	OperationalState string   `xml:"operState,attr,omitempty"`
	Rn               string   `xml:"rn,attr,omitempty"`
	Size             int      `xml:"size,attr,omitempty"`
	Used             int      `xml:"used,attr,omitempty"`
}

// NetworkElement represents a physical network element, such as a Fabric Interconnect.
type NetworkElement struct {
	XMLName                   xml.Name             `xml:"networkElement"`
	AdminEvacState            string               `xml:"adminEvacState,attr,omitempty"`
	AdminInbandInterfaceState string               `xml:"adminInbandIfState,attr,omitempty"`
	ChildAction               string               `xml:"childAction,attr,omitempty"`
	DiffMemory                int                  `xml:"diffMemory,attr,omitempty"`
	Dn                        string               `xml:"dn,attr,omitempty"`
	ExpectedMemory            int                  `xml:"expectedMemory,attr,omitempty"`
	FltAggr                   int                  `xml:"fltAggr,attr,omitempty"`
	ForceEvac                 string               `xml:"forceEvac,attr,omitempty"`
	Id                        string               `xml:"id,attr,omitempty"`
	InbandInterfaceGateway    net.IP               `xml:"inbandIfGw,attr,omitempty"`
	InbandInterfaceIp         net.IP               `xml:"inbandIfIp,attr,omitempty"`
	InbandInterfaceNetmask    net.IP               `xml:"inbandIfMask,attr,omitempty"`
	InbandInterfaceVnet       int                  `xml:"inbandIfVnet,attr,omitempty"`
	InventoryStatus           string               `xml:"inventoryStatus,attr,omitempty"`
	MinActiveFan              int                  `xml:"minActiveFan,attr,omitempty"`
	Model                     string               `xml:"model,attr,omitempty"`
	OobInterfaceGateway       net.IP               `xml:"oobIfGw,attr,omitempty"`
	OobInterfaceIp            net.IP               `xml:"oobIfIp,attr,omitempty"`
	OobInterfaceNetmask       net.IP               `xml:"oobIfMask,attr,omitempty"`
	OobInterfaceMac           string               `xml:"oobIfMac,attr,omitempty"`
	OperEvacState             string               `xml:"operEvacState,attr,omitempty"`
	Operability               string               `xml:"operability,attr,omitempty"`
	Revision                  string               `xml:"revision,attr,omitempty"`
	Serial                    string               `xml:"serial,attr,omitempty"`
	ShutdownFanRemoval        string               `xml:"shutdownFanRemoveal,attr,omitempty"`
	Thermal                   string               `xml:"thermal,attr,omitempty"`
	TotalMemory               int                  `xml:"totalMemory,attr,omitempty"`
	Vendor                    string               `xml:"vendor,attr,omitempty"`
	FanModules                []EquipmentFanModule `xml:"equipmentFanModule"`
	ManagementController      ManagementController `xml:"mgmtController"`
	StorageItems              []StorageItem        `xml:"storageItem"`
}

// LsServer represents Service Profile Mo
type LsServer struct {
	XMLName                        xml.Name `xml:"lsServer"`
	AgentPolicyName                string   `xml:"agentPolicyName,attr,omitempty"`
	AssetTag                       string   `xml:"assetTag,attr,omitempty"`
	AssignState                    string   `xml:"assignState,attr,omitempty"`
	AssocState                     string   `xml:"assocState,attr,omitempty"`
	BiosProfileName                string   `xml:"biosProfileName,attr,omitempty"`
	BootPolicyName                 string   `xml:"bootPolicyName,attr,omitempty"`
	ConfigQualifier                string   `xml:"configQualifier,attr,omitempty"`
	ConfigState                    string   `xml:"configState,attr,omitempty"`
	Descr                          string   `xml:"descr,attr,omitempty"`
	Dn                             string   `xml:"dn,attr,omitempty"`
	DynamicConPolicyName           string   `xml:"dynamicConPolicyName,attr,omitempty"`
	ExtIPPoolName                  string   `xml:"extIPPoolName,attr,omitempty"`
	ExtIPState                     string   `xml:"extIPState,attr,omitempty"`
	FltAggr                        string   `xml:"fltAggr,attr,omitempty"`
	FsmDescr                       string   `xml:"fsmDescr,attr,omitempty"`
	FsmFlags                       string   `xml:"fsmFlags,attr,omitempty"`
	FsmPrev                        string   `xml:"fsmPrev,attr,omitempty"`
	FsmProgr                       string   `xml:"fsmProgr,attr,omitempty"`
	FsmRmtInvErrCode               string   `xml:"fsmRmtInvErrCode,attr,omitempty"`
	FsmRmtInvErrDescr              string   `xml:"fsmRmtInvErrDescr,attr,omitempty"`
	FsmRmtInvRslt                  string   `xml:"fsmRmtInvRslt,attr,omitempty"`
	FsmStageDescr                  string   `xml:"fsmStageDescr,attr,omitempty"`
	FsmStamp                       string   `xml:"fsmStamp,attr,omitempty"`
	FsmStatus                      string   `xml:"fsmStatus,attr,omitempty"`
	FsmTry                         string   `xml:"fsmTry,attr,omitempty"`
	GraphicsCardPolicyName         string   `xml:"graphicsCardPolicyName,attr,omitempty"`
	HostFwPolicyName               string   `xml:"hostFwPolicyName,attr,omitempty"`
	IdentPoolName                  string   `xml:"identPoolName,attr,omitempty"`
	IntId                          string   `xml:"intId,attr,omitempty"`
	KvmMgmtPolicyName              string   `xml:"kvmMgmtPolicyName,attr,omitempty"`
	LocalDiskPolicyName            string   `xml:"localDiskPolicyName,attr,omitempty"`
	MaintPolicyName                string   `xml:"maintPolicyName,attr,omitempty"`
	MgmtAccessPolicyName           string   `xml:"mgmtAccessPolicyName,attr,omitempty"`
	MgmtFwPolicyName               string   `xml:"mgmtFwPolicyName,attr,omitempty"`
	Name                           string   `xml:"name,attr,omitempty"`
	OperBiosProfileName            string   `xml:"operBiosProfileName,attr,omitempty"`
	OperBootPolicyName             string   `xml:"operBootPolicyName,attr,omitempty"`
	OperDynamicConPolicyName       string   `xml:"operDynamicConPolicyName,attr,omitempty"`
	OperExtIPPoolName              string   `xml:"operExtIPPoolName,attr,omitempty"`
	OperGraphicsCardPolicyName     string   `xml:"operGraphicsCardPolicyName,attr,omitempty"`
	OperHostFwPolicyName           string   `xml:"operHostFwPolicyName,attr,omitempty"`
	OperIdentPoolName              string   `xml:"operIdentPoolName,attr,omitempty"`
	OperKvmMgmtPolicyName          string   `xml:"operKvmMgmtPolicyName,attr,omitempty"`
	OperLocalDiskPolicyName        string   `xml:"operLocalDiskPolicyName,attr,omitempty"`
	OperMaintPolicyName            string   `xml:"operMaintPolicyName,attr,omitempty"`
	OperMgmtFwPolicyName           string   `xml:"operMgmtFwPolicyName,attr,omitempty"`
	OperPersistentMemoryPolicyName string   `xml:"operPersistentMemoryPolicyName,attr,omitempty"`
	OperPowerPolicyName            string   `xml:"operPowerPolicyName,attr,omitempty"`
	OperPowerSyncPolicyName        string   `xml:"operPowerSyncPolicyName,attr,omitempty"`
	OperScrubPolicyName            string   `xml:"operScrubPolicyName,attr,omitempty"`
	OperSolPolicyName              string   `xml:"operSolPolicyName,attr,omitempty"`
	OperState                      string   `xml:"operState,attr,omitempty"`
	OperStatsPolicyName            string   `xml:"operStatsPolicyName,attr,omitempty"`
	OperVconProfileName            string   `xml:"operVconProfileName,attr,omitempty"`
	OperVmediaPolicyName           string   `xml:"operVmediaPolicyName,attr,omitempty"`
	Owner                          string   `xml:"owner,attr,omitempty"`
	PersistentMemoryPolicyName     string   `xml:"persistentMemoryPolicyName,attr,omitempty"`
	PnDn                           string   `xml:"pnDn,attr,omitempty"`
	PolicyLevel                    string   `xml:"policyLevel,attr,omitempty"`
	PolicyOwner                    string   `xml:"policyOwner,attr,omitempty"`
	PowerPolicyName                string   `xml:"powerPolicyName,attr,omitempty"`
	PowerSyncPolicyName            string   `xml:"powerSyncPolicyName,attr,omitempty"`
	PropAcl                        string   `xml:"propAcl,attr,omitempty"`
	ResolveRemote                  string   `xml:"resolveRemote,attr,omitempty"`
	ScrubPolicyName                string   `xml:"scrubPolicyName,attr,omitempty"`
	SolPolicyName                  string   `xml:"solPolicyName,attr,omitempty"`
	SrcTemplName                   string   `xml:"srcTemplName,attr,omitempty"`
	StatsPolicyName                string   `xml:"statsPolicyName,attr,omitempty"`
	Status                         string   `xml:"status,attr,omitempty"`
	SvnicConfig                    string   `xml:"svnicConfig,attr,omitempty"`
	Type                           string   `xml:"type,attr,omitempty"`
	UsrLbl                         string   `xml:"usrLbl,attr,omitempty"`
	Uuid                           string   `xml:"uuid,attr,omitempty"`
	UuidSuffix                     string   `xml:"uuidSuffix,attr,omitempty"`
	VconProfileName                string   `xml:"vconProfileName,attr,omitempty"`
	VmediaPolicyName               string   `xml:"vmediaPolicyName,attr,omitempty"`
}

// FaultInst is getting injected into Mos in case of faults
type FaultInst struct {
	XMLName         xml.Name
	Ack             string   `xml:"ack,attr,omitempty"`
	Cause           string   `xml:"cause,attr,omitempty"`
	ChangeSet       string   `xml:"changeSet,attr,omitempty"`
	ChildAction     string   `xml:"childAction,attr,omitempty"`
	Code            string   `xml:"code,attr,omitempty"`
	Created         string   `xml:"created,attr,omitempty"`
	Descr           string   `xml:"descr,attr,omitempty"`
	HighestSeverity string   `xml:"highestSeverity,attr,omitempty"`
	Id              string   `xml:"id,attr,omitempty"`
	LastTransition  string   `xml:"lastTransition,attr,omitempty"`
	Lc              string   `xml:"lc,attr,omitempty"`
	Occur           string   `xml:"occur,attr,omitempty"`
	OrigSeverity    string   `xml:"origSeverity,attr,omitempty"`
	PrevSeverity    string   `xml:"prevSeverity,attr,omitempty"`
	Rn              string   `xml:"rn,attr,omitempty"`
	Rule            string   `xml:"rule,attr,omitempty"`
	Severity        string   `xml:"severity,attr,omitempty"`
	Status          string   `xml:"status,attr,omitempty"`
	Tags            string   `xml:"tags,attr,omitempty"`
	Type            string   `xml:"type,attr,omitempty"`
}

// LsBinding represents Service Profile Binding Mo
type LsBinding struct {
	XMLName            xml.Name   `xml:"lsBinding"`
	AdminAction        string     `xml:"adminAction,attr,omitempty"`
	AdminActionTrigger string     `xml:"adminActionTrigger,attr,omitempty"`
	AssignedToDn       string     `xml:"assignedToDn,attr,omitempty"`
	ComputeEpDn        string     `xml:"computeEpDn,attr,omitempty"`
	Dn                 string     `xml:"dn,attr,omitempty"`
	Issues             string     `xml:"issues,attr,omitempty"`
	Name               string     `xml:"name,attr,omitempty"`
	OperState          string     `xml:"operState,attr,omitempty"`
	PnDn               string     `xml:"pnDn,attr,omitempty"`
	PropAcl            string     `xml:"propAcl,attr,omitempty"`
	RestrictMigration  string     `xml:"restrictMigration,attr,omitempty"`
	FaultInst          *FaultInst `xml:"faultInst,omitempty"`
}

// LsBindingMo represents Service Profile Binding Mo to retrive from OutConfigs
type LsBindingMo struct {
	XMLName xml.Name
	LsBinding LsBinding `xml:"lsBinding"`
}

// LsBindingPair representes pair structure for lsBinding class
// returned in OutConfigs by ConfigConfMos
type LsBindingPair struct {
	XMLName xml.Name
	Key       string  `xml:"key,attr"`
	LsBinding LsBinding
}

// LsBindingPairs represents Pairs array for lsBinding class
type LsBindingPairs struct {
	XMLName xml.Name
        Pairs [] LsBindingPair `xml:"pair"`
}

// LsRequirement represents lsRequirement Mo
type LsRequirement struct {
	XMLName             xml.Name `xml:"lsRequirement"`
	AdminAction         string   `xml:"adminAction,attr,omitempty"`
	AdminActionTrigger  string   `xml:"adminActionTrigger,attr,omitempty"`
	AssignedToDn        string   `xml:"assignedToDn,attr,omitempty"`
	ComputeEpDn         string   `xml:"computeEpDn,attr,omitempty"`
	Dn                  string   `xml:"dn,attr,omitempty"`
	Issues              string   `xml:"issues,attr,omitempty"`
	Name                string   `xml:"name,attr,omitempty"`
	OperName            string   `xml:"operName,attr,omitempty"`
	OperState           string   `xml:"operState,attr,omitempty"`
	PnDn                string   `xml:"pnDn,attr,omitempty"`
	PnPoolDn            string   `xml:"pnPoolDn,attr,omitempty"`
	Qualifier           string   `xml:"qualifier,attr,omitempty"`
	RestrictMigration   string   `xml:"restrictMigration,attr,omitempty"`
	Status              string   `xml:"status,attr,omitempty"`
}

// LsIssues represents lsIssues Mo
type LsIssues struct {
	XMLName                xml.Name `xml:"lsIssues"`
	ConfigWarnings         string   `xml:"configWarnings,attr,omitempty"`
	Dn                     string   `xml:"dn,attr,omitempty"`
	IscsiConfigIssues      string   `xml:"iscsiConfigIssues,attr,omitempty"`
	NetworkConfigIssues    string   `xml:"networkConfigIssues,attr,omitempty"`
	PropAcl                string   `xml:"propAcl,attr,omitempty"`
	ServerConfigIssues     string   `xml:"serverConfigIssues,attr,omitempty"`
	ServerExtdConfigIssues string   `xml:"serverExtdConfigIssues,attr,omitempty"`
	Status                 string   `xml:"status,attr,omitempty"`
	StorageConfigIssues    string   `xml:"storageConfigIssues,attr,omitempty"`
	VnicConfigIssues       string   `xml:"vnicConfigIssues,attr,omitempty"`
}

// LsServerPair representes structure for lsServer class
// returned in OutConfigs/pair by orgResolveElements call and given Dn
type LsServerPair struct {
	XMLName xml.Name
	Key      string  `xml:"key,attr"`
	LsServer LsServer
}

// LsServerPairs represents Pairs array for lsServer class
type LsServerPairs struct {
	XMLName xml.Name
        Pairs   []LsServerPair `xml:"pair"`
}

// LsIssuesPairs representes structure for lsIssues class
// returned in OutConfigs/pair
type LsIssuesPair struct {
	XMLName xml.Name
	Key      string  `xml:"key,attr"`
	LsIssues LsIssues
}

// LsServerPairs represents Pairs array for lsServer class
type LsIssuesPairs struct {
	XMLName xml.Name
        Pairs [] LsIssuesPair `xml:"pair"`
}

// Represents generic Mo pair
type MoPair struct {
	XMLName xml.Name
	Key string `xml:"key,attr"`
	Mo  Any
}

// Represents generic Mo pairs
type MoPairs struct {
	XMLName xml.Name
        Pairs [] MoPair `xml:"pair"`
}

// Blades represents structure  for computeBlade class
// returned in OutConfigs by ConfigResolveClass call
type Blades struct {
	XMLName xml.Name
	ComputeBlades  []ComputeBlade `xml:"computeBlade"`
}

// LsPower represents structure for lsPower class
// submitted in OutConfigs for ConfigConfMol call
type LsPower struct {
	XMLName           xml.Name
	Dn                string   `xml:"dn,attr,omitempty"`
	PropAcl           string   `xml:"propAcl,attr,omitempty"`
	SoftShutdownTimer string   `xml:"softShutdownTimer,attr,omitempty"`
	State             string   `xml:"state,attr,omitempty"`
}

// LsPower represents structure for lsPower class
// returned in OutConfigs by ConfigResolveChildren call
type LsPowerMo struct {
	XMLName xml.Name
	LsPower LsPower `xml:"lsPower"`
}

// LsServerMo represents structure for lsServer class
// used in configConfMo by ConfigConfMo call
type LsServerMo struct {
	XMLName xml.Name
	LsServer LsServer `xml:"lsServer"`
}

// MemoryQual represents structure for memoryQual class
// returned in OutConfigs by ConfigResolveChildren call
type MemoryQual struct {
	XMLName xml.Name
	Memory struct {
		Clock   string `xml:"clock,attr,omitempty"`
		Dn      string `xml:"dn,attr,omitempty"`
		Latency string `xml:"latency,attr,omitempty"`
		MaxCap  string `xml:"maxCap,attr,omitempty"`
		MinCap  string `xml:"minCap,attr,omitempty"`
		Speed   string `xml:"speed,attr,omitempty"`
		Units   string `xml:"units,attr,omitempty"`
		Width   string `xml:"width,attr,omitempty"`
	} `xml:"memoryQual"`
}

// ProcessorQual represents structure for processorQual class
// returned in OutConfigs by ConfigResolveChildren call
type ProcessorQual struct {
	XMLName xml.Name
	Processor struct {
		Arch       string `xml:"arch,attr,omitempty"`
		Dn         string `xml:"dn,attr,omitempty"`
		MaxCores   string `xml:"maxCores,attr,omitempty"`
		MaxProcs   string `xml:"maxProcs,attr,omitempty"`
		MaxThreads string `xml:"maxThreads,attr,omitempty"`
		MinCores   string `xml:"minCores,attr,omitempty"`
		MinProcs   string `xml:"minProcs,attr,omitempty"`
		MinThreads string `xml:"minThreads,attr,omitempty"`
		Model      string `xml:"model,attr,omitempty"`
		Speed      string `xml:"speed,attr,omitempty"`
		Stepping   string `xml:"stepping,attr,omitempty"`
	} `xml:"processorQual"`
}

// ComputePhysicalQual represents structure for computePhysicalQual class
// returned in OutConfigs by ConfigResolveChildren call
type ComputePhysicalQual struct {
	XMLName xml.Name
	Physical struct {
		Dn      string `xml:"dn,attr,omitempty"`
		Model   string `xml:"model,attr,omitempty"`
		PropAcl string `xml:"propAcl,attr,omitempty"`
	} `xml:"computePhysicalQual"`
}

// VnicsEither is used to obtain Ether vnics via ConfigResolveChildren call
type VnicsEther struct {
	XMLName xml.Name
	Vnics []VnicEther `xml:"vnicEther"`
}

// VnicEther represents structure for vnicEther class
// returned in OutConfigs by ConfigResolveChildren call
type VnicEther struct {
	AdaptorProfileName       string `xml:"adaptorProfileName,attr,omitempty"`
	Addr                     string `xml:"addr,attr,omitempty"`
	AdminCdnName             string `xml:"adminCdnName,attr,omitempty"`
	AdminHostPort            string `xml:"adminHostPort,attr,omitempty"`
	AdminVcon                string `xml:"adminVcon,attr,omitempty"`
	BootDev                  string `xml:"bootDev,attr,omitempty"`
	CdnPropInSync            string `cdnPropInSyncxml:",attr,omitempty"`
	CdnSource                string `xml:"cdnSource,attr,omitempty"`
	ConfigQualifier          string `xml:"configQualifier,attr,omitempty"`
	ConfigState              string `xml:"configState,attr,omitempty"`
	Dn                       string `xml:"dn,attr,omitempty"`
	DynamicId                string `xml:"dynamicId,attr,omitempty"`
	EquipmentDn              string `xml:"equipmentDn,attr,omitempty"`
	FltAggr                  string `xml:"fltAggr,attr,omitempty"`
	IdentPoolName            string `xml:"identPoolName,attr,omitempty"`
	InstType                 string `xml:"instType,attr,omitempty"`
	Mtu                      string `xml:"mtu,attr,omitempty"`
	Name                     string `xml:"name,attr,omitempty"`
	NwCtrlPolicyName         string `xml:"nwCtrlPolicyName,attr,omitempty"`
	NwTemplName              string `xml:"nwTemplName,attr,omitempty"`
	OperAdaptorProfileName   string `xml:"operAdaptorProfileName,attr,omitempty"`
	OperCdnName              string `xml:"operCdnName,attr,omitempty"`
	OperHostPort             string `xml:"operHostPort,attr,omitempty"`
	OperIdentPoolName        string `xml:"operIdentPoolName,attr,omitempty"`
	OperNwCtrlPolicyName     string `xml:"operNwCtrlPolicyName,attr,omitempty"`
	OperNwTemplName          string `xml:"operNwTemplName,attr,omitempty"`
	OperOrder                string `xml:"operOrder,attr,omitempty"`
	OperPinToGroupName       string `xml:"operPinToGroupName,attr,omitempty"`
	OperQosPolicyName        string `xml:"operQosPolicyName,attr,omitempty"`
	OperSpeed                string `xml:"operSpeed,attr,omitempty"`
	OperStatsPolicyName      string `xml:"operStatsPolicyName,attr,omitempty"`
	OperVcon                 string `xml:"operVcon,attr,omitempty"`
	Order                    string `xml:"order,attr,omitempty"`
	Owner                    string `xml:"owner,attr,omitempty"`
	PfDn                     string `xml:"pfDn,attr,omitempty"`
	PinToGroupName           string `xml:"pinToGroupName,attr,omitempty"`
	PropAcl                  string `xml:"propAcl,attr,omitempty"`
	Purpose                  string `xml:"purpose,attr,omitempty"`
	QosPolicyName            string `xml:"qosPolicyName,attr,omitempty"`
	RedundancyPairType       string `xml:"redundancyPairType,attr,omitempty"`
	RedundancyPeer           string `xml:"redundancyPeer,attr,omitempty"`
	StatsPolicyName          string `xml:"statsPolicyName,attr,omitempty"`
	SwitchId                 string `xml:"switchId,attr,omitempty"`
	Type                     string `xml:"type,attr,omitempty"`
	VirtualizationPreference string `xml:"virtualizationPreference,attr,omitempty"`
}

// VnicIPv4IscsiAddr represents nested structure inside of VnicIPv4If
type VnicIPv4IscsiAddr struct {
	XMLName xml.Name
	Addr        string `xml:"addr,attr,omitempty"`
	ChildAction string `xml:"childAction,attr,omitempty"`
	DefGw       string `xml:"defGw,attr,omitempty"`
	PrimDns     string `xml:"primDns,attr,omitempty"`
	Rn          string `xml:"rn,attr,omitempty"`
	SecDns      string `xml:"secDns,attr,omitempty"`
	Subnet      string `xml:"subnet,attr,omitempty"`
}

// VnicIPv4Dhcp represents nested structure inside of VnicIPv4If
type VnicIPv4Dhcp struct {
	XMLName xml.Name
	Addr   string `xml:"addr,attr,omitempty"`
	DefGw  string `xml:"defGw,attr,omitempty"`
	Dn     string `xml:"dn,attr,omitempty"`
	Status string `xml:"status,attr,omitempty"`
	Subnet string `xml:"subnet,attr,omitempty"`
}

// VnicIPv4If represents nested structure inside of VnicVlan
type VnicIPv4If struct {
	XMLName xml.Name
	Addr                string            `xml:"addr,attr,omitempty"`
	ChildAction         string            `xml:"childAction,attr,omitempty"`
	ConfigQualifier     string            `xml:"configQualifier,attr,omitempty"`
	Name                string            `xml:"name,attr,omitempty"`
	OperPrimaryVnetDn   string            `xml:"operPrimaryVnetDn,attr,omitempty"`
	OperPrimaryVnetName string            `xml:"operPrimaryVnetName,attr,omitempty"`
	OperState           string            `xml:"operState,attr,omitempty"`
	OperVnetDn          string            `xml:"operVnetDn,attr,omitempty"`
	OperVnetName        string            `xml:"operVnetName,attr,omitempty"`
	Owner               string            `xml:"owner,attr,omitempty"`
	PropAcl             string            `xml:"propAcl,attr,omitempty"`
	PubNwId             string            `xml:"pubNwId,attr,omitempty"`
	Rn                  string            `xml:"rn,attr,omitempty"`
	Sharing             string            `xml:"sharing,attr,omitempty"`
	SwitchId            string            `xml:"switchId,attr,omitempty"`
	Type                string            `xml:"type,attr,omitempty"`
	VnicIPv4IscsiAddr   VnicIPv4IscsiAddr `xml:"vnicIPv4IscsiAddr"`
	VnicIPv4Dhcp        VnicIPv4Dhcp      `xml:"vnicIPv4Dhcp,omitempty"`
}

// VnicLun represents nested structure inside of VnicIScsiStaticTargetIf
type VnicLun struct {
	XMLName xml.Name
	Bootable    string `xml:"bootable,attr,omitempty"`
	ChildAction string `xml:"childAction,attr,omitempty"`
	Id          string `xml:"id,attr,omitempty"`
	Rn          string `xml:"rn,attr,omitempty"`
}

// VnicIScsiStaticTargetIf represents nested structure inside of VnicVlan
type VnicIScsiStaticTargetIf struct {
	XMLName xml.Name
	AuthProfileName     string    `xml:"authProfileName,attr,omitempty"`
	ChildAction         string    `xml:"childAction,attr,omitempty"`
	IpAddress           string    `xml:"ipAddress,attr,omitempty"`
	Name                string    `xml:"name,attr,omitempty"`
	OperAuthProfileName string    `xml:"operAuthProfileName,attr,omitempty"`
	Port                string    `xml:"port,attr,omitempty"`
	Priority            string    `xml:"priority,attr,omitempty"`
	PropAcl             string    `xml:"propAcl,attr,omitempty"`
	Rn                  string    `xml:"rn,attr,omitempty"`
	Status              string    `xml:"status,attr,omitempty"`
	VnicLuns            []VnicLun `xml:"vnicLun,omitempty"`
}

// VnicVlan represents nested structure inside of VnicIScsi
type VnicVlan struct {
	XMLName xml.Name
	ChildAction              string                    `xml:"childAction,attr,omitempty"`
	ConfigQualifier          string                    `xml:"configQualifier,attr,omitempty"`
	FltAggr                  string                    `xml:"fltAggr,attr,omitempty"`
	Name                     string                    `xml:"name,attr,omitempty"`
	OperPrimaryVnetDn        string                    `xml:"operPrimaryVnetDn,attr,omitempty"`
	OperPrimaryVnetName      string                    `xml:"operPrimaryVnetName,attr,omitempty"`
	OperState                string                    `xml:"operState,attr,omitempty"`
	OperVnetDn               string                    `xml:"operVnetDn,attr,omitempty"`
	OperVnetName             string                    `xml:"operVnetName,attr,omitempty"`
	Owner                    string                    `xml:"owner,attr,omitempty"`
	PubNwId                  string                    `xml:"pubNwId,attr,omitempty"`
	Rn                       string                    `xml:"rn,attr,omitempty"`
	Sharing                  string                    `xml:"sharing,attr,omitempty"`
	SwitchId                 string                    `xml:"switchId,attr,omitempty"`
	Type                     string                    `xml:"type,attr,omitempty"`
	VlanName                 string                    `xml:"vlanName,attr,omitempty"`
	Vnet                     string                    `xml:"vnet,attr,omitempty"`
	VnicIPv4If               VnicIPv4If                `xml:"vnicIPv4If"`
	VnicIScsiStaticTargets   []VnicIScsiStaticTargetIf `xml:"vnicIScsiStaticTargetIf"`
}

// VnicIScsi represents structure for vnicIScsi class
// returned in OutConfigs by ConfigResolveChildren call and
// used in InCounfigs by ConfigConfMos call to set iSCSI interfaces
type VnicIScsi struct {
	XMLName                xml.Name  `xml:"vnicIScsi"`
	AdaptorProfileName     string    `xml:"adaptorProfileName,attr,omitempty"`
	Addr                   string    `xml:"addr,attr,omitempty"`
	AdminCdnName           string    `xml:"adminCdnName,attr,omitempty"`
	AdminHostPort          string    `xml:"adminHostPort,attr,omitempty"`
	AdminVcon              string    `xml:"adminVcon,attr,omitempty"`
	AuthProfileName        string    `xml:"authProfileName,attr,omitempty"`
	BootDev                string    `xml:"bootDev,attr,omitempty"`
	CdnPropInSync          string    `xml:"cdnPropInSync,attr,omitempty"`
	CdnSource              string    `xml:"cdnSource,attr,omitempty"`
	ChildAction            string    `xml:"childAction,attr,omitempty"`
	ConfigQualifier        string    `xml:"configQualifier,attr,omitempty"`
	ConfigState            string    `xml:"configState,attr,omitempty"`
	Dn                     string    `xml:"dn,attr,omitempty"`
	EquipmentDn            string    `xml:"equipmentDn,attr,omitempty"`
	EthEpDn                string    `xml:"ethEpDn,attr,omitempty"`
	ExtIPState             string    `xml:"extIPState,attr,omitempty"`
	FltAggr                string    `xml:"fltAggr,attr,omitempty"`
	IdentPoolName          string    `xml:"identPoolName,attr,omitempty"`
	InitNameSuffix         string    `xml:"initNameSuffix,attr,omitempty"`
	InitiatorName          string    `xml:"initiatorName,attr,omitempty"`
	InstType               string    `xml:"instType,attr,omitempty"`
	IqnIdentPoolName       string    `xml:"iqnIdentPoolName,attr,omitempty"`
	Name                   string    `xml:"name,attr,omitempty"`
	NwTemplName            string    `xml:"nwTemplName,attr,omitempty"`
	OperAdaptorProfileName string    `xml:"operAdaptorProfileName,attr,omitempty"`
	OperAuthProfileName    string    `xml:"operAuthProfileName,attr,omitempty"`
	OperCdnName            string    `xml:"operCdnName,attr,omitempty"`
	OperHostPort           string    `xml:"operHostPort,attr,omitempty"`
	OperIdentPoolName      string    `xml:"operIdentPoolName,attr,omitempty"`
	OperIqnIdentPoolName   string    `xml:"operIqnIdentPoolName,attr,omitempty"`
	OperOrder              string    `xml:"operOrder,attr,omitempty"`
	OperSpeed              string    `xml:"operSpeed,attr,omitempty"`
	OperStatsPolicyName    string    `xml:"operStatsPolicyName,attr,omitempty"`
	OperVcon               string    `xml:"operVcon,attr,omitempty"`
	Order                  string    `xml:"order,attr,omitempty"`
	Owner                  string    `xml:"owner,attr,omitempty"`
	PinToGroupName         string    `xml:"pinToGroupName,attr,omitempty"`
	PropAcl                string    `xml:"propAcl,attr,omitempty"`
	QosPolicyName          string    `xml:"qosPolicyName,attr,omitempty"`
	RedundancyPairType     string    `xml:"redundancyPairType,attr,omitempty"`
	RedundancyPeer         string    `xml:"redundancyPeer,attr,omitempty"`
	StatsPolicyName        string    `xml:"statsPolicyName,attr,omitempty"`
	SwitchId               string    `xml:"switchId,attr,omitempty"`
	Type                   string    `xml:"type,attr,omitempty"`
	VnicDefType            string    `xml:"vnicDefType,attr,omitempty"`
	VnicName               string    `xml:"vnicName,attr,omitempty"`
	VnicVlan               VnicVlan  `xml:"vnicVlan"`
}

// VnicsIScsi is used to obtain iSCSI vnics via ConfigResolveChildren call
type VnicsIScsi struct {
	XMLName xml.Name
	Vnics []VnicIScsi `xml:"vnicIScsi"`
}

// VnicsIScsiPair representes structure submitted in InConfigs
// for ConfigConfMos call to set iSCSI interfaces
type VnicIScsiPair struct {
	XMLName xml.Name
	Key        string     `xml:"key,attr"`
	VnicIScsi VnicIScsi
}

// VnicIScsiPairs represents Pairs array submitted in InConfigs and
// ConfigConfMos result retrieved in OutConfigs
type VnicIScsiPairs struct {
	XMLName xml.Name
        Pairs   []VnicIScsiPair `xml:"pair"`
}
