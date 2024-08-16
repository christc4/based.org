# RAID (Redundant Array of Independent Disks)

Up: [Hard Disk Drives](hard_disk_drives)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Computer Systems

### Different types

#### RAID 0 (speed)

Data is stripped across all disks (each disk holds a different part of the same piece of data)

Can be good for performance but if one disk goes down, all data becomes corrupted

#### RAID 1 (safety)

Data is mirrored between disks, 2nd disk acts as a backup

Halves storage capacity

#### RAID 5

N-1 disks worth of storage for the data with the final 'disk' acting as a parity so that the data can be restored using parity (the parity is striped across all disks)

If the parity disk goes out, its GG once another goes out


#### RAID 6

Same as RAID 5 but 2 disks are used for parity and N-2 used for actual storage.

What is the difference between RAID0, 1, and 5? #flashcard 
RAID0 - stripes data across all disks (each disk holds a different part of the same piece of data)
RAID1 - Mirrors data between disks, halves storage capacity
RAID5 - N-1 disk space, the final 'disk' acts as parity for data so that it can be recovered
<!--ID: 1701369224380-->


































#### Why:
#### How:









