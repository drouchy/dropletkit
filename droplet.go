package dropletkit

type Region struct {
  slug string
  name string
  size []string
  available bool
  features []string
}

type Image struct {
  Id int
  Name string
  Distribution string
  Slug string
  Public bool
  Regions  []string
  Created_at string
}

type Kernel struct {
  Id int
  Name string
  Version string
}

type Droplet struct {
  Id int
  Name string
  Memory int
  Vcpus int
  Disk int
  Region Region
  Image Image
  SizeSlug string
  Locked bool
  Status string
  Kernel Kernel
  Networks string
  CreateAt string
  Features []string
  BackupIds []string
  SnapshotIds []int
  ActionIds []string
}
// {
//       "networks": {
//         "v4": [
//           {
//             "ip_address": "10.0.0.19",
//             "netmask": "255.255.0.0",
//             "gateway": "10.0.0.1",
//             "type": "private"
//           },
//           {
//             "ip_address": "127.0.0.19",
//             "netmask": "255.255.255.0",
//             "gateway": "127.0.0.20",
//             "type": "public"
//           }
//         ],
//         "v6": [
//           {
//             "ip_address": "2001::13",
//             "cidr": 124,
//             "gateway": "2400:6180:0000:00D0:0000:0000:0009:7000",
//             "type": "public"
//           }
//         ]
//       },
