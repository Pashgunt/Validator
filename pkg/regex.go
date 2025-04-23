package pkg

type RegexType string

const (
	email      = `^[a-zA-Z0-9._]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`
	url        = `^(https?:\/\/)?([\w-]+(\.[\w-]+)+)(:[0-9]{1,5})?(\/[^\s]*)?$`
	macAddress = `^([0-9A-Fa-f]{2}[:-]){5}[0-9A-Fa-f]{2}$`
	uuid       = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`
	hostname   = `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)*[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?$`
	ipv4       = `^((25[0-5]|2[0-4][0-9]|1?[0-9]{1,2})\.){3}(25[0-5]|2[0-4][0-9]|1?[0-9]{1,2})$`
)

const (
	Email      RegexType = "Email"
	Url        RegexType = "Url"
	MacAddress RegexType = "MacAddress"
	Uuid       RegexType = "Uuid"
	Hostname   RegexType = "Hostname"
	IPv4       RegexType = "IPv4"
)

var RegexTypeAssoc = map[RegexType]string{
	Email:      email,
	Url:        url,
	MacAddress: macAddress,
	Uuid:       uuid,
	Hostname:   hostname,
	IPv4:       ipv4,
}
