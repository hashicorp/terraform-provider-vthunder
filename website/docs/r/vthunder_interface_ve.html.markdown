---
layout: "vthunder"
page_title: "vthunder: vthunder_interface_ve"
sidebar_current: "docs-vthunder-resource-interface-ve"
description: |-
	Provides details about vthunder interface ve resource for A10
---

# vthunder\_interface\_ve

`vthunder_interface_ve` Provides details about vthunder interface ve


## Argument Reference

* `action` - ‘enable’: Enable Router Advertisements on this interface; ‘disable’: Disable Router Advertisements on this interface;
* `ifnum` - Virtual ethernet interface number
* `l3_vlan_fwd_disable` - Disable L3 forwarding between VLANs for incoming packets on this interface
* `mtu` - OSPF interface MTU (MTU size)
* `name` - Name for the interface
* `trap_source` - The trap source
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `inside` - DDoS inside (trusted) or outside (untrusted) interface
* `map_t_inside` - Configure MAP inside interface (connected to MAP domains)
* `map_t_outside` - Configure MAP outside interface
* `outside` - DDoS inside (trusted) or outside (untrusted) interface
* `lockup` - Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)
* `lockup_period` - Lockup period (second)
* `normal` - Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit
* `circuit_type` - ‘level-1’: Level-1 only adjacencies are formed; ‘level-1-2’: Level-1-2 adjacencies are formed; ‘level-2-only’: Level-2 only adjacencies are formed;
* `lsp_interval` - Set LSP transmission interval (LSP transmission interval (milliseconds))
* `network` - ‘broadcast’: Specify IS-IS broadcast multi-access network; ‘point-to-point’: Specify IS-IS point-to-point network;
* `padding` - Add padding to IS-IS hello packets
* `retransmit_interval` - Time between retransmitting lost link state advertisements (Seconds)
* `level` - ‘level-1’: Speficy interval for level-1 CSNPs; ‘level-2’: Specify interval for level-2 CSNPs;
* `priority` - Router priority
* `hello_interval_minimal` - Set Hello holdtime 1 second, interval depends on multiplier
* `blocked` - Block LSPs on this interface
* `value` - ‘message-digest’: Use message-digest authentication; ‘null’: Use no authentication;
* `bfd` - Bidirectional Forwarding Detection (BFD)
* `disable` - Disable BFD
* `password` - Key String
* `key_chain` - Authentication key-chain (Name of key-chain)
* `mode` - ‘md5’: Keyed message digest; ‘text’: Clear text authentication;
* `send_only` - Authentication send-only
* `wide_metric` - Configure the wide metric for interface
* `hello_interval` - Time between HELLO packets (Seconds)
* `hello_multiplier` - Set multiplier for Hello holding time (Hello multiplier value)
* `metric` - Configure the metric for interface (Default metric)
* `csnp_interval` - Set CSNP interval in seconds (CSNP interval value)
* `demand` - Demand mode
* `echo` - Enable BFD Echo
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `key_id` - Key ID
* `method` - ‘md5’: Keyed MD5; ‘meticulous-md5’: Meticulous Keyed MD5; ‘meticulous-sha1’: Meticulous Keyed SHA1; ‘sha1’: Keyed SHA1; ‘simple’: Simple Password;
* `interval` - Transmit interval between BFD packets (Milliseconds)
* `min_rx` - Minimum receive interval capability (Milliseconds)
* `multiplier` - Multiplier value used to compute holddown (value used to multiply the interval)
* `allow_promiscuous_vip` - Allow traffic to be associated with promiscuous VIP
* `client` - Client facing interface for IPv4/v6 traffic
* `dhcp` - Use DHCP to configure IP address
* `generate_membership_query` - Enable Membership Query
* `max_resp_time` - Maximum Response Time (Max Response Time (Default is 100))
* `query_interval` - 1 - 255 (Default is 125)
* `server` - Server facing interface for IPv4/v6 traffic
* `slb_partition_redirect` - Redirect SLB traffic across partition
* `ttl_ignore` - Ignore TTL decrement for a received packet
* `ipv4_address` - IP address
* `ipv4_netmask` - IP subnet mask
* `tag` - ISO routing area tag
* `helper_address` - Helper address for DHCP packets (IP address)
* `access_list` - Access-list for traffic from the outside
* `acl_id` - ACL id
* `class_list` - Class List (Class List Name)
* `receive_packet` - Enable receiving packet through the specified interface
* `send_packet` - Enable sending packets through the specified interface
* `receive` - Advertisement reception
* `version` - ‘1’: RIP version 1; ‘2’: RIP version 2; ‘1-compatible’: RIPv1-compatible; ‘1-2’: RIP version 1 & 2;
* `state` - ‘poisoned’: Perform split horizon with poisoned reverse; ‘disable’: Disable split horizon; ‘enable’: Perform split horizon without poisoned reverse;
* `string` - The RIP authentication string
* `send` - Advertisement transmission
* `authentication` - Enable authentication
* `authentication_key` - Authentication password (key) (The OSPF password (key))
* `cost` - Interface cost
* `database_filter` - ‘all’: Filter all LSA;
* `dead_interval` - Interval after which a neighbor is declared dead (Seconds)
* `ip_addr` - Address of interface
* `mtu_ignore` - Ignores the MTU in DBD packets
* `out` - Outgoing LSA
* `transmit_delay` - Link state transmit delay (Seconds)
* `md5_value` - The OSPF password (1-16)
* `message_digest_key` - Message digest authentication password (key) (Key id)
* `broadcast` - Specify OSPF broadcast multi-access network
* `non_broadcast` - Specify OSPF NBMA network
* `p2mp_nbma` - Specify non-broadcast point-to-multipoint network
* `point_to_multipoint` - Specify OSPF point-to-multipoint network
* `point_to_point` - Specify OSPF point-to-point network
* `bind_type` - ‘inside’: This interface is connected to NPTv6 inside networks; ‘outside’: This interface is connected to NPTv6 outside networks;
* `domain_name` - NPTv6 domain name
* `lockup_period_v6` - Lockup period (second)
* `lockup_v6` - Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)
* `normal_v6` - Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit
* `acl_name` - Access-list Name
* `counters1` - ‘all’: all; ‘num_pkts’: Input packets; ‘num_total_bytes’: Input bytes; ‘num_unicast_pkts’: Received unicasts; ‘num_broadcast_pkts’: Received braodcasts; ‘num_multicast_pkts’: Received multicasts; ‘num_tx_pkts’: Transmitted packtes; ‘num_total_tx_bytes’: Transmitte bytes; ‘num_unicast_tx_pkts’: Trasnmitted unicasts; ‘num_broadcast_tx_pkts’: Transmitted broadcasts; ‘num_multicast_tx_pkts’: Transmitted multicasts; ‘rate_pkt_sent’: Packet sent rate packets/sec; ‘rate_byte_sent’: Byte sent rate bits/sec; ‘rate_pkt_rcvd’: Packet received rate packets/sec; ‘rate_byte_rcvd’: Byte received rate bits/sec; ‘load_interval’: Load Interval;
* `inbound` - ACL applied on incoming packets to this interface
* `ipv6_enable` - Enable IPv6 processing
* `v6_acl_name` - Apply ACL rules to incoming packets on this interface (Named Access List)
* `address_type` - ‘anycast’: Configure an IPv6 anycast address; ‘link-local’: Configure an IPv6 link local address;
* `ipv6_addr` - Set the IPv6 address of an interface
* `rip` - RIP Routing for IPv6
* `area_id_addr` - OSPF area ID in IP address format
* `area_id_num` - OSPF area ID as a decimal value
* `instance_id` - Specify the interface instance ID
* `broadcast_type` - ‘broadcast’: Specify OSPF broadcast multi-access network; ‘non-broadcast’: Specify OSPF NBMA network; ‘point-to-point’: Specify OSPF point-to-point network; ‘point-to-multipoint’: Specify OSPF point-to-multipoint network;
* `network_instance_id` - Specify the interface instance ID
* `neig_inst` - Specify the interface instance ID
* `neighbor` - OSPFv3 neighbor (Neighbor IPv6 address)
* `neighbor_cost` - OSPF cost for point-to-multipoint neighbor (metric)
* `neighbor_poll_interval` - OSPF dead-router polling interval (Seconds)
* `neighbor_priority` - OSPF priority of non-broadcast neighbor
* `adver_mtu` - Set Router Advertisement MTU Option
* `adver_mtu_disable` - Disable Router Advertisement MTU Option
* `adver_vrid` - Vrid
* `adver_vrid_default` - Default VRRP-A vrid
* `default_lifetime` - Set Router Advertisement Default Lifetime (default: 1800) (Default Lifetime (seconds))
* `floating_ip` - Use a floating IP as the source address for Router advertisements
* `floating_ip_default_vrid` - Use a floating IP as the source address for Router advertisements
* `hop_limit` - Set Router Advertisement Hop Limit (default: 255)
* `managed_config_action` - ‘enable’: Enable the Managed Address Configuration flag; ‘disable’: Disable the Managed Address Configuration flag (default);
* `max_interval` - Set Router Advertisement Max Interval (default: 600) (Max Router Advertisement Interval (seconds))
* `min_interval` - Set Router Advertisement Min Interval (default: 200) (Min Router Advertisement Interval (seconds))
* `other_config_action` - ‘enable’: Enable the Other Stateful Configuration flag; ‘disable’: Disable the Other Stateful Configuration flag (default);
* `rate_limit` - Rate Limit the processing of incoming Router Solicitations (Max Number of Router Solicitations to process per second)
* `reachable_time` - Set Router Advertisement Reachable ime (default: 0) (Reachable Time (milliseconds))
* `retransmit_timer` - Set Router Advertisement Retransmit Timer (default: 0)
* `use_floating_ip` - Use a floating IP as the source address for Router advertisements
* `use_floating_ip_default_vrid` - Use a floating IP as the source address for Router advertisements
* `not_autonomous` - Specify that the Prefix is not usable for autoconfiguration (default:autonomous)
* `not_on_link` - Specify that the Prefix is not On-Link (default: on-link)
* `preferred_lifetime` - Specify Prefix Preferred Lifetime (default:604800) (Prefix Advertised Preferred Lifetime (default: 604800))
* `prefix` - Set Router Advertisement On-Link Prefix (IPv6 On-Link Prefix)
* `valid_lifetime` - Specify Valid Lifetime (default:2592000) (Prefix Advertised Valid Lifetime (default: 2592000))

