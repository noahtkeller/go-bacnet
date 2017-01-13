package bacnet

type Service interface {
	Encode() Service
	Decode() Service
}

func NewBACNET_ADDRESS() *BACNET_ADDRESS {
	self := &BACNET_ADDRESS{}
	self.Mac = make([]byte, MAX_MAC_LEN)
	self.Adr = make([]byte, MAX_MAC_LEN)
	return self
}

type BACNET_ADDRESS struct {
				   /* mac_len = 0 is a broadcast address */
	Mac_len byte
				   /* note: MAC for IP addresses uses 4 bytes for addr, 2 bytes for port */
				   /* use de/encode_unsigned32/16 for re/storing the IP address */
	Mac     []byte
				   /* DNET,DLEN,DADR or SNET,SLEN,SADR */
				   /* the following are used if the device is behind a router */
				   /* net = 0 indicates local */
	Net     uint16 /* BACnet network number */
				   /* LEN = 0 denotes broadcast MAC ADR and ADR field is absent */
				   /* LEN > 0 specifies length of ADR field */
	Len     byte   /* length of MAC address */
	Adr     []byte /* hwaddr (MAC) address */
}


/* define a MAC address for manipulation */
type BACNET_MAC_ADDRESS struct {
	Len byte /* length of MAC address */
	Adr []byte
}

/* note: with microprocessors having lots more code space than memory,
   it might be better to have a packed encoding with a library to
   easily access the data. */
type BACNET_OBJECT_ID struct {
	Type     uint16
	Instance uint32
}

type BACNET_NPDU_DATA struct {
	Protocol_version      byte
								 /* parts of the control octet: */
	Data_expecting_reply  bool
	Network_layer_message bool   /* false if APDU */
	Priority              byte
								 /* optional network message info */
	Network_message_type  byte   /* optional */
	Vendor_id             uint16 /* optional, if net message type is > 0x80 */
	Hop_count             byte
}

type BACNET_ROUTER_PORT struct {
	Dnet     uint16              /**< The DNET number that identifies this port. */
	Id       byte                /**< Either 0 or some ill-defined, meaningless value. */
	Info     []byte              /**< Info like 'modem dialing string' */
	Info_len byte                /**< Length of info[]. */
	Next     *BACNET_ROUTER_PORT /**< Point to next in linked list */
}

//todo: start find right value for these
const MAX_APDU = 1476
const MAX_CHARACTER_STRING_BYTES = MAX_APDU - 6
const CHARACTER_STRING_CAPACITY = MAX_CHARACTER_STRING_BYTES - 1
//todo: end   find right value for these

const BACNET_STATUS_OK int = 0
const BACNET_STATUS_ERROR int = -1
const BACNET_STATUS_ABORT int = -2
const BACNET_STATUS_REJECT int = -3

const BACNET_BROADCAST_NETWORK uint16 = 0xFFFF
const BACNET_MAX_INSTANCE int32 = 0x3FFFFF

const BACNET_INSTANCE_BITS int = 22
const BACNET_MAX_OBJECT uint32 = 0x3FF

const HOP_COUNT_DEFAULT byte = 0xFF
const ROUTER_PORT_INFO_LEN byte = 0x02

const MAX_MAC_LEN = 0x07

const BIT0 byte = 0x01
const BIT1 byte = 0x02
const BIT2 byte = 0x04
const BIT3 byte = 0x08
const BIT4 byte = 0x10
const BIT5 byte = 0x20
const BIT6 byte = 0x40
const BIT7 byte = 0x80
